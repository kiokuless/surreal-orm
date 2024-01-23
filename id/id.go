package id

import (
	"encoding/json"

	"errors"
)

// {
//     "id": {
// 	       "Thing": {
// 		       "id": {
// 			       "String": "41hkwf1qnr4925w2iqg4"
// 		       },
// 	           "tb": "record"
// 	       }
//     }
// }

type idJSON struct {
	Inner innerIDJSON `json:"id"`
}

type innerIDJSON struct {
	Inner innerInnerIDJSON `json:"Thing"`
}

type innerInnerIDJSON struct {
	ID idStringJSON `json:"id"`
	TB string       `json:"tb"`
}

type idStringJSON struct {
	String string `json:"String"`
}

type IDThing struct {
	ID string
	TB string
}

func GetID(jsonBytes []byte) (*IDThing, error) {
	var root idJSON
	err := json.Unmarshal(jsonBytes, &root)
	if err != nil {
		return nil, errors.Join(errors.New("failed to unmarshal json"), err)
	}

	r := IDThing{}
	r.ID = root.Inner.Inner.ID.String
	r.TB = root.Inner.Inner.TB

	return &r, nil
}
