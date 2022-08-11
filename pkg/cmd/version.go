package cmd

import "github.com/spf13/cobra"

func VersionCommand() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: " version",
		Long: `
		quickctl version 
`,
	}
	return versionCmd
}
