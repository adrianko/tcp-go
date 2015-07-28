package main

import (
    "bytes"
    "io"
    "log"
    "net"
    "strconv"
    "time"
)

const (
    CONN_HOST = "localhost"
    CONN_PORT = "9999"
    CONN_TYPE = "tcp"
)

func check(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func handle(conn net.Conn) {
    log.Printf("handle")
    var buf bytes.Buffer
    io.Copy(&buf, conn)
    log.Printf("Recieved: %s", buf)
    body := "<h1>Hello World</h1>"
    response := "HTTP/1.1 200 OK\n"
    response += "Server Maester/0.6\n"
    response += "Date: " + time.Now().Format(time.RFC1123) + "\n"
    response += "Content-Type: text/html\n"
    response += "Content-Length: " + strconv.Itoa(len(body)) + "\n"
    response += "\n"
    response += body
    conn.Write([]byte(response))
    conn.Close()
}

func main() {
    listener, err := net.Listen(CONN_TYPE, CONN_HOST + ":" + CONN_PORT)
    check(err)
    defer listener.Close()
    log.Printf("Listening on %s:%s", CONN_HOST, CONN_PORT)

    for {
        conn, err := listener.Accept()
        check(err)
        go handle(conn)
    }
}
