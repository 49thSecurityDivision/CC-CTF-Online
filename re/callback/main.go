package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./main <key>")
		os.Exit(1)
	}

	arg := os.Args[1]

	resp, err := callBack(arg)

	if err != nil {
	    log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := body(resp.Body)

	if err != nil {

	    log.Fatal(err)
	}

	fmt.Println(string(body))
}
