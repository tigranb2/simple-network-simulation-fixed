# simple-network-simulation-fixed

# To Run:

First, open up the config.json file in the messages folder, and set the desired ip and port for each id.
Feel free to add id numbers, but do not skip lines. 
   
In one terminal instance, open to the directory and type:
```
go run server.go ip:port
```
ip:port should allign with the ip and port you entered for the processes in the config.json file.   

In other terminal instances, type:
```
go run client.go
```
You will be prompted to input an id number for each terminal instances open to client.go  

In order to send messages, type messages with the following format in the client.go terminal instances:
```
send 2 hello
```
2 should be replaced with the destination id. hello should be replaced with the content of the message.

# Design
## File Structure: 
The application has the following structure:
```
   -message
      -config.go
      -message.go
   -client.go
   -server.go
   -config.json
```

## Network:
The tcp server is made and handled in server.go. It makes use of a goroutine so that it can handle multiple connections at once.  
There are two channels withing server.go:
```
   conn := make(chan net.Conn)
	unicast := make(chan message.Message)
```
- conn handles incoming messages from connections to the server. It also assigns an id to each connection.  
- unicast checks if the destination exists, and sends the message to the intended connection.

## Client: 
client.go handles client connections to the server. It has 5 helper functions:
```
   func definePocess(processId string) message.Config
   func dial(processInfo message.Config) net.Conn
   func listen(c net.Conn)
   func constructMessage(processId string, msg string, message message.Message) message.Message
   func delay(delays string)
```
- defineProcess uses the id provided by the user and the info in config.json in order to figure out which server to dial
- dial connects to the tcp server using info from config.json
- listen recieves messages from the server and prints them
- constructMessage assigns the user input to fields of the Message struct declared in message.go
- delay simulates a network delay. The delay is within two integers declared in config.json

client.go uses a goroutine to handle messages incoming from the server.  

## Message:
message.go creates a custom Message data type: 
```
type Message struct {
	FromId, ToId, Content string
}
```

## Config:
config.go uses a bufio scanner to read the config.json file and store it's information. 

## Citations: 
- TCP Server info: https://www.linode.com/docs/development/go/developing-udp-and-tcp-clients-and-servers-in-go/    
- Golang documentation: https://golang.org/   
- Generating a random int: https://golang.cafe/blog/golang-random-number-generator.html   


