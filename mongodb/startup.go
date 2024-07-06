package mongodb

import (
	"github.com/shanluzhineng/app/web"
)

func init() {
	web.ConfigureService(initMongodbConfigurator)
}
