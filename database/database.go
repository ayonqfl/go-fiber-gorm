package database

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database instances for different databases
type Databases struct {
	QtraderDB *gorm.DB
	TradeDB   *gorm.DB
}

// Global database instance
var DB Databases

// ConnectDatabases initializes all database connections
func ConnectDatabases() {
	var err error

	// Connect to QtraderDB
	qtraderURI := os.Getenv("QTRADER_DB_URI")
	if qtraderURI == "" {
		log.Fatal("QTRADER_DB_URI not set in environment")
	}

	DB.QtraderDB, err = connectDB(qtraderURI, "QtraderDB")
	if err != nil {
		log.Fatalf("Failed to connect to QtraderDB: %v", err)
	}

	// Connect to TradeDB
	tradeURI := os.Getenv("TRADE_DB_URI")
	if tradeURI == "" {
		log.Fatal("TRADE_DB_URI not set in environment")
	}

	DB.TradeDB, err = connectDB(tradeURI, "TradeDB")
	if err != nil {
		log.Fatalf("Failed to connect to TradeDB: %v", err)
	}

	log.Println("All databases connected successfully")
}

// connectDB creates a single database connection
func connectDB(dsn string, dbName string) (*gorm.DB, error) {
	// Open database connection with disabled prepared statements
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // Disable prepared statements
	}), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: false,
	})

	if err != nil {
		return nil, err
	}

	// Get underlying SQL database
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// Configure connection pool
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	// Test connection
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	log.Printf("Connected to %s successfully", dbName)
	return db, nil
}

// CloseDatabases closes all database connections gracefully
func CloseDatabases() error {
	// Close QtraderDB
	if DB.QtraderDB != nil {
		sqlDB, err := DB.QtraderDB.DB()
		if err == nil {
			if err := sqlDB.Close(); err != nil {
				log.Printf("Error closing QtraderDB: %v", err)
			} else {
				log.Println("QtraderDB closed successfully")
			}
		}
	}

	// Close TradeDB
	if DB.TradeDB != nil {
		sqlDB, err := DB.TradeDB.DB()
		if err == nil {
			if err := sqlDB.Close(); err != nil {
				log.Printf("Error closing TradeDB: %v", err)
			} else {
				log.Println("TradeDB closed successfully")
			}
		}
	}

	log.Println("All database connections closed")
	return nil
}

// GetQtraderDB returns the QtraderDB instance
func GetQtraderDB() *gorm.DB {
	if DB.QtraderDB == nil {
		log.Fatal("QtraderDB is not initialized")
	}
	return DB.QtraderDB
}

// GetTradeDB returns the TradeDB instance
func GetTradeDB() *gorm.DB {
	if DB.TradeDB == nil {
		log.Fatal("QtraderDB is not initialized")
	}
	return DB.TradeDB
}