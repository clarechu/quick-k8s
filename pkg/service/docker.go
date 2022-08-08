package service

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io"
	"io/fs"
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
			reader, err := d.client.ImageSave(ctx, image.RepoTags)
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
			break
		}
	}

	return nil
}

func (d *DockerClient) Push(ctx context.Context, target string) error {
	images, err := d.client.ImageList(ctx, types.ImageListOptions{
		All: true,
	})
	if err != nil {
		return err
	}
	for _, image := range images {
		for _, tag := range image.RepoTags {
			if strings.Contains(tag, target) {
				reader, err := d.client.ImagePush(ctx, image.RepoTags[0], types.ImagePushOptions{
					// RegistryAuth:
				})
				if err != nil {
					return err
				}
				io.Copy(os.Stdout, reader)
				break
			}
		}
	}
	return err
}

// Tag 给单个容器打一个tag
// Tag 需要打tag 的容器id
func (d *DockerClient) Tag(ctx context.Context, source string, target string) error {
	images, err := d.client.ImageList(ctx, types.ImageListOptions{
		All: true,
	})
	if err != nil {
		return err
	}
	for _, image := range images {
		for _, repo := range image.RepoTags {
			// todo 获取镜像的host 转换成 target
			host := getHost(repo)
			if host != source && source != "*" {
				continue
			}
			targetRepo := toSourceRepo(repo, target)
			err = d.client.ImageTag(ctx, image.ID, targetRepo)
			if err != nil {
				log.Errorf("images tag error:%s", err.Error())
				continue
			}
			break
		}

	}
	return nil
}

func (d *DockerClient) LoadAll(ctx context.Context, dir string) error {
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return err
		}
		log.Infof("docker load images :%s", info.Name())
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		_, err = d.client.ImageLoad(ctx, file, false)
		return err
	})
	return err
}

const (
	DockerDefaultHost = "docker.io"
)

// 获取
func getHost(host string) string {
	names := strings.Split(host, "/")
	if len(names) == 1 {
		return DockerDefaultHost
	}
	if !strings.Contains(names[0], ".") {
		return DockerDefaultHost
	}
	return names[0]
}

func toSourceRepo(repo string, target string) string {
	names := strings.Split(repo, "/")
	if len(names) == 1 {
		return fmt.Sprintf("%s/%s", target, repo)
	}
	if !strings.Contains(names[0], ".") {
		return fmt.Sprintf("%s/%s", target, repo)
	}
	return strings.ReplaceAll(repo, getHost(repo), target)
}
