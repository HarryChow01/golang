package http_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestSimpleRequest(t *testing.T) {
	resp, err := http.Get("https://baidu.com")
	if err != nil {
		panic(err)

	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	fmt.Printf(string(s))
}
