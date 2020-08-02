package xorm

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"gopkg.in/yaml.v2"
	"io/ioutil"

	"github.com/go-xorm/core"
	"github.com/x-chain/baasconsole/baas-core/common/log"
	"time"
)

var level = log.ERROR
var logger = log.GetLogger("xorm", level)

type Xorm struct {
	Config *MysqlConfig `yaml:"xorm"`
}

type MysqlConfig struct {
	Drivename string `yaml:"drivename"`
	Database  string `yaml:"database"`
	Ip        string `yaml:"ip"`
	Port      string `yaml:"port"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	Showsql   bool   `yaml:"showsql"`
	Maxidle   int    `yaml:"maxidle"`
	Maxopen   int    `yaml:"maxopen"`
}

func newXorm() *Xorm {
	return &Xorm{
		Config: &MysqlConfig{},
	}
}
func loadConfig(file string) *MysqlConfig {
	cfg, err := ioutil.ReadFile(file)
	if err != nil {
		logger.Error(err.Error())
	}
	var xorm = newXorm()
	err = yaml.Unmarshal(cfg, xorm)
	if err != nil {
		logger.Error(err.Error())
	}
	return xorm.Config
}

func GetEngine(configFile string) *xorm.Engine {
	config := loadConfig(configFile)
	//conn string
	conn := config.User + ":" + config.Password + "@tcp(" + config.Ip + ":" + config.Port + ")/" + config.Database + "?charset=utf8"
	engine, err := xorm.NewEngine(config.Drivename, conn)
	if err != nil {
		logger.Error(err.Error())
	}
	// 打印sql
	xormLogger := &OrmLogger{
		logger: logger,
		level:  core.LogLevel(level),
	}
	engine.SetLogger(xormLogger)
	engine.ShowSQL(config.Showsql)
	engine.SetMaxIdleConns(config.Maxidle)
	engine.SetMaxOpenConns(config.Maxopen)
	//连接生存时间半个小时
	engine.SetConnMaxLifetime(1800 * time.Second)
	return engine
}
