package database

import (
	"fmt"
	"github.com/yugendra/TransactionsRoutine/database/models"
	"github.com/yugendra/TransactionsRoutine/transactionsroutine/interactor"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

/*database implements interactor.DatabaseInteractor
 *Provides mainly create and read operations for account and transactions.
 */
type database struct {
	conn *gorm.DB

	dbName     string
	dbUsername string
	dbPassword string
	dbPort     int
}

/*NewDatabase initializes and returns database connection
 *Panic out, if initialization fails
 */
func NewDatabase() interactor.DatabaseInteractor {
	db := &database{}
	dBUser := os.Getenv("DBUSER")
	dBPass := os.Getenv("DBPASS")
	dBPort := os.Getenv("DBPORT")
	DBName := os.Getenv("DBNAME")

	var conn *gorm.DB
	var err error
	dsn := fmt.Sprintf("host=database user=%s password=%s dbname=%s port=%s sslmode=disable",
		dBUser, dBPass, DBName, dBPort)
	for i := 0; i < 5; i++ {
		conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		time.Sleep(time.Second)
	}
	if err != nil {
		panic("Failed to connect to database")
	}

	/*Initial migration for project
	 *As writing database DDL is out of scope of this assignment,
	 *using this method to create table in database.
	 *But this is not an ideal way to create table in DB.
	 */
	err = conn.AutoMigrate(&models.Account{})
	err = conn.AutoMigrate(&models.Transaction{})
	if err != nil {
		panic("Failed to create tables in database")
	}

	db.conn = conn
	return db
}

//Close the database
func (db *database) Close() error {
	sqlDB, err := db.conn.DB()
	if err != nil {
		return err
	}
	sqlDB.Close()

	return nil
}
