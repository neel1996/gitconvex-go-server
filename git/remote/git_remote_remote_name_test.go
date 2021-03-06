package remote

import (
	"fmt"
	git2go "github.com/libgit2/git2go/v31"
	"github.com/stretchr/testify/suite"
	"os"
	"path/filepath"
	"testing"
)

type GetRemoteNameTestSuite struct {
	suite.Suite
	getRemoteName Name
	repo          *git2go.Repository
	noHeadRepo    *git2go.Repository
}

func (suite *GetRemoteNameTestSuite) SetupTest() {
	r, err := git2go.OpenRepository(os.Getenv("GITCONVEX_TEST_REPO"))
	if err != nil {
		fmt.Println(err)
	}
	noHeadPath := os.Getenv("GITCONVEX_TEST_REPO") + string(filepath.Separator) + "no_head"
	noHeadRepo, _ := git2go.OpenRepository(noHeadPath)

	suite.repo = r
	suite.noHeadRepo = noHeadRepo
	suite.getRemoteName = NewGetRemoteName(suite.repo, "https://github.com/neel1996/gitconvex-test.git")
}

func TestGetRemoteNameTestSuite(t *testing.T) {
	suite.Run(t, new(GetRemoteNameTestSuite))
}

func (suite *GetRemoteNameTestSuite) TestGetRemoteName_ShouldReturnRemoteName_WhenRemoteUrlIsValid() {
	expectedRemote := "origin"

	actualRemote := suite.getRemoteName.GetRemoteNameWithUrl()

	suite.Equal(expectedRemote, actualRemote)
}

func (suite *GetRemoteNameTestSuite) TestGetRemoteName_ShouldReturnEmptyString_WhenRepoIsNil() {
	expectedRemote := ""
	getRemoteName := NewGetRemoteName(nil, "https://github.com/neel1996/gitconvex-test.git")

	actualRemote := getRemoteName.GetRemoteNameWithUrl()

	suite.Equal(expectedRemote, actualRemote)
}

func (suite *GetRemoteNameTestSuite) TestGetRemoteName_ShouldReturnEmptyString_WhenRepoHasNoRemotes() {
	expectedRemote := ""
	getRemoteName := NewGetRemoteName(suite.noHeadRepo, "https://github.com/neel1996/gitconvex-test.git")

	actualRemote := getRemoteName.GetRemoteNameWithUrl()

	suite.Equal(expectedRemote, actualRemote)
}
