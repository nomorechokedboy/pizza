package usecases_test

import (
	"api/src/category/domain"
	"api/src/category/domain/usecases"
	"api/src/category/repository"
	"api/src/common"
	"api/src/utils"
	"testing"

	"github.com/stretchr/testify/suite"
)

type FindCategoryTestSuite struct {
	suite.Suite
	useCase usecases.FindCategoryUseCase
	repo    repository.CategoryInMemoryRepo
}

type FindCategoryTestCase struct {
	input    domain.CategoryQuery
	expected []domain.Category
	TestName string
	initData []domain.Category
}

func (s *FindCategoryTestSuite) SetupTest() {
	s.repo = repository.CategoryInMemoryRepo{Data: make([]domain.Category, 0), IsErr: false}
	s.useCase = usecases.FindCategoryUseCase{Repo: &s.repo}
}

func (s *FindCategoryTestSuite) TearDownTest() {
	s.useCase.Repo = &repository.CategoryInMemoryRepo{Data: make([]domain.Category, 0), IsErr: false}
}

func (s *FindCategoryTestSuite) TestFindUnknownError() {
	s.repo.IsErr = true
	category, err := s.useCase.Execute(nil)

	s.Assertions.Nil(category)
	s.Assertions.EqualError(err, "unknown error")
}

func (s *FindCategoryTestSuite) TestFindUseCases() {
	initData := []domain.Category{{ID: 1, Name: "Test", Description: utils.GetDataTypeAddress("123")}, {ID: 2, Name: "Test123", Description: utils.GetDataTypeAddress("Traitor's requiem")}}
	testCases := []FindCategoryTestCase{
		{
			input: domain.CategoryQuery{
				BaseQuery: common.BaseQuery{
					Page:     utils.GetDataTypeAddress(uint(3)),
					PageSize: utils.GetDataTypeAddress(uint(2)),
				},
			},
			expected: make([]domain.Category, 0),
			TestName: "Exceed Page Number",
			initData: initData,
		},
		{
			input: domain.CategoryQuery{},
			expected: []domain.Category{
				{
					ID:          1,
					Name:        "Test",
					Description: utils.GetDataTypeAddress("123"),
				},
				{
					ID: 2, Name: "Test123", Description: utils.GetDataTypeAddress("Traitor's requiem"),
				},
			},
			TestName: "Happy Case",
			initData: initData,
		},
		{
			input: domain.CategoryQuery{
				BaseQuery: common.BaseQuery{
					Page:     utils.GetDataTypeAddress(uint(0)),
					PageSize: utils.GetDataTypeAddress(uint(1)),
				},
			},
			expected: []domain.Category{initData[0]},
			initData: initData,
			TestName: "Pagination",
		},
		{
			input: domain.CategoryQuery{
				BaseQuery: common.BaseQuery{
					Page:     utils.GetDataTypeAddress(uint(1)),
					PageSize: utils.GetDataTypeAddress(uint(1)),
				},
			},
			expected: []domain.Category{initData[1]},
			initData: initData,
			TestName: "Pagination",
		},
		{
			input: domain.CategoryQuery{
				BaseQuery: common.BaseQuery{
					Q: utils.GetDataTypeAddress("requiem"),
				},
			},
			expected: []domain.Category{initData[1]},
			initData: initData,
			TestName: "Search Query Happy Case",
		},
		{
			input: domain.CategoryQuery{
				BaseQuery: common.BaseQuery{
					Page:     utils.GetDataTypeAddress(uint(1)),
					PageSize: utils.GetDataTypeAddress(uint(1)),
					Q:        utils.GetDataTypeAddress("requiem"),
				},
			},
			expected: []domain.Category{},
			initData: initData,
			TestName: "Search Query With Other Param",
		},
		{
			input: domain.CategoryQuery{
				BaseQuery: common.BaseQuery{
					Q: utils.GetDataTypeAddress("You never gonna get me lalalalala"),
				},
			},
			expected: []domain.Category{},
			initData: initData,
			TestName: "Not Found Search Query",
		},
	}

	for _, c := range testCases {
		s.Run(c.TestName, func() {
			s.repo.Data = c.initData
			categories, err := s.useCase.Execute(&c.input)

			s.Assertions.Nil(err)
			s.Assertions.Equal(c.expected, *categories)
		})
	}
}

func TestFindCategoryTestSuite(t *testing.T) {
	suite.Run(t, new(FindCategoryTestSuite))
}
