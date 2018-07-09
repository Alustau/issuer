package request

//Issue ...
type Issue struct {
	Action string                 `json:"action"`
	Issue  map[string]interface{} `json:"issue"`
}
