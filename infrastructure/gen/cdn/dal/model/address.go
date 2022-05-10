package model

import (
	jsoniter "github.com/json-iterator/go"
	"gorm.io/gorm"
	"io/ioutil"
)

type Address struct {
	Id        string `gorm:"column:id;type:varchar(50);primary_key;not null"`                   //id
	ParentId  string `gorm:"column:parent_id;type:varchar(50);comment:父级ID"`                    //父级ID
	CityCode  string `gorm:"column:city_code;type:varchar(50);comment:城市ID"`                    //城市ID
	Name      string `gorm:"column:name;type:varchar(50);comment:城市名称"`                         //城市名称
	Level     int    `gorm:"column:level;type:int;comment:等级"`                                  //等级
	CreatedAt int64  `gorm:"column:created_at;type:int8;index:idx_address_created_at;not null"` //创建时间
	UpdatedAt int64  `gorm:"column:updated_at;type:int8;index:idx_address_updated_at;not null"` //修改时间，新记录保持和创建时间一致
}

func (This *Address) TableName() string {
	return "address"
}

type Data struct {
	Data []Info
}

type Info struct {
	AddressInfo
	Districts []AddressInfo
}

type AddressInfo struct {
	Id       string
	ParentId string
	CityCode string
	Name     string
}

func (m *Address) Init(db *gorm.DB) error {
	//自动化迁移
	if err := db.AutoMigrate(&Address{}); err != nil {
		panic(err)
	}
	var Count int64
	db = db.Model(&Address{}).Count(&Count)
	if Count == 0 {
		var (
			bytes []byte
			err   error
		)

		if bytes, err = ioutil.ReadFile("region.json"); err != nil {
			return err
		}
		var Data Data
		err = jsoniter.Unmarshal(bytes, &Data)
		if err != nil {
			return err
		}
		db = db.Begin()
		for _, AddressInfo := range Data.Data {
			//添加省份
			if err := db.Model(&Address{}).Create(&Address{
				Id:       AddressInfo.Id,
				ParentId: AddressInfo.ParentId,
				CityCode: AddressInfo.CityCode,
				Name:     AddressInfo.Name,
				Level:    1,
			}).Error; err != nil {
				db.Rollback()
				return err
			}

			//添加省份下面的城市
			for _, CitiInfo := range AddressInfo.Districts {
				if err := db.Model(&Address{}).Create(&Address{
					Id:       CitiInfo.Id,
					ParentId: CitiInfo.ParentId,
					CityCode: CitiInfo.CityCode,
					Name:     CitiInfo.Name,
					Level:    2,
				}).Error; err != nil {
					db.Rollback()
					return err
				}
			}
		}
		db.Commit()
	}
	return nil
}
