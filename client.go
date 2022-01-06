package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"bufio"
)

func getArgs() []int {
  var s []int
	if len(os.Args) != 3 {
		fmt.Printf("Usage: go run client.go <portnumber>\n")
		os.Exit(1)
	} else {
		  fmt.Printf("#DEBUG ARGS Port Number : %s\n", os.Args[1])
		  portNumber, err := strconv.Atoi(os.Args[1])
      s= append(s,portNumber)
		  if err != nil {
			  fmt.Printf("Usage: go run client.go <portnumber>\n")
			  os.Exit(1)
      }
      
      para, err := strconv.Atoi(os.Args[2])
		  if err != nil {
			  os.Exit(1)

        }
      s=append(s,para)
      }
	return s
}

func main() {
	para := getArgs()
	fmt.Printf("#DEBUG DIALING TCP Server on port %d\n", para[0])
	portString := fmt.Sprintf("127.0.0.1:%s", strconv.Itoa(para[0]))
	fmt.Printf("#DEBUG MAIN PORT STRING |%s|\n", portString)

	conn, err := net.Dial("tcp", portString)
	if err != nil {
		fmt.Printf("#DEBUG MAIN could not connect\n")
		os.Exit(1)
	} else {

        defer conn.Close()
		    fmt.Printf("#DEBUG MAIN connected\n")
	    fmt.Println(strconv.Itoa(para[1]))
        io.WriteString(conn,strconv.Itoa(para[1])+"\n")
        reader := bufio.NewReader(conn)
        for {
            resultString, err := reader.ReadString('\n')
            if (err != nil){
                fmt.Printf("DEBUG MAIN could not read from server")
                os.Exit(1)
            	}
            fmt.Println(resultString)
						}
				}
}
