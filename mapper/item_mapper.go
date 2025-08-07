package mapper

import (
	"github.com/dickysetiawan031000/go-backend/dto/item"
	"github.com/dickysetiawan031000/go-backend/model"
)

func ToItemModel(input interface{}) model.Item {
	switch v := input.(type) {
	case item.CreateItemRequest:
		return model.Item{
			Name:        v.Name,
			Description: v.Description,
			Price:       v.Price,
			Stock:       v.Stock,
		}
	case item.UpdateItemRequest:
		return model.Item{
			Name:        v.Name,
			Description: v.Description,
			Price:       v.Price,
			Stock:       v.Stock,
		}
	default:
		return model.Item{}
	}
}

func ToItemResponse(m model.Item) item.ItemResponse {
	return item.ItemResponse{
		ID:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		Price:       m.Price,
		Stock:       m.Stock,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

func ToItemResponses(items []model.Item) []item.ItemResponse {
	var responses []item.ItemResponse
	for _, i := range items {
		responses = append(responses, ToItemResponse(i))
	}
	return responses
}
