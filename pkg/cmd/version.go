package cmd

import "github.com/spf13/cobra"

func VersionCommand() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "quick-k8s version",
		Long: `
		quick-k8s version 
`,
	}
	return versionCmd
}
