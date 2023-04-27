package service

import (
	"errors"
	"fmt"

	"github.com/kenykendf/go-restful/internal/app/model"
	"github.com/kenykendf/go-restful/internal/app/repository"
	"github.com/kenykendf/go-restful/internal/app/schema"

	log "github.com/sirupsen/logrus"
)

type ProductService struct {
	repo repository.IProductRepo
}

func NewProductService(repo repository.IProductRepo) *ProductService {
	return &ProductService{repo: repo}
}

func (ps *ProductService) Create(req schema.CreateProductReq) error {
	var insertData model.Product

	insertData.Name = req.Name
	insertData.Description = req.Description
	insertData.Currency = req.Currency
	insertData.Price = req.Price
	insertData.TotalStock = req.TotalStock
	insertData.IsActive = req.IsActive
	insertData.CategoryID = req.CategoryID

	err := ps.repo.Create(insertData)
	if err != nil {
		log.Error(fmt.Errorf("error ProductService - Create : %w", err))
		return errors.New("service unable create new product")
	}
	return nil
}

func (ps *ProductService) BrowseAll() ([]schema.GetProductResp, error) {
	var resp []schema.GetProductResp

	products, err := ps.repo.Browse()
	if err != nil {
		return nil, errors.New("server error, unable to fetch products")
	}

	for _, value := range products {
		var respData schema.GetProductResp
		respData.ID = value.ID
		respData.Name = value.Name
		respData.Description = value.Description
		respData.Currency = value.Currency
		respData.Price = value.Price
		respData.TotalStock = value.TotalStock
		respData.CategoryID = value.CategoryID
		resp = append(resp, respData)
	}

	return resp, nil
}
