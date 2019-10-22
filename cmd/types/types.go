package types

//UserInfo contains all user related information
type UserInfo struct {
	Name        string `json:"name"`
	Location    string `json:"location"`
	PublicRepos int    `json:"public_repos"`
}

//PublicGist contains basic info of publi gist
type PublicGist struct {
	Owner       map[string]interface{} `jsob:"owner"`
	URL         string                 `json:"html_url"`
	CreatedAt   string                 `json:"created_at"`
	Description string                 `json:"description"`
}
