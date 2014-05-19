package main

import (
	"fmt"
	"github.com/imeyer/go-utilities/hostnameutils"
	"log"
	"os"
)

func ReverseDemo() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Panic("Can not call os.Hostname(): ", err)
	}

	fmt.Println("##### Reverse Demo #####")
	fmt.Printf("HOSTNAME BEFORE: %v\n", hostname)
	fmt.Printf("HOSTNAME AFTER:  %v\n", hostnameutils.Reverse(hostname))
}

func ReverseOffsetDemo() {
	// Set a manual hostname
	hostname := "this.is.a.test.domain.com"

	fmt.Println("##### Reverse Offset Demo of 2 #####")
	fmt.Printf("HOSTNAME BEFORE: %v\n", hostname)
	fmt.Printf("HOSTNAME AFTER:  %v\n", hostnameutils.ReverseOffset(hostname, 2))
}

func main() {
	ReverseDemo()
	fmt.Println()
	ReverseOffsetDemo()
}
