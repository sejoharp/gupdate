package sub

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Authentication Authentication `yaml:"authentication"`
	Me             User           `yaml:"me"`
	Users          []User         `yaml:"users"`
	Teams          []Team         `yaml:"teams"`
}

func (c *Config) GetConf(path string) *Config {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}

func (c *Config) Header() string {
	var names []string
	names = append(names, c.Me.Username)
	for _, u := range c.Users {
		names = append(names, u.Username)
	}
	for _, t := range c.Teams {
		names = append(names, t.Teamname)
	}

	return fmt.Sprintf("Updating %s and %s...", strings.Join(names[:len(names)-1], ", "), names[len(names)-1])
}

func (c *Config) HasMinimalFields() (bool, error) {
	if c.Authentication.TokenVariable == "" && c.Authentication.Token == "" && c.Authentication.TokenFile == "" {
		return false, errors.New("specify token, tokenfile or tokenvariable for authentication")
	}
	return true, nil
}

func (c *Config) ToAuthentication() ValidAuthentication {
	var token string
	if c.Authentication.TokenFile != "" {
		content, err := ioutil.ReadFile(c.Authentication.TokenFile)
		if err != nil {
			log.Fatal(err)
		}
		token = strings.TrimSuffix(string(content), "\n")
	} else if c.Authentication.TokenVariable != "" {
		token = os.Getenv(c.Authentication.TokenVariable)
	} else if c.Authentication.Token != "" {
		token = c.Authentication.Token
	}

	return ValidAuthentication{
		Username: c.Authentication.Username,
		Token:    token,
	}
}
