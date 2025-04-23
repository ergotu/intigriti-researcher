# intigriti-researcher

A Go client library for interacting with the [Intigriti Researcher API](https://kb.intigriti.com/en/articles/8529303-intigriti-researcher-api).

## Installation

`go get github.com/ergotu/intigriti-researcher`

## Usage

```go
package main

import (
	"fmt"
	"log"
	"os"

	intigriti "github.com/ergotu/intigriti-researcher/client"
)

func main() {
	authToken := os.Getenv("INTIGRITI_AUTH_TOKEN")
	if authToken == "" {
		log.Fatal("INTIGRITI_AUTH_TOKEN environment variable not set")
	}

	client := intigriti.New(authToken)

	programs, err := client.GetPrograms(intigriti.GetProgramsOptions{
		Limit: intigriti.Int(5),
	})
	if err != nil {
		log.Fatalf("Error fetching programs: %v", err)
	}

	fmt.Printf("Fetched %d programs:\n", len(programs))
	for _, program := range programs {
		fmt.Printf("- %s (%s)\n", program.Name, program.Handle)
	}
}
```

## Authentication

The client requires an intigriti API authentication token. It's recommended to provide this via an environment variable (e.g., `INTIGRITI_AUTH_TOKEN`) as shown in the usage example.

## Implemented Endpoints

This library currently supports the following intigriti API v1 endpoints:

* `GET /programs` (supports optional filters: statusId, typeId, following, limit, offset)
* `GET /programs/{programId}`
* `GET /programs/{programId}/domains/{versionId}`
* `GET /programs/{programId}/rules-of-engagements/{versionId}`
* `GET /programs/activities` (supports optional filters: createdSince, following, limit, offset)
