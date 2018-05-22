package Connectors

import (
	"Constant"
	"bytes"
	"net/url"
	"testing"
)

func TestNearbySearch(t *testing.T) {
	t.Logf("start %v test", Constant.Api_Method_NearbySearch)
	var buffer bytes.Buffer
	buffer.WriteString(Constant.Api_Server)
	buffer.WriteString(Constant.Api_Method_Detail)
	buffer.WriteString("json?placeid=ChIJKzfU3sSrQjQRCzAF0PADjQE&key=AIzaSyBUtCV0pLaG0ovuQSQIM0RyoEILkWWquqo")
	u, err := url.Parse(buffer.String())

}
func TestDetail(t *testing.T) {
	t.Logf("start %v test", Constant.Api_Method_Detail)
}
