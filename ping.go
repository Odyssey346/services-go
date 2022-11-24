package main

import (
	"log"
	"net/http"
)

func Ping(domain string) int {
	resp, err := http.Get(domain)
	if err != nil {
		log.Fatal(err)
	}
	return resp.StatusCode
}
