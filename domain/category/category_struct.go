package category

import (
	"context"
	"github.com/noxymon/ecomm-go/domain/pkg/dao/model"
	"github.com/noxymon/ecomm-go/domain/pkg/dao/query"
)

type ServiceCategory struct {
	Query *query.Query
}

func newServiceCategory(q *query.Query) *ServiceCategory {
	return &ServiceCategory{q}
}

func (category ServiceCategory) getCategoryById(ctx context.Context, id int) (*model.CategoriesLvl1, error) {
	firstRecord, err := category.Query.WithContext(ctx).CategoriesLvl1.First()
	if err != nil {
		return nil, err
	}
	return firstRecord, nil
}
