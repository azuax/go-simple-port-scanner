package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

func isOpen(host string, port int) bool {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), time.Millisecond*500)

	if err != nil {
		return false
	}

	conn.Close()
	return true
}

func printStatus(p int, status bool) {
	if status {
		fmt.Printf("%d - open\n", p)
	} else {
		fmt.Printf("%d - closed\n", p)
	}
}

func scan(host string, portList []int) {
	for _, p := range portList {
		open := isOpen(host, p)

		printStatus(p, open)
	}
}

func scanGoroutines(nGoroutines int, host string, portList []int) {
	portChan := make(chan int, nGoroutines)
	wg := new(sync.WaitGroup)

	for _, p := range portList {
		portChan <- p
		wg.Add(1)
		//we use a function literal to create goroutine
		go func(h string, pChan chan int, WG *sync.WaitGroup) {
			defer WG.Done()
			p := <-pChan
			printStatus(p, isOpen(h, p))
		}(host, portChan, wg)
	}

	wg.Wait()
}

func main() {
	nGoroutines := 20
	host := "192.168.0.1"

	// nmap top 100, got with: awk '$2~/tcp$/' /usr/share/nmap/nmap-services | sort -r -k3 | head -n 100 | awk -F '/tcp' '{print $1}' | awk '{print $2}' | tr '\n' ','
	portList := []int{80, 23, 443, 21, 22, 25, 3389, 110, 445, 139, 143, 53, 135, 3306, 8080, 1723, 111, 995, 993, 5900, 1025, 587, 8888, 199, 1720, 465, 548, 113, 81, 6001, 10000, 514, 5060, 179, 1026, 2000, 8443, 8000, 32768, 554, 26, 1433, 49152, 2001, 515, 8008, 49154, 1027, 5666, 646, 5000, 5631, 631, 49153, 8081, 2049, 88, 79, 5800, 106, 2121, 1110, 49155, 6000, 513, 990, 5357, 427, 49156, 543, 544, 5101, 144, 7, 389, 8009, 3128, 444, 9999, 5009, 7070, 5190, 3000, 5432, 1900, 3986, 13, 1029, 9, 5051, 6646, 49157, 1028, 873, 1755, 2717, 4899, 9100, 119, 37}

	// If we don't pass any argument, we execute the linear scan first
	if len(os.Args) < 2 {
		// linear scan
		start1 := time.Now()
		scan(host, portList)
		elapsed1 := time.Since(start1)
		fmt.Printf("\nTotal time linear scan: %s\n\n", elapsed1)

		fmt.Println("****************************")
	}

	// goroutine scan
	start2 := time.Now()
	scanGoroutines(nGoroutines, host, portList)
	elapsed2 := time.Since(start2)
	fmt.Printf("\nTotal time with %d goroutines: %s\n", nGoroutines, elapsed2)
}
