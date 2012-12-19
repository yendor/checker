// You can edit this code!
// Click here and start typing.
package main

import "fmt"
import "io/ioutil"
import "os"
import "flag"
import "crypto/md5"
import "io"
import "bytes"
import "launchpad.net/gommap"

func main() {
    flag.Parse()

    start := flag.Arg(0)
_, err := os.Stat(start)
    if err != nil {
        usage();
        os.Exit(1);
        panic(err)
    }

    FileChecksums(flag.Arg(0))
}

func FileChecksums(dir string) {
    stat, err := os.Stat(dir)
    if err != nil {
        panic(err)
    }

    hash := "abcdef"

    if (stat.IsDir()) {
        files, err := ioutil.ReadDir(dir)

        if err != nil {
            panic(err)
        }

        for i := 0; i < len(files); i++ {
            fullpath := dir + "/" + files[i].Name()

            if files[i].IsDir() {
                FileChecksums(fullpath);
            } else {
                hash = GetFileHash(fullpath)
                fmt.Printf( "%s  %s\n", hash, fullpath)
            }
        }
    } else {
        hash = GetFileHash(dir)
        fmt.Printf( "%s  %s\n", hash, dir)
    }
}

func GetFileHash(filepath string) string {

    fi, err := os.Open(filepath)
    if err != nil { panic(err) }
    defer fi.Close()

    running_hash := md5.New()
    //io.Copy(running_hash, fi)
    mmap, err := gommap.Map(fi.Fd(), gommap.PROT_READ, gommap.MAP_PRIVATE)
    if err != nil {
        panic(err)
    }
    //mmapReader := bytes.NewReader(mmap)
    //io.Copy(running_hash, mmapReader)
    buf := bytes.NewBuffer(mmap)
    io.Copy(running_hash, buf)

    return fmt.Sprintf("%x", running_hash.Sum(nil))
}

func usage() {
    fmt.Printf("Usage checker directorypath\n")
}
