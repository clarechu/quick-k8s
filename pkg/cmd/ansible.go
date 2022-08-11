package cmd

import (
	"fmt"
	"github.com/clarechu/quick-k8s/pkg/models"
	"github.com/clarechu/quick-k8s/pkg/service"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	log "k8s.io/klog/v2"
	"os"
	"path/filepath"
)

type Cluster struct {
	Name      string
	Playbooks []string
	Inventory string
	Envs      []string
	Step      string
}

func ClusterCommand(args []string) *cobra.Command {
	cluster := &Cluster{}
	clusterCmd := &cobra.Command{
		Use:   "cluster",
		Short: "集群设置",
		Long: `
操作集群
      prepare            准备CA证书和kube-config以及其他系统设置
      etcd               设置etcd集群
      runtime            设置容器运行时（docker或containerd）
      kubernetes-master  设置master 节点
      kube-node          设置node节点
      network            设置网络插件
      manifests          给集群添加附加组件
      setup              运行所有的步骤
      clean              清理节点的k8s及runtime
      harbor             安装新的harbor服务器或与现有服务器集成
EXAMPLE:
quickctl cluster 
`,
		PreRun: func(cmd *cobra.Command, args []string) {
			//集群名称
			if len(args) == 1 {
				cluster.Name = args[0]
			} else {
				log.Fatal("请设置集群名称\n quickctl cluster <CLUSTER_NAME>")
			}

			if cluster.Step == "" {
				cluster.Step = setStep()
			}
			if cluster.Step == "new" {
				log.Infof("new cluster to %s", filepath.Join(models.WorkDir, "clusters", cluster.Name))
				err := service.NewCluster(cluster.Name)
				if err != nil {
					log.Fatalf("新建集群失败：:%s", err.Error())
				}
				log.Infof("================")
				log.Infof("创建集群 ✅")
				os.Exit(0)
			}
			cluster.setCluster()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cluster.Start()
		},
	}
	addClusterFlags(clusterCmd, cluster)
	return clusterCmd
}

func addClusterFlags(cmd *cobra.Command, cluster *Cluster) {
	cmd.Flags().StringSliceVarP(&cluster.Envs, "envs", "e",
		[]string{}, "环境变量 文件例如config.yaml")
	cmd.Flags().StringSliceVarP(&cluster.Playbooks, "playbooks", "p",
		[]string{}, "运行ansible的playbooks")
	cmd.Flags().StringVarP(&cluster.Inventory, "inventory", "i",
		"", "环境变量 文件例如config.yaml")
	cmd.Flags().StringVarP(&cluster.Step, "step", "s",
		"", "当前对集群的操作步骤")
}

func (c *Cluster) setCluster() {
	inventory := fmt.Sprintf("%s.yml", c.Step)
	if len(c.Envs) == 0 {
		c.Envs = []string{filepath.Join(models.WorkDir, "cluster", c.Name, "config.yml")}
	}
	if len(c.Playbooks) == 0 {
		c.Playbooks = []string{filepath.Join(models.WorkDir, "playbooks", inventory)}
	}
	if c.Inventory == "" {
		c.Inventory = filepath.Join(models.WorkDir, "cluster", c.Name, "hosts")
	}
	if _, err := os.Stat(c.Inventory); err != nil {
		log.Fatalf("集群cluster:%s 不存在", c.Name)
	}
}

func (c *Cluster) Start() error {
	var envs []string
	for _, env := range c.Envs {
		envs = append(envs, fmt.Sprintf("@%s", env))
	}
	client := service.NewAnsibleClient(c.Playbooks, c.Inventory, envs...)
	return client.Run()
}

var iterms = []string{"new", "clean", "prepare", "runtime", "kubernetes-master",
	"addmaster", "addnode", "manifests", "setup"}

func setStep() string {
	prompt := promptui.Select{
		Label: "Select step",
		Items: iterms,
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("select error :%s", err.Error())
	}
	return result
}
