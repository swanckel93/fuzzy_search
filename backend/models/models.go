package models

type SearchRequest struct {
	FileID string `json:"file_id"`
	Query  string `json:"query"`
}

type ExpandContextRequest struct {
	FileID string `json:"file_id"`
	Index  int    `json:"index"`
}

type SearchResult struct {
	Sentence string `json:"sentence"`
	Index    int    `json:"index"`
	Match    string `json:"match"`
	Distance int    `json:"distance"`
}
