package contracts

import (
	"encoding/json"
	"io"
)

type City struct {
	Name string
}

func (f *City) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(f)
}
