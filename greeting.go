package greeting

import "fmt"
import "errors"

func Hello(name string) (string , error){
  if(name =="") {
    return "", errors.New("empty name")
  }
  msg := fmt.Sprintf("Hi %v", name)
  return msg, nil
}
