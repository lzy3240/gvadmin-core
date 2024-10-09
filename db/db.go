package db

import (
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gvadmin_v3/core/config"
	"gvadmin_v3/core/log"
	defaultLog "log"
	"os"
	"time"
)

var (
	db  *gorm.DB
	err error
)

func Instance() *gorm.DB {
	if db == nil {
		InitConn()
	}
	return db
}

func InitConn() {
	m := config.Instance().DB
	dsn := m.DBUser + ":" + m.DBPwd + "@tcp(" + m.DBHost + ":" + m.DBPort + ")/" + m.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			// 启用更新时间戳功能
			NowFunc: func() time.Time {
				return time.Now().Local()
			},
			// 启用单数表和前缀
			NamingStrategy: schema.NamingStrategy{
				//TablePrefix:   "pre_", // 表前缀
				SingularTable: true, // 禁用表名复数
			},
			// 启用数据库日志
			Logger: logger.New(defaultLog.New(os.Stdout, "\r\n", defaultLog.LstdFlags),
				logger.Config{
					SlowThreshold: time.Second,
					LogLevel:      logger.Info,
					Colorful:      true,
				}),
		})
	if err != nil {
		log.Instance().Error("MySQL Init Error..." + err.Error())
		os.Exit(0)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(300)
}

//func initMyCallbacks() {
//	_ = db.Callback().Create().Replace("gorm:update_time_stamp", beforeCreate)
//	_ = db.Callback().Update().Replace("gorm:update_time_stamp", beforeUpdate)
//}

//func beforeCreate(scope *gorm.Scope) {
//	if !scope.HasError() {
//		now := time.Now()
//		if createdAtField, ok := scope.FieldByName("CreatedAt"); ok {
//			if createdAtField.IsBlank { // 判断该字段的值是否为空
//				createdAtField.Set(now)
//			}
//		}
//		if updatedAtField, ok := scope.FieldByName("UpdatedAt"); ok {
//			if updatedAtField.IsBlank {
//				updatedAtField.Set(now)
//			}
//		}
//	}
//}

//func beforeUpdate(scope *gorm.Scope) {
//	if updatedAtField, ok := scope.FieldByName("UpdatedAt"); ok {
//		if updatedAtField.IsBlank {
//			updatedAtField.Set(time.Now())
//		}
//	}
//}
