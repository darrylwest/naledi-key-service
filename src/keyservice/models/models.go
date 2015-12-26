package models

import (
	"time"
	"encoding/json"
	"fmt"
	"errors"
)

var ModelStatus *ModelStatusType

type ModelStatusType struct {
	Active string
	Inactive string
	Deleted string
	Banned string
	Valid string
	Expired string
	Canceled string
}

func init() {
	ModelStatus = new(ModelStatusType)

	ModelStatus.Active = "active"
	ModelStatus.Inactive = "inactive"
	ModelStatus.Deleted = "deleted"
	ModelStatus.Banned = "baned"
	ModelStatus.Valid = "valid"
	ModelStatus.Expired = "expired"
	ModelStatus.Canceled = "canceled"
}

type ChallengeCode struct {
	doi           DocumentIdentifier
	challengeType string // Document, Access
	sendTo        string
	sendDate      time.Time
	expires       time.Time
	status        string // Active, Canceled, Expired
}


// TODO move this to JSON helper delegate
func ParseJSONDate(json map[string]interface{}, node string, dflt time.Time) (time.Time, error) {
	if dts, ok := json[node].(string); ok {
		if dt, err := time.Parse(time.RFC3339Nano, dts); err == nil {
			return dt, nil
		} else {
			return dflt, err
		}
	}

	msg := fmt.Sprintf("node: %s is not a string type", node)

	return dflt, errors.New(msg)
}

func FormatJSONDate(dt time.Time) string {
	dts := dt.Format( time.RFC3339Nano )

	return dts
}

func MapToJSON(v map[string]interface{}) ([]byte, error) {
	for k, v := range v {
		// fmt.Println(k, v)
		switch t := v.(type) {
		case float64, string, bool:
			fmt.Sprintf("%T\n", t )
		default:
			fmt.Printf("error: %s is a %T\n", k, t )
		}
	}

	return json.MarshalIndent( v, "", "  ")
}

func MapFromJSON(bytes []byte) (map[string]interface{}, error) {
	hash := make(map[string]interface{})
	err := json.Unmarshal(bytes, &hash)

	return hash, err
}
