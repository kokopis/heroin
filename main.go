package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
  file, err := os.ReadFile(os.Getenv("APP_LIST"))

  if err != nil {
    log.Panic(err)
  }

  a := getAppsByJson(string(file))

  for _, v := range a {
    ver := v.getRedisVersionString()
    if ver != "" {
      fmt.Println(v.Name)
      fmt.Println(ver)
    }
  }
}
