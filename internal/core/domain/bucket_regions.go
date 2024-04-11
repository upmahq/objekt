package domain

import (
	"bytes"
	"encoding/json"
	"errors"
)

type BucketRegion int

const (
	InvalidRegion BucketRegion = iota
	Ashburn
	Frankfurt
	London
	Phoenix
	Singapore
	Sydney
)

func (br BucketRegion) String() string {
	return regionToString[br]
}

var regionToString = map[BucketRegion]string{
	Ashburn:   "ashburn",
	Frankfurt: "frankfurt",
	London:    "london",
	Phoenix:   "phoenix",
	Singapore: "singapore",
	Sydney:    "sydney",
}

var regionToID = map[string]BucketRegion{
	"ashburn":   Ashburn,
	"frankfurt": Frankfurt,
	"london":    London,
	"phoenix":   Phoenix,
	"singapore": Singapore,
	"sydney":    Sydney,
}

// MarshalJSON marshals the enum as a quoted json string
func (br BucketRegion) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(regionToString[br])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (br *BucketRegion) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'AWS' in this case.
	*br = regionToID[j]
	return nil
}

func (br BucketRegion) Value() (string, error) {
	if br == InvalidRegion {
		return "", errors.New("invalid bucket region")
	}
	return regionToString[br], nil
}

func (br *BucketRegion) Scan(value any) error {
	switch value := value.(type) {
	case string:
		*br = regionToID[value]
	case []byte:
		*br = regionToID[string(value)]
	case int:
		*br = BucketRegion(value)
	default:
		return errors.New("incompatible type for BucketRegion")
	}
	return nil
}
