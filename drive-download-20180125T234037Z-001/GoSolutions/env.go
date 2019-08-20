package env

import (
        "fmt"
        "os"
)

func Vars(){
    fmt.Println(os.Environ())
}

