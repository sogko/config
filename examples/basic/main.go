package main

import (
	"fmt"
	"github.com/sogko/config"
)

func main() {

	// Load configuration
	cfg := config.Load()

	// Get configuration value
	foo := cfg.GetString("foo")
	// Output: foo: bar

	// Print configuration value
	fmt.Printf("foo: %s\n", foo)

}
