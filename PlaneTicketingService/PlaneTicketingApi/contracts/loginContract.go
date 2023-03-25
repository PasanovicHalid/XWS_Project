package contracts

import (
	"encoding/json"
	"io"
)

type LoginContract struct {
	Username string
	Password string
}

func (l *LoginContract) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(l)
}
