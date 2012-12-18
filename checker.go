// You can edit this code!
// Click here and start typing.
package main

import "fmt"
import "io/ioutil"
import "os"
import "flag"
import "crypto/md5"
import "io"
//import "bytes"

func main() {
    flag.Parse()

    start := flag.Arg(0)

    stat, err := os.Stat(start)
    if err != nil {
        usage();
        os.Exit(1);
        panic(err)
    }

    if (stat.IsDir()) {
        FileChecksums(flag.Arg(0))
    } else {
        usage();
    }
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
            hash = GetFileHash(fullpath)
            fmt.Printf( "%s  %s\n", hash, fullpath)
        }
    }
}

func GetFileHash(filepath string) string {

    fi, err := os.Open(filepath)
    if err != nil { panic(err) }
    defer fi.Close()

    running_hash := md5.New()
    io.Copy(running_hash, fi)
    return fmt.Sprintf("%x", running_hash.Sum(nil))
}

func usage() {
    fmt.Printf("Usage checker directorypath\n")
}
