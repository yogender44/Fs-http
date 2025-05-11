package main
import (
 "fmt"
 "net"
)

func main() {
  listener, err := net.listen("tcp", ":1729")
  if err != nil {
    log.Fatal(err)
  }
  fmt.println("Hello world!", listener)
}
