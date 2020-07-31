package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
)

type Report struct {
	Port   string
	Status string
	Type   string
}

type ToDo struct {
	Port string
	Type string
}

func worker(Address string, ToDo <-chan ToDo, Done chan<- Report, wg *sync.WaitGroup, f *os.File, Verbose bool) {
	for Job := range ToDo {
		_, err := net.Dial(Job.Type, Address+":"+Job.Port)
		if err == nil {
			Done <- Report{Port: Job.Port, Status: "open", Type: Job.Type}
			fmt.Println("Port " + Job.Port + " is open")
			f.WriteString("Port " + Job.Port + " is open \n")
		} else {
			Done <- Report{Port: Job.Port, Status: "closed", Type: Job.Type}

			if Verbose == true {
				fmt.Println("Port " + Job.Port + " is closed")
				f.WriteString("Port " + Job.Port + " is closed \n")
			}
		}
		wg.Done()
	}
}

func main() {
	fmt.Println("[+] Welcome to Go-PortScan [+]")
	fmt.Println("[-] Coded by zer0 p1k4chu [-] \n")
	Address := flag.String("ip", "false", "IP Address to be scanned")
	help := flag.Bool("h", false, "Put this flag to print help")
	ScanType := flag.String("t", "tcp", "Scan Type : TCP or UDP")
	Full := flag.Bool("a", false, "Scan all 65K ports? Default: 1-1024 ports.")
	Log := flag.String("lf", "", "LogFile location? default: log.txt")
	Verbose := flag.Bool("v", false, "Verbose Output")
	Threads := flag.Int("f", 100, "Number of Threads to run")
	flag.Parse()
	if *Address == "false" {
		flag.PrintDefaults()
		return
	}
	if *help {
		flag.PrintDefaults()
		return
	}
	// LogFile := Log //Logfile location
	// fmt.Println(*Log)
	f, _ := os.Create(*Log)
	var Result = make(chan Report, 1024)
	var Init = make(chan ToDo, 1024)
	var wg sync.WaitGroup
	var Limit int
	// fmt.Println(*Full)
	if *Full == false {
		Limit = 1024
	} else {
		Limit = 65355
	}
	for i := 1; i < *Threads; i++ {
		go worker(*Address, Init, Result, &wg, f, *Verbose)
	}
	for i := 0; i < Limit; i++ {
		Init <- ToDo{Port: strconv.Itoa(i), Type: *ScanType}
		wg.Add(1)
	}
	close(Init)
	wg.Wait()
	f.Sync()
	f.Close()
}
