package decode

import (
	"encoding/json"
	"io"
)

func FromJSON(data interface{}, r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(data)
}
