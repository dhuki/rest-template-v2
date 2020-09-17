package config

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/dhuki/rest-template-v2/common"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// gorm is still powefull because it's orm framework with bunch of functionality
// like preloading, erager_loading and many more
// but it's not fast enough than sqlx, or pq lib.
// downside of sqlx and pq is you have to do manully all of stuff to deal with database 
// because lack functionality that available in that orm lib.
func NewDatabase() (*gorm.DB, error) {
	// by default sslmode=enable, so you have to connect with ssl
	// since your server doesn't provide it
	// just use sslmode=disable
	dbURI := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable TimeZone=Asia/Jakarta", common.DbHost, common.DbPort, common.DbName, common.DbUsername, common.DbPassword)
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // disable auto log gorm v2
	}) // configuration gorm v2 (already release on 28 Aug 2020) src : https://gorm.io/docs/v2_release_note.html
	if err != nil {
		return nil, err
	}
	// disable auto logging from gorm lib v1
	// db.LogMode(false)

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// checking if connection to db is still alive
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	// Set maxActive (maximum number of active/connection (in-use or idle) ) that available in pool
	// if all connection already in-use (in this case 10 connection in-use) and need new connection
	// it will force to wait until at least one connection idle
	// by default postgres set max 100 connection if there is one connection want to establish
	// it will return "pq: sorry, too many clients already".
	sqlDB.SetMaxOpenConns(3)

	// Set maxidle connection that retained
	// if it set to 5, so there are 5 retained connection to db
	// and the other 5 is not retain connection it will close connection after used
	sqlDB.SetMaxIdleConns(3)

	// Set maximum lifetime of connection before retiring it.
	// if idle connection has been reached max lifetime it'll destroy connection
	// but if connection is in use and has reach max lifetime 
	// it'll wait until connection back again to the
	// pool and then destory it.
	sqlDB.SetConnMaxLifetime(10 * time.Minute)

	return db, nil
}

func NewTesting(sql *sql.DB) (*gorm.DB, error) {
	// dbURI := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable TimeZone=Asia/Jakarta", common.DbHost, common.DbPort, common.DbName, common.DbUsername, common.DbPassword)
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sql,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}) // configuration gorm v2 (it's unstable yet)
	if err != nil {
		return nil, err
	}

	return db, nil
}
