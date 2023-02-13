package utils

import (
	"api/src/config"
	"strings"

	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func EscapeLike(left, right, word string) string {
	var n int
	for i := range word {
		if c := word[i]; c == '%' || c == '_' || c == '\\' {
			n++
		}
	}
	// No characters to escape.
	if n == 0 {
		return left + word + right
	}
	var b strings.Builder
	b.Grow(len(word) + n)
	for _, c := range word {
		if c == '%' || c == '_' || c == '\\' {
			b.WriteByte('\\')
		}
		b.WriteRune(c)
	}
	return left + b.String() + right
}

func SetupGormIntegrationTest[T any](s *suite.Suite, entity T) *gorm.DB {
	config, err := config.LoadEnv()
	s.Require().NoError(err)
	db, err := gorm.Open(postgres.Open(GetDbURI(config)), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	s.Require().NoError(err)
	err = db.AutoMigrate(&entity)
	s.Require().NoError(err)

	return db
}
