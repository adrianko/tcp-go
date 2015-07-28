package main

import (
    "bytes"
    "io"
    "log"
    "net"
)

const (
    CONN_HOST = "localhost"
    CONN_PORT = "3333"
    CONN_TYPE = "tcp"
)

func check(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func handle(conn net.Conn) {
    var buf bytes.Buffer
    io.Copy(&buf, conn)
    log.Printf("Recieved: %s", buf.String())
    conn.Write([]byte("Message recieved."))
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
