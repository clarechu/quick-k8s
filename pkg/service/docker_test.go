package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_DockerPull(t *testing.T) {
	client := NewNewDockerClient()
	err := client.Pull(context.TODO(), "nginx")
	assert.Equal(t, nil, err)
}

func Test_DockerSave(t *testing.T) {
	client := NewNewDockerClient()
	dir, err := os.Getwd()
	assert.Equal(t, nil, err)

	err = client.SaveAll(context.TODO(), dir)
	assert.Equal(t, nil, err)
}

func Test_DockerLoad(t *testing.T) {
	client := NewNewDockerClient()
	dir, err := os.Getwd()
	dir = "/Users/clare/go/src/github.com/clarechu/quick-k8s/offline/images"
	assert.Equal(t, nil, err)
	err = client.LoadAll(context.TODO(), dir)
	assert.Equal(t, nil, err)
}

func Test_DockerTag(t *testing.T) {
	client := NewNewDockerClient()
	err := client.Tag(context.TODO(), "*", "haror.aa.aa")
	assert.Equal(t, nil, err)
}
