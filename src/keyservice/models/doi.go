package models

import (
	"code.google.com/p/go-uuid/uuid"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type DocumentIdentifier struct {
	id          string
	dateCreated time.Time
	lastUpdated time.Time
	version     int64
}

func NewModelId() string {
	parts := strings.Split(uuid.New(), "-")
	return strings.Join(parts, "")
}

func NewDocumentIdentifier() DocumentIdentifier {
	doi := DocumentIdentifier{
		id:          NewModelId(),
		dateCreated: time.Now().UTC(),
		lastUpdated: time.Now().UTC(),
		version:     int64(0),
	}

	return doi
}

func (doi DocumentIdentifier) GetId() string {
	return doi.id
}

func (doi DocumentIdentifier) GetDateCreated() time.Time {
	return doi.dateCreated
}

func (doi DocumentIdentifier) GetLastUpdated() time.Time {
	return doi.lastUpdated
}

func (doi DocumentIdentifier) GetVersion() int64 {
	return doi.version
}

func (doi DocumentIdentifier) String() string {
	return fmt.Sprintf("%s, %v, %v, %d", doi.id, doi.dateCreated, doi.lastUpdated, doi.version)
}

func (doi *DocumentIdentifier) UpdateVersion() int64 {
	if doi.dateCreated.Year() == 1 {
		doi.dateCreated = time.Now().UTC()
	}

	doi.version++
	doi.lastUpdated = time.Now().UTC()

	return doi.version
}

func (doi *DocumentIdentifier) ToMap() map[string]interface{} {
	var mp = map[string]interface{}{
		"id": doi.id,
	}

	mp["dateCreated"] = FormatJSONDate(doi.dateCreated)
	mp["lastUpdated"] = FormatJSONDate(doi.lastUpdated)
	mp["version"] = float64(doi.version)

	return mp
}

func (doi *DocumentIdentifier) FromMap(v map[string]interface{}) error {
	doi.id = v["id"].(string)

	dflt := time.Now().UTC()

	if dt, err := ParseJSONDate(v, "dateCreated", dflt); err == nil {
		doi.dateCreated = dt
	}

	if dt, err := ParseJSONDate(v, "lastUpdated", dflt); err == nil {
		doi.lastUpdated = dt
	}

	version := v["version"]

	switch version.(type) {
	case float64:
		doi.version = int64(version.(float64))
	case int64:
		doi.version = version.(int64)
	case string:
		if val, err := strconv.ParseInt(version.(string), 10, 64); err == nil {
			doi.version = val
		}
	}

	return nil
}
