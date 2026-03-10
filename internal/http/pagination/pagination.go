package pagination

import (
	"net/http"
	"strconv"
)

const (
	DefaultLimit = 20
	MaxLimit     = 100
)

type Pagination struct {
	Limit  int
	Offset int
}

func ParsePagination(r *http.Request) Pagination {
	limit := DefaultLimit
	offset := 0

	if l := r.URL.Query().Get("limit"); l != "" {
		if parsed, err := strconv.Atoi(l); err == nil {
			limit = parsed
		}
	}

	if o := r.URL.Query().Get("offset"); o != "" {
		if parsed, err := strconv.Atoi(o); err == nil {
			offset = parsed
		}
	}

	if limit > MaxLimit {
		limit = MaxLimit
	}

	if limit <= 0 {
		limit = DefaultLimit
	}

	if offset < 0 {
		offset = 0
	}

	return Pagination{
		Limit:  limit,
		Offset: offset,
	}
}
