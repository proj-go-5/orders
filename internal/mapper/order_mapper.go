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
