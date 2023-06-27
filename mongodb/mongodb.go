package mongodb

import (
	"fmt"
	"time"

	"abmp.cc/abmp/pkg/log"
	"abmp.cc/app/pkg/app"
	"abmp.cc/app/pkg/app/web"
	"github.com/abmpio/configurationx"
	"github.com/abmpio/configurationx/options/mongodb"
	"github.com/abmpio/mongodbr"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func initMongodbConfigurator(wa web.WebApplication) {
	if app.HostApplication.SystemConfig().App.IsRunInCli {
		return
	}
	if mongodbr.DefaultClient != nil {
		//已经初始化完成，则直接返回
		return
	}
	initMongodb()
}

// 初始化mongodb
func initMongodb() {
	var defaultMongodbOptions *mongodb.MongodbOptions
	if len(configurationx.GetInstance().Mongodb.MongodbList) >= 0 {
		defaultMongodbOptions = configurationx.GetInstance().Mongodb.GetDefaultOptions()
	}
	if defaultMongodbOptions == nil {
		err := fmt.Errorf("无法初始化mongodb,没有配置好Mongodb的字符串连接")
		log.Logger.Error(err.Error())
		panic(err)
	}

	_, err := mongodbr.SetupDefaultClient(defaultMongodbOptions.Uri, func(co *options.ClientOptions) {})
	if err != nil {
		log.Logger.Error(err.Error())
		panic(err)
	}
	//测试ping
	for {
		err := mongodbr.Ping()
		if err == nil {
			break
		}
		log.Logger.Warn(err.Error())
		log.Logger.Warn("2s后重新测试...")
		time.Sleep(2 * time.Second)
	}
}
