package sub

import (
	"os"
	"os/exec"
)

func GitPull(dir string) (string, error) {
	cmd := exec.Command("git", "pull", "-r")
	cmd.Dir = dir
	out, err := cmd.Output()
	return string(out), err
}

func GitClone(r Repository, dir string, targetName string) (string, error) {
	cmd := exec.Command("git", "clone", r.SshUrl, targetName)
	cmd.Dir = dir
	out, err := cmd.Output()
	return string(out), err
}

func DoesDirectoryExist(dir string) bool {
	_, err := os.Stat(dir)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
