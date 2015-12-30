package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

var ModelStatus *ModelStatusType

type ModelStatusType struct {
	Active   string
	Inactive string
	Deleted  string
	Banned   string
	Valid    string
	Expired  string
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

type DataModelType interface {
	GetDOI() DocumentIdentifier
	Validate() (list []error, ok bool)
	ToMap() map[string]interface{}
	ToJSON() ([]byte, error)
	// UpdateVersion(doi *DocumentIdentifier)
	// FromMap(map[string]interface{}) (*DataModelType, error)
	// FromJSON(json []byte) (*DataModelType, error)
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
	dts := dt.Format(time.RFC3339Nano)

	return dts
}

func FilterModelMap(hash map[string]interface{}) error {
	for k, v := range hash {
		// fmt.Println(k, v)
		switch t := v.(type) {
		case float64, string, bool:
			fmt.Sprintf("%T\n", t)
		case time.Time:
			hash[k] = FormatJSONDate(v.(time.Time))
		default:
			s := fmt.Sprintf("error: %s is a %T\n", k, t)
			return errors.New(s)
		}
	}

	return nil
}

func MapToJSON(v map[string]interface{}) ([]byte, error) {
	if err := FilterModelMap(v); err != nil {
		return nil, err
	}

	return json.Marshal(v)
}

func MapFromJSON(bytes []byte) (map[string]interface{}, error) {
	hash := make(map[string]interface{})
	err := json.Unmarshal(bytes, &hash)

	return hash, err
}
