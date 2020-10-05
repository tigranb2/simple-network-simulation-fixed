package message
import (
	"bufio"
	"fmt"
	"os"
)

type Config struct {
	Id, Ip, Port string
}


func readConfig() (map[int]string, string){
	line := 0
	delays := ""
	file, err := os.Open("config.json") //opens the file
	if err != nil {
		fmt.Println(err) //handles error
	}

	defer file.Close()

	scanner := bufio.NewScanner(file) //new scanner for scanning file word by word

	processes := make(map[int]string)
	for scanner.Scan() {
		if line != 0 { //scans id info
			processes[line] = scanner.Text() //each line is an individual id; uses id as key
		} else { //scans delay info
			delays = scanner.Text()
		}
		
		line++
	}
	return processes, delays
}

var Processes, Delays = readConfig()
