package resource

import (
	"time"

	"github.com/google/uuid"
)

type Resource struct{
	Uuid uuid.UUID
	Versions []*resourceVersion
	ResourceType string
}

type resourceVersion struct {
	VersionId string
	CreatedAt time.Time
	Data resourceData
}

type resourceData struct {
	Name *string
	Description *string
	Notes *string
}

func (r *Resource) CreateCleanVersion () (*resourceVersion, error) {
	 return nil, nil
}

func (r *Resource) CreateVersion () (*resourceVersion, error) {
	return nil, nil
}