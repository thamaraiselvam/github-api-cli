package types

//UserInfo encapsulates user information
type UserInfo struct {
	Name        string `json:"name"`
	Location    string `json:"location"`
	PublicRepos int    `json:"public_repos"`
}

//Follower encapsulates follower meta
type follower struct {
	Name    string `json:"login"`
	HTMLURL string `json:"html_url"`
}

//Followers represents list of followers
type Followers []follower

//pRInfo contains URL information for PR
type pRInfo struct {
	URL string `json:"html_url"`
}

//pRItem contains details of each PR
type pRItem struct {
	Number      int    `json:"number"`
	Title       string `json:"title"`
	PullRequest pRInfo `json:"pull_request"`
	Body        string `json:"body"`
}

//PRItemList List of all PR items
type PRItemList struct {
	Items []pRItem `json:"items"`
}

//followingUser contains following user information
type followingUser struct {
	Name string `json:"login"`
	URL  string `json:"html_url"`
}

//FollowingUsers stores a list of FollowingUser
type FollowingUsers = []followingUser

//Gists contains gists information
type Gists []gist

type gist struct {
	URL         string                 `json:"html_url"`
	CreatedAt   string                 `json:"created_at"`
	Description string                 `json:"description"`
	Files       map[string]interface{} `json:"files"`
}
