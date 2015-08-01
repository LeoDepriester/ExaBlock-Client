package main

import (
  "fmt"
  "os"
)

func display_help(){
  fmt.Println("Usage: exablock OPTIONS [ARGUMENTS]")
  fmt.Println("Options:")
  fmt.Println("   help        => Display this message")
  fmt.Println("   version     => Print version of this program")
  fmt.Println("   password    => Change password")
  fmt.Println("   list        => Display the list of file on cloud")
  fmt.Println("   send FILE   => Send file on cloud")
  fmt.Println("   get ID      => Get file from cloud")
  fmt.Println("   delete ID   => Delete file on cloud")
  os.Exit(0)
}

func display_version(){
  fmt.Println("ExaBlock version 0.0.1\nWritten by Léo Depriester <leo.depriester@exadot.fr>")
  os.Exit(0)
}


func display_list_files(){
  fmt.Println("ID            NAME")
  ids, names := db_list_files()
  for i := 0; i < len(ids); i++{
    fmt.Println(ids[i], "  ", names[i])
  }
}

func display_error(error_id string, data string){
  if error_id == "AD_ARGV"{
    fmt.Println("The option \"",data,"\" is not available.")
  }else if error_id == "MISSING_ARGV"{
    fmt.Println("You have to add an argument with the command \""+data+"\".\nPlease read help message.")
  }
}


func display_informations(data string){
  fmt.Println("[*]", data)
}

func display_choice(question string, choices []string) (string){
  var choice string
  var running bool = true
  for running{
    fmt.Print(question," ")
    fmt.Scanf("%s", &choice)
    for _, av_choice := range choices{
      if choice == av_choice{
        running = false
      }
    }
    if running{
      fmt.Print("Choose : ")
      for index, av_choice := range choices{
        if len(choices)-1 == index{
          fmt.Println(av_choice)
        }else{
          fmt.Print(av_choice, " or ")
        }
      }
    }
  }
  return choice
}
