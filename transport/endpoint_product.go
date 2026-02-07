package transport

import (
	"context"
	"fmt"

	"github.com/Mrhb787/hospital-ward-manager/common"
	"github.com/Mrhb787/hospital-ward-manager/model"
	"github.com/Mrhb787/hospital-ward-manager/service/http/database"
)

func MakeUpsertProductEndpoint(dbService database.Service) common.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UpsertProductRequest)

		fmt.Println("MakeUpsertProductEndpoint", req)
		product := model.Product{
			ID:            req.ID,
			Name:          req.Name,
			Image:         req.Image,
			Price:         req.Price,
			UnitOfMeasure: req.UnitOfMeasure,
			TotalQty:      req.TotalQty,
			Description:   req.Description,
			Rating:        req.Rating,
			RatingCount:   req.RatingCount,
			Tag:           req.Tag,
		}

		err = dbService.UpsertProduct(product)
		if err != nil {
			return nil, err
		}

		return "success", nil
	}
}
