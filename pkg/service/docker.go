package service

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io"
	"os"
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

func (d *DockerClient) Save(name string, path string) error {
	return nil
}
