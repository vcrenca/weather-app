package http

type ErrorResponse struct {
	Title    string              `json:"title"`
	Status   int                 `json:"status"`
	Detail   string              `json:"detail,omitempty"`
	Instance string              `json:"instance"`
	Errors   []map[string]string `json:"errors,omitempty"`
}
