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

func newResourceLabel(db *gorm.DB) resourceLabel {
	_resourceLabel := resourceLabel{}

	_resourceLabel.resourceLabelDo.UseDB(db)
	_resourceLabel.resourceLabelDo.UseModel(&model.ResourceLabel{})

	tableName := _resourceLabel.resourceLabelDo.TableName()
	_resourceLabel.ALL = field.NewField(tableName, "*")
	_resourceLabel.Id = field.NewString(tableName, "id")
	_resourceLabel.Key = field.NewString(tableName, "key")
	_resourceLabel.CreateId = field.NewString(tableName, "create_id")
	_resourceLabel.CreatedAt = field.NewInt64(tableName, "created_at")
	_resourceLabel.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_resourceLabel.fillFieldMap()

	return _resourceLabel
}

type resourceLabel struct {
	resourceLabelDo resourceLabelDo

	ALL       field.Field
	Id        field.String
	Key       field.String
	CreateId  field.String
	CreatedAt field.Int64
	UpdatedAt field.Int64

	fieldMap map[string]field.Expr
}

func (r resourceLabel) Table(newTableName string) *resourceLabel {
	r.resourceLabelDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r resourceLabel) As(alias string) *resourceLabel {
	r.resourceLabelDo.DO = *(r.resourceLabelDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *resourceLabel) updateTableName(table string) *resourceLabel {
	r.ALL = field.NewField(table, "*")
	r.Id = field.NewString(table, "id")
	r.Key = field.NewString(table, "key")
	r.CreateId = field.NewString(table, "create_id")
	r.CreatedAt = field.NewInt64(table, "created_at")
	r.UpdatedAt = field.NewInt64(table, "updated_at")

	r.fillFieldMap()

	return r
}

func (r *resourceLabel) WithContext(ctx context.Context) *resourceLabelDo {
	return r.resourceLabelDo.WithContext(ctx)
}

func (r resourceLabel) TableName() string { return r.resourceLabelDo.TableName() }

func (r resourceLabel) Alias() string { return r.resourceLabelDo.Alias() }

func (r *resourceLabel) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *resourceLabel) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 5)
	r.fieldMap["id"] = r.Id
	r.fieldMap["key"] = r.Key
	r.fieldMap["create_id"] = r.CreateId
	r.fieldMap["created_at"] = r.CreatedAt
	r.fieldMap["updated_at"] = r.UpdatedAt
}

func (r resourceLabel) clone(db *gorm.DB) resourceLabel {
	r.resourceLabelDo.ReplaceDB(db)
	return r
}

type resourceLabelDo struct{ gen.DO }

func (r resourceLabelDo) Debug() *resourceLabelDo {
	return r.withDO(r.DO.Debug())
}

func (r resourceLabelDo) WithContext(ctx context.Context) *resourceLabelDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r resourceLabelDo) Clauses(conds ...clause.Expression) *resourceLabelDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r resourceLabelDo) Returning(value interface{}, columns ...string) *resourceLabelDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r resourceLabelDo) Not(conds ...gen.Condition) *resourceLabelDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r resourceLabelDo) Or(conds ...gen.Condition) *resourceLabelDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r resourceLabelDo) Select(conds ...field.Expr) *resourceLabelDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r resourceLabelDo) Where(conds ...gen.Condition) *resourceLabelDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r resourceLabelDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *resourceLabelDo {
	return r.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (r resourceLabelDo) Order(conds ...field.Expr) *resourceLabelDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r resourceLabelDo) Distinct(cols ...field.Expr) *resourceLabelDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r resourceLabelDo) Omit(cols ...field.Expr) *resourceLabelDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r resourceLabelDo) Join(table schema.Tabler, on ...field.Expr) *resourceLabelDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r resourceLabelDo) LeftJoin(table schema.Tabler, on ...field.Expr) *resourceLabelDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r resourceLabelDo) RightJoin(table schema.Tabler, on ...field.Expr) *resourceLabelDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r resourceLabelDo) Group(cols ...field.Expr) *resourceLabelDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r resourceLabelDo) Having(conds ...gen.Condition) *resourceLabelDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r resourceLabelDo) Limit(limit int) *resourceLabelDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r resourceLabelDo) Offset(offset int) *resourceLabelDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r resourceLabelDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *resourceLabelDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r resourceLabelDo) Unscoped() *resourceLabelDo {
	return r.withDO(r.DO.Unscoped())
}

func (r resourceLabelDo) Create(values ...*model.ResourceLabel) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r resourceLabelDo) CreateInBatches(values []*model.ResourceLabel, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r resourceLabelDo) Save(values ...*model.ResourceLabel) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r resourceLabelDo) First() (*model.ResourceLabel, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResourceLabel), nil
	}
}

func (r resourceLabelDo) Take() (*model.ResourceLabel, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResourceLabel), nil
	}
}

func (r resourceLabelDo) Last() (*model.ResourceLabel, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResourceLabel), nil
	}
}

func (r resourceLabelDo) Find() ([]*model.ResourceLabel, error) {
	result, err := r.DO.Find()
	return result.([]*model.ResourceLabel), err
}

func (r resourceLabelDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ResourceLabel, err error) {
	buf := make([]*model.ResourceLabel, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r resourceLabelDo) FindInBatches(result *[]*model.ResourceLabel, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r resourceLabelDo) Attrs(attrs ...field.AssignExpr) *resourceLabelDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r resourceLabelDo) Assign(attrs ...field.AssignExpr) *resourceLabelDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r resourceLabelDo) Joins(fields ...field.RelationField) *resourceLabelDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r resourceLabelDo) Preload(fields ...field.RelationField) *resourceLabelDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r resourceLabelDo) FirstOrInit() (*model.ResourceLabel, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResourceLabel), nil
	}
}

func (r resourceLabelDo) FirstOrCreate() (*model.ResourceLabel, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResourceLabel), nil
	}
}

func (r resourceLabelDo) FindByPage(offset int, limit int) (result []*model.ResourceLabel, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r resourceLabelDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r *resourceLabelDo) withDO(do gen.Dao) *resourceLabelDo {
	r.DO = *do.(*gen.DO)
	return r
}
