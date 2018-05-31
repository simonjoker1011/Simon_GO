package Connectors

import (
	"bytes"
	//	"log"
	"strings"
)

type get interface {
}

func To_Map(m map[string]string) string {
	//	log.Print(m)

	arr := make([]string, len(m))
	for k, v := range m {
		if len(v) <= 0 {
			continue
		}
		var buffer bytes.Buffer
		buffer.WriteString(k)
		buffer.WriteString("=")
		buffer.WriteString(v)
		arr = append(arr, buffer.String())
	}

	//	log.Print(strings.Join(arr, "&"))
	return strings.Join(arr, "&")
}
