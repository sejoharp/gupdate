package cmd

import (
	"fmt"
	"gupdate/sub"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/spf13/cobra"
)

func CreateDirectory(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}
}

var configPath string
var Verbose bool

func init() {
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "$HOME/.gupdate.yaml", "path to config file")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "enable verbose logging")
}

func fillEnvVariables(input string) string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	new1 := strings.Replace(input, "$HOME", homeDir, -1)
	new2 := strings.Replace(new1, "${HOME}", homeDir, -1)
	return new2
}

var rootCmd = &cobra.Command{
	Use:   "gupdate",
	Short: "gupdate keeps local git repositories up-to-date",
	Long:  `long explanation`,
	Run:   execute,
}

func execute(cmd *cobra.Command, args []string) {

	cleanConfigPath := fillEnvVariables(configPath)
	var c sub.Config
	c.GetConf(cleanConfigPath)

	if c, err := c.HasMinimalFields(); !c {
		log.Fatal(err)
	}

	auth := c.ToAuthentication()

	resultChan := make(chan sub.Result, 10000)
	parallelityChan := make(chan int, 20)
	var wg sync.WaitGroup

	fmt.Println(c.Header())

	CreateDirectory(c.Me.Directory)
	repos := c.Me.ListRepositories(auth, true)
	sub.UpdateRepositories(Verbose, repos, c.Me.ShouldBeUpdated, "", c.Me.Directory, &wg, resultChan, parallelityChan)

	for _, s := range c.Users {
		CreateDirectory(s.Directory)
		repos := s.ListRepositories(auth, false)
		sub.UpdateRepositories(Verbose, repos, s.ShouldBeUpdated, "", s.Directory, &wg, resultChan, parallelityChan)
	}

	for _, t := range c.Teams {
		CreateDirectory(t.Dir)
		var emptyList []sub.Repository
		repos := t.ListRepositories(auth, emptyList, 1)
		sub.UpdateRepositories(Verbose, repos, t.ShouldBeUpdated, t.Prefix, t.Dir, &wg, resultChan, parallelityChan)
	}

	for _, o := range c.Organizations {
		CreateDirectory(o.Dir)
		var emptyList []sub.Repository
		repos := o.ListRepositories(auth, emptyList, 1)
		sub.UpdateRepositories(Verbose, repos, o.ShouldBeUpdated, "", o.Dir, &wg, resultChan, parallelityChan)
	}

	wg.Wait()
	input := []sub.Result{}
	close(resultChan)
	for x := range resultChan {
		input = append(input, x)
	}
	fmt.Println(sub.Consolidate(input))
}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
