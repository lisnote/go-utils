package main

import (
    "fmt"
    "log"
    "net"
    "os"
    "speed-test/config"
    "time"
)

func main() {
    if len(os.Args) < 2 {
        log.Fatal("usage: server <host:port>")
    }
    addr := os.Args[1]

    ln, err := net.Listen("tcp", addr)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("listening on %s\n", addr)

    for {
        conn, err := ln.Accept()
        if err != nil {
            log.Println("accept error:", err)
            continue
        }
        go handleConn(conn)
    }
}

func handleConn(c net.Conn) {
    defer c.Close()

    buf := make([]byte, config.BufSize)
    totalRead, totalWrite := 0, 0
    endTime := time.Now().Add(config.Duration)

    for time.Now().Before(endTime) {
        // 读数据
        c.SetReadDeadline(time.Now().Add(config.Timeout))
        n, err := c.Read(buf)
        if n > 0 {
            totalRead += n
            // 回写给客户端
            c.SetWriteDeadline(time.Now().Add(config.Timeout))
            wn, werr := c.Write(buf[:n])
            totalWrite += wn
            if werr != nil {
                log.Println("write error:", werr)
                break
            }
        }
        if err != nil {
            if ne, ok := err.(net.Error); ok && ne.Timeout() {
                continue
            }
            log.Println("read error:", err)
            break
        }
    }

    fmt.Printf("client %s finished: totalRead=%d bytes, totalWrite=%d bytes\n", c.RemoteAddr(), totalRead, totalWrite)
}
