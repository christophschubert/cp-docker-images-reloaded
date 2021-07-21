package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
	"text/template"
	"time"
)

func ensureAtLeastOne(keys []string) bool {
	for _, varName := range keys {
		if _, found := os.LookupEnv(varName); found {
			return true
		}
	}
	return false
}

/**
Converts an environment variable name to a property-name according to the following rules:
- a single underscore (_) is replaced with a .
- a double underscore (__) is replaced with a single underscore
- a triple underscore (___) is replaced with a dash
Moreover, the whole string is converted to lower-case.
The behavior of sequences of four or more underscores is undefined.
*/
func convertKey(key string) string {
	re := regexp.MustCompile("[^_]_[^_]")
	singleReplaced := re.ReplaceAllStringFunc(key, replaceUnderscores)
	singleTripleReplaced := strings.ReplaceAll(singleReplaced, "___", "-")
	return strings.ToLower(strings.ReplaceAll(singleTripleReplaced, "__", "_"))
}

//helper function to replace every underscore '_' by a dot '.'
func replaceUnderscores(s string) string {
	return strings.ReplaceAll(s, "_", ".")
}

type ConfigSpec struct {
	Prefixes map[string]bool   `json:"prefixes"`
	Excludes []string          `json:"excludes"`
	Renamed  map[string]string `json:"renamed"`
	Defaults map[string]string `json:"defaults"`
}

func Contains(slice []string, element string) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

func ListToMap(kvList []string) map[string]string {
	m := make(map[string]string)
	for _, l := range kvList {
		parts := strings.Split(l, "=")
		if len(parts) == 2 {
			m[parts[0]] = parts[1]
		}
	}
	return m
}

func KvStringToMap(kvString string, sep string) map[string]string {
	return ListToMap(strings.Split(kvString, sep))
}
func GetEnvironment() map[string]string {
	return ListToMap(os.Environ())
}

func parse(spec ConfigSpec, environment map[string]string) map[string]string {
	config := make(map[string]string)
	for key, value := range spec.Defaults {
		config[key] = value
	}
	for envKey, envValue := range environment {
		if newKey, found := spec.Renamed[envKey]; found {
			config[newKey] = envValue
		} else {
			if !Contains(spec.Excludes, envKey) {
				for prefix, keep := range spec.Prefixes {
					if strings.HasPrefix(envKey, prefix) {
						var effectiveKey string
						if keep {
							effectiveKey = envKey
						} else {
							effectiveKey = envKey[len(prefix)+1:]
						}
						config[convertKey(effectiveKey)] = envValue
					}
				}
			}
		}
	}
	return config
}

func printConfig(config map[string]string) {
	fmt.Printf("# created by 'ub' from environment variables on %s\n", time.Now().String())
	// Go randomizes iterations over map by design. We want to sort properties by name to ease debugging.
	sortedNames := make([]string, 0, len(config))
	for name := range config {
		sortedNames = append(sortedNames, name)
	}
	sort.Strings(sortedNames)
	for _, n := range sortedNames {
		fmt.Printf("%s=%s\n", n, config[n])
	}
}

func listenersFromAdvertisedListeners(listeners string) string {
	re := regexp.MustCompile("://(.*?):")
	return re.ReplaceAllString(listeners, "://0.0.0.0:")
}

func printProperty(pathToSpec string) {
	jsonFile, err := os.Open(pathToSpec)
	if err != nil {
		panic(err)
	}
	bytes, err2 := ioutil.ReadAll(jsonFile)
	if err2 != nil {
		panic(err2)
	}

	var spec ConfigSpec

	err3 := json.Unmarshal(bytes, &spec)
	if err3 != nil {
		panic(err3)
	}
	config := parse(spec, GetEnvironment())
	printConfig(config)
}

type LoggerSpec struct {
	RootLevel string            `json:"rootLevel"`
	Loggers   map[string]string `json:"loggers"`
}

func buildLoggerSpec(defaultsPath string, rootLoggerEnvVar string, levelEnvVar string) LoggerSpec {
	jsonFile, err := os.Open(defaultsPath)
	if err != nil {
		panic(err) // TODO: write to stderr instead of break
	}
	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}
	var loggerSpec LoggerSpec
	err = json.Unmarshal(bytes, &loggerSpec)
	if err != nil {
		panic(err)
	}
	if rootLevel, found := os.LookupEnv(rootLoggerEnvVar); found {
		loggerSpec.RootLevel = rootLevel
	}
	loggers := KvStringToMap(os.Getenv(levelEnvVar), ",")
	for logger, level := range loggers {
		loggerSpec.Loggers[logger] = level
	}
	return loggerSpec
}

func formatLogger(templatePath string, spec LoggerSpec) {
	templateFile, err := os.Open(templatePath)
	if err != nil {
		panic(err) // TODO: write to stderr instead of break
	}
	bytes, err := ioutil.ReadAll(templateFile)
	if err != nil {
		panic(err)
	}

	t := template.Must(template.New("tmpl").Parse(string(bytes)))

	t.Execute(os.Stdout, spec)
}

func checkDeprecate(deprecatedEnv string, deprecatedProperty string, replacement string) {
	if _, found := os.LookupEnv(deprecatedEnv); found {
		fmt.Printf("'%s' is deprecated. Use '%s' instead.\n", deprecatedProperty, replacement)
		os.Exit(1)
	}
}

func main() {
	if os.Args[1] == "check-deprecated" {
		checkDeprecate(os.Args[2], os.Args[3], os.Args[4])
	}
	// used, eg, for schema registry  and all admin-properties
	if os.Args[1] == "propertiesFromEnvPrefix" {
		envPrefix := os.Args[2]
		spec := ConfigSpec{
			Prefixes: map[string]bool{envPrefix: false},
			Excludes: []string{},
			Renamed:  map[string]string{},
			Defaults: map[string]string{},
		}
		config := parse(spec, GetEnvironment())
		printConfig(config)
	}
	if os.Args[1] == "propertiesFromEnv" {
		printProperty(os.Args[2])
	}
	if os.Args[1] == "formatLogger" {
		formatLogger(os.Args[2], buildLoggerSpec(os.Args[3], os.Args[4], os.Args[5]))
	}
	if os.Args[1] == "listeners" {
		fmt.Println(listenersFromAdvertisedListeners(os.Args[2]))
	}
	if os.Args[1] == "ensureAtLeastOne" {
		if !ensureAtLeastOne(os.Args[2:]) {
			os.Exit(1)
		}
	}
}
