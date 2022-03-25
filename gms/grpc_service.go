package gms

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"xorm.io/xorm"
)

var Logger *logrus.Logger

type gmService struct {
	dbContext interface{}
	settings  Configuration
}

func (me *gmService) GetDBContext() interface{} {
	return me.dbContext
}

func (me *gmService) AddLogger() {
}

func (me *gmService) Configure(appSettings Configuration) error {
	if err := me.configureSettings(appSettings); err != nil {
		return err
	}

	if err := me.configureLogger(); err != nil {
		return err
	}

	if err := me.configureDB(); err != nil {
		return err
	}

	return nil
}

func (me *gmService) configureSettings(appSettings Configuration) error {
	viper.AddConfigPath(".")
	viper.SetConfigName("server")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	viper.ReadInConfig()
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&appSettings); err != nil {
		return err
	}
	me.settings = appSettings
	return nil
}

func (me *gmService) configureLogger() error {
	var logger = logrus.New()

	var logLevelMap = map[string]logrus.Level{
		"trace": logrus.TraceLevel,
		"debug": logrus.DebugLevel,
		"info":  logrus.InfoLevel,
		"warn":  logrus.WarnLevel,
		"error": logrus.ErrorLevel,
	}

	level, ok := logLevelMap[me.settings.GetLogLevel()]
	if !ok {
		return errors.New("can't make logger")
	}
	logger.SetLevel(level)
	logger.SetFormatter(&logrus.JSONFormatter{})

	Logger = logger
	Logger.Info("Logger", "Initialization", "OK")
	return nil
}

func (me *gmService) configureDB() error {
	if me.settings.GetDBType() == "mongodb" {
		return me.configureMongoDB()
	}
	return me.configureRDB()
}

func (me *gmService) configureMongoDB() error {

	connString := fmt.Sprintf("mongodb+srv://%s:%s@%s?retryWrites=true&w=majority",
		me.settings.GetDBUser(),
		me.settings.GetDBPassword(),
		me.settings.GetDBHost())
	clientOptions := options.Client().ApplyURI(connString)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)

	// uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", config.User, config.Password, config.Host, config.Port)
	// utils.Logger.Info("DB", "Connection", uri)
	// clientOptions := options.Client().ApplyURI(uri)
	// Connect to MongoDB
	//client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return err
	}
	// Check the connection
	// err = client.Ping(context.TODO(), nil)
	// if err != nil {
	// 	return nil, err
	// }

	db := client.Database(me.settings.GetDatabase())
	me.dbContext = db
	Logger.Info("DB", "Initialization", "OK")
	return nil
}

func (me *gmService) configureRDB() error {
	//connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslrootcert=AWS-ECF.pem sslmode=verify-full",
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		me.settings.GetDBHost(),
		me.settings.GetDBPort(),
		me.settings.GetDBUser(),
		me.settings.GetDBPassword(),
		me.settings.GetDatabase())

	db, err := xorm.NewEngine(me.settings.GetDBType(), connString)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	db.ShowSQL(true)
	db.SetMaxIdleConns(me.settings.GetDBMaxIdelConns())
	db.SetMaxOpenConns(me.settings.GetDBMaxOpenConns())

	me.dbContext = db
	Logger.Info("DB", "Initialization", "OK")
	return nil
}

func (me *gmService) StartUp(fn RegisterFunc) {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(me.settings.GetBindPort()))
	if err != nil {
		Logger.Fatal("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	fn(grpcServer)
	//userpb.RegisterUserServer(grpcServer, &userServer{})

	Logger.Info("start gRPC server on %s port", me.settings.GetBindPort())
	if err := grpcServer.Serve(lis); err != nil {
		Logger.Fatal("failed to serve: %s", err)
	}
}

func NewService() *gmService {
	return &(gmService{})
}

type RegisterFunc func(*grpc.Server)
