package models

import (
	"time"
    "strings"
    "code.google.com/p/go-uuid/uuid"
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

func NewDocumentIdentifier() *DocumentIdentifier {
    doi := new(DocumentIdentifier)

    doi.id = NewModelId()
    doi.dateCreated = time.Now().UTC()
    doi.lastUpdated = time.Now().UTC()
    doi.version = 0

    return doi
}

func (doi *DocumentIdentifier) GetId() string {
    return doi.id
}

func (doi *DocumentIdentifier) GetDateCreated() time.Time {
    return doi.dateCreated
}

func (doi *DocumentIdentifier) GetLastUpdated() time.Time {
    return doi.lastUpdated
}

func (doi *DocumentIdentifier) GetVersion() int64 {
    return doi.version
}

func (doi *DocumentIdentifier) updateVersion() int64 {
    doi.version++
    doi.lastUpdated = time.Now().UTC()

    return doi.version
}

func (doi *DocumentIdentifier) ToMap() map[string]interface{} {
    var mp = map[string]interface{} {
        "id": doi.id,
        "dateCreated": doi.dateCreated,
        "lastUpdated": doi.lastUpdated,
        "version": doi.version,
    }

    return mp
}

func (doi *DocumentIdentifier) FromMap(v map[string]interface{}) error {
    doi.id = v["id"].(string)

    if dt, ok := v["dateCreated"].(time.Time); ok {
        doi.dateCreated = dt
    }

    if dt, ok := v["lastUpdated"].(time.Time); ok {
        doi.lastUpdated = dt
    }

    doi.version = v["version"].(int64)

    return nil
}
