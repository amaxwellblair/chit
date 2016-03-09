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

// Client gets messages from chit server
func Client(ip string) {
	for {
		response, err := http.Get("http://" + ip + ":9000/chat")
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

// REPL reads a given string evaluates and prints to the server via post
func REPL(ip, user string) {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		v := url.Values{}
		v.Add("body", user+": "+text)
		_, err := http.PostForm("http://"+ip+":9000/chat", v)
		if err != nil {
			fmt.Println(err)
		}
	}
}

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
	if *client == true {
		Client(*ip)
	}
	if *user == "" {
		fmt.Println("You have not entered in a username using the flag -user")
		return
	}
	if *repl == true {
		REPL(*ip, *user)
	}
}
