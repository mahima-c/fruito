package transport

import (
	"context"

	"github.com/mahima-c/fruito/common"
	"github.com/mahima-c/fruito/model"
	"github.com/mahima-c/fruito/service/http/database"
)

func MakeUpsertProductEndpoint(dbService database.Service) common.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		reqs := request.([]UpsertProductRequest)

		var products []model.Product

		for _, req := range reqs {
			products = append(products, model.Product{
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
			})
		}

		err = dbService.UpsertProducts(products)
		if err != nil {
			return nil, err
		}

		return "success", nil
	}
}

func MakeGetProductByIdEndpoint(dbService database.Service) common.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetProductByIdRequest)

		product, err := dbService.GetProductById(req.ID)
		if err != nil {
			return nil, err
		}

		return product, nil
	}
}

func MakeGetAllProductsEndpoint(dbService database.Service) common.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetAllProductsRequest)

		products, err := dbService.GetAllProducts(req.Limit, req.Offset)
		if err != nil {
			return nil, err
		}

		if products == nil {
			products = []model.Product{}
		}

		return products, nil
	}
}
