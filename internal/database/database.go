package database

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
    db    *gorm.DB
    err   error
    dbErr error
)

type Database struct{
    *gorm.DB
}

// Log writer
func getWriter() io.Writer {
    file, err := os.OpenFile(
        viper.GetString("logger.database_log_name"),
        os.O_CREATE | os.O_WRONLY | os.O_APPEND, 
        0666)
    if err != nil {
        return os.Stdout
    } else {
        return file
    }
}

// DB will return Database Connection
func DB() *gorm.DB {
    return db
}

// Return DB Error
func DBErr() error {
    return dbErr
}


// Init database and forward saved ref to the Database struct 
func Setup() error {
    var database *gorm.DB

    // Extract config data from '.config.yaml'
    driver   := viper.GetString("database.driver")
    dbname   := viper.GetString("database.database")
    username := viper.GetString("database.username")
    password := viper.GetString("database.password")
    host     := viper.GetString("database.host")
    port     := viper.GetString("database.port")
    // sslmode  := viper.GetBool("database.sslmode")
    logmode  := viper.GetBool("database.logmode")


    // Database logging
    loglevel := logger.Silent
    if logmode {
        loglevel = logger.Info
    }

    DBLogger := logger.New(
        log.New(getWriter(), "\r\n", log.LstdFlags),
        logger.Config{
            SlowThreshold             : time.Second,  // Slow threshold
            LogLevel                  : loglevel,     // Log level '{Silent, Error, Warn, Info}'
            IgnoreRecordNotFoundError : true,         // Ignore logger if record not found
            Colorful                  : false,        // Disable log color feature
        },
    )

    switch driver {
    case "postgres":
        conn := fmt.Sprintf(
            "host=%s port=%s user=%s password=%s dbname=%s", // sslmode=%t",
            host, port, username, password, dbname, // sslmode,
        )
        database, err = gorm.Open(postgres.Open(conn), &gorm.Config{Logger: DBLogger})
    case "mysql":
        conn := fmt.Sprintf(
            "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", 
            username, password, host, port, dbname,
        )
        database, err = gorm.Open(mysql.Open(conn), &gorm.Config{Logger: DBLogger})
    case "sqlite":
        conn := fmt.Sprintf("%s.db", dbname)
        database, err = gorm.Open(sqlite.Open(conn), &gorm.Config{Logger: DBLogger})
    default:
        return errors.New(
            "Database unsuported. Supported database: 'postgresql', 'mysql', and 'sqlite'")
    }

    if err != nil {
        dbErr = err
        return err
    }

    // NOTE:
    // On testing, you can do 'automigrate' here, or at the module `router` 
    // for 'production', better run it via 'SQL' command

    // database.automigrate(
    //     &yourModel1{},
    //     &yourModel2{},
    //     &etcModel{},
    // )

    // pass the database connection to 'db' var
    db = database

    return nil
}
