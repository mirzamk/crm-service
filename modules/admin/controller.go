package admin

type ControllerAdmin interface {
}

type controllerAdmin struct {
	adminUseCase AdminUseCase
}
