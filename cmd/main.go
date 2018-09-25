package main
  
import (
    "net/http"
    "github.com/goframe/log"
    "github.com/goframe/conf"
    "github.com/goframe/service"
    "github.com/goframe/dao"
    "fmt"
    "flag"
    "github.com/goframe/cache"
)

func hello(w http.ResponseWriter, r *http.Request) {
    log.Info("Hello world!")
    fmt.Fprint(w, "Hello world !")
}
func main() {
    flag.Parse()
    if err := conf.Init(); err != nil {
        log.Error("conf.Init() err:%+v", err)
    }
    innerConf := conf.Conf
    log.Init(innerConf.Log)

    redisConf := innerConf.Redis
    mysqlConf := innerConf.Mysql

    redis := dao.NewRedis(redisConf.Addr, redisConf.Password)
    mysql := dao.NewMysql(mysqlConf.UserName, mysqlConf.Password, mysqlConf.IpHost, mysqlConf.DbName)
    localCache := cache.NewCache(mysql, conf.Conf.GetCronFreq())
    localCache.Init()

    http.Handle("/", &service.Proxy{DataService: &service.DataService{Mysql: mysql, Redis: redis, LocalCache: localCache}, Conf: conf.Conf})
    http.HandleFunc("/hello", hello)

    err := http.ListenAndServe(innerConf.SonoServer.Port, nil)
    if err != nil {
        log.Info("ListenAndServe:%+v", err)
    }
    mysql.Close()
    redis.Close()
}
