package main

import (
        "database/sql"
	"crud_go/repository"
	"fmt"
	"log"
	"os"
)


const fileName = "persons.db"

func main(){

    os.Remove(fileName)
  
    db, err := sql.Open("sqlite3", fileName)
    if err != nil {
      log.Fatal(err)
    }

    personsRepository := repository.NewSQLiteRepository(db)

    if err := personsRepository.Migrate(); err != nil {
      log.Fatal(err)
    }

    mockperson1 := repository.Person{
      Id: 0,
      GivenName: "Jan",
      FamilyName: "Novak",
      Affiliation: "BEZPP",
      Gender: 1,
      Foreigner: 0,
      Labels: "soustruznik",
    }

    mockperson2 := repository.Person{
      Id: 0,
      GivenName: "Pepa",
      FamilyName: "Dvorak",
      Affiliation: "BEZPP",
      Gender: 1,
      Foreigner: 0,
      Labels: "fylosof",
    }

    // CREATE
    createPerson1, err := personsRepository.Create(mockperson1)
    if err != nil {
      log.Fatal(err)
    }
    fmt.Println(createPerson1)

    createPerson2, err := personsRepository.Create(mockperson2)
    if err != nil {
      log.Fatal(err)
    }
    fmt.Println(createPerson2)

    // GET ALL
    getAllPersons, err := personsRepository.All()
    if err != nil {
      log.Fatal(err)
    }
    fmt.Println(getAllPersons)

}

