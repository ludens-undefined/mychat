package model

import (
	"context"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var globalIsRelated = true // 全局预加载

var db *gorm.DB

// prepare for other
type _BaseMgr struct {
	*gorm.DB
	ctx       context.Context
	cancel    context.CancelFunc
	timeout   time.Duration
	isRelated bool
}

func generateMysqlDsn(config map[string]string) (dsn string) {
	dsnConfig := mysql.Config{
		User:                 config["username"],
		Passwd:               config["password"],
		Net:                  "tcp",
		Addr:                 config["host"] + ":" + config["port"],
		DBName:               config["database"],
		Params:               map[string]string{"charset": config["charset"]},
		Loc:                  time.Local,
		ParseTime:            true,
		AllowNativePasswords: true,
	}
	return dsnConfig.FormatDSN()
}

func init() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("conf")
	err := viper.ReadInConfig()

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	var dsn string

	if viper.IsSet("mysql") {
		dsn = generateMysqlDsn(viper.GetStringMapString("mysql"))
	}

	db, err = gorm.Open(gmysql.Open(dsn), &gorm.Config{})
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}

// SetTimeOut SetCtx set context
func (obj *_BaseMgr) SetTimeOut(timeout time.Duration) {
	obj.ctx, obj.cancel = context.WithTimeout(context.Background(), timeout)
	obj.timeout = timeout
}

// SetCtx set context
func (obj *_BaseMgr) SetCtx(c context.Context) {
	obj.ctx = c
}

// GetDB get gorm.DB info
func (obj *_BaseMgr) GetDB() *gorm.DB {
	return obj.DB
}

// UpdateDB update gorm.DB info
func (obj *_BaseMgr) UpdateDB(db *gorm.DB) {
	obj.DB = db
}

// GetIsRelated Query foreign key Association.获取是否查询外键关联(gorm.Related)
func (obj *_BaseMgr) GetIsRelated() bool {
	return obj.isRelated
}

// SetIsRelated Query foreign key Association.设置是否查询外键关联(gorm.Related)
func (obj *_BaseMgr) SetIsRelated(b bool) {
	obj.isRelated = b
}

// New new gorm.新gorm
func (obj *_BaseMgr) New() *gorm.DB {
	return obj.DB.Session(&gorm.Session{Context: obj.ctx})
}

// OpenRelated 打开全局预加载
func OpenRelated() {
	globalIsRelated = true
}

// CloseRelated 关闭全局预加载
func CloseRelated() {
	globalIsRelated = true
}
