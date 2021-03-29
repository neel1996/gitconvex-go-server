// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AddRepoParams struct {
	RepoID  string `json:"repoId"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type BranchDeleteStatus struct {
	Status string `json:"status"`
}

// Details about the repo data stored in gitconvex JSON data file
type FetchRepoParams struct {
	// UUID based ID generated for the repo
	RepoID []*string `json:"repoId"`
	// Name of the repo given while adding the repo
	RepoName []*string `json:"repoName"`
	// Path of the git repository
	RepoPath []*string `json:"repoPath"`
	// Timestamp at which the repo was added to gitconvex
	TimeStamp []*string `json:"timeStamp"`
}

type FetchResult struct {
	Status       string    `json:"status"`
	FetchedItems []*string `json:"fetchedItems"`
}

type GitFolderContentResults struct {
	TrackedFiles     []*string `json:"trackedFiles"`
	FileBasedCommits []*string `json:"fileBasedCommits"`
}

// Basic information about the target repository
type GitRepoStatusResults struct {
	// The remote repos available in the target repository
	GitRemoteData *string `json:"gitRemoteData"`
	// Name of the repo
	GitRepoName *string `json:"gitRepoName"`
	// List of all the local branches available in the repo
	GitBranchList []*string `json:"gitBranchList"`
	// List of all the local branches + remote branches available in the repo
	GitAllBranchList []*string `json:"gitAllBranchList"`
	// The current branch
	GitCurrentBranch *string `json:"gitCurrentBranch"`
	// Remote host name based on the remote URL (E.g: https://github.com/github/repo.git -> Github)
	GitRemoteHost *string `json:"gitRemoteHost"`
	// Total number of commits tracked by the current branch
	GitTotalCommits *float64 `json:"gitTotalCommits"`
	// The latest commit (HEAD commit)
	GitLatestCommit *string `json:"gitLatestCommit"`
	// Total number of files tracked by git repo
	GitTotalTrackedFiles *int `json:"gitTotalTrackedFiles"`
}

// Returns the host OS and the current version of gitconvex
type HealthCheckParams struct {
	// OS on which gitconvex is running
	Os string `json:"os"`
	// Current version of gitconvex
	Gitconvex string `json:"gitconvex"`
}

type PullResult struct {
	Status      string    `json:"status"`
	PulledItems []*string `json:"pulledItems"`
}

type UnPushedCommitResult struct {
	IsNewBranch bool          `json:"isNewBranch"`
	GitCommits  []*GitCommits `json:"gitCommits"`
}

type BranchCompareResults struct {
	Date    string        `json:"date"`
	Commits []*GitCommits `json:"commits"`
}

type CodeFileType struct {
	FileData []*string `json:"fileData"`
}

type DeleteStatus struct {
	Status string `json:"status"`
	RepoID string `json:"repoId"`
}

type FileLineChangeResult struct {
	DiffStat string    `json:"diffStat"`
	FileDiff []*string `json:"fileDiff"`
}

type GitChangeResults struct {
	GitUntrackedFiles []*string `json:"gitUntrackedFiles"`
	GitChangedFiles   []*string `json:"gitChangedFiles"`
	GitStagedFiles    []*string `json:"gitStagedFiles"`
}

type GitCommitFileResult struct {
	Type     string `json:"type"`
	FileName string `json:"fileName"`
}

type GitCommitLogResults struct {
	TotalCommits *float64      `json:"totalCommits"`
	Commits      []*GitCommits `json:"commits"`
}

type GitCommits struct {
	Hash             *string `json:"hash"`
	Author           *string `json:"author"`
	CommitTime       *string `json:"commitTime"`
	CommitMessage    *string `json:"commitMessage"`
	CommitFilesCount *int    `json:"commitFilesCount"`
}

type RemoteDetails struct {
	RemoteName string `json:"remoteName"`
	RemoteURL  string `json:"remoteUrl"`
}

type RemoteMutationResult struct {
	Status string `json:"status"`
}

type SettingsDataResults struct {
	SettingsDatabasePath string `json:"settingsDatabasePath"`
	SettingsPortDetails  string `json:"settingsPortDetails"`
}
