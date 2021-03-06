package main

import (
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
	// "encoding/json"
	"github.com/bitly/go-simplejson"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("eg: yd hello")
		fmt.Println("eg: yd hello world")
		return
	}
	var q, sep string
	for i := 1; i < len(os.Args); i++ {
		q += sep + os.Args[i]
		sep = " "
	}
	base_url := "http://fanyi.youdao.com/openapi.do?keyfrom=learnGo&key=50779724&type=data&doctype=json&version=1.1&q="
	request_url := base_url + q
	resp, err := http.Get(request_url)
	if err != nil {
		fmt.Sprint(err)
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Sprintf("while reading %s: %v", request_url, err)
		os.Exit(1)
	}
	// fmt.Printf("%s", b)
	js, err := simplejson.NewJson(b)
	if err != nil {
		fmt.Sprint(err)
		return
  }
  // explains := js.Get("translation").MustArray()
  if len(os.Args) == 2 {
  	explains := js.Get("basic").Get("explains").MustArray()
  	for i := 0; i < len(explains); i++ {
  		fmt.Println(explains[i])
 		}
  }else{
  	explains := js.Get("translation").MustArray()
  	fmt.Println(explains[0])
  }
}

