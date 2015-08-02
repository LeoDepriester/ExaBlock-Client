package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
  "os/user"
)

func db_create(){
  usr, err := user.Current()

  db, err := sql.Open("sqlite3", usr.HomeDir+"/.exablock/exablock.db")
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  sqlStmt_files := `
  CREATE TABLE IF NOT EXISTS files(
    id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
    identifier INTEGER,
    name TEXT,
    hfn TEXT,
    host TEXT
    )`
  _, err = db.Exec(sqlStmt_files)
  if err != nil {
    log.Fatal(err)
  }
	sqlStmt_infos := `
	CREATE TABLE IF NOT EXISTS infos(
		id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
		password TEXT
		)`
	_, err = db.Exec(sqlStmt_infos)
	if err != nil{
		log.Fatal(err)
	}
}

func db_list_files() (ids []int, names []string){

  usr, err := user.Current()
  if err != nil{log.Fatal(err)}

  db, err := sql.Open("sqlite3", usr.HomeDir+"/.exablock/exablock.db")
  if err != nil {log.Fatal(err)}
  defer db.Close()

  rows, err := db.Query("SELECT identifier, name FROM files")
  if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()

  for rows.Next(){
    var id int
    var name string
    rows.Scan(&id, &name)
    ids = append(ids, id)
    names = append(names, name)
  }
  rows.Close()
  return
}

func db_available_identifier(identifier string) (bool){
  usr, err := user.Current()
  if err != nil{log.Fatal(err)}

  db, err := sql.Open("sqlite3", usr.HomeDir+"/.exablock/exablock.db")
  if err != nil {log.Fatal(err)}
  defer db.Close()


	stmt, err := db.Prepare("SELECT id FROM files WHERE identifier = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var id int
	err = stmt.QueryRow(identifier).Scan(&id)
  if id == 0{
    return false
  }else{
    return true
  }
}
func db_insert_password(password string){
	usr, err := user.Current()
  if err != nil{log.Fatal(err)}

  db, err := sql.Open("sqlite3", usr.HomeDir+"/.exablock/exablock.db")
  if err != nil {log.Fatal(err)}
  defer db.Close()


	stmt, err := db.Prepare("INSERT INTO infos(password) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var id int
	err = stmt.QueryRow(password).Scan(&id)
}
