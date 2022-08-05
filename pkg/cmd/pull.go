package cmd

import (
	"context"
	"fmt"
	"github.com/clarechu/quick-k8s/pkg/service"
	"github.com/spf13/cobra"
	log "k8s.io/klog/v2"
)

func PullCommand() *cobra.Command {
	pullCmd := &cobra.Command{
		Use:   "pull",
		Short: "pull package",
		Long: `
下载安装k8s所需要的安装包(image, rpm, dep, helm chart)
EXAMPLE:
quick-k8s pull 
`,
		Run: func(cmd *cobra.Command, args []string) {
			Pull()
		},
	}
	pullCmd.Flags().StringVar(&filePath, "path", "/etc/quick-k8s/config.yaml", "配置文件的默认路径")

	return pullCmd
}

func Pull() {
	config, err := service.GetConfig(filePath)
	if err != nil {
		log.Fatalf("获取配置文件失败 %v", err)
	}
	fmt.Println("\n[Kubernetes Image]")
	dockerClient := service.NewNewDockerClient()
	ctx := context.TODO()
	// 获取 kubernetes 镜像
	for _, k8s := range config.KubernetesImages {
		log.Infof("开始拉取镜像 --> %s", k8s.Repository)
		err = dockerClient.Pull(ctx, k8s.Repository)
		if err != nil {
			log.Fatalf("拉取镜像失败 %v", err)
		}
	}
	fmt.Println("\n[Addon Image]")
	// 获取 kubernetes 镜像
	for _, k8s := range config.AddonImages {
		log.Infof("开始拉取镜像 --> %s", k8s.Repository)
		err = dockerClient.Pull(ctx, k8s.Repository)
		if err != nil {
			log.Fatalf("拉取镜像失败 %v", err)
		}
	}
}
