// grab price quote from MT4 server (which using "web plugin")
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	// MT4 server
	host string = "207.228.252.22:443"
)

func printQuotes(symbols []string) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(conn, "WQUOTES-%s,\n", strings.Join(symbols, ","))
	r := bufio.NewReader(conn)

	i := 0
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			panic(err)
		}
		fmt.Print(line)
		if i == len(symbols) {
			break
		}
		i++
	}
}

func main() {
	symbols := []string{"EURUSD", "EURGBP", "USDJPY"}
	c := time.Tick(2 * time.Second)
	for t := range c {
		go func() {
			printQuotes(symbols)
			fmt.Printf("%v\n\n", t)
		}()
	}

}
