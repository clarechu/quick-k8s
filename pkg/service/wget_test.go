package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Wget(t *testing.T) {
	file := NewFile("./", "linux.zip")
	uri := "https://github.com/krishpranav/gowget/releases/download/v2/linux.zip"
	err := file.Wget(uri)
	assert.Equal(t, nil, err)
}
