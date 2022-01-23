# AcmeShield

![alt text](ExampleImage.png "Title")
A secured HTTP proxy that forwards requests to a remote service(Postman). This is a tool (or more) that will proxy simple HTTPS requests to an external HTTP endpoint (service in the image above)

## Description

The HTTP proxy operates between the sending Web server and your receiving Web client. It processes the HTTP protocol line-by-line for any potentially harmful content before sending it to an internal Web client. It also acts as a buffer between your Web server and potentially harmful Web clients by enforcing HTTP RFC compliance and preventing potential buffer overflow.
## Getting Started

### Dependencies

* GoLang Standard Library

### Installing

* Make sure to have GoLang installed in your enviroment.
* Make sure to have Git installed on your enviroment.
* Run the following commad to install the program:
```
go install github.com/tavikano/acmeshield@latest
```
* This downloads acmeshield and all of its dependencies, builds the program, and installs the binary in your $GOPATH/bin directory.
* Run the application on your terminal by running: proxy

### Executing program

1. Execute program
2. Using BASH/ZSHELL on a Terminal/Powershell app, run the following command:
```
curl ​http://localhost:8080/test
```

## Help

Port Configuration
```
To Configure the listening port, update line 11 of the acmeshield.go file to the desired port number.
```

## Authors

Octavio Cano

## Version History

* v1.0
    * Initial Release
