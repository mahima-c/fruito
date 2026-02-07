package transport

import (
	"context"
	"encoding/json"
	"net/http"
)

func DecodeUpsertProductRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	req := UpsertProductRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}
