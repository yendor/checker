// You can edit this code!
// Click here and start typing.
package main

import "fmt"
import "io/ioutil"
import "os"
import "flag"
import "crypto/md5"
//import "io"
//import "bytes"

func main() {
    flag.Parse()

    // TODO: Make sure the supplied argument is a directory
    FileChecksums(flag.Arg(0))
}

func FileChecksums(dir string) {
    _, err := os.Stat(dir)
    if err != nil {
        panic(err)
    }

    files, err := ioutil.ReadDir(dir)

    if err != nil {
        panic(err)
    }

    hash := "abcdef"

    for i := 0; i < len(files); i++ {
        fullpath := dir + "/" + files[i].Name()

        if files[i].IsDir() {
            FileChecksums(fullpath);
        } else {
            // TODO: Calculate the hash of the file
            hash = FileChecksum(fullpath)
            fmt.Printf( "%s  %s\n", hash, fullpath)
        }
    }
}

func FileChecksum(filepath string) string {
    contents, err := ioutil.ReadFile(filepath)
    if err != nil {
        panic(err)
    }
    running_hash := md5.New()
    running_hash.Write(contents)
    return fmt.Sprintf("%x", running_hash.Sum(nil))
}
