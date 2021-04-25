package websocket

import (

	"fmt"
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)


type HashStruct struct {
	Hash string
}

func SelectWord () string {
	words := [3]string{"hello", "hello", "hello"}
	rand.Seed(time.Now().UnixNano())
	return words[rand.Intn(len(words) - 0) + 0]
	//fmt.Println(words[rand.Intn(len(words) - 0) + 0])

}


func MakeItChunky(pool *Pool, hashstring string)  {
	rand.Seed(time.Now().UnixNano());
	ultimateMax := 32;
	startingMax := 6;
	startingMin := 1;

	i := 0;
	lastEnd := 0;
	newHash := HashStruct{};
	for i <= ultimateMax {
		time.Sleep(time.Second * 1);
		randInt := rand.Intn(startingMax - startingMin) + startingMin;
		endval := i + randInt;
		if ((i + randInt) >= ultimateMax) {
				//fmt.Println("Toobig");
				randInt = ultimateMax - i;
				i = lastEnd;
				endval = ultimateMax;
				fmt.Println(hashstring[i:endval]);
				newHash = HashStruct{Hash: hashstring[i:endval]}
				pool.Broadcast <- newHash;
				break;
		} else {
				//fmt.Printf("i = %d & endval = %d\n", i, endval);
				fmt.Println(hashstring[i:endval]);
				newHash = HashStruct{Hash: hashstring[i:endval]}
				pool.Broadcast <- newHash;
				lastEnd = endval;
				i += randInt;
		}
	}

}


func GenHash(pool *Pool, word string) {

	algorithm := md5.New()
		pool.SendHash <- word
		algorithm.Write([]byte(word))
		hashString := hex.EncodeToString(algorithm.Sum(nil))
		fmt.Printf("The secret word is: %s and the hash is: %s\n", word, hashString);
	for {
		MakeItChunky(pool, hashString);
	}
}