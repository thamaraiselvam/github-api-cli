package types

//UserInfo contains all user related information
type UserInfo struct {
	Name        string `json:"name"`
	Location    string `json:"location"`
	PublicRepos int    `json:"public_repos"`
}
