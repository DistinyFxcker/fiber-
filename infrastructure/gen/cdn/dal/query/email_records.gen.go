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

func newEmailRecord(db *gorm.DB) emailRecord {
	_emailRecord := emailRecord{}

	_emailRecord.emailRecordDo.UseDB(db)
	_emailRecord.emailRecordDo.UseModel(&model.EmailRecord{})

	tableName := _emailRecord.emailRecordDo.TableName()
	_emailRecord.ALL = field.NewField(tableName, "*")
	_emailRecord.Id = field.NewString(tableName, "id")
	_emailRecord.Uid = field.NewString(tableName, "uid")
	_emailRecord.Type = field.NewInt(tableName, "type")
	_emailRecord.OperatorId = field.NewString(tableName, "operatorId")
	_emailRecord.Mail = field.NewString(tableName, "ip")
	_emailRecord.CreatedAt = field.NewInt64(tableName, "created_at")

	_emailRecord.fillFieldMap()

	return _emailRecord
}

type emailRecord struct {
	emailRecordDo emailRecordDo

	ALL        field.Field
	Id         field.String
	Uid        field.String
	Type       field.Int
	OperatorId field.String
	Mail       field.String
	CreatedAt  field.Int64

	fieldMap map[string]field.Expr
}

func (e emailRecord) Table(newTableName string) *emailRecord {
	e.emailRecordDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e emailRecord) As(alias string) *emailRecord {
	e.emailRecordDo.DO = *(e.emailRecordDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *emailRecord) updateTableName(table string) *emailRecord {
	e.ALL = field.NewField(table, "*")
	e.Id = field.NewString(table, "id")
	e.Uid = field.NewString(table, "uid")
	e.Type = field.NewInt(table, "type")
	e.OperatorId = field.NewString(table, "operatorId")
	e.Mail = field.NewString(table, "ip")
	e.CreatedAt = field.NewInt64(table, "created_at")

	e.fillFieldMap()

	return e
}

func (e *emailRecord) WithContext(ctx context.Context) *emailRecordDo {
	return e.emailRecordDo.WithContext(ctx)
}

func (e emailRecord) TableName() string { return e.emailRecordDo.TableName() }

func (e emailRecord) Alias() string { return e.emailRecordDo.Alias() }

func (e *emailRecord) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *emailRecord) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 6)
	e.fieldMap["id"] = e.Id
	e.fieldMap["uid"] = e.Uid
	e.fieldMap["type"] = e.Type
	e.fieldMap["operatorId"] = e.OperatorId
	e.fieldMap["ip"] = e.Mail
	e.fieldMap["created_at"] = e.CreatedAt
}

func (e emailRecord) clone(db *gorm.DB) emailRecord {
	e.emailRecordDo.ReplaceDB(db)
	return e
}

type emailRecordDo struct{ gen.DO }

func (e emailRecordDo) Debug() *emailRecordDo {
	return e.withDO(e.DO.Debug())
}

func (e emailRecordDo) WithContext(ctx context.Context) *emailRecordDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e emailRecordDo) Clauses(conds ...clause.Expression) *emailRecordDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e emailRecordDo) Returning(value interface{}, columns ...string) *emailRecordDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e emailRecordDo) Not(conds ...gen.Condition) *emailRecordDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e emailRecordDo) Or(conds ...gen.Condition) *emailRecordDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e emailRecordDo) Select(conds ...field.Expr) *emailRecordDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e emailRecordDo) Where(conds ...gen.Condition) *emailRecordDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e emailRecordDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *emailRecordDo {
	return e.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (e emailRecordDo) Order(conds ...field.Expr) *emailRecordDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e emailRecordDo) Distinct(cols ...field.Expr) *emailRecordDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e emailRecordDo) Omit(cols ...field.Expr) *emailRecordDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e emailRecordDo) Join(table schema.Tabler, on ...field.Expr) *emailRecordDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e emailRecordDo) LeftJoin(table schema.Tabler, on ...field.Expr) *emailRecordDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e emailRecordDo) RightJoin(table schema.Tabler, on ...field.Expr) *emailRecordDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e emailRecordDo) Group(cols ...field.Expr) *emailRecordDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e emailRecordDo) Having(conds ...gen.Condition) *emailRecordDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e emailRecordDo) Limit(limit int) *emailRecordDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e emailRecordDo) Offset(offset int) *emailRecordDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e emailRecordDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *emailRecordDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e emailRecordDo) Unscoped() *emailRecordDo {
	return e.withDO(e.DO.Unscoped())
}

func (e emailRecordDo) Create(values ...*model.EmailRecord) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e emailRecordDo) CreateInBatches(values []*model.EmailRecord, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e emailRecordDo) Save(values ...*model.EmailRecord) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e emailRecordDo) First() (*model.EmailRecord, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.EmailRecord), nil
	}
}

func (e emailRecordDo) Take() (*model.EmailRecord, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.EmailRecord), nil
	}
}

func (e emailRecordDo) Last() (*model.EmailRecord, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.EmailRecord), nil
	}
}

func (e emailRecordDo) Find() ([]*model.EmailRecord, error) {
	result, err := e.DO.Find()
	return result.([]*model.EmailRecord), err
}

func (e emailRecordDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EmailRecord, err error) {
	buf := make([]*model.EmailRecord, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e emailRecordDo) FindInBatches(result *[]*model.EmailRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e emailRecordDo) Attrs(attrs ...field.AssignExpr) *emailRecordDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e emailRecordDo) Assign(attrs ...field.AssignExpr) *emailRecordDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e emailRecordDo) Joins(fields ...field.RelationField) *emailRecordDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e emailRecordDo) Preload(fields ...field.RelationField) *emailRecordDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e emailRecordDo) FirstOrInit() (*model.EmailRecord, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.EmailRecord), nil
	}
}

func (e emailRecordDo) FirstOrCreate() (*model.EmailRecord, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.EmailRecord), nil
	}
}

func (e emailRecordDo) FindByPage(offset int, limit int) (result []*model.EmailRecord, count int64, err error) {
	result, err = e.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = e.Offset(-1).Limit(-1).Count()
	return
}

func (e emailRecordDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e *emailRecordDo) withDO(do gen.Dao) *emailRecordDo {
	e.DO = *do.(*gen.DO)
	return e
}