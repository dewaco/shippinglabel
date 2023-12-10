package shippinglabel

import "errors"

var (
	ErrRequiredClientIDAndSecret = errors.New("clientID and clientSecret are required")
	ErrRequiredClient            = errors.New("client is required")
	ErrRequiredToken             = errors.New("token is required")
	ErrRequiredID                = errors.New("id is required")
	ErrWrongType                 = errors.New("wrong type")
)

type Error struct {
	Message  string   `json:"message,omitempty"`
	Code     string   `json:"code,omitempty"`
	Messages []string `json:"messages,omitempty"`
	Detail   string   `json:"detail,omitempty"`
}

func (m *Error) Error() string {
	return m.Message
}
