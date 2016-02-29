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
	client := flag.Bool("client", false, "CLIENT recieves messaages from other chit-chat clients")
	repl := flag.Bool("repl", false, "REPL allows you send messages")
	ip := flag.String("ip", "", "IP is the address Client and REPL use to communicate to the server")
	flag.Parse()

	if *ip == "" {
		fmt.Println("You have not entered in an ip address using the flag -ip")
		return
	}

	if *client == true {
		for {
			response, err := http.Get("http://" + *ip + ":9000/chat")
			if err != nil {
				fmt.Println(err.Error())
			}

			bs, err := ioutil.ReadAll(response.Body)

			tr := string(bs)
			fmt.Println(tr)
		}
	}

	if *repl == true {
		for {
			fmt.Println(*ip)
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter text: ")
			text, _ := reader.ReadString('\n')
			v := url.Values{}
			v.Add("body", text)
			_, _ = http.PostForm("http://"+*ip+":9000/", v)
		}
	}
}

// client - recieves messages
// long polling for messages
// coordinator - collects messages from client and responds with the same message
// only sends messages when a CLI sends a message to the coordinator
// CLI - sends messages
