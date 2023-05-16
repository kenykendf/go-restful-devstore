package service

import (
	"errors"
	"fmt"
	"mime/multipart"
	"strconv"

	"github.com/kenykendf/go-restful/internal/app/model"
	"github.com/kenykendf/go-restful/internal/app/repository"
	"github.com/kenykendf/go-restful/internal/app/schema"
	"github.com/kenykendf/go-restful/internal/pkg/reason"

	log "github.com/sirupsen/logrus"
)

type ImageUploader interface {
	UploadImage(input *multipart.FileHeader) (string, error)
}

type ProductService struct {
	productRepo  repository.IProductRepo
	categoryRepo repository.ICategoryRepo
	uploader     ImageUploader
}

func NewProductService(productRepo repository.IProductRepo, categoryRepo repository.ICategoryRepo, uploader ImageUploader) *ProductService {
	return &ProductService{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
		uploader:     uploader,
	}
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

	categoryID := strconv.Itoa(req.CategoryID)
	_, err := ps.categoryRepo.Detail(categoryID)
	if err != nil {
		return errors.New(reason.ProductCannotCreate)
	}

	// upload file to cloudinary
	imageURL, err := ps.uploader.UploadImage(req.Image)
	if err != nil {
		return errors.New(reason.ProductCannotCreate)
	}

	insertData.ImageURL = &imageURL

	err = ps.productRepo.Create(insertData)
	if err != nil {
		log.Error(fmt.Errorf("error ProductService - Create : %w", err))
		return errors.New(reason.ProductCannotCreate)
	}

	return nil
}

func (ps *ProductService) BrowseAll() ([]schema.GetProductResp, error) {
	var resp []schema.GetProductResp

	products, err := ps.productRepo.Browse()
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
