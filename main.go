package main

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"github.com/yookoala/realpath"
)


func check_config_directory(){
	usr, _ := user.Current()
	if _, err := os.Stat(usr.HomeDir+"/.exablock/"); os.IsNotExist(err) {
		os.Mkdir(usr.HomeDir+"/.exablock", 0777)
		os.Mkdir(usr.HomeDir+"/.exablock/files", 0777)
		db_create()
		var password string = display_get_pass("Enter a password (remember it, it will encrypt your files) : ")
		var hashed_password string = make_hash(password)
		db_insert_password(hashed_password)
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
