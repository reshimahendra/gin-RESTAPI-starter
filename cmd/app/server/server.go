/*
    Package server is our app starting point
*/
package server 

import (
	"fmt"
	"strings"

	"github.com/reshimahendra/gin-starter/internal/config"
	"github.com/reshimahendra/gin-starter/internal/database"
	"github.com/reshimahendra/gin-starter/pkg/logger"
	"github.com/spf13/viper"
)

// Run is for running the server
func Run() {
    /* 
        TODO:
            1. Initialize configuration
            2. Initialize database
            3. Initialize router
            4. Run the server
    */
    fmt.Println("Masuk init")
    // Setup server 
    if err := config.Setup(); err != nil {
        logger.Fatalf("Initializing server fail: %s", err)
    }

    // setup database
    if err := database.Setup(); err != nil {
        logger.Fatalf("Initializing database fail: %s", err)
    }

    // Initialize and prepare Database
    db := database.GetDB()

    // Initialize and prepare to load main router 
    r := Router(db)

    // setup server host
    host := "127.0.0.1"
    if hostname := viper.GetString("host"); hostname != "" {
        host = hostname
    }

    baseURL := host + ":" + viper.GetString("server.port") 

    logger.Infof("Server is running at %s", baseURL)
    welcome("Lotus BaliWeb", "http://" + baseURL, "-", 46)
    logger.Fatalf("%v", r.Run(baseURL))
}

// 'Welcome' is for console 'greetings' when the server is starting
func welcome(greetings, server string, fillerChar string, lineLength int) {
    fmt.Println()
    fmt.Println(printFiller(fillerChar, lineLength))
    fmt.Println(formatGreeting(greetings, lineLength))
    fmt.Println(formatGreeting(server, lineLength))
    fmt.Println(printFiller(fillerChar, lineLength))
    fmt.Println()
}

// formatGreeting will formating our server greeting text
func formatGreeting(greeting string, lineLength int) string {
    var fillerGreeting string
    var fillerLen int

    if len(greeting)%2 == 0 {
        fillerLen = (lineLength - len(greeting)) / 2
    } else {
        fillerLen = ((lineLength - len(greeting)) / 2) - 1
    }
    for i := 0; i < fillerLen; i++ {
        fillerGreeting += " "
    }

    return fmt.Sprintf("%s%s",fillerGreeting, strings.ToUpper(greeting))
}

// printFiller will creating a horizontal line based on given string length
func printFiller(fillerChar string, lineLength int) string {
    // print filler char 
    var filler string
    for i := 0; i <= lineLength; i++ {
            filler += fillerChar
    }
    return filler
}
