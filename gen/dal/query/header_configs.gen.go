// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"go_web_example/gen/dal/model"
)

func newHeaderConfig(db *gorm.DB) headerConfig {
	_headerConfig := headerConfig{}

	_headerConfig.headerConfigDo.UseDB(db)
	_headerConfig.headerConfigDo.UseModel(&model.HeaderConfig{})

	tableName := _headerConfig.headerConfigDo.TableName()
	_headerConfig.ALL = field.NewField(tableName, "*")
	_headerConfig.Id = field.NewString(tableName, "id")
	_headerConfig.DomainId = field.NewString(tableName, "domain_id")
	_headerConfig.Parameter = field.NewString(tableName, "parameter")
	_headerConfig.Value = field.NewString(tableName, "value")
	_headerConfig.CreatedAt = field.NewInt64(tableName, "created_at")
	_headerConfig.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_headerConfig.fillFieldMap()

	return _headerConfig
}

type headerConfig struct {
	headerConfigDo headerConfigDo

	ALL       field.Field
	Id        field.String
	DomainId  field.String
	Parameter field.String
	Value     field.String
	CreatedAt field.Int64
	UpdatedAt field.Int64

	fieldMap map[string]field.Expr
}

func (h headerConfig) Table(newTableName string) *headerConfig {
	h.headerConfigDo.UseTable(newTableName)
	return h.updateTableName(newTableName)
}

func (h headerConfig) As(alias string) *headerConfig {
	h.headerConfigDo.DO = *(h.headerConfigDo.As(alias).(*gen.DO))
	return h.updateTableName(alias)
}

func (h *headerConfig) updateTableName(table string) *headerConfig {
	h.ALL = field.NewField(table, "*")
	h.Id = field.NewString(table, "id")
	h.DomainId = field.NewString(table, "domain_id")
	h.Parameter = field.NewString(table, "parameter")
	h.Value = field.NewString(table, "value")
	h.CreatedAt = field.NewInt64(table, "created_at")
	h.UpdatedAt = field.NewInt64(table, "updated_at")

	h.fillFieldMap()

	return h
}

func (h *headerConfig) WithContext(ctx context.Context) *headerConfigDo {
	return h.headerConfigDo.WithContext(ctx)
}

func (h headerConfig) TableName() string { return h.headerConfigDo.TableName() }

func (h headerConfig) Alias() string { return h.headerConfigDo.Alias() }

func (h *headerConfig) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := h.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (h *headerConfig) fillFieldMap() {
	h.fieldMap = make(map[string]field.Expr, 6)
	h.fieldMap["id"] = h.Id
	h.fieldMap["domain_id"] = h.DomainId
	h.fieldMap["parameter"] = h.Parameter
	h.fieldMap["value"] = h.Value
	h.fieldMap["created_at"] = h.CreatedAt
	h.fieldMap["updated_at"] = h.UpdatedAt
}

func (h headerConfig) clone(db *gorm.DB) headerConfig {
	h.headerConfigDo.ReplaceDB(db)
	return h
}

type headerConfigDo struct{ gen.DO }

func (h headerConfigDo) Debug() *headerConfigDo {
	return h.withDO(h.DO.Debug())
}

func (h headerConfigDo) WithContext(ctx context.Context) *headerConfigDo {
	return h.withDO(h.DO.WithContext(ctx))
}

func (h headerConfigDo) Clauses(conds ...clause.Expression) *headerConfigDo {
	return h.withDO(h.DO.Clauses(conds...))
}

func (h headerConfigDo) Returning(value interface{}, columns ...string) *headerConfigDo {
	return h.withDO(h.DO.Returning(value, columns...))
}

func (h headerConfigDo) Not(conds ...gen.Condition) *headerConfigDo {
	return h.withDO(h.DO.Not(conds...))
}

func (h headerConfigDo) Or(conds ...gen.Condition) *headerConfigDo {
	return h.withDO(h.DO.Or(conds...))
}

func (h headerConfigDo) Select(conds ...field.Expr) *headerConfigDo {
	return h.withDO(h.DO.Select(conds...))
}

func (h headerConfigDo) Where(conds ...gen.Condition) *headerConfigDo {
	return h.withDO(h.DO.Where(conds...))
}

