package encode

import (
	"encoding/json"
	"io"
)

func ToJSON(data interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(data)
}
