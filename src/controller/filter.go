package controller

import (
	"api/model"
	"fmt"
	"strings"
)

func parseFilters(query string) ([]model.Filter, error) {

	if query == "" {
		return []model.Filter{}, nil
	}

	filterParts := strings.Split(query, "*")
	var filters []model.Filter

	for _, part := range filterParts {
		filter := strings.SplitN(part, "[", 2)
		if len(filter) < 2 {
			return []model.Filter{}, fmt.Errorf("Invalid filter format")
		}
		operationAndValue := strings.SplitN(filter[1], "]", 2)
		if len(operationAndValue) < 2 {
			return []model.Filter{}, fmt.Errorf("Invalid filter format")
		}
		filters = append(filters, model.Filter{
			Field:    filter[0],
			Operator: operationAndValue[0],
			Value:    operationAndValue[1],
		})
	}
	return filters, nil
}
