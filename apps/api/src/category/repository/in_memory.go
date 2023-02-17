package repository

import (
	"api/src/category/domain"
	"api/src/common"
	"errors"
	"strings"
)

type CategoryInMemoryRepo struct {
	Data  []*domain.Category
	IsErr bool
}

func (repo *CategoryInMemoryRepo) Insert(req *domain.WriteCategoryBody) (*domain.Category, error) {
	if repo.IsErr {
		return nil, errors.New("unknown error")
	}

	for _, product := range repo.Data {
		if product.Name == req.Name {
			return nil, errors.New("resource exist")
		}
	}

	Id := uint(len(repo.Data) + 1)
	newCategory := domain.Category{ID: Id, Description: &req.Description, Name: req.Name}
	repo.Data = append(repo.Data, &newCategory)

	return &newCategory, nil
}

func (repo *CategoryInMemoryRepo) Update(id *int, req *domain.WriteCategoryBody) (*domain.Category, error) {
	if repo.IsErr {
		return nil, errors.New("unknown error")
	}

	for i := range repo.Data {
		category := repo.Data[i]
		if category.ID == uint(*id) {
			category.Description = &req.Description
			category.Name = req.Name
			return category, nil
		}
	}

	return nil, nil
}

func (repo *CategoryInMemoryRepo) Delete(req *int) (*domain.Category, error) {
	if repo.IsErr {
		return nil, errors.New("unknown error")
	}

	var res domain.Category
	pos := -1
	filteredList := make([]*domain.Category, 0)

	for i, inventory := range repo.Data {
		if inventory.ID == uint(*req) {
			res = *inventory
			pos = i
			continue
		}

		filteredList = append(filteredList, inventory)
	}
	repo.Data = filteredList

	if pos < 0 {
		return nil, nil
	}

	return &res, nil
}

func (repo *CategoryInMemoryRepo) FindOne(id *int) (*domain.Category, error) {
	if repo.IsErr {
		return nil, errors.New("unknown error")
	}

	var res *domain.Category

	for _, inventory := range repo.Data {
		if inventory.ID == uint(*id) {
			res = inventory
			break
		}
	}

	return res, nil
}

func (repo *CategoryInMemoryRepo) Find(req *domain.CategoryQuery) (common.BasePaginationResponse[domain.Category], error) {
	if repo.IsErr {
		return common.BasePaginationResponse[domain.Category]{}, errors.New("unknown error")
	}

	Page := req.GetPage()
	PageSize := req.GetPageSize()
	res := repo.Data
	if req.Q != nil {
		res = make([]*domain.Category, 0)
		for _, category := range repo.Data {
			q := strings.ToLower(*req.Q)
			entityContainsQ := strings.Contains(strings.ToLower(*category.Description), q) || strings.Contains(strings.ToLower(category.Name), q)

			if entityContainsQ {
				res = append(res, category)
			}
		}
	}

	start := uint(req.GetOffset())
	end := uint(start + PageSize)
	lenght := uint(len(res))
	res = res[min(start, lenght):min(end, lenght)]

	return common.BasePaginationResponse[domain.Category]{
		Items:    res,
		Page:     Page,
		PageSize: PageSize,
		Total:    uint(len(repo.Data)),
	}, nil
}

func min(x, y uint) uint {
	if y < x {
		return y
	}
	return x
}
