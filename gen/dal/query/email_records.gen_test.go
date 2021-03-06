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
	err := db.AutoMigrate(&model.EmailRecord{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.EmailRecord{}) fail: %s", err)
	}
}

func Test_emailRecordQuery(t *testing.T) {
	emailRecord := newEmailRecord(db)
	emailRecord = *emailRecord.As(emailRecord.TableName())
	do := emailRecord.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(emailRecord.TableName(), clause.PrimaryKey)
	_, err := do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <email_records> fail:", err)
		return
	}

	_, ok := emailRecord.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from emailRecord success")
	}

	err = do.Create(&model.EmailRecord{})
	if err != nil {
		t.Error("create item in table <email_records> fail:", err)
	}

	err = do.Save(&model.EmailRecord{})
	if err != nil {
		t.Error("create item in table <email_records> fail:", err)
	}

	err = do.CreateInBatches([]*model.EmailRecord{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <email_records> fail:", err)
	}

	_, err = do.Select(emailRecord.ALL).Take()
	if err != nil {
		t.Error("Take() on table <email_records> fail:", err)
	}

	_, err = do.First()
	if err != nil {
		t.Error("First() on table <email_records> fail:", err)
	}

	_, err = do.Last()
	if err != nil {
		t.Error("First() on table <email_records> fail:", err)
	}

	_, err = do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <email_records> fail:", err)
	}

	err = do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.EmailRecord{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <email_records> fail:", err)
	}

	_, err = do.Select(emailRecord.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <email_records> fail:", err)
	}

	_, err = do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <email_records> fail:", err)
	}

	_, err = do.Select(emailRecord.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <email_records> fail:", err)
	}

	_, err = do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <email_records> fail:", err)
	}

	_, err = do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <email_records> fail:", err)
	}

	_, _, err = do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <email_records> fail:", err)
	}

	_, err = do.ScanByPage(&model.EmailRecord{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <email_records> fail:", err)
	}

	_, err = do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <email_records> fail:", err)
	}

	_, err = do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <email_records> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), clause.PrimaryKey)

	err = do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <email_records> fail:", err)
	}

	err = do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <email_records> fail:", err)
	}

	_, err = do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <email_records> fail:", err)
	}
}
