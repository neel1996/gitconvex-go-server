package tests

import (
	git "github.com/go-git/go-git/v5"
	git2 "github.com/neel1996/gitconvex-server/git"
	"os"
	"path"
	"strings"
	"testing"
)

func TestRemoteData(t *testing.T) {
	remoteChan := make(chan *git2.RemoteDataModel)
	cwd, _ := os.Getwd()
	r, _ := git.PlainOpen(path.Join(cwd, ".."))

	type args struct {
		repo       *git.Repository
		remoteChan chan *git2.RemoteDataModel
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Git remote data test case", args: struct {
			repo       *git.Repository
			remoteChan chan *git2.RemoteDataModel
		}{repo: r, remoteChan: remoteChan}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go git2.RemoteData(tt.args.repo, tt.args.remoteChan)
			remoteData := <-remoteChan

			rHost := remoteData.RemoteHost
			rURL := remoteData.RemoteURL

			if !strings.Contains(*rURL[0], "github") || *rHost != "github" {
				t.Error("Expected remote data not received")
			}
		})
	}
}
