package initialize

import (
	"edushiwei-service/global"
	"edushiwei-service/initialize/internal"
	"edushiwei-service/model"
	"os"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//@author: SliverHorn
//@function: Gorm
//@description: 初始化数据库并产生数据库全局变量
//@return: *gorm.DB
func Gorm() *gorm.DB {
	switch global.COURSE_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

//@author: SliverHorn
//@function: GormMysql
//@description: 初始化Mysql数据库
//@return: *gorm.DB
func GormMysql() *gorm.DB {
	m := global.COURSE_CONFIG.Mysql
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig(m.LogMode)); err != nil {
		global.COURSE_LOG.Error("MySQL启动异常", zap.Any("err", err))
		os.Exit(0)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

//@author: SliverHorn
//@function: gormConfig
//@description: 根据配置决定是否开启日志
//@param: mod bool
//@return: *gorm.Config
func gormConfig(mod bool) *gorm.Config {
	var config = &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	switch global.COURSE_CONFIG.Mysql.LogZap {
	case "silent", "Silent":
		config.Logger = internal.Default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = internal.Default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = internal.Default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = internal.Default.LogMode(logger.Info)
	case "zap", "Zap":
		config.Logger = internal.Default.LogMode(logger.Info)
	default:
		if mod {
			config.Logger = internal.Default.LogMode(logger.Info)
			break
		}
		config.Logger = internal.Default.LogMode(logger.Silent)
	}
	return config
}

// MysqlTables
//@author: SliverHorn
//@function: MysqlTables
//@description: 注册数据库表专用
//@param: db *gorm.DB
func MysqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.Accout{},
		model.Area{},
		model.Audit{},
		model.Category{},
		model.ChapterLike{},
		model.ChapterLive{},
		model.ChapterRead{},
		model.ChapterResource{},
		model.ChapterUser{},
		model.ChapterVod{},
		model.Chapter{},
		model.Content{},
		model.Consult{},
		model.ConsultLike{},
		model.Course{},
		model.CourseCategory{},
		model.CourseFavorite{},
		model.CoursePackage{},
		model.CourseRating{},
		model.CourseRelated{},
		model.CourseTopic{},
		model.CourseUser{},
		model.Danmu{},
		model.Help{},
		model.ImFriendGroup{},
		model.ImFriendUser{},
		model.ImGroup{},
		model.ImGroupUser{},
		model.ImMessage{},
		model.ImNotice{},
		model.ImUser{},
		model.Learning{},
		model.Nav{},
		model.Online{},
		model.Order{},
		model.OrderStatus{},
		model.Package{},
		model.Page{},
		model.RefoundStatus{},
		model.Refund{},
		model.Resource{},
		model.Review{},
		model.ReviewLike{},
		model.Reward{},
		model.Role{},
		model.Setting{},
		model.Slide{},
		model.Task{},
		model.Topic{},
		model.Trade{},
		model.TradeStatus{},
		model.Upload{},
		model.UserSession{},
		model.UserToken{},
		model.User{},
		model.Vip{},
		model.WechatSubscribe{},
	)
	if err != nil {
		global.COURSE_LOG.Error("register table failed", zap.Any("err", err))
		os.Exit(0)
	}
	global.COURSE_LOG.Info("register table success")
}
