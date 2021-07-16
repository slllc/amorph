package amorph

import "encoding/json"

func JSONEncode(a Amorph) (string, error) {
	enc, err := json.Marshal(a)
	return string(enc), err
}
