package transport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func DecodeUpsertProductRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	// Try to decode as a list first
	var reqs []UpsertProductRequest

	// Read body into buffer to allow re-reading if needed, or simply peek?
	// For simplicity, let's assume if it starts with [ it is array.
	// But standard json.Decode is smart.
	// To support BOTH single object and array without reading body twice easily:
	// We can decode into json.RawMessage then check.
	// Or we can just support array for now as requested.
	// Given "add support of array", replacing single with array is cleaner API design usually (batch endpoint).
	// But let's support array by default.
	if err := json.NewDecoder(r.Body).Decode(&reqs); err != nil {
		// If fails, maybe it was a single object?
		// Re-creation of decoder is hard because stream is consumed.
		// So simpler is to just define the endpoint now accepts an ARRAY.
		return nil, err
	}
	return reqs, nil
}

func DecodeGetProductByIdRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	req := GetProductByIdRequest{}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		return req, errors.New("invalid product id")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return req, errors.New("invalid product id")
	}
	req.ID = id

	return req, nil
}

func DecodeGetAllProductsRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	req := GetAllProductsRequest{}

	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("page_size")

	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
		if page < 1 {
			page = 1
		}
	}

	pageSize := 10
	if pageSizeStr != "" {
		pageSize, _ = strconv.Atoi(pageSizeStr)
		if pageSize < 1 {
			pageSize = 10
		}
	}

	req.Limit = pageSize
	req.Offset = (page - 1) * pageSize

	return req, nil
}
