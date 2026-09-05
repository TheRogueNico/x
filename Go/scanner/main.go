package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

func main() {
	host := flag.String("host", "localhost", "target host")
	startPort := flag.Int("start", 1, "start port")
	endPort := flag.Int("end", 1024, "end port")
	timeout := flag.Duration("timeout", 500*time.Millisecond, "dial timeout")
	workers := flag.Int("workers", 100, "concurrent workers")
	flag.Parse()

	ports := make(chan int, *workers)
	var wg sync.WaitGroup

	for range *workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for port := range ports {
				address := net.JoinHostPort(*host, strconv.Itoa(port))
				conn, err := net.DialTimeout("tcp", address, *timeout)
				if err != nil {
					continue
				}
				conn.Close()
				fmt.Printf("Port %d: open\n", port)
			}
		}()
	}

	for p := *startPort; p <= *endPort; p++ {
		ports <- p
	}
	close(ports)

	wg.Wait()
}
