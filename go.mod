module github.com/duobiex/entity

go 1.20

require (
	abmp.cc/abmp v0.0.0-00010101000000-000000000000
	abmp.cc/app v0.0.0-00010101000000-000000000000
	github.com/abmpio/configurationx v0.0.0-20230614115922-7e2f5818b58d
	github.com/abmpio/libx v0.0.0-20230625025307-2f8abba069d6
	github.com/abmpio/mongodbr v0.0.0-20230611075214-8c9508eae6a4
	go.mongodb.org/mongo-driver v1.11.0
)

require (
	github.com/deckarep/golang-set v1.8.0 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/kataras/golog v0.1.8 // indirect
	github.com/kataras/pio v0.0.11 // indirect
	github.com/klauspost/compress v1.15.12 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible // indirect
	github.com/lestrrat-go/strftime v1.0.6 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/montanaflynn/stats v0.6.6 // indirect
	github.com/pelletier/go-toml/v2 v2.0.6 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/spf13/afero v1.9.3 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.15.0 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.1 // indirect
	github.com/xdg-go/stringprep v1.0.3 // indirect
	github.com/youmark/pkcs8 v0.0.0-20201027041543-1326539a0a0a // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	golang.org/x/crypto v0.5.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	abmp.cc/abmp => ../../../cloudmobile2.0_beta/lib/abmp.core
	abmp.cc/app => ../../../cloudmobile2.0_beta/lib/abmp.app
)
