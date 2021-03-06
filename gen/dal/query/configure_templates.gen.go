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

func newConfigureTemplate(db *gorm.DB) configureTemplate {
	_configureTemplate := configureTemplate{}

	_configureTemplate.configureTemplateDo.UseDB(db)
	_configureTemplate.configureTemplateDo.UseModel(&model.ConfigureTemplate{})

	tableName := _configureTemplate.configureTemplateDo.TableName()
	_configureTemplate.ALL = field.NewField(tableName, "*")
	_configureTemplate.Id = field.NewString(tableName, "id")
	_configureTemplate.Name = field.NewString(tableName, "name")
	_configureTemplate.RelationName = field.NewString(tableName, "relation_name")
	_configureTemplate.RelationId = field.NewString(tableName, "relation_id")
	_configureTemplate.RelationType = field.NewInt(tableName, "relation_type")
	_configureTemplate.CreatedId = field.NewString(tableName, "created_id")
	_configureTemplate.CreatedRole = field.NewInt(tableName, "created_role")
	_configureTemplate.Type = field.NewInt(tableName, "type")
	_configureTemplate.ConnectTime = field.NewInt(tableName, "connect_time")
	_configureTemplate.MinConnNum = field.NewInt(tableName, "min_conn_num")
	_configureTemplate.MaxMemoryCache = field.NewInt(tableName, "max_memory_cache")
	_configureTemplate.MaxDiskCache = field.NewInt(tableName, "max_disk_cache")
	_configureTemplate.DiskCacheType = field.NewInt(tableName, "disk_cache_type")
	_configureTemplate.CompressionGrade = field.NewInt(tableName, "compression_grade")
	_configureTemplate.CompressionType = field.NewString(tableName, "compression_type")
	_configureTemplate.CompressionSizeUpper = field.NewInt(tableName, "compression_size_upper")
	_configureTemplate.CompressionSizeLower = field.NewInt(tableName, "compression_size_lower")
	_configureTemplate.CompressionUpperType = field.NewInt(tableName, "compression_upper_type")
	_configureTemplate.CompressionLowerType = field.NewInt(tableName, "compression_lower_type")
	_configureTemplate.FrequencyLimit = field.NewInt(tableName, "frequency_limit")
	_configureTemplate.MinuteSetting = field.NewInt(tableName, "minute_setting")
	_configureTemplate.MinuteTotalTime = field.NewInt(tableName, "minute_total_time")
	_configureTemplate.ConnectNum = field.NewInt(tableName, "connect_num")
	_configureTemplate.LocationInfo = field.NewString(tableName, "location_info")
	_configureTemplate.Remark = field.NewString(tableName, "remark")
	_configureTemplate.BusinessType = field.NewString(tableName, "business_type")
	_configureTemplate.DomainSourceSortType = field.NewInt(tableName, "domain_sourceSortType")
	_configureTemplate.DomainSourceType = field.NewInt(tableName, "domain_source_type")
	_configureTemplate.DomainSourceList = field.NewString(tableName, "domain_source_list")
	_configureTemplate.DomainForward = field.NewInt(tableName, "domain_forward")
	_configureTemplate.DomainForwardHost = field.NewString(tableName, "domain_forward_host")
	_configureTemplate.DomainForwardHeader = field.NewString(tableName, "domain_forward_header")
	_configureTemplate.CacheConfig = field.NewString(tableName, "cache_config")
	_configureTemplate.CacheSource = field.NewInt(tableName, "cache_source")
	_configureTemplate.HeaderList = field.NewString(tableName, "header_list")
	_configureTemplate.ErrorPageList = field.NewString(tableName, "error_page_list")
	_configureTemplate.DomainAgreementType = field.NewInt(tableName, "domain_agreement_type")
	_configureTemplate.Agreement = field.NewString(tableName, "agreement")
	_configureTemplate.EmptyReferer = field.NewInt(tableName, "empty_referer")
	_configureTemplate.BlackListDomain = field.NewString(tableName, "black_list_domain")
	_configureTemplate.WhiteListDomain = field.NewString(tableName, "white_list_domain")
	_configureTemplate.AccessControl = field.NewInt(tableName, "access_control")
	_configureTemplate.AntiTheftChainSwitch = field.NewInt(tableName, "anti_theft_chain_switch")
	_configureTemplate.BlackListSwitch = field.NewInt(tableName, "black_list_switch")
	_configureTemplate.WhiteListSwitch = field.NewInt(tableName, "white_list_switch")
	_configureTemplate.BlackList = field.NewString(tableName, "black_list")
	_configureTemplate.WhiteList = field.NewString(tableName, "white_list")
	_configureTemplate.DomainForwardList = field.NewString(tableName, "domain_forward_list")
	_configureTemplate.IgnoreURL = field.NewInt(tableName, "ignore_url")
	_configureTemplate.GzipCompress = field.NewInt(tableName, "gzip_compress")
	_configureTemplate.GzipCompressSwitch = field.NewInt(tableName, "gzip_compress_switch")
	_configureTemplate.CreatedAt = field.NewInt64(tableName, "created_at")
	_configureTemplate.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_configureTemplate.fillFieldMap()

	return _configureTemplate
}

