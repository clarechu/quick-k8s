package service

import (
	"github.com/clarechu/quick-k8s/pkg/models"
	yaml "gopkg.in/yaml.v3"
	log "k8s.io/klog/v2"
	"os"
)

func GetConfig(path string) (*models.ClusterConfiguration, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("读文件失败:%v", err)
	}
	config := &models.ClusterConfiguration{}
	err = yaml.Unmarshal(file, config)
	return config, err
}
