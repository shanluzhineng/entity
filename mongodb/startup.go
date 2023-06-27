package mongodb

import (
	"abmp.cc/app/pkg/app/web"
)

func init() {
	web.ConfigureService(initMongodbConfigurator)
}
