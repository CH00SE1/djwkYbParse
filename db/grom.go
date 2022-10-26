package db

/**
 * @title grom
 * @author CH00SE1
 * @date 2022-10-26 15:18:16
 */

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// 数据库连接
func Goorm() (*gorm.DB, error) {
	return gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:11098319@tcp(192.168.10.87:3306)/djwk_test?charset=utf8mb4&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         170,                                                                                        // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                                       // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                                       // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                                       // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                                      // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "djwk_",
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
}

// 创建表
func CreateTableGrom(TableName interface{}) error {
	goorm, _ := Goorm()
	return goorm.Migrator().CreateTable(&TableName)
}
