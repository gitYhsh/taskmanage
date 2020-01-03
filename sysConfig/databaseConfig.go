package sysConfig

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
)

var redisClient *redis.Pool

func GetCon() redis.Conn {
	return redisClient.Get()
}

func InitDatabase() {
	/**数据库地址**/
	dbHost := beego.AppConfig.String("db.host")
	/**账户**/
	dbUser := beego.AppConfig.String("db.user")

	dbPwd := beego.AppConfig.String("db.password")

	dbPort := beego.AppConfig.String("db.port")

	dbName := beego.AppConfig.String("db.name")

	/**dbZone := beego.AppConfig.String("db.timezone")**/
	zone, _ := time.LoadLocation("Asia/Shanghai")
	orm.SetDataBaseTZ("default", zone)
	orm.RegisterDataBase("default", "mysql", dbUser+":"+dbPwd+"@tcp("+dbHost+":"+
		dbPort+")/"+dbName+"?charset=utf8", 30)

}

func newPool() {

	server := beego.AppConfig.String("redis.host")
	password := beego.AppConfig.String("redis.user")
	db, _ := beego.AppConfig.Int("redis.db")

	redisClient = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server, redis.DialDatabase(db))
			if err != nil {
				return nil, err
			}
			if password != "" {
				_, err := c.Do("AUTH", password)
				if err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