type configureTemplate struct {
	configureTemplateDo configureTemplateDo

	ALL                  field.Field
	Id                   field.String
	Name                 field.String
	RelationName         field.String
	RelationId           field.String
	RelationType         field.Int
	CreatedId            field.String
	CreatedRole          field.Int
	Type                 field.Int
	ConnectTime          field.Int
	MinConnNum           field.Int
	MaxMemoryCache       field.Int
	MaxDiskCache         field.Int
	DiskCacheType        field.Int
	CompressionGrade     field.Int
	CompressionType      field.String
	CompressionSizeUpper field.Int
	CompressionSizeLower field.Int
	CompressionUpperType field.Int
	CompressionLowerType field.Int
	FrequencyLimit       field.Int
	MinuteSetting        field.Int
	MinuteTotalTime      field.Int
	ConnectNum           field.Int
	LocationInfo         field.String
	Remark               field.String
	BusinessType         field.String
	DomainSourceSortType field.Int
	DomainSourceType     field.Int
	DomainSourceList     field.String
	DomainForward        field.Int
	DomainForwardHost    field.String
	DomainForwardHeader  field.String
	CacheConfig          field.String
	CacheSource          field.Int
	HeaderList           field.String
	ErrorPageList        field.String
	DomainAgreementType  field.Int
	Agreement            field.String
	EmptyReferer         field.Int
	BlackListDomain      field.String
	WhiteListDomain      field.String
	AccessControl        field.Int
	AntiTheftChainSwitch field.Int
	BlackListSwitch      field.Int
	WhiteListSwitch      field.Int
	BlackList            field.String
	WhiteList            field.String
	DomainForwardList    field.String
	IgnoreURL            field.Int
	GzipCompress         field.Int
	GzipCompressSwitch   field.Int
	CreatedAt            field.Int64
	UpdatedAt            field.Int64

	fieldMap map[string]field.Expr
}

