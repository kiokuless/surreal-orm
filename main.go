package main

import (
	"encoding/json"
	"fmt"
)

type RecordJSON struct {
	RequestedAt StrandJSON `json:"requested_at"`
	RespondedAt StrandJSON `json:"responded_at"`
}

type StrandJSON struct {
	Strand string `json:"Strand"`
}

type Record struct {
	RequestedAt string
	RespondedAt string
}

// RecordJSON to Record
func (r *Record) UnmarshalJSON(data []byte) error {
	var RecordJSON RecordJSON
	if err := json.Unmarshal(data, &RecordJSON); err != nil {
		return err
	}
	r.RequestedAt = RecordJSON.RequestedAt.Strand
	r.RespondedAt = RecordJSON.RespondedAt.Strand
	return nil
}

func main() {
	buf := []byte(`{"requested_at":{"Strand":"2019-01-01T00:00:00Z"},"responded_at":{"Strand":"2019-01-01T00:00:00Z"}}`)
	var Record Record
	if err := Record.UnmarshalJSON(buf); err != nil {
		fmt.Println(err)
	}
	fmt.Println(Record)
}
