# valente - websocket for golang webapp

[![license](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)]()
[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/trumae/valente)
[![Go Report Card](https://goreportcard.com/badge/github.com/trumae/valente)](https://goreportcard.com/report/github.com/trumae/valente)
[![wercker status](https://app.wercker.com/status/5a8ce77ffcb15c1e5271849db344fc12/s "wercker status")](https://app.wercker.com/project/bykey/5a8ce77ffcb15c1e5271849db344fc12)

valente is an experiment with Golang webapp using WebSockets. There are similar solutions in another languages/platform:

 * The Wt WebFramework has an experimental features with all comunications between server and browser using websockets. 
 * N2O and Nitrogen are frameworks with that feature coded in Erlang. 

The use of asynchronism is mandatory for this solutions. Traditional threads aren't viable, due to high memory consumed for each connection. 
The Wt Framework is using Boost::asio to handle connections. In Go and Erlang, the languages features should make that scheme simple and scalable. 
I don't know :)

The valente is based on Nitrogen ideas.


## Installing

To start using valente, install Go and run go get:

```bash
$ go get github.com/trumae/valente/...
```

This will retrieve the library and install the valente command line utility into your $GOBIN path.

## How to use 

```bash
$ valente new appsample
$ cd appsample
$ go get
$ go build
$ ./appsample
```

Access http://localhost:8000/ with your browser for boilerplate demo.


