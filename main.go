package main


import (
  "fmt"
  "log"
  "os"
  "database"
  "github.com/mattn/go-sqlite3"
)


const fileName = "persons.db"

func main(){

    os.Remove(fileName)
  
    db, err := sqlOpen("sqlite3", fileName)
    if err != nil {
      log.Fatal(err)
    }

}

