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

	"cdnFiber/infrastructure/gen/cdn/dal/model"
)

func newDocManage(db *gorm.DB) docManage {
	_docManage := docManage{}

	_docManage.docManageDo.UseDB(db)
	_docManage.docManageDo.UseModel(&model.DocManage{})

	tableName := _docManage.docManageDo.TableName()
	_docManage.ALL = field.NewField(tableName, "*")
	_docManage.Id = field.NewUint64(tableName, "id")
	_docManage.Title = field.NewString(tableName, "title")
	_docManage.Content = field.NewString(tableName, "content")
	_docManage.Reading = field.NewUint64(tableName, "reading")
	_docManage.Type = field.NewInt(tableName, "type")
	_docManage.Label = field.NewString(tableName, "label")
	_docManage.CreatedAt = field.NewInt64(tableName, "created_at")
	_docManage.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_docManage.fillFieldMap()

	return _docManage
}

type docManage struct {
	docManageDo docManageDo

	ALL       field.Field
	Id        field.Uint64
	Title     field.String
	Content   field.String
	Reading   field.Uint64
	Type      field.Int
	Label     field.String
	CreatedAt field.Int64
	UpdatedAt field.Int64

	fieldMap map[string]field.Expr
}

func (d docManage) Table(newTableName string) *docManage {
	d.docManageDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d docManage) As(alias string) *docManage {
	d.docManageDo.DO = *(d.docManageDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *docManage) updateTableName(table string) *docManage {
	d.ALL = field.NewField(table, "*")
	d.Id = field.NewUint64(table, "id")
	d.Title = field.NewString(table, "title")
	d.Content = field.NewString(table, "content")
	d.Reading = field.NewUint64(table, "reading")
	d.Type = field.NewInt(table, "type")
	d.Label = field.NewString(table, "label")
	d.CreatedAt = field.NewInt64(table, "created_at")
	d.UpdatedAt = field.NewInt64(table, "updated_at")

	d.fillFieldMap()

	return d
}

func (d *docManage) WithContext(ctx context.Context) *docManageDo {
	return d.docManageDo.WithContext(ctx)
}

func (d docManage) TableName() string { return d.docManageDo.TableName() }

func (d docManage) Alias() string { return d.docManageDo.Alias() }

func (d *docManage) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *docManage) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 8)
	d.fieldMap["id"] = d.Id
	d.fieldMap["title"] = d.Title
	d.fieldMap["content"] = d.Content
	d.fieldMap["reading"] = d.Reading
	d.fieldMap["type"] = d.Type
	d.fieldMap["label"] = d.Label
	d.fieldMap["created_at"] = d.CreatedAt
	d.fieldMap["updated_at"] = d.UpdatedAt
}

func (d docManage) clone(db *gorm.DB) docManage {
	d.docManageDo.ReplaceDB(db)
	return d
}

type docManageDo struct{ gen.DO }

func (d docManageDo) Debug() *docManageDo {
	return d.withDO(d.DO.Debug())
}

func (d docManageDo) WithContext(ctx context.Context) *docManageDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d docManageDo) Clauses(conds ...clause.Expression) *docManageDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d docManageDo) Returning(value interface{}, columns ...string) *docManageDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d docManageDo) Not(conds ...gen.Condition) *docManageDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d docManageDo) Or(conds ...gen.Condition) *docManageDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d docManageDo) Select(conds ...field.Expr) *docManageDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d docManageDo) Where(conds ...gen.Condition) *docManageDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d docManageDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *docManageDo {
	return d.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (d docManageDo) Order(conds ...field.Expr) *docManageDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d docManageDo) Distinct(cols ...field.Expr) *docManageDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d docManageDo) Omit(cols ...field.Expr) *docManageDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d docManageDo) Join(table schema.Tabler, on ...field.Expr) *docManageDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d docManageDo) LeftJoin(table schema.Tabler, on ...field.Expr) *docManageDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d docManageDo) RightJoin(table schema.Tabler, on ...field.Expr) *docManageDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d docManageDo) Group(cols ...field.Expr) *docManageDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d docManageDo) Having(conds ...gen.Condition) *docManageDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d docManageDo) Limit(limit int) *docManageDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d docManageDo) Offset(offset int) *docManageDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d docManageDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *docManageDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d docManageDo) Unscoped() *docManageDo {
	return d.withDO(d.DO.Unscoped())
}

func (d docManageDo) Create(values ...*model.DocManage) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d docManageDo) CreateInBatches(values []*model.DocManage, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d docManageDo) Save(values ...*model.DocManage) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d docManageDo) First() (*model.DocManage, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DocManage), nil
	}
}

func (d docManageDo) Take() (*model.DocManage, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DocManage), nil
	}
}

func (d docManageDo) Last() (*model.DocManage, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DocManage), nil
	}
}

func (d docManageDo) Find() ([]*model.DocManage, error) {
	result, err := d.DO.Find()
	return result.([]*model.DocManage), err
}

func (d docManageDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DocManage, err error) {
	buf := make([]*model.DocManage, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d docManageDo) FindInBatches(result *[]*model.DocManage, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d docManageDo) Attrs(attrs ...field.AssignExpr) *docManageDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d docManageDo) Assign(attrs ...field.AssignExpr) *docManageDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d docManageDo) Joins(fields ...field.RelationField) *docManageDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d docManageDo) Preload(fields ...field.RelationField) *docManageDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d docManageDo) FirstOrInit() (*model.DocManage, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DocManage), nil
	}
}

func (d docManageDo) FirstOrCreate() (*model.DocManage, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DocManage), nil
	}
}

func (d docManageDo) FindByPage(offset int, limit int) (result []*model.DocManage, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d docManageDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d *docManageDo) withDO(do gen.Dao) *docManageDo {
	d.DO = *do.(*gen.DO)
	return d
}
