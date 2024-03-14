// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"github.com/noxymon/ecomm-go/domain/pkg/dao/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newProduct(db *gorm.DB, opts ...gen.DOOption) product {
	_product := product{}

	_product.productDo.UseDB(db, opts...)
	_product.productDo.UseModel(&model.Product{})

	tableName := _product.productDo.TableName()
	_product.ALL = field.NewAsterisk(tableName)
	_product.ID = field.NewInt64(tableName, "id")
	_product.Name = field.NewString(tableName, "name")
	_product.Description = field.NewString(tableName, "description")
	_product.MerchantID = field.NewString(tableName, "merchant_id")
	_product.Status = field.NewInt32(tableName, "status")
	_product.CreatedDate = field.NewTime(tableName, "created_date")
	_product.UpdatedDate = field.NewTime(tableName, "updated_date")
	_product.CreatedBy = field.NewString(tableName, "created_by")
	_product.UpdatedBy = field.NewString(tableName, "updated_by")

	_product.fillFieldMap()

	return _product
}

type product struct {
	productDo productDo

	ALL         field.Asterisk
	ID          field.Int64
	Name        field.String
	Description field.String
	MerchantID  field.String
	Status      field.Int32
	CreatedDate field.Time
	UpdatedDate field.Time
	CreatedBy   field.String
	UpdatedBy   field.String

	fieldMap map[string]field.Expr
}

func (p product) Table(newTableName string) *product {
	p.productDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p product) As(alias string) *product {
	p.productDo.DO = *(p.productDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *product) updateTableName(table string) *product {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.Name = field.NewString(table, "name")
	p.Description = field.NewString(table, "description")
	p.MerchantID = field.NewString(table, "merchant_id")
	p.Status = field.NewInt32(table, "status")
	p.CreatedDate = field.NewTime(table, "created_date")
	p.UpdatedDate = field.NewTime(table, "updated_date")
	p.CreatedBy = field.NewString(table, "created_by")
	p.UpdatedBy = field.NewString(table, "updated_by")

	p.fillFieldMap()

	return p
}

func (p *product) WithContext(ctx context.Context) *productDo { return p.productDo.WithContext(ctx) }

func (p product) TableName() string { return p.productDo.TableName() }

func (p product) Alias() string { return p.productDo.Alias() }

func (p product) Columns(cols ...field.Expr) gen.Columns { return p.productDo.Columns(cols...) }

func (p *product) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *product) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 9)
	p.fieldMap["id"] = p.ID
	p.fieldMap["name"] = p.Name
	p.fieldMap["description"] = p.Description
	p.fieldMap["merchant_id"] = p.MerchantID
	p.fieldMap["status"] = p.Status
	p.fieldMap["created_date"] = p.CreatedDate
	p.fieldMap["updated_date"] = p.UpdatedDate
	p.fieldMap["created_by"] = p.CreatedBy
	p.fieldMap["updated_by"] = p.UpdatedBy
}

func (p product) clone(db *gorm.DB) product {
	p.productDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p product) replaceDB(db *gorm.DB) product {
	p.productDo.ReplaceDB(db)
	return p
}

type productDo struct{ gen.DO }

func (p productDo) Debug() *productDo {
	return p.withDO(p.DO.Debug())
}

func (p productDo) WithContext(ctx context.Context) *productDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p productDo) ReadDB() *productDo {
	return p.Clauses(dbresolver.Read)
}

func (p productDo) WriteDB() *productDo {
	return p.Clauses(dbresolver.Write)
}

func (p productDo) Session(config *gorm.Session) *productDo {
	return p.withDO(p.DO.Session(config))
}

func (p productDo) Clauses(conds ...clause.Expression) *productDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p productDo) Returning(value interface{}, columns ...string) *productDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p productDo) Not(conds ...gen.Condition) *productDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p productDo) Or(conds ...gen.Condition) *productDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p productDo) Select(conds ...field.Expr) *productDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p productDo) Where(conds ...gen.Condition) *productDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p productDo) Order(conds ...field.Expr) *productDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p productDo) Distinct(cols ...field.Expr) *productDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p productDo) Omit(cols ...field.Expr) *productDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p productDo) Join(table schema.Tabler, on ...field.Expr) *productDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p productDo) LeftJoin(table schema.Tabler, on ...field.Expr) *productDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p productDo) RightJoin(table schema.Tabler, on ...field.Expr) *productDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p productDo) Group(cols ...field.Expr) *productDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p productDo) Having(conds ...gen.Condition) *productDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p productDo) Limit(limit int) *productDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p productDo) Offset(offset int) *productDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p productDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *productDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p productDo) Unscoped() *productDo {
	return p.withDO(p.DO.Unscoped())
}

func (p productDo) Create(values ...*model.Product) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p productDo) CreateInBatches(values []*model.Product, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p productDo) Save(values ...*model.Product) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p productDo) First() (*model.Product, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Product), nil
	}
}

func (p productDo) Take() (*model.Product, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Product), nil
	}
}

func (p productDo) Last() (*model.Product, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Product), nil
	}
}

func (p productDo) Find() ([]*model.Product, error) {
	result, err := p.DO.Find()
	return result.([]*model.Product), err
}

func (p productDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Product, err error) {
	buf := make([]*model.Product, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p productDo) FindInBatches(result *[]*model.Product, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p productDo) Attrs(attrs ...field.AssignExpr) *productDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p productDo) Assign(attrs ...field.AssignExpr) *productDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p productDo) Joins(fields ...field.RelationField) *productDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p productDo) Preload(fields ...field.RelationField) *productDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p productDo) FirstOrInit() (*model.Product, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Product), nil
	}
}

func (p productDo) FirstOrCreate() (*model.Product, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Product), nil
	}
}

func (p productDo) FindByPage(offset int, limit int) (result []*model.Product, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p productDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p productDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p productDo) Delete(models ...*model.Product) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *productDo) withDO(do gen.Dao) *productDo {
	p.DO = *do.(*gen.DO)
	return p
}