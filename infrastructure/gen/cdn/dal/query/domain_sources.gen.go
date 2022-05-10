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

func newDomainSource(db *gorm.DB) domainSource {
	_domainSource := domainSource{}

	_domainSource.domainSourceDo.UseDB(db)
	_domainSource.domainSourceDo.UseModel(&model.DomainSource{})

	tableName := _domainSource.domainSourceDo.TableName()
	_domainSource.ALL = field.NewField(tableName, "*")
	_domainSource.Id = field.NewString(tableName, "id")
	_domainSource.DomainId = field.NewString(tableName, "domain_id")
	_domainSource.CurrenNumber = field.NewInt(tableName, "curren_number")
	_domainSource.Context = field.NewString(tableName, "context")
	_domainSource.Port = field.NewInt(tableName, "port")
	_domainSource.SerialNumber = field.NewInt(tableName, "serial_number")
	_domainSource.CreatedAt = field.NewInt64(tableName, "created_at")
	_domainSource.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_domainSource.fillFieldMap()

	return _domainSource
}

type domainSource struct {
	domainSourceDo domainSourceDo

	ALL          field.Field
	Id           field.String
	DomainId     field.String
	CurrenNumber field.Int
	Context      field.String
	Port         field.Int
	SerialNumber field.Int
	CreatedAt    field.Int64
	UpdatedAt    field.Int64

	fieldMap map[string]field.Expr
}

func (d domainSource) Table(newTableName string) *domainSource {
	d.domainSourceDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d domainSource) As(alias string) *domainSource {
	d.domainSourceDo.DO = *(d.domainSourceDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *domainSource) updateTableName(table string) *domainSource {
	d.ALL = field.NewField(table, "*")
	d.Id = field.NewString(table, "id")
	d.DomainId = field.NewString(table, "domain_id")
	d.CurrenNumber = field.NewInt(table, "curren_number")
	d.Context = field.NewString(table, "context")
	d.Port = field.NewInt(table, "port")
	d.SerialNumber = field.NewInt(table, "serial_number")
	d.CreatedAt = field.NewInt64(table, "created_at")
	d.UpdatedAt = field.NewInt64(table, "updated_at")

	d.fillFieldMap()

	return d
}

func (d *domainSource) WithContext(ctx context.Context) *domainSourceDo {
	return d.domainSourceDo.WithContext(ctx)
}

func (d domainSource) TableName() string { return d.domainSourceDo.TableName() }

func (d domainSource) Alias() string { return d.domainSourceDo.Alias() }

func (d *domainSource) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *domainSource) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 8)
	d.fieldMap["id"] = d.Id
	d.fieldMap["domain_id"] = d.DomainId
	d.fieldMap["curren_number"] = d.CurrenNumber
	d.fieldMap["context"] = d.Context
	d.fieldMap["port"] = d.Port
	d.fieldMap["serial_number"] = d.SerialNumber
	d.fieldMap["created_at"] = d.CreatedAt
	d.fieldMap["updated_at"] = d.UpdatedAt
}

func (d domainSource) clone(db *gorm.DB) domainSource {
	d.domainSourceDo.ReplaceDB(db)
	return d
}

type domainSourceDo struct{ gen.DO }

func (d domainSourceDo) Debug() *domainSourceDo {
	return d.withDO(d.DO.Debug())
}

func (d domainSourceDo) WithContext(ctx context.Context) *domainSourceDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d domainSourceDo) Clauses(conds ...clause.Expression) *domainSourceDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d domainSourceDo) Returning(value interface{}, columns ...string) *domainSourceDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d domainSourceDo) Not(conds ...gen.Condition) *domainSourceDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d domainSourceDo) Or(conds ...gen.Condition) *domainSourceDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d domainSourceDo) Select(conds ...field.Expr) *domainSourceDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d domainSourceDo) Where(conds ...gen.Condition) *domainSourceDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d domainSourceDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *domainSourceDo {
	return d.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (d domainSourceDo) Order(conds ...field.Expr) *domainSourceDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d domainSourceDo) Distinct(cols ...field.Expr) *domainSourceDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d domainSourceDo) Omit(cols ...field.Expr) *domainSourceDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d domainSourceDo) Join(table schema.Tabler, on ...field.Expr) *domainSourceDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d domainSourceDo) LeftJoin(table schema.Tabler, on ...field.Expr) *domainSourceDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d domainSourceDo) RightJoin(table schema.Tabler, on ...field.Expr) *domainSourceDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d domainSourceDo) Group(cols ...field.Expr) *domainSourceDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d domainSourceDo) Having(conds ...gen.Condition) *domainSourceDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d domainSourceDo) Limit(limit int) *domainSourceDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d domainSourceDo) Offset(offset int) *domainSourceDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d domainSourceDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *domainSourceDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d domainSourceDo) Unscoped() *domainSourceDo {
	return d.withDO(d.DO.Unscoped())
}

func (d domainSourceDo) Create(values ...*model.DomainSource) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d domainSourceDo) CreateInBatches(values []*model.DomainSource, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d domainSourceDo) Save(values ...*model.DomainSource) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d domainSourceDo) First() (*model.DomainSource, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DomainSource), nil
	}
}

func (d domainSourceDo) Take() (*model.DomainSource, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DomainSource), nil
	}
}

func (d domainSourceDo) Last() (*model.DomainSource, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DomainSource), nil
	}
}

func (d domainSourceDo) Find() ([]*model.DomainSource, error) {
	result, err := d.DO.Find()
	return result.([]*model.DomainSource), err
}

func (d domainSourceDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DomainSource, err error) {
	buf := make([]*model.DomainSource, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d domainSourceDo) FindInBatches(result *[]*model.DomainSource, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d domainSourceDo) Attrs(attrs ...field.AssignExpr) *domainSourceDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d domainSourceDo) Assign(attrs ...field.AssignExpr) *domainSourceDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d domainSourceDo) Joins(fields ...field.RelationField) *domainSourceDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d domainSourceDo) Preload(fields ...field.RelationField) *domainSourceDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d domainSourceDo) FirstOrInit() (*model.DomainSource, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DomainSource), nil
	}
}

func (d domainSourceDo) FirstOrCreate() (*model.DomainSource, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DomainSource), nil
	}
}

func (d domainSourceDo) FindByPage(offset int, limit int) (result []*model.DomainSource, count int64, err error) {
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

func (d domainSourceDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d *domainSourceDo) withDO(do gen.Dao) *domainSourceDo {
	d.DO = *do.(*gen.DO)
	return d
}
