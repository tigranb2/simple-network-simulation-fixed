package main

import (
	"fmt"
	"net"
	"os"
	"log"
	"encoding/gob"
	"simple-network-simulation-fixed/message"
)

func main(){
	//recieves ip and port from user
	arguments := os.Args 
	if len(arguments) == 1 {
			fmt.Println("Please provide ip:port.") 
			return
	}
	
	//create channel on the specified ip:port
	CONNECT := arguments[1]
	ln, err := net.Listen("tcp", CONNECT)
	if err != nil {
		log.Println(err.Error())
	}

	conn := make(chan net.Conn)
	mconn := make(map[string]net.Conn) //used to reference connections by id
	unicast := make(chan message.Message)
	id := ""
	m := message.Message{}
	
	//Accepts connections and sends it to conn channel
	go func () {
		for {
			c, err := ln.Accept()
			
			if err != nil {
				log.Println(err.Error())
			}
			
			conn <- c
			}
		}()

	for {
		select {
		//assigns id to connection, recieves messages from connecton
		case c := <-conn:

			idDec := gob.NewDecoder(c) //recives process id
			idDec.Decode(&id)
			mconn[id] = c

			go func(c net.Conn) {
				decoder := gob.NewDecoder(c)
			 	for {
					decoder.Decode(&m)
					unicast <- m 		
				}
			}(c)

		//sends message to correct connection 
		case msg := <-unicast:
			_, ok := mconn[msg.ToId]

			if ok { //only send message if destination exisits
				encoder := gob.NewEncoder(mconn[msg.ToId])
				encoder.Encode(msg)
			} else {
				fmt.Println("This process does not exist.")
			}
				
			
			} 
			
 		}

}

