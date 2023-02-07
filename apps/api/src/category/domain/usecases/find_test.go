package usecases_test

import (
	"api/src/category/domain"
	"api/src/category/domain/usecases"
	"api/src/category/repository"
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
	initData := []domain.Category{{Id: 1, Name: "Test", Description: "123"}, {Id: 2, Name: "Test123", Description: "Traitor's requiem"}}
	testCases := []FindCategoryTestCase{
		{
			input: domain.CategoryQuery{
				Page:     3,
				PageSize: 2,
			},
			expected: make([]domain.Category, 0),
			TestName: "Exceed Page Number",
			initData: initData,
		},
		{
			input: domain.CategoryQuery{},
			expected: []domain.Category{
				{
					Id: 1, Name: "Test", Description: "123",
				},
				{
					Id: 2, Name: "Test123", Description: "Traitor's requiem",
				},
			},
			TestName: "Happy Case",
			initData: initData,
		},
		{
			input: domain.CategoryQuery{
				Page: 0, PageSize: 1,
			},
			expected: []domain.Category{initData[0]},
			initData: initData,
			TestName: "Pagination",
		},
		{
			input: domain.CategoryQuery{
				Page:     1,
				PageSize: 1,
			},
			expected: []domain.Category{initData[1]},
			initData: initData,
			TestName: "Pagination",
		},
		{
			input:    domain.CategoryQuery{Q: &[]string{"requiem"}[0]},
			expected: []domain.Category{initData[1]},
			initData: initData,
			TestName: "Search Query Happy Case",
		},
		{
			input: domain.CategoryQuery{
				Page:     1,
				PageSize: 1,
				Q:        &[]string{"requiem"}[0],
			},
			expected: []domain.Category{},
			initData: initData,
			TestName: "Search Query With Other Param",
		},
		{
			input: domain.CategoryQuery{
				Q: &[]string{"You never gonna get me lalalalala"}[0],
			},
			expected: []domain.Category{},
			initData: initData,
			TestName: "Not Found Search Query",
		},
	}

	for _, c := range testCases {
		s.T().Run(c.TestName, func(t *testing.T) {
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
