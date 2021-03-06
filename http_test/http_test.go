package http_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestRequest1(t *testing.T) {
	resp, err := http.Get("https://baidu.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	fmt.Printf(string(s))
}

func TestRequest2(t *testing.T) {
	params := url.Values{}

	Url, err := url.Parse("http://baidu.com?fd=fdsf")
	if err != nil {
		panic(err.Error())
	}
	params.Set("a", "fdfds")
	params.Set("id", string("1"))
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := http.Get(urlPath)
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(s))
}
