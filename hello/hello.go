package main

import (
	"rsc.io/quote"
	quoteV3 "rsc.io/quote/v3"
)

func hello() string {
	//return "Hello World"
	return quote.Hello()
}

func Proverb() string {
	return quoteV3.Concurrency()
}
