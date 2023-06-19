package model

import (
	"bufio"
	"fmt"
	"os"
)

const (
	baukastenTagLocationPattern = "spec.components[name=%s].properties.tag"
	valuesLocation              = "%s/apps/env/%s/%s"
)

type Config struct {
	Development               bool     `json:"development"`
	BaseDir                   string   `json:"base_dir"`
	GitUrl                    string   `json:"git_url"`
	Username                  string   `json:"username"`
	Email                     string   `json:"email"`
	Reponame                  string   `json:"reponame"`
	Branch                    string   `json:"branch"`
	SshConfigDir              string   `json:"ssh_dir" `
	RepoToken                 string   `json:"repo_token"`
	InfraRepoSuffix           string   `json:"infra_repo_suffix"`
	ImageTag                  string   `json:"image_tag"`
	AppConfigFile             string   `json:"image_tag_file_name"`
	TagLocation               string   `json:"tag_location"`
	Stages                    []string `json:"stages"`
	Env                       string   `json:"env"`
	FromBranch                string   `json:"from_branch"`
	ToBranch                  string   `json:"to_branch"`
	Force                     bool     `json:"force"`
	ResourcesOnly             bool     `json:"resources_only"`
	Descriptor                string   `json:"descriptor"`
	DefaultDescriptorLocation string   `json:"default_descriptor_location"`
	CommitRef                 string   `json:"commit_ref"`
	Component                 string   `json:"component"`
}

func (c *Config) ApplicationClonePath() string {

	return fmt.Sprintf("%s%s", c.BaseDir, c.Reponame)
}

func (c *Config) InfrastructureClonePath() string {
	return fmt.Sprintf("%s%s%s", c.BaseDir, c.Reponame, c.InfraRepoSuffix)
}

func (c *Config) IsPushEnabled() bool {
	return !c.Development
}

func (c *Config) GetTagLocation() string {
	if c.Component == "" {
		return c.TagLocation
	}

	return fmt.Sprintf(baukastenTagLocationPattern, c.Component)
}

func (c *Config) BuildAppConfigFilePath(rootPath string, env string) (string, error) {
	if c.Component == "" {
		return fmt.Sprintf(valuesLocation, rootPath, env, c.AppConfigFile), nil
	}

	baukastenIndicatorFile := fmt.Sprintf(valuesLocation, rootPath, env, ".baukasten")
	appConfigFileBK, err := readFirstLine(baukastenIndicatorFile)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(valuesLocation, rootPath, env, appConfigFileBK), nil
}

func readFirstLine(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	return scanner.Text(), nil
}
