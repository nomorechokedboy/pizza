package utils

import "github.com/gosimple/slug"

func Slug(s string) string {
	slug := slug.Make(s)

	return slug
}
