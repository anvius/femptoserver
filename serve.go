package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {

	var port string

	portFlag := flag.Int("port", 5555, "Port to use in localhost")

	flag.Parse()

	fmt.Println("\nFemptoServer SERVE v0.1 : License MIT, 2018, Antonio Villamarin\n")

	if *portFlag < 1000 || *portFlag > 10000 {
		fmt.Println("Error 1: port must be in 1000-10000")
		fmt.Println("Format: serve [--port=NNNN]")
		os.Exit(1)
	} else if len(flag.Args()) != 0 {
		fmt.Println("Error 2: too many arguments")
		fmt.Println("Format: serve [--port=NNNN]")
		os.Exit(2)
	} else {
		port = strconv.Itoa(*portFlag)
	}

	fmt.Println("Serve initialized at http://localhost:" + port)
	fmt.Println("Ctrl-C to cancel.\n")
	fmt.Println("Working...")

	http.ListenAndServe(":"+port, http.FileServer(http.Dir(".")))

}
