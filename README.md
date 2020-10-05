# gorest
Web programming examples in Golang.

- Each folder is meant to be run not from the root of the repo, but,
from within each individual folder as:

```
// The repo structure is the following:
	~/go/src/github.com/ren-zxcyq/gorest$ tree
	.
	├── muxRouter
	│   └── main.go
	├── queryParameters
	│   └── main.go
	├── README.md

go run ~/go/src/github.com/ren-zxcq/gorest/muxRouter/main.go
```

- if you want to know what a "module" does, in respect to the repo structure shown above:
```
go doc gorest/muxRouter
```
