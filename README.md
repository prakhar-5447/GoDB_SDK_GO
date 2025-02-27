# GoDB SDK for Golang

GoDB SDK is a Go client for interacting with the GoDB gRPC database service. It allows easy database operations similar to MongoDB clients.

## Installation

You can install the SDK using `go get`:

```sh
go get github.com/prakhar-5447/GoDB_SDK_GO
```

Alternatively, you can clone the SDK and use a local module replacement in your `go.mod`:

```sh
git clone https://github.com/prakhar-5447/GoDB_SDK_GO.git
```

Modify your `go.mod` to use the local SDK:

```go
require github.com/prakhar-5447/GoDB_SDK_GO v1.0.1

replace github.com/prakhar-5447/GoDB_SDK_GO => ../GoDB_SDK_GO
```
../GoDB_SDK_GO --> This is path to cloned repository

## Usage Example

```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	godb "github.com/prakhar-5447/GoDB_SDK_GO"
)

func main() {
	// Replace with the actual address of your GoDB gRPC service
	address := "localhost:50051"

	// Create a new GoDB client
	client, err := godb.NewGoDBClient(address)
	if err != nil {
		log.Fatalf("Failed to create GoDB client: %v", err)
	}
	defer client.Close()

	// Set up a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Set connection string
	connectionString := "grpc://john:secret123/ecommerceDB"
	client.SetConnectionString(connectionString)

	// Create a table with a defined schema
	columns := map[string]string{
		"id":    "INTEGER PRIMARY KEY",
		"name":  "TEXT",
		"price": "REAL",
	}
	msg, err := client.CreateTable(ctx, "products", columns, connectionString)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
	fmt.Println("CreateTable:", msg)

	// Insert a record
	insertMsg, err := client.Insert(ctx).
		Table("products").
		Values(map[string]string{
			"id":    "1",
			"name":  "Gaming",
			"price": "1500.00",
		}).Exec()
	if err != nil {
		log.Fatalf("Error inserting record: %v", err)
	}
	fmt.Println("InsertRecord:", insertMsg)

	// Query records
	records, err := client.Query(ctx).
		Table("products").
		Columns("name").
		Equal("name", "Gaming").
		Limit(10).
		Offset(0).
		Exec()
	if err != nil {
		log.Fatalf("Error fetching records: %v", err)
	}

	fmt.Println("Fetched Records from 'products' table:", records)
}
```

## License
This project is licensed under the MIT License. See the `LICENSE` file for details.