package main

import "fmt"
import "github.com/google/uuid"
import "github.com/dgrijalva/jwt-go"

const ver = "7.0.0"

func main() {
	id := uuid.New()
	fmt.Printf("Hello Berlinski version for stage v=%s\nid=%s\n", ver, id)
	ExampleNewWithClaims_standardClaims()
}

func ExampleNewWithClaims_standardClaims() {
	mySigningKey := []byte("AllYourBase")

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: 15000,
		Issuer:    "test",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", ss, err)
	//Output: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MDAwLCJpc3MiOiJ0ZXN0In0.QsODzZu3lUZMVdhbO76u3Jv02iYCvEHcYVUI1kOWEU0 <nil>
}
