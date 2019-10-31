package types

//UserInfo encapsulates user information
type UserInfo struct {
	Name        string `json:"name"`
	Location    string `json:"location"`
	PublicRepos int    `json:"public_repos"`
}

//Gist contains basic info of publi gist
type Gist struct {
	Owner       map[string]interface{} `jsob:"owner"`
	URL         string                 `json:"html_url"`
	CreatedAt   string                 `json:"created_at"`
	Description string                 `json:"description"`
}

//Follower encapsulates follower meta
type follower struct {
	Name    string `json:"login"`
	HTMLURL string `json:"html_url"`
}

//Followers represents list of followers
type Followers []follower

//FollowingUser contains following user information
type FollowingUser struct {
	Name string `json:"login"`
	URL  string `json:"html_url"`
}

//FollowingUsers stores a list of FollowingUser
type FollowingUsers = []FollowingUser
