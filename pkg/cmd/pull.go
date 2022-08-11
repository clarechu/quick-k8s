package cmd

import (
	"context"
	"fmt"
	"github.com/clarechu/quick-k8s/pkg/models"
	"github.com/clarechu/quick-k8s/pkg/service"
	"github.com/spf13/cobra"
	git "gopkg.in/src-d/go-git.v4"
	"io"
	log "k8s.io/klog/v2"
	"os"
	"path/filepath"
)

var path = ""

const (
	BinPath   = "offline/bin"
	ImagePath = "offline/images"
	DebPath   = "offline/packages/deb"
	RpmPath   = "offline/packages/rpm"
)

func PullCommand() *cobra.Command {
	pullCmd := &cobra.Command{
		Use:   "save",
		Short: "save package",
		Long: `
下载安装k8s所需要的安装包(image, rpm, dep, helm chart)
EXAMPLE:
quickctl save 
`,
		Run: func(cmd *cobra.Command, args []string) {
			config, err := service.GetConfig(filePath)
			if err != nil {
				log.Fatalf("获取配置文件失败 %v", err)
			}
			ImagePull(config)
			DownloadBin(config)
			DownloadRpm(config)
			DownloadDeb(config)
			err = SaveImage(filepath.Join(path, ImagePath))
			if err != nil {
				log.Fatalf("save %v", err)
			}
		},
	}
	pullCmd.Flags().StringVar(&filePath, "config", "/etc/quick-k8s/config.yaml", "配置文件的默认路径")
	pullCmd.Flags().StringVar(&path, "path", "/etc/quick-k8s", "镜像存储根路径")

	return pullCmd
}

func ImagePull(config *models.ClusterConfiguration) {
	fmt.Println("\n[Kubernetes Image]")
	dockerClient := service.NewNewDockerClient()
	ctx := context.TODO()
	// 获取 kubernetes 镜像
	for _, k8s := range config.KubernetesImages {
		log.Infof("开始拉取镜像 --> %s", k8s.Repository)
		err := dockerClient.Pull(ctx, k8s.Repository)
		if err != nil {
			log.Fatalf("拉取镜像失败 %v", err)
		}
	}
	fmt.Println("\n[Addon Image]")
	// 获取 kubernetes 镜像
	for _, k8s := range config.AddonImages {
		log.Infof("开始拉取镜像 --> %s", k8s.Repository)
		err := dockerClient.Pull(ctx, k8s.Repository)
		if err != nil {
			log.Fatalf("拉取镜像失败 %v", err)
		}
	}
}

func DownloadBin(config *models.ClusterConfiguration) {
	for _, bin := range config.BinaryURI {
		err := service.NewFile(filepath.Join(path, BinPath), bin.Name).Wget(bin.URI)
		if err != nil {
			log.Errorf("wget file error:%s", err.Error())
			continue
		}
	}
}

func DownloadRpm(config *models.ClusterConfiguration) {
	for _, bin := range config.RedHatPackageManagerURI {
		err := service.NewFile(filepath.Join(path, RpmPath), bin.Name).Wget(bin.URI)
		if err != nil {
			log.Errorf("wget file error:%s", err.Error())
			continue
		}
	}
}

func DownloadDeb(config *models.ClusterConfiguration) {
	for _, bin := range config.DebianPackageManagerURI {
		err := service.NewFile(filepath.Join(path, DebPath), bin.Name).Wget(bin.URI)
		if err != nil {
			log.Errorf("wget file error:%s", err.Error())
			continue
		}
	}
}

func SaveImage(path string) error {
	dockerClient := service.NewNewDockerClient()
	ctx := context.TODO()
	return dockerClient.SaveAll(ctx, path)
}

func Clone(github, path string) error {
	// Clones the repository into the worktree (fs) and storer all the .git
	// content into the storer
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL: github,
	})
	if err != nil {
		return err
	}

	// Prints the content of the CHANGELOG file from the cloned repository
	changelog, err := os.Open(filepath.Join(path, "CHANGELOG"))
	if err != nil {
		return err
	}

	io.Copy(os.Stdout, changelog)

	// Output: Initial changelog
	return err
}
