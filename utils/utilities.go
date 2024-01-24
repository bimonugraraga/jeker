package utils

import "github.com/bdn/jeker/dto/requests"

func Ternary[T any](condition bool, a, b T) T {
	if condition {
		return a
	}
	return b
}

func DefaultPagination(params requests.GeneralPaginationQuery) requests.GeneralPaginationQuery {
	if params.Page == 0 {
		params.Page = 1
	}
	if params.Limit == 0 {
		params.Limit = 10
	}
	if params.OrderWith == "" {
		params.OrderWith = "updated_at"
	}
	if params.OrderBy == "" {
		params.OrderBy = "DESC"
	}

	return params
}
