package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_DockerPull(t *testing.T) {
	client := NewNewDockerClient()
	err := client.Pull(context.TODO(), "nginx")
	assert.Equal(t, nil, err)
}
