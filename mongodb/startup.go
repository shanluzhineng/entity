package mongodb

import (
	"github.com/abmpio/abmp/app/web"
)

func init() {
	web.ConfigureService(initMongodbConfigurator)
}
