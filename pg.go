package main

import (
  "database/sql"
  "fmt"
  "log"
  "os"
  _ "github.com/lib/pq"
)

const (
  dbhost = "DBHOST"
  dbport = "DBPORT"
  dbuser = "DBUSER"
  dbpass = "DBPASS"
  dbname = "DBNAME"
 )


func dbConfig() map[string]string {
    conf := make(map[string]string)
    host, ok := os.LookupEnv(dbhost)
    if !ok {
        panic("DBHOST environment variable required but not set")
    }
    port, ok := os.LookupEnv(dbport)
    if !ok {
        panic("DBPORT environment variable required but not set")
    }
    user, ok := os.LookupEnv(dbuser)
    if !ok {
        panic("DBUSER environment variable required but not set")
    }
    password, ok := os.LookupEnv(dbpass)
    if !ok {
        panic("DBPASS environment variable required but not set")
    }
    name, ok := os.LookupEnv(dbname)
    if !ok {
        panic("DBNAME environment variable required but not set")
    }
    conf[dbhost] = host
    conf[dbport] = port
    conf[dbuser] = user
    conf[dbpass] = password
    conf[dbname] = name
    return conf
}

func main() {
  // get the configuration from the environment
  conf := dbConfig()

  // build the string
  psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
    "password=%s dbname=%s sslmode=disable",
    conf[dbhost], conf[dbport], conf[dbuser], conf[dbpass], conf[dbname])

  // open the connection
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()

  // ping to see if we have any errors
  err = db.Ping()
  if err != nil {
    panic(err)
  }

  // successfuly connected
  log.Println("Successfully connected!")

//    sqlStatement := `
// INSERT INTO users (age, email, first_name, last_name)
// VALUES ($1, $2, $3, $4)
// RETURNING id`
//   id := 0
//   err = db.QueryRow(sqlStatement, 30, "ylaz@phd.io", "Yiannis", "Lazarides").Scan(&id)
//   if err != nil {
//     panic(err)
//   }
//   fmt.Println("New record ID is:", id)

var first_name, last_name string
var age int
err = db.QueryRow("select first_name, last_name, age FROM users where last_name = $1", "Antoniades").Scan(&first_name, &last_name, &age)
if err != nil {
  log.Fatal(err)
}
fmt.Println("Success found...", age, first_name, last_name)
}