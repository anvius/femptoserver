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

	portFlag := flag.Int("port", 5555, "Port in localhost to serve files")

	flag.Parse()

	if *portFlag < 1000 || *portFlag > 10000 {
		fmt.Println("FemptoServer SERVE v0.1")
		fmt.Println("Error port value: must be in 1000-10000")
		fmt.Println("Format: serve [--port=NNNN]")
		os.Exit(1)
	} else if len(flag.Args()) != 0 {
		fmt.Println("FemptoServer SERVE v0.1")
		fmt.Println("Error arguments: too many arguments")
		fmt.Println("Format: serve [--port=NNNN]")
		os.Exit(2)
	} else {
		port = strconv.Itoa(*portFlag)
	}

	fmt.Println("FemptoServer SERVE v0.1")
	fmt.Println("FemptoServer initialized at http://localhost:" + port)
	fmt.Println("Ctrl-C to cancel...")

	http.ListenAndServe(":"+port, http.FileServer(http.Dir(".")))

}
