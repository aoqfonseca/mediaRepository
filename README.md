##mediaRepository
==============
The mediaRepository is a server write in Go for upload file. I use this for learning Go.

##Contribuing
=============

This tool assumes you are working in a standard Go workspace,
as described in http://golang.org/doc/code.html. We require Go 1.1
or newer to build godep itself, but you can use it on any project
that works with Go 1 or newer.

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

### Build

In your root project dir, run this command:

    $go build webserver/server.go

This will create runable file called server. To run:

    $ ./server --config config.yml

Voilá, your serving is up and runing
