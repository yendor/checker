// You can edit this code!
// Click here and start typing.
package main

import "fmt"
import "io/ioutil"
import "os"
import "flag"

func main() {
    flag.Parse()

    checksums := FileChecksums(flag.Arg(0))
    for k, v := range checksums {
        fmt.Printf( "The checksum of %s is %s\n", k, v)
    }
}

func FileChecksums(dir string) map[string]string {
    files, err := ioutil.ReadDir(dir)
    if err != nil {
        fmt.Print(err)
        os.Exit(1)
    }
    hashes := make(map[string]string)
    hash := "abcdef"


    for i := 0; i < len(files); i++ {
        fullpath := dir + "/" + files[i].Name()
        if files[i].IsDir() {
            for k, v := range FileChecksums(fullpath) {
                hashes[k] = v
            }
        } else {
            hashes[fullpath] = hash
        }
    }

    return hashes
}
