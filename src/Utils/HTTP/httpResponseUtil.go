package Utils

import (
	"encoding/json"
	"net/http"
)

func BuildJsonResp(w http.ResponseWriter, obj interface{}) http.ResponseWriter {
	w.Header().Add("Content-Type", "application/json")
	b, _ := json.Marshal(obj)
	w.Write(b)

	return w
}
