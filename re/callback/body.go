package main

import (
	"io"
	"io/ioutil"
)

func body(resBody io.ReadCloser) (body []byte, err error){
	body, err = ioutil.ReadAll(resBody)
	return body, err
}
