package sub

import (
	"io/ioutil"
	"log"
	"reflect"
	"testing"
)

func readInput(path string) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func TestParseJson(t *testing.T) {
	input := readInput("../test-resources/git-response.json")
	result := ParseJson(input)
	expected := []Repository{
		{Name: "test_repo_1", Archived: false, SshUrl: "git@github.com:test_user/test_repo_1.git", FullName: "test_user/test_repo_1"},
		{Name: "test_repo_2", Archived: false, SshUrl: "git@github.com:test_user/test_repo_2.git", FullName: "test_user/test_repo_2"}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Configs are not equal. Got %v, wanted %v", result, expected)
	}
}
