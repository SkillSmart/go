package main

import (
	//"LearnGoProject/03_applications/01_codeOrg/modules/check"
	//"LearnGoProject/03_applications/01_codeOrg/modules/config/db"

	"LearnGoProject/03_applications/01_codeOrg/controllers/front"
	"LearnGoProject/03_applications/01_codeOrg/modules/config"
	"database/sql"
	"net/http"
)

// Set global variables for the program
var dbSQL, dbMongo *sql.DB

func init() {

	// Intialize the configuration variables
	Cfg := config.NewConfig()
	Cfg.ParseConfigfile("config.yaml")
	Cfg.ParseEnvVars()

	// Configure the packages through the config package API
	LogConfig := config.NewLogConfig()
	LogConfig.SetLogString("%s:%s")
	LogConfig.SetLogLevel(config.LogLevels.Debug)

	////// Instantiate the database connection by passing the appropriate Connector value
	////dbSQL = check.DBConnect(db.NewSQL("mysql", &db.MySQLConnector{
	////	// Working with nested types we need to manually instantiate the
	////	// SQLConnector for attribute reuse
	////	SQLConnector: db.SQLConnector{
	////		// TODO: Import from configuration file
	////		Username: "username",
	////		Password: "password",
	////		DBname: "database",
	////	},
	////}))
	////
	////// Doing this with a NoSQL connection it would look like this
	////dbMongo = check.DBConnect(db.NewSQL("mongo", &db.MongoDBConnector{
	////	SQLConnector: db.SQLConnector{
	////		// TODO: Import from configuration file
	////		Username: "mongoUname",
	////		Password: "mongoPW",
	////		DBname:   "mongoDB",
	////	},
	//}))

	//// Initialize templates
	//Tpl = template.Must(template.ParseGlob("/templates/*.gohtml"))

}

func main() {
	// Define the handlers
	http.HandleFunc("/", front.IndexHandler)
	// Get the server instance
	http.ListenAndServe(":8080", nil)
}
