package util

import (
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type DBType string

var (
	envLoaded bool
	//保证某个操作在程序生命周期内仅执行一次（线程安全）
	envOnce sync.Once
)

const (
	DBTypeSQLite   DBType = "sqlite"
	DBTypeMysql    DBType = "mysql"
	DBTypePostgres DBType = "postgres"
)

func loadEnv() {
	envOnce.Do(func() {

		_, currentFile, _, ok := runtime.Caller(0)
		if !ok {
			return
		}
		utilDir := filepath.Dir(currentFile)
		envDir := filepath.Dir(utilDir)
		envPath := filepath.Join(envDir, ".env")
		if err := godotenv.Load(envPath); err != nil {
			// .env file is optional, so we don't fail if it doesn't exist
			// Environment variables can still be set directly via system environment
			return
		}
		envLoaded = true

	})
}
func NewUtilDB() (*gorm.DB, error) {

	dbType := getDBType()
	switch dbType {
	case DBTypeSQLite:
		return newSQLiteDB()
	case DBTypePostgres:
		return newPostgresDB()
	case DBTypeMysql:
		return newMySQLDB()
	default:
		return newMySQLDB()
	}

}

func newPostgresDB() (*gorm.DB, error) {
	panic("unimplemented")
}

func newSQLiteDB() (*gorm.DB, error) {
	panic("unimplemented")
}

func getDBType() DBType {
	loadEnv()
	dbType := os.Getenv("TEST_DB_TYPE")
	switch dbType {
	case "sqlite":
		return DBTypeSQLite
	case "mysql":
		return DBTypeMysql
	case "postgres":
		return DBTypePostgres
	}
	return DBTypeMysql // 默认使用 MySQL
}

func newMySQLDB() (*gorm.DB, error) {
	loadEnv()
	dsn := os.Getenv("TEST_MYSQL_DSN")
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),

		// NamingStrategy: Customize table and column naming
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: false,
			NoLowerCase:   false,
		},
	})

}
