package mapper

import (
	"orders/internal/dto"
	"orders/internal/models"
)

func ConvertOrderDTOToModel(dto *dto.Order) *models.Order {
	return &models.Order{
		CustomerID: dto.CustomerID,
		Products:   ConvertOrderProductsDTOToModel(dto.Products),
	}
}

func ConvertOrderProductsDTOToModel(orderProductsDTO []dto.OrderProduct) []models.OrderProduct {
	var products []models.OrderProduct
	for _, dtoProduct := range orderProductsDTO {
		products = append(products, models.OrderProduct{
			ProductID: dtoProduct.ProductID,
			Quantity:  dtoProduct.Quantity,
		})
	}
	return products
}

func ConvertOrderStatusDTOtoModels(dto *dto.OrderStatus) (*models.OrderHistory, *models.Order) {
	return &models.OrderHistory{
			Status:  dto.Status,
			Comment: dto.Comment,
		}, &models.Order{
			Status: dto.Status,
		}
}
