# YAPS - Yet another port scanner
Yet another port scanner built to give quick results.

## Usage: 

### Detailed Usage: 
```
[+] Welcome to Go-PortScan [+]
[-] Coded by zer0 p1k4chu [-]

  -a    Scan all 65K ports? Default: 1-1024 ports.
  -f int
        Number of Threads to run (default 100)
  -h    Put this flag to print help
  -ip string
        IP Address to be scanned (default "false")
  -lf string
        LogFile location? default: log.txt
  -t string
        Scan Type : TCP or UDP (default "tcp")
  -v    Verbose Output

```

### For the impatient:
```
go run scanner.go  -ip=192.168.0.7 -t=tcp -a=true -f=50 -lf=logg.txt
```
### Issues/Ideas/Just anything else? 

Open an issue. /-.-\ 

