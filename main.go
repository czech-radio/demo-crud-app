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
    Id: 1,
    GivenName: "Pepa",
    FamilyName: "Dvorak",
    Affiliation: "BEZPP",
    Gender: 1,
    Foreigner: 0,
    Labels: "fyzik",
  }

  // CREATE
  createdPerson1, err := personsRepository.Create(mockperson1)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(createdPerson1)

  createdPerson2, err := personsRepository.Create(mockperson2)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(createdPerson2)

  fmt.Println("CREATE SUCCESSFUL")


  // RETURN ALL
  getAllPersons, err := personsRepository.All()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(getAllPersons)

  // RETURN ONE
  getOnePerson, err := personsRepository.GetByFamilyName("Dvorak")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(getOnePerson)

  fmt.Println("RETURN SUCCESSFUL")

  // UPDATE
  createdPerson1.Affiliation = "Pirati"
  if _, err := personsRepository.Update(createdPerson1.Id, *createdPerson1); err != nil {
    log.Fatal(err)
  }
  fmt.Println(personsRepository.All())
  fmt.Println("UPDATE SUCCESSFUL")

  // DELETE
  if err := personsRepository.Delete(createdPerson1.Id); err != nil {
    log.Fatal(err)
  }
  fmt.Println(personsRepository.All())
  
  if err := personsRepository.Delete(createdPerson2.Id); err != nil {
    log.Fatal(err)
  }
  fmt.Println(personsRepository.All())
  fmt.Println("DELETE SUCCESSFUL")

}

