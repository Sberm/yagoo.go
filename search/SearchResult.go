package search

type SearchResult struct {
	ID    string `json:"id"`
	Dir   string `json:"dir"`
	Text  string `json:"text"`
	Token string `json:"token"`
	Url   string `json:"url"`
}
