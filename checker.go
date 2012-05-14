// You can edit this code!
// Click here and start typing.
package main

import "fmt"
import "io/ioutil"
import "os"

func main() {
    fmt.Println("Hello, ..")
    files, err := ioutil.ReadDir(".")
    if err != nil {
        fmt.Print(err)
        os.Exit(1)
    }
    for i := 0; i < len(files); i++ {
        fmt.Printf("%s\n", files[i].Name())
    }
    os.Exit(0)
}