func (c configureTemplate) Table(newTableName string) *configureTemplate {
	c.configureTemplateDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c configureTemplate) As(alias string) *configureTemplate {
	c.configureTemplateDo.DO = *(c.configureTemplateDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *configureTemplate) updateTableName(table string) *configureTemplate {
	c.ALL = field.NewField(table, "*")
	c.Id = field.NewString(table, "id")
	c.Name = field.NewString(table, "name")
	c.RelationName = field.NewString(table, "relation_name")
	c.RelationId = field.NewString(table, "relation_id")
	c.RelationType = field.NewInt(table, "relation_type")
	c.CreatedId = field.NewString(table, "created_id")
	c.CreatedRole = field.NewInt(table, "created_role")
	c.Type = field.NewInt(table, "type")
	c.ConnectTime = field.NewInt(table, "connect_time")
	c.MinConnNum = field.NewInt(table, "min_conn_num")
	c.MaxMemoryCache = field.NewInt(table, "max_memory_cache")
	c.MaxDiskCache = field.NewInt(table, "max_disk_cache")
	c.DiskCacheType = field.NewInt(table, "disk_cache_type")
	c.CompressionGrade = field.NewInt(table, "compression_grade")
	c.CompressionType = field.NewString(table, "compression_type")
	c.CompressionSizeUpper = field.NewInt(table, "compression_size_upper")
	c.CompressionSizeLower = field.NewInt(table, "compression_size_lower")
	c.CompressionUpperType = field.NewInt(table, "compression_upper_type")
	c.CompressionLowerType = field.NewInt(table, "compression_lower_type")
	c.FrequencyLimit = field.NewInt(table, "frequency_limit")
	c.MinuteSetting = field.NewInt(table, "minute_setting")
	c.MinuteTotalTime = field.NewInt(table, "minute_total_time")
	c.ConnectNum = field.NewInt(table, "connect_num")
	c.LocationInfo = field.NewString(table, "location_info")
	c.Remark = field.NewString(table, "remark")
	c.BusinessType = field.NewString(table, "business_type")
	c.DomainSourceSortType = field.NewInt(table, "domain_sourceSortType")
	c.DomainSourceType = field.NewInt(table, "domain_source_type")
	c.DomainSourceList = field.NewString(table, "domain_source_list")
	c.DomainForward = field.NewInt(table, "domain_forward")
	c.DomainForwardHost = field.NewString(table, "domain_forward_host")
	c.DomainForwardHeader = field.NewString(table, "domain_forward_header")
	c.CacheConfig = field.NewString(table, "cache_config")
	c.CacheSource = field.NewInt(table, "cache_source")
	c.HeaderList = field.NewString(table, "header_list")
	c.ErrorPageList = field.NewString(table, "error_page_list")
	c.DomainAgreementType = field.NewInt(table, "domain_agreement_type")
	c.Agreement = field.NewString(table, "agreement")
	c.EmptyReferer = field.NewInt(table, "empty_referer")
	c.BlackListDomain = field.NewString(table, "black_list_domain")
	c.WhiteListDomain = field.NewString(table, "white_list_domain")
	c.AccessControl = field.NewInt(table, "access_control")
	c.AntiTheftChainSwitch = field.NewInt(table, "anti_theft_chain_switch")
	c.BlackListSwitch = field.NewInt(table, "black_list_switch")
	c.WhiteListSwitch = field.NewInt(table, "white_list_switch")
	c.BlackList = field.NewString(table, "black_list")
	c.WhiteList = field.NewString(table, "white_list")
	c.DomainForwardList = field.NewString(table, "domain_forward_list")
	c.IgnoreURL = field.NewInt(table, "ignore_url")
	c.GzipCompress = field.NewInt(table, "gzip_compress")
	c.GzipCompressSwitch = field.NewInt(table, "gzip_compress_switch")
	c.CreatedAt = field.NewInt64(table, "created_at")
	c.UpdatedAt = field.NewInt64(table, "updated_at")

	c.fillFieldMap()

	return c
}

func (c *configureTemplate) WithContext(ctx context.Context) *configureTemplateDo {
	return c.configureTemplateDo.WithContext(ctx)
}

func (c configureTemplate) TableName() string { return c.configureTemplateDo.TableName() }

func (c configureTemplate) Alias() string { return c.configureTemplateDo.Alias() }

func (c *configureTemplate) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *configureTemplate) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 53)
	c.fieldMap["id"] = c.Id
	c.fieldMap["name"] = c.Name
	c.fieldMap["relation_name"] = c.RelationName
	c.fieldMap["relation_id"] = c.RelationId
	c.fieldMap["relation_type"] = c.RelationType
	c.fieldMap["created_id"] = c.CreatedId
	c.fieldMap["created_role"] = c.CreatedRole
	c.fieldMap["type"] = c.Type
	c.fieldMap["connect_time"] = c.ConnectTime
	c.fieldMap["min_conn_num"] = c.MinConnNum
	c.fieldMap["max_memory_cache"] = c.MaxMemoryCache
	c.fieldMap["max_disk_cache"] = c.MaxDiskCache
	c.fieldMap["disk_cache_type"] = c.DiskCacheType
	c.fieldMap["compression_grade"] = c.CompressionGrade
	c.fieldMap["compression_type"] = c.CompressionType
	c.fieldMap["compression_size_upper"] = c.CompressionSizeUpper
	c.fieldMap["compression_size_lower"] = c.CompressionSizeLower
	c.fieldMap["compression_upper_type"] = c.CompressionUpperType
	c.fieldMap["compression_lower_type"] = c.CompressionLowerType
	c.fieldMap["frequency_limit"] = c.FrequencyLimit
	c.fieldMap["minute_setting"] = c.MinuteSetting
	c.fieldMap["minute_total_time"] = c.MinuteTotalTime
	c.fieldMap["connect_num"] = c.ConnectNum
	c.fieldMap["location_info"] = c.LocationInfo
	c.fieldMap["remark"] = c.Remark
	c.fieldMap["business_type"] = c.BusinessType
	c.fieldMap["domain_sourceSortType"] = c.DomainSourceSortType
	c.fieldMap["domain_source_type"] = c.DomainSourceType
	c.fieldMap["domain_source_list"] = c.DomainSourceList
	c.fieldMap["domain_forward"] = c.DomainForward
	c.fieldMap["domain_forward_host"] = c.DomainForwardHost
	c.fieldMap["domain_forward_header"] = c.DomainForwardHeader
	c.fieldMap["cache_config"] = c.CacheConfig
	c.fieldMap["cache_source"] = c.CacheSource
	c.fieldMap["header_list"] = c.HeaderList
	c.fieldMap["error_page_list"] = c.ErrorPageList
	c.fieldMap["domain_agreement_type"] = c.DomainAgreementType
	c.fieldMap["agreement"] = c.Agreement
	c.fieldMap["empty_referer"] = c.EmptyReferer
	c.fieldMap["black_list_domain"] = c.BlackListDomain
	c.fieldMap["white_list_domain"] = c.WhiteListDomain
	c.fieldMap["access_control"] = c.AccessControl
	c.fieldMap["anti_theft_chain_switch"] = c.AntiTheftChainSwitch
	c.fieldMap["black_list_switch"] = c.BlackListSwitch
	c.fieldMap["white_list_switch"] = c.WhiteListSwitch
	c.fieldMap["black_list"] = c.BlackList
	c.fieldMap["white_list"] = c.WhiteList
	c.fieldMap["domain_forward_list"] = c.DomainForwardList
	c.fieldMap["ignore_url"] = c.IgnoreURL
	c.fieldMap["gzip_compress"] = c.GzipCompress
	c.fieldMap["gzip_compress_switch"] = c.GzipCompressSwitch
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
}

