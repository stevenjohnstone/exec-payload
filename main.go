package main

import (
  "os"
  "encoding/base64"
  "fmt"
)

func main() {
  command := os.Args[1]

  fmt.Printf("bash -c {echo, %s}|{base64, -d}|{bash,-i}", base64.StdEncoding.EncodeToString([]byte(command)))
}
