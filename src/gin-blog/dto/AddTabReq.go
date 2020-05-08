package dto

import (
	"time"
)

type AddTagEntity struct {
	Name        string    `json:"name" form:"name" validate:"required"`
	CreatedBy   string    `json:"create_by" form:"create_by"`
	ModifiedBy  string    `json:"modified_by" form:"modified_by"`
	Summary     string    `json:"summary" form:"summary"`
	State       int       `json:"state" form:"state" validate:"required,max=5,min=1"`
	DeletedTime *time.Time `json:"deleted_time" form:"deleted_time"`
}

type AddTagEntities struct {
	Params []AddTagEntity  `json:"params" form:"params" validate:"required"`
}
