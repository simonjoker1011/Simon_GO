package Utils

import (
	"testing"
)

type conf struct {
	A string            `yaml:"attr1"`
	B int               `yaml:"attr2"`
	C bool              `yaml:"attr3"`
	D map[string]string `yaml:"attr4"`
}

func Test_Loading(t *testing.T) {
	var c conf
	Load_settings(&c, "test.yaml")
	t.Log(c)
}
