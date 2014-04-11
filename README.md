##mediaRepository
The mediaRepository is a server write in Go for upload file. I use this for learning Go.

##Contribuing

This tool assumes you are working in a standard Go workspace,
as described in http://golang.org/doc/code.html. We require Go 1.1
or newer to build this itself, but you can use it on any project
that works with Go 1 or newer.


##GIT Config

. Please, use rebase for pull and merge.

    $ git pull -rebase origin master

. When you finish some feature, put, in commit message, ISSUE NUMBER associate with it.


###Configuring GOPATH

In your bash, export environment variable to define GOPATH. This path will be the local whre Go put all libs, bin, etc.

    $ EXPORT GOPATH=$HOME/.go

###Install godeps

    $go get github.com/tools/godep

###Creating symbol

Create a symbol link in $GOPATH/src for your project.

    $ cd $GOPATH/src
    $ mkdir -p github.com
    $ cd github.com
    $ ln -s $HOME/Project/mediaRepository mediaRepository

### Install Deps and run testing

Go to your project dir, and running godeps to instal all dependencies

    $godeps restore

When it's finished:

    $go test ./...

### MAKE SETUP

All steps showed above, you can do calling:

    $ make setup

### Build

In your root project dir, run this command:

    $go build webserver/server.go

This will create runable file called server. To run:

    $ ./server --config config.yml

Voilá, your serving is up and runing


##Benchmark


    $ Server Software:
    $ Server Hostname:        localhost
    $ Server Port:            4321
    $
    $ Document Path:          /upload
    $ Document Length:        7 bytes
    $
    $ Concurrency Level:      450
    $ Time taken for tests:   31.300 seconds
    $ Complete requests:      50000
    $ Failed requests:        0
    $ Write errors:           0
    $ Total transferred:      7100000 bytes
    $ Total body sent:        1265100000
    $ HTML transferred:       350000 bytes
    $ Requests per second:    1597.47 [#/sec] (mean)
    $ Time per request:       281.696 [ms] (mean)
    $ Time per request:       0.626 [ms] (mean, across all concurrent requests)
    $ Transfer rate:          221.52 [Kbytes/sec] received
    $                         39471.80 kb/s sent
    $                         39693.33 kb/s total
    $
    $ Connection Times (ms)
    $               min  mean[+/-sd] median   max
    $ Connect:        0  103 446.2      0    7016
    $ Processing:     1  173 175.4    142    3321
    $ Waiting:        1  172 175.0    141    3321
    $ Total:          1  277 479.2    150    7430
