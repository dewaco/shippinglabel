package types

type Status struct {
	Code  int    `json:"code,omitempty"`
	Error *Error `json:"error,omitempty"`
}
