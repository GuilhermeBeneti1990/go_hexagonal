package application

import (
	"errors"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type IProduct interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetId() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

type IProductService interface {
	Get(id string) (IProduct, error)
	Create(name string, price float64) (IProduct, error)
	Enable(product IProduct) (IProduct, error)
	Disable(product IProduct) (IProduct, error)
}

type IProductReader interface {
	Get(id string) (IProduct, error)
}

type IProductWriter interface {
	Save(product IProduct) (IProduct, error)
}

type IProductPersistence interface {
	IProductReader
	IProductWriter
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	ID     string  `valid: "uuidv4"`
	Name   string  `valid: "required"`
	Price  float64 `valid: "float,optional"`
	Status string  `valid: "required"`
}

func NewProduct() *Product {
	product := Product{
		ID:     uuid.NewV4().String(),
		Status: DISABLED,
	}

	return &product
}

func (product *Product) IsValid() (bool, error) {
	if product.Status == "" {
		product.Status = DISABLED
	}

	if product.Status != ENABLED && product.Status != DISABLED {
		return false, errors.New("status must be enabled or disabled")
	}

	if product.Price < 0 {
		return false, errors.New("the price must be greater or equal to 0")
	}

	_, err := govalidator.ValidateStruct(product)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (product *Product) Enable() error {
	if product.Price > 0 {
		product.Status = ENABLED
		return nil
	}

	return errors.New("the price must be greater than 0 to enable the product")
}

func (product *Product) Disable() error {
	if product.Price == 0 {
		product.Status = DISABLED
		return nil
	}

	return errors.New("the price must be zero in order to have the product disabled")
}

func (product *Product) GetId() string {
	return product.ID
}

func (product *Product) GetName() string {
	return product.Name
}

func (product *Product) GetPrice() float64 {
	return product.Price
}

func (product *Product) GetStatus() string {
	return product.Status
}
