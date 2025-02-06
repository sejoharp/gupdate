package sub

type Repository struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Archived bool   `json:"archived"`
	SshUrl   string `json:"ssh_url"`
}
type OrgTeam struct {
	Name            string `json:"name"`
	Id              int    `json:"id"`
	Slug            string `json:"slug"`
	RepositoriesUrl string `json:"repositories_url"`
}

type Organization struct {
	Name          string `yaml:"name"`
	Dir           string `yaml:"dir"`
	CloneArchived bool   `yaml:"clone_archived"`
}

type User struct {
	Username      string `yaml:"username"`
	Directory     string `yaml:"dir"`
	CloneArchived bool   `yaml:"clone_archived"`
	Private       bool   `yaml:"private"`
}
type Team struct {
	Teamname      string `yaml:"teamname"`
	Prefix        string `yaml:"prefix"`
	Dir           string `yaml:"dir"`
	Org           string `yaml:"org"`
	CloneArchived bool   `yaml:"clone_archived"`
}
type Authentication struct {
	Token         string `yaml:"token"`
	TokenVariable string `yaml:"token_env_variable"`
	TokenFile     string `yaml:"token_file"`
}

type ValidAuthentication struct {
	Username string
	Token    string
}
