package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"bufio"
    "strings"
    "time"
)

func getArgs() []int {
  var s []int
	if len(os.Args) != 6 {
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
      for k := 2 ; k < 6 ; k++ {
        para, err := strconv.Atoi(os.Args[k])
        s= append(s,para)
		    if err != nil {
			    os.Exit(1)
        }
      }
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
        reader := bufio.NewReader(conn)
		fmt.Printf("#DEBUG MAIN connected\n")
        for i:= 1; i < 5; i++{

            io.WriteString(conn,strconv.Itoa(para[i])+"\n")
        }
        for {
            resultString, err := reader.ReadString('\n')
            if (err != nil){
                fmt.Printf("DEBUG MAIN could not read from server")
                os.Exit(1)
            }
            resultString = strings.TrimSuffix(resultString, "\n")
            fmt.Printf("#DEBUG server replied : |%s|\n", resultString)
            

	}

}
