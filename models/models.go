package models

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/Alter/blog/pkg/setting"
	"time"

	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	Ctime int `json:"ctime"`
	Mtime int `json:"mtime"`
	DeletedOn int `json:"deleted_on"`

}
// Setup initializes the database instance
func Setup() {
	var err error
	db, err = gorm.Open(
		setting.DatabaseSetting.Type,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}

func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if scope.HasError() {
		return
	}
	nowTime := time.Now().Unix()
	if createTimeField, ok := scope.FieldByName("Ctime"); ok {
		if createTimeField.IsBlank {
			createTimeField.Set(nowTime)
		}
	}
	if modifyTimeField, ok := scope.FieldByName("Mtime"); ok {
		if modifyTimeField.IsBlank {
			modifyTimeField.Set(nowTime)
		}
	}

}

// scope.Get(...) 根据入参获取设置了字面值的参数，例如本文中是 gorm:update_column ，
// 它会去查找含这个字面值的字段属性
//cope.SetColumn(...) 假设没有指定 update_column 的字段，我们默认在更新回调设置 ModifiedOn 的值
// updateTimeStampForUpdateCallback will set `Mtime` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if scope.HasError() {
		return
	}
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("Mtime", time.Now().Unix())
	}
}

func deleteCallback(scope *gorm.Scope) {
	if scope.HasError() {
		return
	}
	var extraOption string
	if str, ok := scope.Get("gorm:delete_option"); ok {
		extraOption = fmt.Sprint(str)
	}
	deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")
	// scope.AddToVars 该方法可以添加值作为SQL的参数，也可用于防范SQL注入
	if !scope.Search.Unscoped && hasDeletedOnField {
		scope.Raw(fmt.Sprintf(
			"UPDATE %v SET %v=%v%v%v",
			scope.QuotedTableName(),
			scope.Quote(deletedOnField.DBName),
			scope.AddToVars(time.Now().Unix()),
			addExtraSpaceIfExist(scope.CombinedConditionSql()),
			addExtraSpaceIfExist(extraOption),
			)).Exec()
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}