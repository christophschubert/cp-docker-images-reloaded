package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

/**
Converts an environment variable name to a property-name according to the following rules:

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

func convertEnvVars(prefix string, keep bool) {
	for _, envVar := range os.Environ() {
		if strings.HasPrefix(envVar, prefix) {
			var effectiveValue string
			if keep {
				effectiveValue = envVar
			} else {
				effectiveValue = envVar[len(prefix)+1:]
			}
			parts := strings.Split(effectiveValue, "=")
			fmt.Printf("%s=%s\n", convertKey(parts[0]), parts[1])
		}
	}
}

func listenersFromAdvertisedListeners(listeners string) string {
	re := regexp.MustCompile("://(.*?):")
	return re.ReplaceAllString(listeners, "://0.0.0.0:")
}

func main() {
	if os.Args[1] == "envToProp" {
		convertEnvVars(os.Args[2], false)
	}
	if os.Args[1] == "envToPropKeepPrefix" {
		convertEnvVars(os.Args[2], true)
	}
	if os.Args[1] == "listeners" {
		fmt.Println(listenersFromAdvertisedListeners(os.Args[2]))
	}
}
