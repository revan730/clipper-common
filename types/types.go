package types

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User represents system's user
type User struct {
	ID            int64  `json:"-"`
	Login         string `sql:",unique" json:"-"`
	Password      string `json:"-"`
	IsAdmin       bool   `json:"-" sql:"default:false"`
	WebhookSecret string `json:"-" sql:"default:''"`
	AccessToken   string `json:"-" sql:"default:''"`
}

// GithubRepo represents GitHub repository
type GithubRepo struct {
	ID       int64  `json:"repoID"`
	FullName string `json:"fullName" sql:",unique"`
	UserID   int64  `json:"-" pg:",fk" sql:"on_delete:CASCADE"`
	User     *User  `json:"-"`
}

// BranchConfig sets CI configuration for specific branch of repo
type BranchConfig struct {
	ID           int64       `json:"-"`
	GithubRepoID int64       `json:"-" pg:",fk" sql:"on_delete:CASCADE"`
	GithubRepo   *GithubRepo `json:"-"`
	Branch       string      `json:"branch"`
	IsCiEnabled  bool        `json:"ci_enabled"`
}

// Build represents CI job process with GCR container push
type Build struct {
	ID            int64       `json:"buildID"`
	GithubRepoID  int64       `json:"-" pg:",fk" sql:"on_delete:CASCADE"`
	GithubRepo    *GithubRepo `json:"-"`
	IsSuccessfull bool        `json:"success"`
	Date          time.Time   `json:"date"`
	Branch        string      `json:"branch"`
	Stdout        string      `json:"stdout"`
}

// BuildArtifact represents docker container pushed to GCR after CI process
type BuildArtifact struct {
	ID      int64  `json:"artifactID"`
	BuildID int64  `json:"-" pg:",fk" sql:"on_delete:CASCADE"`
	Build   *Build `json:"-"`
	// Name is a complete container name, with version
	Name string `json:"name"`
}

// Authenticate checks if provided password matches
// for this user
func (u User) Authenticate(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

type DBClientConfig struct {
	DBAddr        string
	DB            string
	DBUser        string
	DBPassword    string
	AdminLogin    string
	AdminPassword string
}
