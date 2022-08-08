package cmd

import (
	"fmt"
	"github.com/clarechu/quick-k8s/pkg/service"
	"github.com/spf13/cobra"
	log "k8s.io/klog/v2"
)

var filePath string

func ListCommand() *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "list package",
		Long: `
查询 需要安装 kubernetes 的包(image, rpm, dep, helm chart)
EXAMPLE:
quick-k8s pull 
`,
		Run: func(cmd *cobra.Command, args []string) {
			List()
		},
	}
	listCmd.Flags().StringVar(&filePath, "config", "/etc/quick-k8s/config.yaml", "配置文件的默认路径")
	return listCmd
}

func List() {
	config, err := service.GetConfig(filePath)
	if err != nil {
		log.Fatalf("获取配置文件错误:%s", err)
	}
	fmt.Println("\n[Kubernetes Image]")
	// 获取 kubernetes 镜像
	for _, k8s := range config.KubernetesImages {
		fmt.Println(k8s.Repository)
	}
	fmt.Println("\n[Addon Image]")
	// 获取 kubernetes 镜像
	for _, addon := range config.AddonImages {
		fmt.Println(addon.Repository)
	}
	fmt.Println("\n[Binary]")
	for _, bin := range config.BinaryURI {
		fmt.Println(bin.Name)
	}
	fmt.Println("\n[rpm]")
	for _, bin := range config.RedHatPackageManagerURI {
		fmt.Println(bin.Name)
	}
	fmt.Println("\n[deb]")
	for _, bin := range config.DebianPackageManagerURI {
		fmt.Println(bin.Name)
	}
}
