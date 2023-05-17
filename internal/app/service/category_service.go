package service

import (
	"errors"
	"fmt"

	"github.com/kenykendf/go-restful/internal/app/model"
	"github.com/kenykendf/go-restful/internal/app/schema"
	"github.com/kenykendf/go-restful/internal/pkg/reason"
	"github.com/sirupsen/logrus"
)

type CategoryService struct {
	repo CategoryRepository
}

func NewCategoryService(repo CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (cs *CategoryService) Create(req *schema.CreateCategoryReq) error {
	var insertData model.Category

	insertData.Name = req.Name
	insertData.Description = req.Description

	err := cs.repo.Create(insertData)
	if err != nil {
		return errors.New(reason.CategoryCannotCreate)
	}
	return nil
}

func (cs *CategoryService) BrowseAll(req *schema.BrowseCategoryReq) ([]schema.GetCategoryResp, error) {
	var resp []schema.GetCategoryResp

	dbSearch := model.BrowseCategory{}
	dbSearch.Page = req.Page
	dbSearch.PageSize = req.PageSize

	categories, err := cs.repo.Browse(dbSearch)
	if err != nil {
		return nil, errors.New(reason.CategoryCannotBrowse)
	}

	for _, value := range categories {
		var respData schema.GetCategoryResp
		respData.ID = value.ID
		respData.Name = value.Name
		respData.Description = value.Description
		resp = append(resp, respData)
	}

	return resp, nil
}

func (cs *CategoryService) DetailCategory(id string) (schema.GetCategoryResp, error) {
	var resp schema.GetCategoryResp

	category, err := cs.repo.Detail(id)
	if err != nil {
		return schema.GetCategoryResp{}, errors.New("server error, unable to fetch category detail")
	}

	resp.ID = category.ID
	resp.Name = category.Name
	resp.Description = category.Description

	return resp, err
}

func (cs *CategoryService) UpdateCategory(id string, req *schema.UpdateCategoryReq) error {
	var updateData model.Category

	category, err := cs.repo.Detail(id)
	if err != nil {
		return errors.New(reason.CategoryNotFound)
	}

	updateData.Name = req.Name
	if req.Name == "" {
		updateData.Name = category.Name
	}
	updateData.Description = req.Description
	if req.Description == "" {
		updateData.Description = category.Description
	}

	err = cs.repo.Update(updateData)
	if err != nil {
		logrus.Error(fmt.Errorf("error updating category : %w", err))
		return errors.New(reason.CategoryCannotUpdate)
	}
	return nil
}

func (cs *CategoryService) DeleteCategory(id string) error {

	_, err := cs.repo.Detail(id)
	if err != nil {
		return errors.New(reason.CategoryNotFound)
	}

	err = cs.repo.Delete(id)
	if err != nil {
		return errors.New(reason.CategoryCannotDelete)
	}
	return nil
}
