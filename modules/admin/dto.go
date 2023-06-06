package admin

import (
	"github.com/mirzamk/crm-service/dto"
	"github.com/mirzamk/crm-service/entity"
)

type ParamAdmin struct {
}

type SuccessCreate struct {
	dto.ResponseMeta
	Data ParamAdmin `json:"data"`
}

type FindAdmin struct {
	dto.ResponseMeta
	Data entity.Admin `json:"data"`
}
