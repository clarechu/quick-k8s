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
