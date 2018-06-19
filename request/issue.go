package request

//IssueRequest ...
type IssueRequest struct {
	Action string `json:"action"`
	Issue  Issue  `json:"issue"`
}

//Issue ...
type Issue struct {
	Number int     `json:"number"`
	Title  string  `json:"title"`
	Labels []Label `json:"labels"`
}

//Label ...
type Label struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}
