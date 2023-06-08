package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/mirzamk/crm-service/repository"
	"gorm.io/gorm"
)

type CustomerRoute struct {
	CustomerHandler CustomerRequestHandler
}

func NewRouter(db *gorm.DB, auth middleware.AuthorizationInterface, validation middleware.ValidationInterface) CustomerRoute {
	return CustomerRoute{
		CustomerHandler: &customerRequestHandler{
			CustomerController: &customerController{
				CustomerUseCase: &useCaseCustomer{
					CustomerRepo: repository.New(db),
					ActorRepo:    repository.ActorNewRepo(db),
				},
			},
			Validation: validation,
			Auth:       auth,
		},
	}
}

func (cr *CustomerRoute) Handle(router *gin.Engine) {
	customerPath := "/customer"
	customerRG := router.Group(customerPath)
	customerRG.POST("", cr.CustomerHandler.CreateCustomer)
	customerRG.GET("/:id", cr.CustomerHandler.GetCustomerById)
	customerRG.GET("/search", cr.CustomerHandler.SearchCustomers)
	customerRG.PUT("/:id", cr.CustomerHandler.UpdateCustomer)
	customerRG.DELETE("/:id", cr.CustomerHandler.DeleteCustomer)
}
