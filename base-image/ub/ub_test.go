package main

import (
	"testing"
)

func TestContains(t *testing.T) {
	if Contains([]string{"hello", "world"}, "hi") {
		t.Error(` Contains([]string{"hello", "world"}, "hi") = true`)
	}
	if !Contains([]string{"hello", "world"}, "hello") {
		t.Error(`Contains([]string{"hello", "world"}, "hello") = false`)
	}
}

func TestBuildProperties(t *testing.T) {
	var testEnv = map[string]string{
		"PATH":                    "thePath",
		"KAFKA_BOOTSTRAP_SERVERS": "localhost:9092",
		"CONFLUENT_METRICS":       "metricsValue",
	}

	var onlyDefaultsCS = ConfigSpec{
		Prefixes: map[string]bool{},
		Excludes: []string{},
		Renamed:  map[string]string{},
		Defaults: map[string]string{
			"default.property.key": "default.property.value",
			"bootstrap.servers":    "unknown",
		},
	}

	var onlyDefaults = BuildProperties(onlyDefaultsCS, testEnv)
	if len(onlyDefaults) != 2 {
		t.Error("Failed to parse defaults.")
	}
	if onlyDefaults["default.property.key"] != "default.property.value" {
		t.Error("default.property.key not parsed correctly")
	}

	var serverCS = ConfigSpec{
		Prefixes: map[string]bool{"KAFKA": false, "CONFLUENT": true},
		Excludes: []string{},
		Renamed:  map[string]string{},
		Defaults: map[string]string{
			"default.property.key": "default.property.value",
			"bootstrap.servers":    "unknown",
		},
	}
	var serverProps = BuildProperties(serverCS, testEnv)
	if len(serverProps) != 3 {
		t.Error("Server props size != 3")
	}
	if serverProps["bootstrap.servers"] != "localhost:9092" {
		t.Error("Dropped prefixed not parsed correctly")
	}
	if serverProps["confluent.metrics"] != "metricsValue" {
		t.Error("Kept prefix not parsed correctly")
	}
}
