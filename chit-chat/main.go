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
	user := flag.String("user", "", "USER is what will precede your messages")
	flag.Parse()

	if *ip == "" {
		fmt.Println("You have not entered in an ip address using the flag -ip")
		return
	}
	if *user == "" {
		fmt.Println("You have not entered in a username using the flag -user")
		return
	}

	if *client == true {
		for {
			response, err := http.Get("http://" + *ip + ":9000/chat")
			defer response.Body.Close()
			if err != nil {
				fmt.Println(err.Error())
			}

			bs, err := ioutil.ReadAll(response.Body)
			if string(bs) != "" {
				tr := string(bs)
				fmt.Println(tr)
			}
		}
	}

	if *repl == true {
		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter text: ")
			text, _ := reader.ReadString('\n')
			v := url.Values{}
			v.Add("body", *user+": "+text)
			_, _ = http.PostForm("http://"+*ip+":9000/", v)
		}
	}
}
