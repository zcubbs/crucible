package internal

import (
	"context"
	"github.com/apenella/go-ansible/pkg/options"
	ansiblePlaybook "github.com/apenella/go-ansible/pkg/playbook"
	"log"
)

// RunLocalPlaybook runs a playbook on the local host.
func RunLocalPlaybook(playbooks []string, extraVars map[string]interface{}) error {
	ansiblePlaybookConnectionOptions := &options.AnsibleConnectionOptions{
		Connection: "local",
	}

	ansiblePlaybookOptions := &ansiblePlaybook.AnsiblePlaybookOptions{
		Inventory: "127.0.0.1,",
		ExtraVars: extraVars,
	}

	p := &ansiblePlaybook.AnsiblePlaybookCmd{
		Playbooks:         playbooks,
		ConnectionOptions: ansiblePlaybookConnectionOptions,
		Options:           ansiblePlaybookOptions,
	}

	err := p.Run(context.TODO())
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

type ConnectionUser struct {
	Username string
	Password string
}

// RunRemotePlaybook runs a playbook on a remote host.
func RunRemotePlaybook(playbooks []string, inventory string, connectionUser ConnectionUser, extraVars map[string]interface{}) error {
	ansiblePlaybookConnectionOptions := &options.AnsibleConnectionOptions{
		AskPass:    true,
		Connection: "ssh",
		User:       connectionUser.Username,
	}

	ansiblePlaybookOptions := &ansiblePlaybook.AnsiblePlaybookOptions{
		Inventory: inventory, // ex: "192.168.1.10,,"
		ExtraVars: extraVars,
	}

	_ = ansiblePlaybookOptions.AddExtraVar("ansible_python_interpreter", "/usr/bin/python3")
	_ = ansiblePlaybookOptions.AddExtraVar("ansible_password", connectionUser.Password)
	_ = ansiblePlaybookOptions.AddExtraVar("ansible_become", "true")
	_ = ansiblePlaybookOptions.AddExtraVar("ansible_become_method", "sudo")
	_ = ansiblePlaybookOptions.AddExtraVar("ansible_become_pass", connectionUser.Password)

	p := &ansiblePlaybook.AnsiblePlaybookCmd{
		Playbooks:         playbooks,
		ConnectionOptions: ansiblePlaybookConnectionOptions,
		Options:           ansiblePlaybookOptions,
	}

	err := p.Run(context.TODO())
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
