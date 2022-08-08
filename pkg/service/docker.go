package service

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io"
	log "k8s.io/klog/v2"
	"os"
	"path/filepath"
	"strings"
)

type DockerClient struct {
	client *client.Client
}

func NewNewDockerClient() *DockerClient {
	cli, err := client.NewClientWithOpts()
	if err != nil {
		panic(err)
	}
	return &DockerClient{
		client: cli,
	}
}

func (d *DockerClient) Pull(ctx context.Context, name string) error {
	reader, err := d.client.ImagePull(ctx, name, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	io.Copy(os.Stdout, reader)
	return nil
}

func (d *DockerClient) SaveAll(ctx context.Context, path string) error {
	images, err := d.client.ImageList(ctx, types.ImageListOptions{
		All: true,
	})
	if err != nil {
		return err
	}
	for _, image := range images {
		for _, tags := range image.RepoTags {
			log.Infof("保存镜像:%s, 保存路径:%s", tags, path)
			repos := strings.Split(tags, "/")
			name := repos[len(repos)-1]
			reader, err := d.client.ImageSave(ctx, []string{image.ID})
			if err != nil {
				return err
			}
			filename := filepath.Join(path, fmt.Sprintf("%s.tar.gz", name))
			if _, err := os.Stat(filename); err == nil {
				continue
			}
			file, err := os.Create(filename)
			if err != nil {
				return err
			}
			io.Copy(file, reader)
		}
	}

	return nil
}
