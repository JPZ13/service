package model

// CreateAuthorRequest is the request body to
// create an author
type CreateAuthorRequest struct {
	ID int `json:"id"`
}

// CreateAuthorResponse is the response body when
// creating an author
type CreateAuthorResponse struct {
	ID int `json:"id"`
}
