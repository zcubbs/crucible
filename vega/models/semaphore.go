package models

type Project struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Alert       bool   `json:"alert,omitempty"`
}

type Repository struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	ProjectId   int    `json:"project_id"`
	GitUrl      string `json:"git_url"`
	GitBranch   string `json:"git_branch"`
	SshKeyId    int    `json:"ssh_key_id"`
	LastUpdated string `json:"last_updated"`
}

type SshKey struct {
	Id        int    `json:"id,omitempty"`
	Name      string `json:"name"`
	Value     string `json:"value"`
	Type      string `json:"type"`
	ProjectId int    `json:"project_id"`
}

type Inventory struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Value       string `json:"value"`
	ProjectId   int    `json:"project_id"`
	SshKeyId    int    `json:"ssh_key_id"`
	Type        string `json:"type"`
	BecomeKeyId int    `json:"become_key_id"`
}

type Environment struct {
	Id        int    `json:"id,omitempty"`
	Name      string `json:"name"`
	ProjectId int    `json:"project_id"`
	Password  string `json:"password,omitempty"`
	Json      string `json:"json,omitempty"`
}

type Template struct {
	Id            int    `json:"id,omitempty"`
	Alias         string `json:"alias"`
	Name          string `json:"name"`
	Playbook      string `json:"playbook"`
	ProjectId     int    `json:"project_id"`
	EnvironmentId int    `json:"environment_id"`
	RepositoryId  int    `json:"repository_id"`
	InventoryId   int    `json:"inventory_id"`
	ViewId        int    `json:"view_id"`
	Arguments     string `json:"arguments"`
	Description   string `json:"description"`
	OverrideArgs  bool   `json:"override_args"`
}
