package contracts

import (
	"encoding/json"
	"io"
)

type ApiKeyGenerationContract struct {
	UserId          string
	Duration        string
	DurationForever bool
}

func (s *ApiKeyGenerationContract) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(s)
}
