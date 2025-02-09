package main

import (
	"fmt"
	"go-server/routes"
	"log"      //used for logging errors(it stops the server if something goes wrong)
	"net/http" //used for creating a web server
)

func main() {
	// port := flag.String("port", "8080", "Port to listen on")
	// flag.Parse()

	// //start the http server
	// serverAddress := ":"+*port
	// fmt.Printf("SErver is starting on port %s...\n", *port)

	// go func(){
	// 	log.Fatal(http.ListenAndServe(serverAddress,routes.InitializeRoutes()))

	// }()

	// 	signalChan := make(chan os.Signal,1)
	// 	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	// 	<-signalChan
	// 	fmt.Println("\n shutting down")

	routes.InitializeRoutes()
	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// // w http.ResponseWriter is used to send data back to the client
// // r *http.Request contains all the details about the incoming request
// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello world")
// }

// func main() {

// 	//the below line registers our handler function for the / router
// 	http.HandleFunc("/", handler)

// 	fmt.Println("server is running on the port 8080...")

// 	//web server start on port 8080
// 	//the second parameter we pass is nil meaning we are usig the default settings
// 	//log.Fatal(...) ensures that is the server crashes, it will print a error message
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }
