package service

import (
	"context"

	"github.com/apenella/go-ansible/pkg/options"
	"github.com/apenella/go-ansible/pkg/playbook"
)

type AnsibleClient struct {
	ansibleClient *playbook.AnsiblePlaybookCmd
}

func NewAnsibleClient(playbooks []string, inventory string, envs ...string) *AnsibleClient {
	ansiblePlaybookConnectionOptions := &options.AnsibleConnectionOptions{}
	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Inventory:     inventory,
		ExtraVarsFile: envs,
	}
	ansiblePlaybookCmd := &playbook.AnsiblePlaybookCmd{
		Playbooks:         playbooks,
		ConnectionOptions: ansiblePlaybookConnectionOptions,
		Options:           ansiblePlaybookOptions,
	}
	return &AnsibleClient{
		ansibleClient: ansiblePlaybookCmd,
	}
}

func (a *AnsibleClient) Run() error {
	err := a.ansibleClient.Run(context.TODO())
	return err
}
