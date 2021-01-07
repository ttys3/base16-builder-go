package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	// sourcesDir scheme and template index source repo
	sourcesDir string
	// color schemes dir
	schemesDir string
	// templates dir
	templatesDir string

	// Define the logger we'll be using
	logVerbose bool
	rawLog     = logrus.New()
	log        = logrus.NewEntry(rawLog)
)

func initFlags() {
	RootCmd.PersistentFlags().StringVar(&sourcesDir, "sources-dir", "./sources/", "Target directory for source repos")
	RootCmd.PersistentFlags().StringVar(&schemesDir, "schemes-dir", "./schemes/", "Target directory for scheme data")
	RootCmd.PersistentFlags().StringVar(&templatesDir, "templates-dir", "./templates/", "Target directory for template data")

	RootCmd.PersistentFlags().BoolVar(&logVerbose, "verbose", false, "Log all debug messages")

	cobra.OnInitialize(initLogger)
}

func initLogger() {
	rawLog.Formatter = &logrus.TextFormatter{
		ForceColors:  true,
		PadLevelText: true,
	}
	rawLog.Level = logrus.InfoLevel
	if logVerbose {
		rawLog.Level = logrus.DebugLevel
	}
}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "base16-builder-go",
	Short: "A simple builder for base16 templates and schemes",
}

func main() {
	initFlags()
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
