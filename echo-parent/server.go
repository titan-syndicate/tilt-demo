//server.go

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	// "os"
	// "strconv"
	// "bytes"
)

func main() {

	http.HandleFunc("/echo-child-1", Child1Handler)
	http.HandleFunc("/echo-child-2", Child2Handler)
	fmt.Println("Server started at port 3333")
	log.Fatal(http.ListenAndServe(":3333", nil))
}

func Child1Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("request received %s!\n", r.URL.Path[1:])

	resp, err := http.Get("http://echo-child-1")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintf(w, string(body))
}

func Child2Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("request received %s!\n", r.URL.Path[1:])

	resp, err := http.Get("http://echo-child-2")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintf(w, string(body))
}

// const (
//     CONN_HOST = ""
//     CONN_PORT = "3333"
//     CONN_TYPE = "tcp"
// )

// func main() {
//     // Listen for incoming connections.
//     l, err := net.Listen(CONN_TYPE, ":"+CONN_PORT)
//     if err != nil {
//         fmt.Println("Error listening:", err.Error())
//         os.Exit(1)
//     }
//     // Close the listener when the application closes.
//     defer l.Close()
//     fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
//     for {
//         // Listen for an incoming connection.
//         conn, err := l.Accept()
//         if err != nil {
//             fmt.Println("Error accepting: ", err.Error())
//             os.Exit(1)
//         }

//         //logs an incoming message
//         fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())

//         // Handle connections in a new goroutine.
//         go handleRequest(conn)
//     }
// }

// Handles incoming requests.
// func handleRequest(conn net.Conn) {
//   // Make a buffer to hold incoming data.
//   buf := make([]byte, 1024)
//   // Read the incoming connection into the buffer.
//   reqLen, err := conn.Read(buf)
//   if err != nil {
//     fmt.Println("Error reading:", err.Error())
//   }
//   // Builds the message.
//   message := "Hi, I received your message! It was "
//   message += strconv.Itoa(reqLen)
//   message += " bytes long and that's what it said: \""
//   n := bytes.Index(buf, []byte{0})
//   message += string(buf[:n-1])
//   message += "\" ! Honestly I have no clue about what to do with your messages, so Bye Bye!\n"

//   // Write the message in the connection channel.
//   conn.Write([]byte(message));
//   // Close the connection when you're done with it.
//   conn.Close()
// }