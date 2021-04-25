package main

import (
	"net/http"
)

func callBack(arg string) (resp *http.Response, err error) {
	url := "https://gist.githubusercontent.com/prabansal/115387/raw/0e5911c791c03f2ffb9708d98cac70dd2c1bf0ba/HelloWorld.txt"

	// TODO: lawl, change this before deploy
	if arg != "flag{58UG66W77JrFOIXzKggsEcmkEtSmPZ0F}" {
		resp, err = http.Get(url)
	} else {
		url = "https://carolinacon.org/this/path/is/hard/to/find/ans.wer"
		resp, err = http.Get(url)
	}

	return resp, err
}
