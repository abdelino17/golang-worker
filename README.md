# Golang Worker

This repository contains a sample Golang worker application.

## Overview

The Golang worker application is designed to perform background tasks. It simulates processing multiple jobs concurrently use channels and goroutines.

## Getting Started

### Prerequisites

- Go 1.16 or later

### Installation

Clone the repository:

```sh
git clone https://github.com/abdelino17/golang-worker.git
cd golang-worker
```

### Running the Application

To run the application, use the following command:

```bash
$ go run main.go
```

### main.go

The `main.go` file contains the main logic for the worker application. Below is a brief overview of its contents:

```go
package main

import (
  "fmt"
  "time"
)

func main() {
  fmt.Println("Worker started")
  for {
    // Simulate processing a task
    fmt.Println("Processing task")
    time.Sleep(2 * time.Second)
  }
}
```

### Docker Image

Build the docker image using the following command:

```bash
$ docker build -t golang-worker:latest .
```
