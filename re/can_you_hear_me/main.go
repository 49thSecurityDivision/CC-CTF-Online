package main

import (
	"bufio"
	"crypto/aes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
)

func getFlag(lore string) {
	key, _ := base64.StdEncoding.DecodeString(lore)
	// start := strings.Index(string(key), "{SOME START}")
	key = []byte(string(key[start : start+32]))
	pt := ""
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Print(err)
		return
	}
	out := make([]byte, len(pt))
	c.Encrypt(out, []byte(pt))
}

func getKey(key []byte, flag string) string {
	text, _ := hex.DecodeString(flag)
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Print("Hey, you got this error... %s", err)
	}

	pt := make([]byte, 38)
	c.Decrypt(pt, text)
	return string(pt[:])
}

func main() {
	lore := "SXQgbWF5IHNvdW5kIHN0cmFuZ2UsIGJ1dCB0aGUgd29yZCDigJhiYXJiZWN1ZeKAmSBpcyBhY3R1YWxseSBhIGRlcml2YXRpdmUgb2YgQ2FyaWJiZWFuIGRpYWxlY3QsIGFjY29yZGluZyB0byBTb3V0aGVybiBPcmFsIEhpc3RvcnkuIEJhcmJlY3VlIGNvbWVzIGZyb20gVGFpbm8sIGEgcHJlLUNvbHVtYmlhbiBDYXJpYmJlYW4gbGFuZ3VhZ2UuIFRoZSB3b3JkLCB0aGUgd2Vic2l0ZSBzdGF0ZXMsIGRlc2NyaWJlcyB0aGUgbmF0aXZlIG1ldGhvZCBvZiBjb29raW5nIHNsaWNlZCBtZWF0cyBvdmVyIGFuIG9wZW4gZmxhbWU7IG90aGVyIHNvdXJjZXMgc2F5IGJhcmJhY29hIHNwZWNpZmljYWxseSByZWZlcnJlZCB0byB0aGUgd29vZGVuIGZyYW1lIG9uIHdoaWNoIHRoZSBtZWF0IHdhcyBzbW9rZWQuCg=="
	if lore == lore {
		// Then it was used :)
		fmt.Print("What is the key:\n")
	}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	flag := "e0fdc9b7d78fd232b9cd1b24df16b1b1"

	input := scanner.Text()
	key := []byte(input)
	if string(getKey(key, flag)[0:4]) == "flag" {
		fmt.Printf("The flag is:\n%s\n", getKey(key, flag))
	} else {
		fmt.Print("Please try again")
	}
}
