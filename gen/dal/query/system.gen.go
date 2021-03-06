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

func newSystem(db *gorm.DB) system {
	_system := system{}

	_system.systemDo.UseDB(db)
	_system.systemDo.UseModel(&model.System{})

	tableName := _system.systemDo.TableName()
	_system.ALL = field.NewField(tableName, "*")
	_system.Id = field.NewInt(tableName, "id")
	_system.Key = field.NewString(tableName, "key")
	_system.Value = field.NewString(tableName, "value")

	_system.fillFieldMap()

	return _system
}

type system struct {
	systemDo systemDo

	ALL   field.Field
	Id    field.Int
	Key   field.String
	Value field.String

	fieldMap map[string]field.Expr
}

func (s system) Table(newTableName string) *system {
	s.systemDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s system) As(alias string) *system {
	s.systemDo.DO = *(s.systemDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *system) updateTableName(table string) *system {
	s.ALL = field.NewField(table, "*")
	s.Id = field.NewInt(table, "id")
	s.Key = field.NewString(table, "key")
	s.Value = field.NewString(table, "value")

	s.fillFieldMap()

	return s
}

func (s *system) WithContext(ctx context.Context) *systemDo { return s.systemDo.WithContext(ctx) }

func (s system) TableName() string { return s.systemDo.TableName() }

func (s system) Alias() string { return s.systemDo.Alias() }

func (s *system) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *system) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 3)
	s.fieldMap["id"] = s.Id
	s.fieldMap["key"] = s.Key
	s.fieldMap["value"] = s.Value
}

func (s system) clone(db *gorm.DB) system {
	s.systemDo.ReplaceDB(db)
	return s
}

type systemDo struct{ gen.DO }

func (s systemDo) Debug() *systemDo {
	return s.withDO(s.DO.Debug())
}

func (s systemDo) WithContext(ctx context.Context) *systemDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s systemDo) Clauses(conds ...clause.Expression) *systemDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s systemDo) Returning(value interface{}, columns ...string) *systemDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s systemDo) Not(conds ...gen.Condition) *systemDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s systemDo) Or(conds ...gen.Condition) *systemDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s systemDo) Select(conds ...field.Expr) *systemDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s systemDo) Where(conds ...gen.Condition) *systemDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s systemDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *systemDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s systemDo) Order(conds ...field.Expr) *systemDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s systemDo) Distinct(cols ...field.Expr) *systemDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s systemDo) Omit(cols ...field.Expr) *systemDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s systemDo) Join(table schema.Tabler, on ...field.Expr) *systemDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s systemDo) LeftJoin(table schema.Tabler, on ...field.Expr) *systemDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s systemDo) RightJoin(table schema.Tabler, on ...field.Expr) *systemDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s systemDo) Group(cols ...field.Expr) *systemDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s systemDo) Having(conds ...gen.Condition) *systemDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s systemDo) Limit(limit int) *systemDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s systemDo) Offset(offset int) *systemDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s systemDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *systemDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s systemDo) Unscoped() *systemDo {
	return s.withDO(s.DO.Unscoped())
}

func (s systemDo) Create(values ...*model.System) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s systemDo) CreateInBatches(values []*model.System, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s systemDo) Save(values ...*model.System) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s systemDo) First() (*model.System, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.System), nil
	}
}

func (s systemDo) Take() (*model.System, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.System), nil
	}
}

func (s systemDo) Last() (*model.System, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.System), nil
	}
}

func (s systemDo) Find() ([]*model.System, error) {
	result, err := s.DO.Find()
	return result.([]*model.System), err
}

func (s systemDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.System, err error) {
	buf := make([]*model.System, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s systemDo) FindInBatches(result *[]*model.System, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s systemDo) Attrs(attrs ...field.AssignExpr) *systemDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s systemDo) Assign(attrs ...field.AssignExpr) *systemDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s systemDo) Joins(fields ...field.RelationField) *systemDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s systemDo) Preload(fields ...field.RelationField) *systemDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s systemDo) FirstOrInit() (*model.System, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.System), nil
	}
}

func (s systemDo) FirstOrCreate() (*model.System, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.System), nil
	}
}

func (s systemDo) FindByPage(offset int, limit int) (result []*model.System, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s systemDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s *systemDo) withDO(do gen.Dao) *systemDo {
	s.DO = *do.(*gen.DO)
	return s
}
