package main 

import (
	"bufio"
	"fmt"
	"os"
	"encoding/gob"
	"strconv"
	"strings"
	"net"
	"time"
	"simple-network-simulation-fixed/message"
	"math/rand"
)

func main() {
	
	fmt.Println("Enter an ID: ")
	ri, _ := bufio.NewReader(os.Stdin).ReadString('\n') //recieves id from user 
	ri = strings.TrimSuffix(ri, "\r\n")

	processId := ri //assigns id to user input
	processInfo := defineProcess(processId)
	delays := message.Delays 
	message := message.Message{}

	c := dial(processInfo) 
	gob.NewEncoder(c).Encode(ri) //send the id to the server
	go listen(c) //go routine for handling incoming messages

	for {
		reader := bufio.NewReader(os.Stdin)	//reads user input
		msg, _ := reader.ReadString('\n')
		message = constructMessage(processId, msg, message) //construct a message using the user input
		
		fmt.Printf("Sent '%s' %s at %s\r\n", message.ToId, message.Content, time.Now())
		delay(delays) //simulate a network delay
		encoder := gob.NewEncoder(c)
		encoder.Encode(message)
		
	}

}

//uses values from config.json for the process 
func defineProcess(processId string) message.Config{
	id, _ := strconv.Atoi(processId)
	_, ok := message.Processes[id] //checks if info is specified for this id
	processInfo := message.Config{}

	if ok { //if there are specs for this info, assign the corresponding ip and port
		pI := strings.Fields(message.Processes[id]) 
		processInfo.Id = pI[0]
		processInfo.Ip = pI[1]
		processInfo.Port = pI[2]
	} else {
		fmt.Println("No info discovered for this process ID. Check config.json file.")
	}

	return processInfo
}

 //recieves messages from the server
func listen(c net.Conn) {
	for {
		decoder := gob.NewDecoder(c)
		message := message.Message{}
		decoder.Decode(&message)
		if message.Content != "" { //only prints if the message is valid
			fmt.Printf("Recieved '%s' from %s, system time is %s.\r\n", message.Content, message.FromId, time.Now())
		}
		
	}
}

 //connects to server
func dial(processInfo message.Config) net.Conn { 
	CONNECT := processInfo.Ip+":"+processInfo.Port
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
	}

	return c
}

//constructs a message using user input
func constructMessage(processId string, msg string, message message.Message) message.Message{
	messageWords := strings.Fields(msg)
	message.FromId = processId
	message.ToId = messageWords[1]
	content := messageWords[2:len(messageWords)] //excludes "send" and destination from the message
	message.Content = strings.Join(content, " ")
	return message
}

//simulates a delay by putting the process to sleep
func delay(delays string) { 
	d := strings.Fields(delays)
	minD, _ := strconv.Atoi(d[0])
	maxD, _ := strconv.Atoi(d[1])
	delay := rand.Intn(maxD - minD) + minD //get random integer between the minimum and maximum delays
	time.Sleep(time.Duration(delay)*time.Millisecond)
}