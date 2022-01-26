package bootstrap

import (
	"fmt"
	"gin-blog/app/models"
	"gin-blog/global"
	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"time"
)

func InitDB() *gorm.DB {
	dbConfig := global.App.Config.DB
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database, dbConfig.Charset,
	)
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}

	if db, err := gorm.Open(
		mysql.New(mysqlConfig),
		&gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			Logger:                                   createGormLogger(),
		},
	); err != nil {
		global.App.Log.Error("mysql connect failed, err:", zap.Any("err", err))
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxOpenConns(global.App.Config.DB.MaxOpenConn)
		sqlDB.SetMaxIdleConns(global.App.Config.DB.MaxIdleConn)
		initTables(db)
		global.App.DB = db
		return db
	}
}

func initTables(db *gorm.DB) {
	err := db.AutoMigrate(
		models.User{},
	)
	if err != nil {
		global.App.Log.Error("数据库表迁移失败", zap.Any("err", err))
	}
}

func createGormLogger() logger.Interface {
	var logLevel logger.LogLevel

	switch global.App.Config.DB.LogLevel {
	case "silent":
		logLevel = logger.Silent
	case "error":
		logLevel = logger.Error
	case "warn":
		logLevel = logger.Warn
	case "info":
		logLevel = logger.Info
	default:
		logLevel = logger.Info
	}

	return logger.New(createGormLogWriter(), logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logLevel,
		IgnoreRecordNotFoundError: false,
		Colorful:                  !global.App.Config.DB.EnableLog, // 彩色打印
	})
}

// 创建一个新的日志写入器
func createGormLogWriter() logger.Writer {
	var writer io.Writer

	if global.App.Config.DB.EnableLog {
		writer = &lumberjack.Logger{
			Filename:   global.App.Config.Log.RootDir + "/" + global.App.Config.DB.LogFilename,
			MaxSize:    global.App.Config.Log.MaxSize,
			MaxAge:     global.App.Config.Log.MaxAge,
			MaxBackups: global.App.Config.Log.MaxBackup,
			Compress:   global.App.Config.Log.Compress,
		}
	} else {
		writer = os.Stdout
	}

	return log.New(writer, "\r\n", log.LstdFlags)
}
