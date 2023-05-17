package service

import (
	"errors"
	"mime/multipart"
	"strconv"

	"github.com/kenykendf/go-restful/internal/app/model"
	"github.com/kenykendf/go-restful/internal/app/schema"
	"github.com/kenykendf/go-restful/internal/pkg/reason"

	log "github.com/sirupsen/logrus"
)

type ImageUploader interface {
	UploadImage(input *multipart.FileHeader) (string, error)
}

type ProductService struct {
	productRepo  ProductRepository
	categoryRepo CategoryRepository
	uploader     ImageUploader
}

func NewProductService(productRepo ProductRepository, categoryRepo CategoryRepository, uploader ImageUploader) *ProductService {
	return &ProductService{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
		uploader:     uploader,
	}
}

func (ps *ProductService) Create(req *schema.CreateProductReq) error {
	var insertData model.Product

	insertData.Name = req.Name
	insertData.Description = req.Description
	insertData.Currency = req.Currency
	insertData.TotalStock = req.TotalStock
	insertData.IsActive = req.IsActive
	insertData.CategoryID = req.CategoryID

	categoryID := strconv.Itoa(req.CategoryID)
	_, err := ps.categoryRepo.Detail(categoryID)
	if err != nil {
		return errors.New(reason.CategoryNotFound)
	}

	// upload file to cloudinary
	imageURL, err := ps.uploader.UploadImage(req.Image)
	if err != nil {
		log.Error("upload image product : %w", err)
		return errors.New(reason.ProductCannotCreate)
	}

	insertData.ImageURL = &imageURL

	// Return productID when create product
	err = ps.productRepo.Create(insertData)
	if err != nil {
		return errors.New(reason.ProductCannotCreate)
	}

	return nil
}

func (ps *ProductService) BrowseAll(req *schema.BrowseProductReq) ([]schema.BrowseProductResp, error) {
	var resp []schema.BrowseProductResp

	dbSearch := model.BrowseProduct{}
	dbSearch.Page = req.Page
	dbSearch.PageSize = req.PageSize

	products, err := ps.productRepo.Browse(dbSearch)
	if err != nil {
		return nil, errors.New(reason.ProductCannotBrowse)
	}

	for _, value := range products {
		respData := schema.BrowseProductResp{
			ID:          value.ID,
			Name:        value.Name,
			Description: value.Description,
			Currency:    value.Currency,
			TotalStock:  value.TotalStock,
			IsActive:    value.IsActive,
			ImageURL:    value.ImageURL,
		}

		resp = append(resp, respData)
	}

	return resp, nil
}

// get detail product
func (ps *ProductService) Detail(id string) (schema.DetailProductResp, error) {
	var resp schema.DetailProductResp

	product, err := ps.productRepo.Detail(id)
	if err != nil {
		return resp, errors.New(reason.ProductCannotGetDetail)
	}

	categoryID := strconv.Itoa(product.CategoryID)
	category, err := ps.categoryRepo.Detail(categoryID)
	if err != nil {
		return resp, errors.New(reason.ProductCannotGetDetail)
	}

	resp = schema.DetailProductResp{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Currency:    product.Currency,
		TotalStock:  product.TotalStock,
		IsActive:    product.IsActive,
		ImageURL:    product.ImageURL,
		Category: schema.Category{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		},
	}

	return resp, nil
}

// update product by id
func (ps *ProductService) UpdateByID(id string, req *schema.UpdateProductReq) error {

	var updateData model.Product

	oldData, err := ps.productRepo.Detail(id)
	if err != nil {
		return errors.New(reason.ProductNotFound)
	}

	updateData.ID = oldData.ID
	updateData.Name = req.Name
	updateData.Description = req.Description
	updateData.Currency = req.Currency
	updateData.TotalStock = req.TotalStock
	updateData.IsActive = req.IsActive
	updateData.CategoryID = req.CategoryID

	err = ps.productRepo.Update(updateData)
	if err != nil {
		return errors.New(reason.ProductCannotUpdate)
	}

	return nil
}

// delete product by id
func (ps *ProductService) Delete(id string) error {

	_, err := ps.productRepo.Detail(id)
	if err != nil {
		return errors.New(reason.ProductNotFound)
	}

	err = ps.productRepo.Delete(id)
	if err != nil {
		return errors.New(reason.ProductCannotDelete)
	}

	return nil
}
