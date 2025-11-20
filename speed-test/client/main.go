package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"net"
	"os"
	"speed-test/config"
	"speed-test/version"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	version.ShowVersionDetect("1.0.0")
	log.Println("client start")
	if len(os.Args) < 2 {
		log.Fatal("usage: client <host:port>")
	}
	addr := os.Args[1]

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	buf := make([]byte, config.BufSize)
	rand.Read(buf) // 初始化随机数据

	var totalWrite int64
	var totalRead int64
	startTime := time.Now()
	endTime := startTime.Add(config.Duration)

	var wg sync.WaitGroup
	wg.Add(2)

	// 上传
	go func() {
		defer wg.Done()
		for time.Now().Before(endTime) {
			conn.SetWriteDeadline(time.Now().Add(config.Timeout))
			n, err := conn.Write(buf)
			if n > 0 {
				atomic.AddInt64(&totalWrite, int64(n))
			}
			if err != nil {
				if ne, ok := err.(net.Error); ok && ne.Timeout() {
					continue
				}
				log.Println("write error:", err)
				break
			}
		}
	}()

	// 下载
	go func() {
		defer wg.Done()
		rbuf := make([]byte, config.BufSize)
		for time.Now().Before(endTime) {
			conn.SetReadDeadline(time.Now().Add(config.Timeout))
			n, err := conn.Read(rbuf)
			if n > 0 {
				atomic.AddInt64(&totalRead, int64(n))
			}
			if err != nil {
				if ne, ok := err.(net.Error); ok && ne.Timeout() {
					continue
				}
				log.Println("read error:", err)
				break
			}
		}
	}()

	wg.Wait() // 等待上传和下载完成
	elapsed := time.Since(startTime).Seconds()

	totalWriteMB := float64(totalWrite) / 1024 / 1024
	totalReadMB := float64(totalRead) / 1024 / 1024

	fmt.Printf("totalWrite=%.2f MB, avgWrite=%.2f MB/s\n", totalWriteMB, totalWriteMB/elapsed)
	fmt.Printf("totalRead=%.2f MB, avgRead=%.2f MB/s\n", totalReadMB, totalReadMB/elapsed)
	fmt.Printf("total time=%.2f s\n", elapsed)
}
