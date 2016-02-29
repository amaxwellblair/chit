package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func main() {
	clientPtr := flag.Bool("client", false, "bar")
	replPtr := flag.Bool("repl", false, "bar")
	flag.Parse()

	if *clientPtr == true {
		for {
			response, err := http.Get("http://localhost:9000/chat")
			if err != nil {
				fmt.Println(err.Error())
			}

			bs, err := ioutil.ReadAll(response.Body)

			tr := string(bs)
			fmt.Println(tr)
		}
	}

	if *replPtr == true {
		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter text: ")
			text, _ := reader.ReadString('\n')
			v := url.Values{}
			v.Add("body", text)
			_, _ = http.PostForm("http://localhost:9000/", v)
		}
	}
}

// client - recieves messages
// long polling for messages
// coordinator - collects messages from client and responds with the same message
// only sends messages when a CLI sends a message to the coordinator
// CLI - sends messages
