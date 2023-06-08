package customer

import (
	"context"
	"fmt"
	"github.com/mirzamk/crm-service/constant"
	"github.com/mirzamk/crm-service/entity"
	"github.com/mirzamk/crm-service/payload"
	"github.com/mirzamk/crm-service/repository"
	"github.com/mirzamk/crm-service/utils/helper"
	"strconv"
)

type useCaseCustomer struct {
	ActorRepo    repository.ActorInterfaceRepo
	CustomerRepo repository.CustomerInterfaceRepository
}

type UseCaseCustomer interface {
	CreateCustomer(ctx context.Context, customer CustomerDto) error
	GetCustomerById(ctx context.Context, id int) (CustomerDto, error)
	SearchCustomer(ctx context.Context, filter map[string]string) (*helper.Pagination, error)
	UpdateCustomer(ctx context.Context, customer payload.UpdateCustomer, id int) error
	DeleteCustomer(ctx context.Context, id int) error
}

func (uc *useCaseCustomer) CreateCustomer(ctx context.Context, customer CustomerDto) error {
	var err error
	user, err := uc.ActorRepo.GetActorById(ctx.Value("adminId").(uint))
	if err != nil {
		return constant.ErrAdminNotFound
	}
	if user.IsActive != constant.True {
		return constant.ErrAdminNotActive
	}
	customerSave := entity.Customer{
		Firstname: customer.Firstname,
		Lastname:  customer.Lastname,
		Email:     customer.Email,
		Avatar:    customer.Avatar,
	}
	err = uc.CustomerRepo.CreateCustomer(customerSave)
	if err != nil {
		return err
	}
	return nil
}
func (uc *useCaseCustomer) GetCustomerById(ctx context.Context, id int) (CustomerDto, error) {
	var err error
	user, err := uc.ActorRepo.GetActorById(ctx.Value("adminId").(uint))
	if err != nil {
		return CustomerDto{}, constant.ErrAdminNotFound
	}
	if user.IsActive != constant.True {
		return CustomerDto{}, constant.ErrAdminNotActive
	}
	get, err := uc.CustomerRepo.GetCustomerById(uint(id))
	if err != nil {
		return CustomerDto{}, constant.ErrCustomerNotFound
	}
	getCust := CustomerDto{
		Firstname: get.Firstname,
		Lastname:  get.Lastname,
		Email:     get.Email,
		Avatar:    get.Avatar,
	}
	return getCust, nil
}
func (uc *useCaseCustomer) SearchCustomer(ctx context.Context, filter map[string]string) (*helper.Pagination, error) {
	var err error
	user, err := uc.ActorRepo.GetActorById(ctx.Value("adminId").(uint))
	if err != nil {
		return nil, constant.ErrAdminNotFound
	}
	if user.IsActive != constant.True {
		return nil, constant.ErrAdminNotActive
	}
	var customers *helper.Pagination
	var totalRows int64

	page, err := strconv.Atoi(filter["page"])
	if err != nil {
		return &helper.Pagination{}, err
	}
	limit, err := strconv.Atoi(filter["limit"])
	if err != nil {
		return &helper.Pagination{}, err
	}
	err = uc.CustomerRepo.CountRowCustomer(&totalRows)
	if err != nil {
		return &helper.Pagination{}, err
	}
	pagination := helper.Pagination{
		Limit:     limit,
		Page:      page,
		Sort:      fmt.Sprintf("%s %s", filter["sortby"], filter["orderby"]),
		TotalRows: totalRows,
	}
	if totalRows == 0 {
		initData, err := helper.DataCustomerInit()
		if err != nil {
			return nil, err
		}
		for _, data := range initData {
			var tmp = entity.Customer{
				Firstname: data.FirstName,
				Lastname:  data.LastName,
				Avatar:    data.Avatar,
				Email:     data.Email,
			}
			err := uc.CustomerRepo.CreateCustomer(tmp)
			if err != nil {
				return nil, err
			}
		}
	}
	if filter["name"] != "" && filter["email"] == "" {
		customers, err = uc.CustomerRepo.SearchCustomerByName(pagination, filter["name"])
		if err != nil {
			return &helper.Pagination{}, err
		}
	} else if filter["email"] != "" && filter["name"] == "" {
		customers, err = uc.CustomerRepo.SearchCustomerByEmail(pagination, filter["email"])
		if err != nil {
			return &helper.Pagination{}, err
		}
	} else if filter["name"] != "" && filter["email"] != "" {
		customers, err = uc.CustomerRepo.SearchCustomerByNameOrEmail(pagination, filter["name"], filter["email"])
		if err != nil {
			return &helper.Pagination{}, err
		}
	} else {
		customers, err = uc.CustomerRepo.GetAllCustomer(pagination)
		if err != nil {
			return &helper.Pagination{}, err
		}
	}
	var customer []CustomerDto
	data := customers.Rows.([]*entity.Customer)
	for _, item := range data {
		var cust = CustomerDto{
			Firstname: item.Firstname,
			Lastname:  item.Lastname,
			Avatar:    item.Avatar,
			Email:     item.Email,
		}
		customer = append(customer, cust)
	}
	customers.Rows = customer
	return customers, nil
}
func (uc *useCaseCustomer) UpdateCustomer(ctx context.Context, customer payload.UpdateCustomer, id int) error {
	user, err := uc.ActorRepo.GetActorById(ctx.Value("adminId").(uint))
	if err != nil {
		return constant.ErrAdminNotFound
	}
	if user.IsActive != constant.True {
		return constant.ErrAdminNotActive
	}
	_, err = uc.CustomerRepo.GetCustomerById(uint(id))
	if err != nil {
		return constant.ErrCustomerNotFound
	}
	customerUpdate := entity.Customer{
		Firstname: customer.Firstname,
		Lastname:  customer.Lastname,
		Email:     customer.Email,
		Avatar:    customer.Avatar,
	}
	err = uc.CustomerRepo.UpdateCustomer(customerUpdate, uint(id))
	if err != nil {
		return err
	}
	return nil
}
func (uc *useCaseCustomer) DeleteCustomer(ctx context.Context, id int) error {
	user, err := uc.ActorRepo.GetActorById(ctx.Value("adminId").(uint))
	if err != nil {
		return constant.ErrAdminNotFound
	}
	if user.IsActive != constant.True {
		return constant.ErrAdminNotActive
	}
	_, err = uc.CustomerRepo.GetCustomerById(uint(id))
	if err != nil {
		return constant.ErrCustomerNotFound
	}
	err = uc.CustomerRepo.DeleteCustomer(uint(id))
	if err != nil {
		return err
	}
	return nil
}
