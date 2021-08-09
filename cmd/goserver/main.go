package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	port  = ""
	route = "./dist"
)

func main() {
	readFlags()
	runServer()
}

func readFlags() {
	var newPort int
	var customPath string

	defaultPort, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		defaultPort = 8080
	}

	// Flag Port (-port)
	// Flag to define a custom port
	flag.IntVar(&newPort, "port", defaultPort, "Custom port")

	// Flag Path (-path)
	// Flag to define a custom Path
	flag.StringVar(&customPath, "path", "", "Define a custom route to serve")

	// Read Flags
	flag.Parse()

	// Set custom port
	newStringPort := string(":" + strconv.Itoa(newPort))
	port = newStringPort

	// Set custom route
	if customPath != "" {
		route = customPath
	}
}

func runServer() {
	fmt.Printf(" * GoServer is Running in http://localhost%s * \n", port)
	http.Handle("/", http.FileServer(http.Dir(route)))
	log.Fatal(http.ListenAndServe(port, nil))
}
