# Website Health Check

This Go program executes a periodicaly health check in a list of endpoints and log them.

## Usage  

1. Set the list of websites in the config.txt as the example  
2. Run the program:  
```bash
go run main.go
```
3. Choose option 1 to start the health check.  
4. Follow up the log file:  
```
$ tail -f website-healthcheck.log
2020/03/07 17:18:23 main.go:67: Monitoring...
2020/03/07 17:18:23 main.go:89: Reading config.txt file...
2020/03/07 17:18:24 main.go:83: https://golang.org/ 200
2020/03/07 17:18:24 main.go:83: https://www.google.com/ 200
2020/03/07 17:18:24 main.go:83: https://www.youtube.com/ 200
2020/03/07 17:18:25 main.go:83: https://www.google.ca/fakepage 404
2020/03/07 17:18:30 main.go:83: https://golang.org/ 200
2020/03/07 17:18:30 main.go:83: https://www.google.com/ 200
```
