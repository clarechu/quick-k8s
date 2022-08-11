package cmd

import (
	"github.com/clarechu/quick-k8s/pkg/models"
	"github.com/spf13/cobra"
)

func RootCommand(args []string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "quickctl",
		Short: "quickctl ...",
		Long: `
Tips  Find more information at: https://github.com/clarechu/quick-k8s
`,
	}
	addRootCommand(rootCmd)
	rootCmd.AddCommand(ListCommand())
	rootCmd.AddCommand(PullCommand())
	rootCmd.AddCommand(SyncCommand())
	rootCmd.AddCommand(VersionCommand())
	rootCmd.AddCommand(ClusterCommand(args))

	return rootCmd
}

func addRootCommand(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&models.WorkDir, "workdir", "w", "/etc/quick-k8s", "项目根目录地址")
}
