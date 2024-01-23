package id

import (
	"testing"
)

func TestGetID(t *testing.T) {
	jsonBytes := []byte(`
	{
		"id": {
	    	"Thing": {
		       "id": {
			       "String": "41hkwf1qnr4925w2iqg4"
		       },
	           "tb": "record"
	       }
		}
	}`)
	r, err := GetID(jsonBytes)
	if err != nil {
		t.Fatal(err)
	}
	if r.ID != "41hkwf1qnr4925w2iqg4" {
		t.Fatalf("expected 41hkwf1qnr4925w2iqg4, got %s", r.ID)
	}
	if r.TB != "record" {
		t.Fatalf("expected record, got %s", r.TB)
	}
}
