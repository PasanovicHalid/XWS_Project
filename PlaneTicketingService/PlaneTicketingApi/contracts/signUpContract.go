package contracts

import (
	"encoding/json"
	"io"
)

type SignUpContract struct {
	Username  string
	Password  string
	Firstname string
	Lastname  string
}

func (s *SignUpContract) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(s)
}
