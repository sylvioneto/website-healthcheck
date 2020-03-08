# Website Health Check

This Go program executes a periodicaly health check in a list of endpoints and log them.

## Usage  

1. Use the config.yaml file to set the list of websites you want to test, the interval and the kind of log you want - console or file.
2. Run the program:  
```bash
$ go run main.go 
2020/03/07 18:46:53 Reading config file...
2020/03/07 18:46:53 Config: {5 [https://golang.org/ https://www.google.com/ https://www.youtube.com/ https://www.google.ca/fakepage] true}
2020/03/07 18:46:53 main.go:47: website-healthcheck: Monitoring...
2020/03/07 18:46:54 main.go:62: website-healthcheck: https://golang.org/ 200
2020/03/07 18:46:54 main.go:62: website-healthcheck: https://www.google.com/ 200
2020/03/07 18:46:54 main.go:62: website-healthcheck: https://www.youtube.com/ 200
2020/03/07 18:46:54 main.go:62: website-healthcheck: https://www.google.ca/fakepage 404
2020/03/07 18:47:00 main.go:62: website-healthcheck: https://golang.org/ 200
```

## Log to file
In case you set `logConsole: false`, the app will write the logs in a file named `website-healthcheck.log`.
```
$ tail -f website-healthcheck.log 
2020/03/07 18:51:12 main.go:63: website-healthcheck: https://golang.org/ 200
2020/03/07 18:51:12 main.go:63: website-healthcheck: https://www.google.com/ 200       
2020/03/07 18:51:12 main.go:63: website-healthcheck: https://www.youtube.com/ 200      
2020/03/07 18:51:12 main.go:63: website-healthcheck: https://www.google.ca/fakepage 404
2020/03/07 18:51:17 main.go:63: website-healthcheck: https://golang.org/ 200
```
