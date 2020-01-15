package main

import "fmt"
import "github.com/google/uuid"

const ver = "6.0.0"

func main() {
	id := uuid.New()
	fmt.Printf("Hello Berlinski version for stage v=%s\nid=%s\n", ver, id)
}
