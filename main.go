package main

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"github.com/yookoala/realpath"
	"io/ioutil"
)


func check_config_directory(){
	usr, _ := user.Current()
	if _, err := os.Stat("/path/to/whatever"); os.IsNotExist(err) {
		os.Mkdir(usr.HomeDir+"/.exablock", 0777)
		os.Mkdir(usr.HomeDir+"/.exablock/files", 0777)
		db_create()
	}
}

func send_file(filename string){
	var identifier string = generate_id()
	realpath, _ := realpath.Realpath(filename)
	filename = path.Base(filename)
	var hfn string = make_hash(filename)
	var host string = "local.dev"
	encrypt_file(realpath, hfn)
	fmt.Println(host, identifier)
}

func change_password(){
	var choice string = display_choice("Do you want change your password ? (y/n)", []string {"y", "n"})
	if choice == "n"{os.Exit(0)}else{
		var choice string = display_choice("Do you want to auto generate a passphrase or write one ? (generate/write)", []string {"generate", "write"})
		if choice == "generate"{
			var password string = generate_passphrase()
			d := []byte(password)
			usr, _ := user.Current()
			err := ioutil.WriteFile(usr.HomeDir+"/.exablock/password.txt", d, 0777)
			if err != nil{
				fmt.Println(err)
			}
		}else{

		}
	}
}

func main(){
	// Check config directory
	check_config_directory()

  if(len(os.Args) == 1){
    display_help()
  }else{
    switch os.Args[1]{
      case "help":
        display_help()
      case "version":
        display_version()
      case "list":
        display_list_files()
			case "password":
				change_password()
			case "send":
				if len(os.Args) == 3{
					send_file(os.Args[2])
				}else{
					display_error("MISSING_ARGV", os.Args[1])
				}
      default:
        display_error("BAD_ARGV", os.Args[1])
    }
  }
}
