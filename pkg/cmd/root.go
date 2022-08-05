package cmd

import "github.com/spf13/cobra"

func RootCommand(args []string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "quick-k8s",
		Short: "quick-k8s ...",
		Long: `
Tips  Find more information at: https://github.com/clarechu/quick-k8s
`,
	}
	rootCmd.AddCommand(ListCommand())
	rootCmd.AddCommand(PullCommand())
	rootCmd.AddCommand(VersionCommand())
	return rootCmd
}
