package model

import (
	"encoding"
	"time"
)

type Modeler interface {
	GetCacheKey() string
	GetCacheDuration() time.Duration
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

func GetAllModels() []interface{} {
	return []interface{}{
		&PacWebSite{},
		&PacContent{},
		&V2rayNode{},
		&Subscription{},
	}
}
