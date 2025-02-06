package sub

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func (u *User) ListRepositories(auth ValidAuthentication, me bool) []Repository { // TODO Paging
	var url string
	if me {
		url = "https://api.github.com/user/repos?per_page=100"
	} else {
		url = "https://api.github.com/users/" + u.Username + "/repos?per_page=100"
	}
	bodyString, _ := request(url, auth)
	repos := ParseJson(bodyString)
	return repos
}

func (t *Team) GetRepositoryUrl(auth ValidAuthentication, page int) (string, error) {
	url := "https://api.github.com/orgs/" + t.Org + "/teams?per_page=100" + "&page=" + strconv.Itoa(page)
	bodyString, err := request(url, auth)
	if err != nil {
		log.Printf("get repository url error: %s", err.Error())
	}
	teams := ParseTeamsJson(bodyString)

	for _, tt := range teams {
		if tt.Name == t.Teamname {
			return tt.RepositoriesUrl, nil
		}
	}
	if len(teams) == 100 {
		return t.GetRepositoryUrl(auth, page+1)
	}
	return "", errors.New("team:[" + t.Teamname + "] not found.")
}

func request(url string, auth ValidAuthentication) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("request with url %s failed with error: %s", url, err.Error())
	}
	req.Header.Add("Authorization", "Basic "+basicAuth(auth.Username, auth.Token))
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("request error with url %s: error message: %s", url, err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	bodyString := string(body)
	return bodyString, nil
}
func (t *Team) ListRepositories(auth ValidAuthentication, repos []Repository, page int) []Repository { // TODO Paging
	repositoryUrl, err := t.GetRepositoryUrl(auth, 1)
	if err != nil {
		log.Fatal(err)
	}
	bodyString, _ := request(repositoryUrl+"?page="+strconv.Itoa(page)+"&per_page=100", auth)
	newRepos := ParseJson(bodyString)
	c := append(repos, newRepos...)
	// FIXME != 0?
	// FIXME archived filter
	if len(newRepos) == 100 {
		return t.ListRepositories(auth, c, page+1)
	}
	return c
}

func (o *Organization) ListRepositories(auth ValidAuthentication, repos []Repository, page int) []Repository {
	url := "https://api.github.com/orgs/" + o.Name + "/repos?per_page=100"
	bodyString, _ := request(url+"&page="+strconv.Itoa(page), auth)
	newRepos := ParseJson(bodyString)
	c := append(repos, newRepos...)
	if len(newRepos) == 100 {
		return o.ListRepositories(auth, c, page+1)
	}
	return c
}

func ParseJson(input string) []Repository {
	var inventory []Repository
	if err := json.Unmarshal([]byte(input), &inventory); err != nil {
		log.Fatal(err)
	}

	return inventory
}

func ParseTeamsJson(input string) []OrgTeam {
	var inventory []OrgTeam
	if err := json.Unmarshal([]byte(input), &inventory); err != nil {
		log.Fatal(err)
	}

	return inventory
}

func UpdateRepository(verbose bool, repository Repository, removePrefix string, directory string, c chan Result, wg *sync.WaitGroup, p chan int) {
	defer wg.Done()
	if verbose {
		fmt.Println("--> " + repository.Name)
	}
	nameWithoutPrefix := strings.ReplaceAll(repository.Name, removePrefix, "")
	repoPath := directory + "/" + nameWithoutPrefix
	if repository.Archived {
		if verbose {
			fmt.Println("Archived " + repository.Name)
		}
		if DoesDirectoryExist(repoPath) {
			c <- LocalArchived{Name: nameWithoutPrefix, Message: repoPath}
		}
	} else {
		if DoesDirectoryExist(repoPath) {
			if verbose {
				fmt.Println("Pulling " + repository.Name)
			}
			_, err := GitPull(repoPath)
			if err != nil {
				c <- Error{Name: nameWithoutPrefix, Message: fmt.Sprintf("%s - %s", repoPath, err.Error())}
			} else {
				c <- Pulled{Name: nameWithoutPrefix, Message: repoPath}
			}
		} else {
			if verbose {
				fmt.Println("Cloning " + repository.Name)
			}
			_, err := GitClone(repository, directory, nameWithoutPrefix)
			if err != nil {
				c <- Error{Name: nameWithoutPrefix, Message: fmt.Sprintf("%s - %s", repoPath, err.Error())}
			} else {
				c <- Cloned{Name: nameWithoutPrefix, Message: repoPath}
			}
		}
	}
	<-p
}

func (u *User) ShouldBeUpdated(repository Repository) bool {
	return strings.HasPrefix(repository.FullName, u.Username) &&
		((u.CloneArchived && repository.Archived) || (!repository.Archived))
}

func (o *Organization) ShouldBeUpdated(repository Repository) bool {
	return (o.CloneArchived && repository.Archived) || (!repository.Archived)
}

func (t *Team) ShouldBeUpdated(repository Repository) bool {
	// FIXME what about clone archived?
	return strings.HasPrefix(repository.Name, t.Prefix)
}

func UpdateRepositories(verbose bool, repositories []Repository, condition func(Repository) bool, removePrefix, dir string, wg *sync.WaitGroup, c chan Result, p chan int) {
	for _, r := range repositories {
		if condition(r) {
			p <- 1
			wg.Add(1)
			go UpdateRepository(verbose, r, removePrefix, dir, c, wg, p)
		}
	}
}

func Find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
