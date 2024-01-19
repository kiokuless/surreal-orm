package main

type Record struct {
	RequestedAt string `surreal:"requested_at"`
	RespondedAt string `surreal:"responded_at"`
}
