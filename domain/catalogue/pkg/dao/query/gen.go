// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                db,
		CategoriesLvl1:    newCategoriesLvl1(db, opts...),
		CategoriesLvl2:    newCategoriesLvl2(db, opts...),
		CategoriesLvl3:    newCategoriesLvl3(db, opts...),
		CategoriesProduct: newCategoriesProduct(db, opts...),
		Product:           newProduct(db, opts...),
		ProductVariant:    newProductVariant(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	CategoriesLvl1    categoriesLvl1
	CategoriesLvl2    categoriesLvl2
	CategoriesLvl3    categoriesLvl3
	CategoriesProduct categoriesProduct
	Product           product
	ProductVariant    productVariant
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                db,
		CategoriesLvl1:    q.CategoriesLvl1.clone(db),
		CategoriesLvl2:    q.CategoriesLvl2.clone(db),
		CategoriesLvl3:    q.CategoriesLvl3.clone(db),
		CategoriesProduct: q.CategoriesProduct.clone(db),
		Product:           q.Product.clone(db),
		ProductVariant:    q.ProductVariant.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                db,
		CategoriesLvl1:    q.CategoriesLvl1.replaceDB(db),
		CategoriesLvl2:    q.CategoriesLvl2.replaceDB(db),
		CategoriesLvl3:    q.CategoriesLvl3.replaceDB(db),
		CategoriesProduct: q.CategoriesProduct.replaceDB(db),
		Product:           q.Product.replaceDB(db),
		ProductVariant:    q.ProductVariant.replaceDB(db),
	}
}

type queryCtx struct {
	CategoriesLvl1    *categoriesLvl1Do
	CategoriesLvl2    *categoriesLvl2Do
	CategoriesLvl3    *categoriesLvl3Do
	CategoriesProduct *categoriesProductDo
	Product           *productDo
	ProductVariant    *productVariantDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		CategoriesLvl1:    q.CategoriesLvl1.WithContext(ctx),
		CategoriesLvl2:    q.CategoriesLvl2.WithContext(ctx),
		CategoriesLvl3:    q.CategoriesLvl3.WithContext(ctx),
		CategoriesProduct: q.CategoriesProduct.WithContext(ctx),
		Product:           q.Product.WithContext(ctx),
		ProductVariant:    q.ProductVariant.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}