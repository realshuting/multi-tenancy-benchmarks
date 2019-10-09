package test

import (
	"flag"
	"testing"

	"k8s.io/kubernetes/test/e2e/framework"
	"k8s.io/kubernetes/test/e2e/framework/config"
)

var viperConfig = flag.String("viper-config", "", "The name of a viper config file (https://github.com/spf13/viper#what-is-viper). All e2e command line parameters can also be configured in such a file. May contain a path and may or may not contain the file suffix. The default is to look for an optional file with `e2e` as base name. If a file is specified explicitly, it must be present.")

// handleFlags sets up all flags and parses the command line.
func handleFlags() {
	config.CopyFlags(config.Flags, flag.CommandLine)
	framework.RegisterCommonFlags(flag.CommandLine)
	framework.RegisterClusterFlags(flag.CommandLine)
	flag.Parse()
}

func init() {
	// Register framework flags, then handle flags.
	handleFlags()
}

func TestE2E(t *testing.T) {
	RunE2ETests(t)
}
