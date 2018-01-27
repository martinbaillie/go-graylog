package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Linker-provided
var (
	BuildTime  string
	Project    string
	Version    string
	APIVersion string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version information",
	Run: func(cmd *cobra.Command, args []string) {
		ShowVersion()
	},
}

func init() {
	graylogCmd.AddCommand(versionCmd)
}

// ShowVersion prints the version information to stdout
func ShowVersion() {
	fmt.Printf("project=%s build_time=%s version=%s api_version=%s\n",
		Project, BuildTime, Version, APIVersion)
}
