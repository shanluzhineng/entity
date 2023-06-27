package mongodb

import (
	"github.com/abmpio/libx/reflector"
	"github.com/abmpio/libx/str"
	"github.com/abmpio/mongodbr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	DatabaseName string
	_db          *mongo.Database
	// key: CollectionName
	// value: RepositoryOption
	_entityRepositoryOptionMap map[string][]mongodbr.RepositoryOption
	_repositoryMapping         map[string]mongodbr.IRepository
}

func NewDatabase(databaseName string) *Database {
	d := &Database{
		DatabaseName: databaseName,
	}
	d._db = mongodbr.DefaultClient.Database(databaseName)
	d._entityRepositoryOptionMap = make(map[string][]mongodbr.RepositoryOption)
	d._repositoryMapping = make(map[string]mongodbr.IRepository)

	return d
}

var (
	_registedDatabase map[string]*Database = make(map[string]*Database)
)

func ensureDatabaseRegisted(databaseName string) *Database {
	d, ok := _registedDatabase[databaseName]
	if ok {
		return d
	}
	d = NewDatabase(databaseName)
	_registedDatabase[databaseName] = d
	return d
}

func GetCollectionName(v interface{}) string {
	return str.ToSnake(reflector.GetName(v))
}

// regist Repository create option
func RegistEntityRepositoryOption[T mongodbr.IEntity](databaseName string, opts ...mongodbr.RepositoryOption) {
	d := ensureDatabaseRegisted(databaseName)
	collectionName := GetCollectionName((new(T)))
	d._entityRepositoryOptionMap[collectionName] = createEntityRepositoryOption[T](opts...)
	d._repositoryMapping[collectionName] = d.ensureCreateRepository(collectionName, opts...)
}

// 获取指定db的IRepository接口
func GetRepository(databaseName string, modelInstance interface{}) mongodbr.IRepository {
	return ensureDatabaseRegisted(databaseName).GetRepository(modelInstance)
}

func createEntityRepositoryOption[T mongodbr.IEntity](opts ...mongodbr.RepositoryOption) []mongodbr.RepositoryOption {
	if len(opts) <= 0 {
		opts = append(opts, mongodbr.WithCreateItemFunc(func() interface{} {
			return new(T)
		}))
		opts = append(opts, mongodbr.WithDefaultSort(func(fo *options.FindOptions) *options.FindOptions {
			return fo.SetSort(bson.D{{Key: "_id", Value: -1}})
		}))
	}
	return opts
}

func (d *Database) GetEntityRepositoryOptionMap() map[string][]mongodbr.RepositoryOption {
	return d._entityRepositoryOptionMap
}

func (d *Database) GetRepository(modelInstance interface{}) mongodbr.IRepository {
	collectionName := GetCollectionName(modelInstance)
	instance, ok := d._repositoryMapping[collectionName]
	if ok {
		return instance
	}
	return nil
}

func (d *Database) ensureCreateRepository(collectionName string, opts ...mongodbr.RepositoryOption) mongodbr.IRepository {
	repository, err := mongodbr.NewRepositoryBase(func() *mongo.Collection {
		return d._db.Collection(collectionName)
	}, opts...)
	if err != nil {
		panic(err)
	}
	if repository == nil {
		panic("cannot create repository for collection" + collectionName)
	}
	return repository
}
