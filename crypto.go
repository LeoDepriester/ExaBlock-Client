package main

import (
  "fmt"
  "github.com/Pallinder/go-randomdata"
  "bytes"
	"golang.org/x/crypto/bcrypt"
	"encoding/base64"
  "os/user"
)


func generate_id() (string){
	var identifier string
  var running bool = true
	for running{
		var buffer bytes.Buffer
	  for i := 0; i < 10; i++{
	    n := fmt.Sprint(randomdata.Number(9))
	    buffer.WriteString(n)
	  }
	  identifier = buffer.String()
		running = db_available_identifier(identifier)
	}
	return identifier
}

func generate_passphrase() (string){
  var paragraph string = randomdata.Paragraph()
  var buffer bytes.Buffer
  for i := 0; i < 20; i++{
    n := fmt.Sprint(randomdata.Number(9))
    buffer.WriteString(n)
  }
  var number string = buffer.String()
  usr, _ := user.Current()
  var username string = usr.Username

  var passphrase string = make_hash(paragraph+number+username)
  return passphrase
}

func make_hfn(filename string) (string){
  var encrypt_fn string = make_hash(filename)
  input := []byte(encrypt_fn)
  var hfn string = base64.StdEncoding.EncodeToString(input)
  return hfn
}

func make_hash(data string) (string){
	bd := []byte(data)
	hash_data, err := bcrypt.GenerateFromPassword(bd, 10)
	if err != nil{
		panic(err)
	}
	return string(hash_data)
}

func encrypt_file(realpath string, hfn string){
  display_informations("Encryption is running ...")

}
