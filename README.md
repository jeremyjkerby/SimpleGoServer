# Simple Go Server

## Overview
This is a simple server that continuously reads data in from some URL containing students and their exam scores. For simplicity, this data is stored in memory and is not persistent. This simple server has a custom built API to retrieve data that is currently in memory that has been read in from URL.

## Server
The server is completely written in Go.

## Mac Setup
Install Go programming language:
```bash
brew install go
```
Install third-party package gorilla/mux:
```bash
go get github.com/gorilla/mux
```

## Usage
Obtain repository:
```bash
git clone git@github.com:jeremyjkerby/SimpleGoServer.git
```
Configure server.go
```bash
L17:    var choice int = [N] // configure, where N is number of records to read in
L18:    var url string = [URL] // configure, where URL is desired destination
```
Compile the code:
```bash
go build server.go controller.go
```
Run the server:
```bash
./server
```

## API Architecture
Once there is a server setup and running you may hit the following API endpoints with cURL, Postman, etc:

Get all students:
```bash
curl -X GET "http://127.0.0.1:9000/api/students"

Example output:
{
    "students": [
        "Tyler.Anderson",
        "Dawson.Lockman",
        "Oleta.Lindgren10",
        ...
    ]
}
```
Get all exam results for given student and the average score:
```bash
curl -X GET "http://127.0.0.1:9000/api/students/{id}"

Example output:
{
    "average": 0.7713795894902912,
    "scores": [
        {
            "Exam": 3398,
            "StudentId": "Tyler.Anderson",
            "Score": 0.7343189587409348
        },
        ...
    ]
}
```
Get all exams:
```bash
curl -X GET "http://127.0.0.1:9000/api/exams"

Example output:
{
    "exams": [
        3399,
        3400,
        3401,
        ...
    ]
}
```
Get all exam results for given exam and the average score:
```bash
curl -X GET "http://127.0.0.1:9000/api/exams/{number}"

Example output:
{
    "average": 0.7346578788593273,
    "scores": [
        {
            "Exam": 3400,
            "StudentId": "Elmo67",
            "Score": 0.6139979704507454
        },
        ...
    ]
}
```

## Unit Tests
To run all unit tests:
```bash
go test -v
```
To generate visualization of unit tests:
```bash
go test -cover -coverprofile=c.out
go tool cover -html=c.out -o coverage.html
```


## Author
Jeremy J. Kerby
