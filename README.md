## Primitive command line chat client

Basic chat server called *chit* along with a REPL and client, called
*chit-chat*, to read and write messages to a chatroom.

### Chit

Install or build the binary on your chosen server. My particular implementation
runs on a [digital ocean](https://www.digitalocean.com/pricing/) droplet (512 MB SSD).
 Chit currently runs on port 9000, and can be changed if necessary.

Once installed simply run the below to start your chit server:
 ```
$ cd chit
$ ./chit
 ```

### Chit-chat

Install or build the binary on your chosen client. It should not be a problem
running the client on most machines. Chit-chat serves as the REPL for entering
 messages and as the view for receiving messages. In order to keep the client
generic all commands using `chit-chat` require the chit server IP address.

Once installed run the below to start your client view:
```
$ cd chit/chit-chat
$ ./chit-chat -client -ip=0.0.0.0
```
This will be blank until you begin to receive messages

To begin your client REPL run the below:
```
$ cd chit/chit-chat
$ ./chit-chat -repl -user=[username] -ip=0.0.0.0
```
This will begin with the `Enter text:`

What your client view and REPL could look like using tmux:
<img width="500px" src="http://www.amaxwellblair.com/images/chat_repl.png" align="center">

### Angular front end

Also check out this implementation with an Angular front end: https://github.com/p4tin/simplegochat
