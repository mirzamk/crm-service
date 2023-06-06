package admin

import "gorm.io/gorm"

type RouterAdmin struct {
	AdminRequestHandler RequestHandlerAdmin
}

func NewRouter(dbCrud *gorm.DB) RouterAdmin {
	return RouterAdmin{AdminRequestHandler: NewAdminRequestHandler(dbCrud)}
}