func (h headerConfigDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *headerConfigDo {
	return h.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (h headerConfigDo) Order(conds ...field.Expr) *headerConfigDo {
	return h.withDO(h.DO.Order(conds...))
}

func (h headerConfigDo) Distinct(cols ...field.Expr) *headerConfigDo {
	return h.withDO(h.DO.Distinct(cols...))
}

func (h headerConfigDo) Omit(cols ...field.Expr) *headerConfigDo {
	return h.withDO(h.DO.Omit(cols...))
}

func (h headerConfigDo) Join(table schema.Tabler, on ...field.Expr) *headerConfigDo {
	return h.withDO(h.DO.Join(table, on...))
}

func (h headerConfigDo) LeftJoin(table schema.Tabler, on ...field.Expr) *headerConfigDo {
	return h.withDO(h.DO.LeftJoin(table, on...))
}

func (h headerConfigDo) RightJoin(table schema.Tabler, on ...field.Expr) *headerConfigDo {
	return h.withDO(h.DO.RightJoin(table, on...))
}

func (h headerConfigDo) Group(cols ...field.Expr) *headerConfigDo {
	return h.withDO(h.DO.Group(cols...))
}

func (h headerConfigDo) Having(conds ...gen.Condition) *headerConfigDo {
	return h.withDO(h.DO.Having(conds...))
}

func (h headerConfigDo) Limit(limit int) *headerConfigDo {
	return h.withDO(h.DO.Limit(limit))
}

func (h headerConfigDo) Offset(offset int) *headerConfigDo {
	return h.withDO(h.DO.Offset(offset))
}

func (h headerConfigDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *headerConfigDo {
	return h.withDO(h.DO.Scopes(funcs...))
}

func (h headerConfigDo) Unscoped() *headerConfigDo {
	return h.withDO(h.DO.Unscoped())
}

func (h headerConfigDo) Create(values ...*model.HeaderConfig) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Create(values)
}

func (h headerConfigDo) CreateInBatches(values []*model.HeaderConfig, batchSize int) error {
	return h.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (h headerConfigDo) Save(values ...*model.HeaderConfig) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Save(values)
}

func (h headerConfigDo) First() (*model.HeaderConfig, error) {
	if result, err := h.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.HeaderConfig), nil
	}
}

func (h headerConfigDo) Take() (*model.HeaderConfig, error) {
	if result, err := h.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.HeaderConfig), nil
	}
}

func (h headerConfigDo) Last() (*model.HeaderConfig, error) {
	if result, err := h.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.HeaderConfig), nil
	}
}

func (h headerConfigDo) Find() ([]*model.HeaderConfig, error) {
	result, err := h.DO.Find()
	return result.([]*model.HeaderConfig), err
}

func (h headerConfigDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.HeaderConfig, err error) {
	buf := make([]*model.HeaderConfig, 0, batchSize)
	err = h.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (h headerConfigDo) FindInBatches(result *[]*model.HeaderConfig, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return h.DO.FindInBatches(result, batchSize, fc)
}

func (h headerConfigDo) Attrs(attrs ...field.AssignExpr) *headerConfigDo {
	return h.withDO(h.DO.Attrs(attrs...))
}

func (h headerConfigDo) Assign(attrs ...field.AssignExpr) *headerConfigDo {
	return h.withDO(h.DO.Assign(attrs...))
}

func (h headerConfigDo) Joins(fields ...field.RelationField) *headerConfigDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Joins(_f))
	}
	return &h
}

func (h headerConfigDo) Preload(fields ...field.RelationField) *headerConfigDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Preload(_f))
	}
	return &h
}

func (h headerConfigDo) FirstOrInit() (*model.HeaderConfig, error) {
	if result, err := h.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.HeaderConfig), nil
	}
}

func (h headerConfigDo) FirstOrCreate() (*model.HeaderConfig, error) {
	if result, err := h.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.HeaderConfig), nil
	}
}

func (h headerConfigDo) FindByPage(offset int, limit int) (result []*model.HeaderConfig, count int64, err error) {
	result, err = h.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = h.Offset(-1).Limit(-1).Count()
	return
}

func (h headerConfigDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = h.Count()
	if err != nil {
		return
	}

	err = h.Offset(offset).Limit(limit).Scan(result)
	return
}

func (h *headerConfigDo) withDO(do gen.Dao) *headerConfigDo {
	h.DO = *do.(*gen.DO)
	return h
}
