package mongodb

import (
	"github.com/abmpio/app/web"
)

func init() {
	web.ConfigureService(initMongodbConfigurator)
}
