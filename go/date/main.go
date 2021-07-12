package main

import (
	"fmt"
	"time"
)

// In today's episode, we'll be playing around with Dates!

func main() {
	// When's now?
	now := time.Now()
	fmt.Println("Now: ", now)

	// How about tomorrow?
	tomorrow := now.Add(time.Hour * 24)
	fmt.Println("Tomorrow:", tomorrow)

	// Can we format the date into ISO 8601?
	iso8601 := "2006-01-02T15:04:05-0700"
	format := now.Format(iso8601)
	fmt.Println("ISO8601:", format)

	// Can we print RFC3339 date format?
	rfc3339 := now.Format(time.RFC3339)
	fmt.Println("RFC3339:", rfc3339)

	// How about another RFC822 date format?
	rfc822 := now.Format(time.RFC822)
	fmt.Println("RFC822:", rfc822)
}
