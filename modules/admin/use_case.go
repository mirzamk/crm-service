package admin

import "github.com/mirzamk/crm-service/repository"

type AdminUseCase interface {
}

type adminUseCase struct {
	adminRepo repository.AdminInterfaceRepo
}
