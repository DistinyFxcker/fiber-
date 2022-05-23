// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"fmt"
	"testing"

	"go_web_example/gen/dal/model"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := db.AutoMigrate(&model.PrimaryDomain{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.PrimaryDomain{}) fail: %s", err)
	}
}

func Test_primaryDomainQuery(t *testing.T) {
	primaryDomain := newPrimaryDomain(db)
	primaryDomain = *primaryDomain.As(primaryDomain.TableName())
	do := primaryDomain.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(primaryDomain.TableName(), clause.PrimaryKey)
	_, err := do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <primary_domain> fail:", err)
		return
	}

	_, ok := primaryDomain.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from primaryDomain success")
	}

	err = do.Create(&model.PrimaryDomain{})
	if err != nil {
		t.Error("create item in table <primary_domain> fail:", err)
	}

	err = do.Save(&model.PrimaryDomain{})
	if err != nil {
		t.Error("create item in table <primary_domain> fail:", err)
	}

	err = do.CreateInBatches([]*model.PrimaryDomain{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <primary_domain> fail:", err)
	}

	_, err = do.Select(primaryDomain.ALL).Take()
	if err != nil {
		t.Error("Take() on table <primary_domain> fail:", err)
	}

	_, err = do.First()
	if err != nil {
		t.Error("First() on table <primary_domain> fail:", err)
	}

	_, err = do.Last()
	if err != nil {
		t.Error("First() on table <primary_domain> fail:", err)
	}

	_, err = do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <primary_domain> fail:", err)
	}

	err = do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.PrimaryDomain{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <primary_domain> fail:", err)
	}

	_, err = do.Select(primaryDomain.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <primary_domain> fail:", err)
	}

	_, err = do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <primary_domain> fail:", err)
	}

	_, err = do.Select(primaryDomain.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <primary_domain> fail:", err)
	}

	_, err = do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <primary_domain> fail:", err)
	}

	_, err = do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <primary_domain> fail:", err)
	}

	_, _, err = do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <primary_domain> fail:", err)
	}

	_, err = do.ScanByPage(&model.PrimaryDomain{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <primary_domain> fail:", err)
	}

	_, err = do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <primary_domain> fail:", err)
	}

	_, err = do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <primary_domain> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), clause.PrimaryKey)

	err = do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <primary_domain> fail:", err)
	}

	err = do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <primary_domain> fail:", err)
	}

	_, err = do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <primary_domain> fail:", err)
	}
}