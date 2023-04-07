package biz_config

import (
	goredis "github.com/go-redis/redis/v8"
	"github.com/isyscore/isc-gobase/extend/etcd"
	"github.com/isyscore/isc-gobase/extend/orm"
	"github.com/isyscore/isc-gobase/extend/redis"
	"github.com/isyscore/isc-gobase/logger"
	"gorm.io/gorm"
)

var Db *gorm.DB
var RedisDb goredis.UniversalClient
var EtcdClient *etcd.EtcdClientWrap

func InitConfig() {
	InitDb()
	InitEtcd()
	InitRedis()
}

func InitDb() {
	dbTem, err := orm.NewGormDb()
	if err != nil {
		logger.Warn("数据库初始化失败, %v", err.Error())
		return
	}
	logger.Info("数据库连接成功")
	Db = dbTem
}

func InitEtcd() {
	etcdClientTem, err := etcd.NewEtcdClient()
	if err != nil {
		logger.Warn("etcd初始化失败, %v", err.Error())
		return
	}
	logger.Info("etcd连接成功")
	EtcdClient = etcdClientTem
}

func InitRedis() {
	RedisDbTem, err := redis.NewClient()
	if err != nil {
		logger.Warn("redis连接异常", err)
	}
	RedisDb = RedisDbTem
	logger.Info("redis连接成功")
}
