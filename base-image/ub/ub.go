package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func convertKey(s string) string {
	//TODO: add handling of '__' and '___'
	return strings.ToLower(strings.ReplaceAll(s, "_", "."))
}

func convertEnvVars(prefix string) {
	for _, envVar := range os.Environ() {
		if strings.HasPrefix(envVar, prefix) {
			woPrefix := envVar[len(prefix)+1:]
			parts := strings.Split(woPrefix, "=")
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
		convertEnvVars(os.Args[2])
	}
	if os.Args[1] == "listeners" {
		fmt.Println(listenersFromAdvertisedListeners(os.Args[2]))
	}
}
