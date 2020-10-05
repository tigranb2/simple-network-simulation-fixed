# simple-network-simulation-fixed

## To Run:

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
