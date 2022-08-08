package cmd

import (
	"context"
	"github.com/clarechu/quick-k8s/pkg/service"
	"github.com/spf13/cobra"
	log "k8s.io/klog/v2"
)

type Sync struct {
	SourceHost  string `yaml:"sourceHost"`
	DestHost    string `yaml:"destHost"`
	IsLoad      bool   `yaml:"isLoad"`
	ImageDirect string `yaml:"imageDirect"`
}

func SyncCommand() *cobra.Command {
	sync := new(Sync)
	syncCmd := &cobra.Command{
		Use:   "sync",
		Short: "sync docker images",
		Long: `
将source docker.io 的同步镜像 同步到本地仓库

EXAMPLE:
quick-k8s sync --source * --dest registry.local 
`,
		Run: func(cmd *cobra.Command, args []string) {
			err := sync.Start()
			if err != nil {
				log.Errorf("同步docker images 失败:%s", err)
				return
			}
			log.Infof("同步镜像成功")
		},
	}

	addSyncCommand(syncCmd, sync)
	return syncCmd
}

func addSyncCommand(cmd *cobra.Command, sync *Sync) {
	cmd.Flags().StringVar(&sync.SourceHost, "source", "*", "源镜像host地址")
	cmd.Flags().StringVar(&sync.DestHost, "dest", "registry.local", "目标镜像地址")
	cmd.Flags().BoolVar(&sync.IsLoad, "is-load", false, "是否load镜像")
	cmd.Flags().StringVar(&sync.ImageDirect, "direct", "/etc/quick-k8s/offline/images", "需要同步镜像的镜像地址")
}

func (s *Sync) Start() error {
	dockerClient := service.NewNewDockerClient()
	ctx := context.TODO()
	if s.IsLoad {
		err := dockerClient.LoadAll(ctx, s.ImageDirect)
		if err != nil {
			return err
		}
	}
	err := dockerClient.Tag(ctx, s.SourceHost, s.DestHost)
	if err != nil {
		return err
	}
	err = dockerClient.Push(ctx, s.DestHost)
	return err
}
