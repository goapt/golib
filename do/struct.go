package do

import (
	"encoding/json"
)

type H map[string]interface{}

func (h H) MustJSON() []byte {
	d, _ := json.Marshal(h)
	return d
}