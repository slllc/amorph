package amorph

import (
	"strconv"
	"time"

	"github.com/slllc/bereos/util/mask64"
)

func SetMapEntry(a Amorph, key string, val interface{}) (ok bool) {
	_, ok = a.(map[string]interface{})
	if !ok {
		return
	}
	a.(map[string]interface{})[key] = val
	return true
}

func SetMapEntryMask64(a Amorph, key string, val *mask64.Mask64) (ok bool) {
	valStr := val.String()
	return SetMapEntryString(a, key, valStr)
}

func SetMapEntryString(a Amorph, key string, val string) (ok bool) {
	return SetMapEntry(a, key, val)
}

func SetMapEntryInt64String(a Amorph, key string, val int64) (ok bool) {
	s := strconv.FormatInt(val, 10)
	return SetMapEntry(a, key, s)
}

func SetMapEntryTime(a Amorph, key string, val time.Time) (ok bool) {
	s := val.Format(time.RFC3339)
	if s == "" {
		return false
	}
	return SetMapEntry(a, key, s)
}

func GetMapEntry(a Amorph, key string) (val Amorph, ok bool) {
	val, ok = a.(map[string]interface{})[key]
	return //
}

func GetMapEntryUint(a Amorph, key string) (val uint, ok bool) {
	var valI interface{}
	valI, ok = GetMapEntry(a, key)
	if !ok {
		return
	}
	val, ok = valI.(uint)
	return //
}

func GetMapEntryInt64(a Amorph, key string) (val int64, ok bool) {
	var valI interface{}
	valI, ok = GetMapEntry(a, key)
	if !ok {
		return
	}
	val, ok = valI.(int64)
	return //
}

func GetMapEntryMask64(a Amorph, key string, l uint) (val *mask64.Mask64, ok bool) {
	maskStr, ok := GetMapEntryString(a, key)
	if !ok {
		return //
	}
	val = mask64.NewMask64(maskStr, l)
	return //
}

func GetMapEntryString(a Amorph, key string) (val string, ok bool) {
	var valI interface{}
	valI, ok = GetMapEntry(a, key)
	if !ok {
		return
	}
	val, ok = valI.(string)
	return //
}

func GetMapEntryTime(a Amorph, key string) (val time.Time, ok bool) {
	timeString, ok := GetMapEntryString(a, key)

	valParsed, err := time.Parse(time.RFC3339, timeString)
	if err != nil {
		return //
	}
	val = valParsed
	return //
}
