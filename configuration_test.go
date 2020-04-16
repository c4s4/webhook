package main

import (
	"testing"
)

func TestLoadConfiguration(t *testing.T) {
	configuration, err := LoadConfiguration("configuration.yml")
	if err != nil {
		t.Fatalf("Error loading configuration: %v", err)
	}
	action := (*configuration)["c4s4/sweetohm"]["push"]
	if action.Key != "XYZ" || action.Command != "command" {
		t.Fatal("Error loading configuration")
	}
}
