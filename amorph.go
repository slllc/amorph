// package amorph provides a wrapper for arbitrary hierarchical
// data. The wrapper type provides:
// 	DeepCopy duplicates an Amorph
//	DeepEqual compares two Amorphs
// 	Diff method that produces a Patch.
//
// 	The Patch can be forward and reverse applied to an Amorph.
//
// Walk allows you to visit all the nodes in an Amorph, and provides methods
// for manipulating the Amorph.

package amorph

// Copyright 2021 Charles J. Luciano and Scalability
// Labs LLC. All rights reserved. Use of this source
// code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"encoding/json"
	"io"
	"os"
)

// internal concrete implementation of an Amorp
type Amorph interface{}

// NewAmorphFromReader builds an Amorph from an io.Reader
func NewAmorphFromReader(rdr io.Reader) (amorphOut Amorph, err error) {
	decoder := json.NewDecoder(rdr)
	err = decoder.Decode(&amorphOut)
	return //
}

// NewAmorphFromFile builds an Amorph from a file specified by name
func NewAmorphFromFile(filename string) (amorphOut Amorph, err error) {
	fp, err := os.Open(filename)
	if err != nil {
		return //
	}
	return NewAmorphFromReader(fp)
}

func NewEmptySliceAmorph() (amorphOut Amorph) {
	return make([]interface{}, 0)
}

func NewEmptyMapAmorph() (amorphOut Amorph) {
	return make(map[string]interface{})
}

// NewAmorphFromJSONString builds an Amorph from a json string
func NewAmorphFromJSONString(js string) (amorphOut Amorph, err error) {
	err = json.Unmarshal([]byte(js), &amorphOut)
	return //
}

func OverlayFromJSONString(amorphIn Amorph, js string) (amorphout Amorph, err error) {
	amorphout, err = NewAmorphFromJSONString(js)
	if err != nil {
		return //
	}
	// Test if this actually does the overlay correctly
	err = json.Unmarshal([]byte(js), &amorphout)
	return //
}

// DeepCopy duplicates an Amorph
func DeepCopy(amorphIn Amorph) (amorphOut Amorph) {
	rdr, wrt := io.Pipe()

	go func() {
		encoder := json.NewEncoder(wrt)
		err := encoder.Encode(amorphIn)
		wrt.CloseWithError(err)
	}()
	decoder := json.NewDecoder(rdr)
	err := decoder.Decode(&amorphOut)
	if err != nil {
		panic("Decoder error " + err.Error())
	}
	return //
}
