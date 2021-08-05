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

// ensureAtLeastOne returns true if at least on of the keys exists as an env var, false otherwise.
func ensureAtLeastOne(keys []string) bool {
	for _, varName := range keys {
		if _, found := os.LookupEnv(varName); found {
			return true
		}
	}
	return false
}

// convertKey Converts an environment variable name to a property-name according to the following rules:
// - a single underscore (_) is replaced with a .
// - a double underscore (__) is replaced with a single underscore
// - a triple underscore (___) is replaced with a dash
// Moreover, the whole string is converted to lower-case.
// The behavior of sequences of four or more underscores is undefined.
func convertKey(key string) string {
	re := regexp.MustCompile("[^_]_[^_]")
	singleReplaced := re.ReplaceAllStringFunc(key, replaceUnderscores)
	singleTripleReplaced := strings.ReplaceAll(singleReplaced, "___", "-")
	return strings.ToLower(strings.ReplaceAll(singleTripleReplaced, "__", "_"))
}

//replaceUnderscores replaces every underscore '_' by a dot '.'
func replaceUnderscores(s string) string {
	return strings.ReplaceAll(s, "_", ".")
}

type ConfigSpec struct {
	Prefixes map[string]bool   `json:"prefixes"`
	Excludes []string          `json:"excludes"`
	Renamed  map[string]string `json:"renamed"`
	Defaults map[string]string `json:"defaults"`
}

//Contains returns true if slice contains element, and false otherwise.
func Contains(slice []string, element string) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

//ListToMap splits each and entry of the kvList argument at '=' into a key/value pair and returns a map of all the k/v pair thus obtained.
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

//GetEnvironment returns the current environment as a map.
func GetEnvironment() map[string]string {
	return ListToMap(os.Environ())
}

//BuildProperties creates a map suitable to be output as Java properties from a ConfigSpec and a map representing an environment.
func BuildProperties(spec ConfigSpec, environment map[string]string) map[string]string {
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
	config := BuildProperties(spec, GetEnvironment())
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
	switch os.Args[1] {
	case "check-deprecated":
		checkDeprecate(os.Args[2], os.Args[3], os.Args[4])
	case "propertiesFromEnvPrefix":
		// used, eg, for schema registry  and all admin-properties
		envPrefix := os.Args[2]
		spec := ConfigSpec{
			Prefixes: map[string]bool{envPrefix: false},
			Excludes: []string{},
			Renamed:  map[string]string{},
			Defaults: map[string]string{},
		}
		config := BuildProperties(spec, GetEnvironment())
		printConfig(config)
	case "propertiesFromEnv":
		printProperty(os.Args[2])
	case "formatLogger":
		formatLogger(os.Args[2], buildLoggerSpec(os.Args[3], os.Args[4], os.Args[5]))
	case "listeners":
		fmt.Println(listenersFromAdvertisedListeners(os.Args[2]))
	case "ensureAtLeastOne":
		if !ensureAtLeastOne(os.Args[2:]) {
			os.Exit(1)
		}
	default:
		fmt.Println(os.Args[0] + ": Unknown option: " + os.Args[1])
		os.Exit(1)
	}
}