func (c configureTemplate) clone(db *gorm.DB) configureTemplate {
	c.configureTemplateDo.ReplaceDB(db)
	return c
}

type configureTemplateDo struct{ gen.DO }

func (c configureTemplateDo) Debug() *configureTemplateDo {
	return c.withDO(c.DO.Debug())
}

func (c configureTemplateDo) WithContext(ctx context.Context) *configureTemplateDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c configureTemplateDo) Clauses(conds ...clause.Expression) *configureTemplateDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c configureTemplateDo) Returning(value interface{}, columns ...string) *configureTemplateDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c configureTemplateDo) Not(conds ...gen.Condition) *configureTemplateDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c configureTemplateDo) Or(conds ...gen.Condition) *configureTemplateDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c configureTemplateDo) Select(conds ...field.Expr) *configureTemplateDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c configureTemplateDo) Where(conds ...gen.Condition) *configureTemplateDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c configureTemplateDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *configureTemplateDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c configureTemplateDo) Order(conds ...field.Expr) *configureTemplateDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c configureTemplateDo) Distinct(cols ...field.Expr) *configureTemplateDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c configureTemplateDo) Omit(cols ...field.Expr) *configureTemplateDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c configureTemplateDo) Join(table schema.Tabler, on ...field.Expr) *configureTemplateDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c configureTemplateDo) LeftJoin(table schema.Tabler, on ...field.Expr) *configureTemplateDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c configureTemplateDo) RightJoin(table schema.Tabler, on ...field.Expr) *configureTemplateDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c configureTemplateDo) Group(cols ...field.Expr) *configureTemplateDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c configureTemplateDo) Having(conds ...gen.Condition) *configureTemplateDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c configureTemplateDo) Limit(limit int) *configureTemplateDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c configureTemplateDo) Offset(offset int) *configureTemplateDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c configureTemplateDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *configureTemplateDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c configureTemplateDo) Unscoped() *configureTemplateDo {
	return c.withDO(c.DO.Unscoped())
}

func (c configureTemplateDo) Create(values ...*model.ConfigureTemplate) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c configureTemplateDo) CreateInBatches(values []*model.ConfigureTemplate, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c configureTemplateDo) Save(values ...*model.ConfigureTemplate) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c configureTemplateDo) First() (*model.ConfigureTemplate, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ConfigureTemplate), nil
	}
}

func (c configureTemplateDo) Take() (*model.ConfigureTemplate, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ConfigureTemplate), nil
	}
}

func (c configureTemplateDo) Last() (*model.ConfigureTemplate, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ConfigureTemplate), nil
	}
}

func (c configureTemplateDo) Find() ([]*model.ConfigureTemplate, error) {
	result, err := c.DO.Find()
	return result.([]*model.ConfigureTemplate), err
}

func (c configureTemplateDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ConfigureTemplate, err error) {
	buf := make([]*model.ConfigureTemplate, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c configureTemplateDo) FindInBatches(result *[]*model.ConfigureTemplate, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c configureTemplateDo) Attrs(attrs ...field.AssignExpr) *configureTemplateDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c configureTemplateDo) Assign(attrs ...field.AssignExpr) *configureTemplateDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c configureTemplateDo) Joins(fields ...field.RelationField) *configureTemplateDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c configureTemplateDo) Preload(fields ...field.RelationField) *configureTemplateDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c configureTemplateDo) FirstOrInit() (*model.ConfigureTemplate, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ConfigureTemplate), nil
	}
}

func (c configureTemplateDo) FirstOrCreate() (*model.ConfigureTemplate, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ConfigureTemplate), nil
	}
}

func (c configureTemplateDo) FindByPage(offset int, limit int) (result []*model.ConfigureTemplate, count int64, err error) {
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

func (c configureTemplateDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c *configureTemplateDo) withDO(do gen.Dao) *configureTemplateDo {
	c.DO = *do.(*gen.DO)
	return c
}
