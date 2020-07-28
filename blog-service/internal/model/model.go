package model

import (
	"fmt"
	otgorm "github.com/eddycjy/opentracing-gorm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-project-example/blog-service/global"
	"go-project-example/blog-service/pkg/setting"
	"time"
)

/**
 *@Author tudou
 *@Date 2020/7/26
 **/

const (
	STATE_CLOSE = iota
	STATE_OPEN
)


type Model struct{
	ID uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"` //创建者
	ModifiedBy string `json:"modified_by"` //修改者
	CreatedOn  uint32 `json:"created_on"` //创建时间
	ModifiedOn uint32 `json:"modified_on"` //修改时间
	DeletedOn  uint32 `json:"deleted_on"` //删除时间
	IsDel      uint8  `json:"is_del"` //删除状态：0：未删除 1：已删除
}


func NewDBEngine(databaseSetting *setting.DatabaseSettingS)(*gorm.DB,error){
	db,err:=gorm.Open(databaseSetting.DBType,fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	))
	if err!=nil{
		return nil,err
	}

	if global.ServerSetting.RunMode=="debug"{
		db.LogMode(true)
	}

	db.SingularTable(true)

	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)
	db.Callback().Create().Replace("gorm:update_time_stamp",updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp",updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete",deleteCallBack)
	otgorm.AddGormCallbacks(db)
	return db,nil
}


func updateTimeStampForCreateCallback(scope *gorm.Scope){
	if !scope.HasError(){
		nowTime:=time.Now().Unix()
		if createTimeField,ok:=scope.FieldByName("CreatedOn");ok{
			if createTimeField.IsBlank{
				_ = createTimeField.Set(nowTime)
			}
		}


		if modifyTimeField,err:=scope.FieldByName("ModifiedOn");err{
			if modifyTimeField.IsBlank{
				_ = modifyTimeField.Set(nowTime)
			}
		}

	}
}


func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _,ok := scope.Get("gorm:update_column");!ok{
		_ = scope.SetColumn("ModifiedOn",time.Now().Unix())
	}
}

func deleteCallBack(scope *gorm.Scope){
	if !scope.HasError(){
		var extraOption string
		if str,ok:=scope.Get("gorm:delete_option");ok{
			extraOption = fmt.Sprint(str)
		}

		deletedOnField,hasDeletedOnField:=scope.FieldByName("DeleteOn")
		isDelField,hasIsDelOnField := scope.FieldByName("IsDel")
		if !scope.Search.Unscoped && hasDeletedOnField && hasIsDelOnField{
			now := time.Now()
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v,%v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(now),
				scope.Quote(isDelField.DBName),
				scope.AddToVars(1),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else{
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}

	}
}


func addExtraSpaceIfExist(str string)string{
	if str != ""{
		return  " "+str
	}
	return ""
}
