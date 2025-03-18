package cli

import (
	"fmt"
	"github.com/jnunes-ds/hexagonal/application"
)

func Run(
	service application.ProductServiceInterface,
	action string,
	productId string,
	productName string,
	productPrice float64,
) (string, error) {
	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, productPrice)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product Id %s with the name %s has been created with the price %f and status %s",
			product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus())
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		res, err := service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s has been enabled", res.GetName())
	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		res, err := service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s has been disabled", res.GetName())
	default:
		res, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID: %s, Name %s, Price %f, Status %s",
			res.GetId(), res.GetName(), res.GetPrice(), res.GetStatus())
	}
	return result, nil
}
