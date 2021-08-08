package sub

import (
	"reflect"
	"testing"
)

func TestParseConfig(t *testing.T) {
	config := Config{}
	config.GetConf("../test-resources/test-config.yaml")
	expected := Config{
		Authentication: Authentication{Username: "my_username", Token: "token"},
		Me:             User{Username: "me", Directory: "/home/user/repositories/me", CloneArchived: false},
		Users:          []User{{Username: "username", Directory: "/home/user/repositories/username", CloneArchived: false}},
		Teams: []Team{{
			Teamname:        "teamname",
			Prefix:          "remove_prefix",
			Dir:             "/home/user/repositories/teamname",
			Org:             "organization",
			CloneArchived:   false,
			AdditionalRepos: []string{"some_repo"},
		}, {
			Teamname:      "another_teamname",
			Prefix:        "",
			Dir:           "/home/user/repositories/another_teamname",
			Org:           "organization",
			CloneArchived: true,
		}},
	}
	if !reflect.DeepEqual(config, expected) {
		t.Errorf("Configs are not equal. Got %v, wanted %v", config, expected)
	}
}

func TestHeader(t *testing.T) {
	config := Config{}
	config.GetConf("../test-resources/test-config.yaml")
	got := config.Header()
	expected := "Updating me, username, teamname and another_teamname..."
	if got != expected {
		t.Errorf("Header is not equal. Got %v, wanted %v", got, expected)
	}
}
