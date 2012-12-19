#!/usr/bin/env python

import hashlib, getopt, sys, os

class Checker:
    def __init__(self):
        self.start_paths = None
        shortargs = ''
        longargs = []
        options, self.start_paths = getopt.getopt(sys.argv[1:], shortargs, longargs)


        if self.start_paths == None:
            print >> sys.stderr, "You must specify the path(s) to scan"
            sys.exit(1)

    def search(self):
        for path in set(self.start_paths):
            full_path = os.path.abspath(path)
            if os.path.exists(full_path):
                self.walk(full_path)
            else:
                print >> sys.stderr, "WARNING: The path %s does not exist" % (full_path)

    def walk(self, start):
        try:
            if os.path.isfile(start):
                (basename, ext) = os.path.splitext(start)
                ext = ext.replace('.', '', 1)
                self.hashfile(start)
            elif os.path.isdir(start) and not os.path.islink(start):
                for name in os.listdir(start):
                    path = os.path.join(start,name)
                    self.walk(path)

        except OSError:
            traceback.print_exc(file=sys.stdout)
            if self.verbose == True:
                print >> sys.stderr, "Permission denied to %s" % (dir)

    def hashfile(self, path):
        chunksize = 65536
        m = hashlib.md5()
        with open(path) as f:
            data = f.read(chunksize)
            while not f.closed:
                m.update(data)
                data = f.read(chunksize)
                if len(data) <= 0:
                    break


        print "%s  %s" % (m.hexdigest(), path)
        sys.stdout.flush()


if __name__ == "__main__":
    s = Checker()
    s.search()
