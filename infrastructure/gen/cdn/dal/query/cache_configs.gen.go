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

func newCacheConfig(db *gorm.DB) cacheConfig {
	_cacheConfig := cacheConfig{}

	_cacheConfig.cacheConfigDo.UseDB(db)
	_cacheConfig.cacheConfigDo.UseModel(&model.CacheConfig{})

	tableName := _cacheConfig.cacheConfigDo.TableName()
	_cacheConfig.ALL = field.NewField(tableName, "*")
	_cacheConfig.Id = field.NewString(tableName, "id")
	_cacheConfig.ConfigureId = field.NewString(tableName, "configure_id")
	_cacheConfig.RelationId = field.NewString(tableName, "relation_id")
	_cacheConfig.RelationType = field.NewInt(tableName, "relation_type")
	_cacheConfig.Name = field.NewString(tableName, "name")
	_cacheConfig.Description = field.NewString(tableName, "description")
	_cacheConfig.Type = field.NewInt(tableName, "type")
	_cacheConfig.TypeContext = field.NewString(tableName, "type_context")
	_cacheConfig.PriorityLevel = field.NewInt(tableName, "priority_level")
	_cacheConfig.CacheTimeV2 = field.NewString(tableName, "cache_time_v2")
	_cacheConfig.CreatedAt = field.NewInt64(tableName, "created_at")
	_cacheConfig.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_cacheConfig.fillFieldMap()

	return _cacheConfig
}

type cacheConfig struct {
	cacheConfigDo cacheConfigDo

	ALL           field.Field
	Id            field.String
	ConfigureId   field.String
	RelationId    field.String
	RelationType  field.Int
	Name          field.String
	Description   field.String
	Type          field.Int
	TypeContext   field.String
	PriorityLevel field.Int
	CacheTimeV2   field.String
	CreatedAt     field.Int64
	UpdatedAt     field.Int64

	fieldMap map[string]field.Expr
}

func (c cacheConfig) Table(newTableName string) *cacheConfig {
	c.cacheConfigDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cacheConfig) As(alias string) *cacheConfig {
	c.cacheConfigDo.DO = *(c.cacheConfigDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cacheConfig) updateTableName(table string) *cacheConfig {
	c.ALL = field.NewField(table, "*")
	c.Id = field.NewString(table, "id")
	c.ConfigureId = field.NewString(table, "configure_id")
	c.RelationId = field.NewString(table, "relation_id")
	c.RelationType = field.NewInt(table, "relation_type")
	c.Name = field.NewString(table, "name")
	c.Description = field.NewString(table, "description")
	c.Type = field.NewInt(table, "type")
	c.TypeContext = field.NewString(table, "type_context")
	c.PriorityLevel = field.NewInt(table, "priority_level")
	c.CacheTimeV2 = field.NewString(table, "cache_time_v2")
	c.CreatedAt = field.NewInt64(table, "created_at")
	c.UpdatedAt = field.NewInt64(table, "updated_at")

	c.fillFieldMap()

	return c
}

func (c *cacheConfig) WithContext(ctx context.Context) *cacheConfigDo {
	return c.cacheConfigDo.WithContext(ctx)
}

func (c cacheConfig) TableName() string { return c.cacheConfigDo.TableName() }

func (c cacheConfig) Alias() string { return c.cacheConfigDo.Alias() }

func (c *cacheConfig) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cacheConfig) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 12)
	c.fieldMap["id"] = c.Id
	c.fieldMap["configure_id"] = c.ConfigureId
	c.fieldMap["relation_id"] = c.RelationId
	c.fieldMap["relation_type"] = c.RelationType
	c.fieldMap["name"] = c.Name
	c.fieldMap["description"] = c.Description
	c.fieldMap["type"] = c.Type
	c.fieldMap["type_context"] = c.TypeContext
	c.fieldMap["priority_level"] = c.PriorityLevel
	c.fieldMap["cache_time_v2"] = c.CacheTimeV2
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
}

func (c cacheConfig) clone(db *gorm.DB) cacheConfig {
	c.cacheConfigDo.ReplaceDB(db)
	return c
}

type cacheConfigDo struct{ gen.DO }

func (c cacheConfigDo) Debug() *cacheConfigDo {
	return c.withDO(c.DO.Debug())
}

func (c cacheConfigDo) WithContext(ctx context.Context) *cacheConfigDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c cacheConfigDo) Clauses(conds ...clause.Expression) *cacheConfigDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c cacheConfigDo) Returning(value interface{}, columns ...string) *cacheConfigDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c cacheConfigDo) Not(conds ...gen.Condition) *cacheConfigDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c cacheConfigDo) Or(conds ...gen.Condition) *cacheConfigDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c cacheConfigDo) Select(conds ...field.Expr) *cacheConfigDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c cacheConfigDo) Where(conds ...gen.Condition) *cacheConfigDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c cacheConfigDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *cacheConfigDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c cacheConfigDo) Order(conds ...field.Expr) *cacheConfigDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c cacheConfigDo) Distinct(cols ...field.Expr) *cacheConfigDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c cacheConfigDo) Omit(cols ...field.Expr) *cacheConfigDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c cacheConfigDo) Join(table schema.Tabler, on ...field.Expr) *cacheConfigDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c cacheConfigDo) LeftJoin(table schema.Tabler, on ...field.Expr) *cacheConfigDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c cacheConfigDo) RightJoin(table schema.Tabler, on ...field.Expr) *cacheConfigDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c cacheConfigDo) Group(cols ...field.Expr) *cacheConfigDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c cacheConfigDo) Having(conds ...gen.Condition) *cacheConfigDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c cacheConfigDo) Limit(limit int) *cacheConfigDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c cacheConfigDo) Offset(offset int) *cacheConfigDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c cacheConfigDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *cacheConfigDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c cacheConfigDo) Unscoped() *cacheConfigDo {
	return c.withDO(c.DO.Unscoped())
}

func (c cacheConfigDo) Create(values ...*model.CacheConfig) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c cacheConfigDo) CreateInBatches(values []*model.CacheConfig, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c cacheConfigDo) Save(values ...*model.CacheConfig) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c cacheConfigDo) First() (*model.CacheConfig, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CacheConfig), nil
	}
}

func (c cacheConfigDo) Take() (*model.CacheConfig, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CacheConfig), nil
	}
}

func (c cacheConfigDo) Last() (*model.CacheConfig, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CacheConfig), nil
	}
}

func (c cacheConfigDo) Find() ([]*model.CacheConfig, error) {
	result, err := c.DO.Find()
	return result.([]*model.CacheConfig), err
}

func (c cacheConfigDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CacheConfig, err error) {
	buf := make([]*model.CacheConfig, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c cacheConfigDo) FindInBatches(result *[]*model.CacheConfig, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c cacheConfigDo) Attrs(attrs ...field.AssignExpr) *cacheConfigDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c cacheConfigDo) Assign(attrs ...field.AssignExpr) *cacheConfigDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c cacheConfigDo) Joins(fields ...field.RelationField) *cacheConfigDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c cacheConfigDo) Preload(fields ...field.RelationField) *cacheConfigDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c cacheConfigDo) FirstOrInit() (*model.CacheConfig, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CacheConfig), nil
	}
}

func (c cacheConfigDo) FirstOrCreate() (*model.CacheConfig, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CacheConfig), nil
	}
}

func (c cacheConfigDo) FindByPage(offset int, limit int) (result []*model.CacheConfig, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c cacheConfigDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c *cacheConfigDo) withDO(do gen.Dao) *cacheConfigDo {
	c.DO = *do.(*gen.DO)
	return c
}
