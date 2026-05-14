package core_http_request

import (
	"fmt"
	"net/http"
	"strconv"

	core_errors "github.com/Mertvyki/book-shop/internal/core/errrors"
)

func GetIntQueryParam(r *http.Request, key string) (*int, error) {
	param := r.URL.Query().Get(key)
	if param == "" {
		return nil, nil
	}

	val, err := strconv.Atoi(param)
	if err != nil {
		return nil, fmt.Errorf("param=%s by key=%s not a valid integer: %v: %w", param, key, err, core_errors.ErrInvalidArgument)
	}

	return &val, nil
}

func GetStringQueryParam(r *http.Request, key string) (*string, error) {
	param := r.URL.Query().Get(key)
	if param == "" {
		return nil, nil
	}

	return &param, nil
}

func GetFloatQueryParam(r *http.Request, key string) (*float64, error) {
	param := r.URL.Query().Get(key)
	if param == "" {
		return nil, nil
	}

	val, err := strconv.ParseFloat(param, 64)
	if err != nil {
		return nil, fmt.Errorf("param=%s by key%s not a valid float: %v: %w", param, key, err, core_errors.ErrInvalidArgument)
	}

	return &val, nil
}
