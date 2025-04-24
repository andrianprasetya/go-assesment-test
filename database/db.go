package database

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
	"strings"
	"sync"
)

type Initializer func(*gorm.DB)

type DialectInitializer func(dsn string) gorm.Dialector

type dialect struct {
	template    string
	initializer DialectInitializer
}

var (
	mu           sync.Mutex
	dbConnection *gorm.DB
	initializers []Initializer
	dialects     = map[string]dialect{}
)

func RegisterDialect(name, template string, initializer DialectInitializer) {
	mu.Lock()
	defer mu.Unlock()
	if _, ok := dialects[name]; ok {
		panic(fmt.Sprintf("Dialect %q already exists", name))
	}
	dialects[name] = dialect{template, initializer}
}

func GetConnection() *gorm.DB {
	mu.Lock()
	defer mu.Unlock()
	if dbConnection == nil {
		dbConnection = newConnection()
	}
	return dbConnection
}

func newConnection() *gorm.DB {
	driver := os.Getenv("DB_CONNECTION") // Using the loaded config

	// Register dialect
	dialect, err := dialects[driver]
	if !err {
		panic(fmt.Sprintf("DB Connection %s not supported, forgotten import?", driver))
	}

	// Build the DSN using the registered dialect
	dsn := dialect.buildDSN()
	// Open connection to DB
	db, errConn := gorm.Open(dialect.initializer(dsn), &gorm.Config{
		PrepareStmt: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	})
	if errConn != nil {
		panic(errConn)
	}

	// Verify DB connection
	_, errConn = db.DB()
	if errConn != nil {
		panic(err)
	}

	// Run initializers
	for _, initializer := range initializers {
		initializer(db)
	}

	db.Logger = logger.Default.LogMode(logger.Info)
	return db
}

// Refactored to use the loaded config
func (d dialect) buildDSN() string {
	connStr := d.template
	connStr = strings.Replace(connStr, "{host}", os.Getenv("DB_HOST"), -1)
	connStr = strings.Replace(connStr, "{port}", os.Getenv("DB_PORT"), -1)
	connStr = strings.Replace(connStr, "{username}", os.Getenv("DB_USER"), -1)
	connStr = strings.Replace(connStr, "{password}", os.Getenv("DB_PASS"), -1)
	connStr = strings.Replace(connStr, "{name}", os.Getenv("DB_NAME"), -1)
	connStr = strings.Replace(connStr, "{options}", "", -1) // Can add more options if needed

	return connStr
}
