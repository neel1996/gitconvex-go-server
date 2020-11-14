package tests

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	git2 "github.com/neel1996/gitconvex-server/git"
	"github.com/neel1996/gitconvex-server/graph/model"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestChangedFiles(t *testing.T) {
	var repoPath string
	var r *git.Repository
	currentEnv := os.Getenv("GOTESTENV")
	fmt.Println("Environment : " + currentEnv)

	if currentEnv == "ci" {
		repoPath = "/home/runner/work/gitconvex-go-server/starfleet"
		r, _ = git.PlainOpen(repoPath)
	}

	untrackedResult := "untracked.txt"
	changedResult := "README.md"
	stagedResult := "README.md"

	expectedResults := &model.GitChangeResults{
		GitUntrackedFiles: []*string{&untrackedResult},
		GitChangedFiles:   []*string{&changedResult},
		GitStagedFiles:    []*string{&stagedResult},
	}

	_ = ioutil.WriteFile(untrackedResult, []byte{byte(63)}, 0755)
	_ = ioutil.WriteFile(changedResult, []byte{byte(65)}, 0755)
	git2.StageItem(r, changedResult)
	_ = ioutil.WriteFile(changedResult, []byte{byte(70)}, 0755)

	type args struct {
		repo *git.Repository
	}
	tests := []struct {
		name string
		args args
		want *model.GitChangeResults
	}{
		{name: "Git changed files test case", args: struct{ repo *git.Repository }{repo: r}, want: expectedResults},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := git2.ChangedFiles(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChangedFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}
