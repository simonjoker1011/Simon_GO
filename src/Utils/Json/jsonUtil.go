package Utils

import (
	"encoding/json"
)

func Marshal(in interface{}) []byte {
	out, err := json.Marshal(in)
	if err != nil {
		panic(err)
	}
	return out
}

func Marshal_To_String(in interface{}) string {
	return string(Marshal(in)[:])
}
