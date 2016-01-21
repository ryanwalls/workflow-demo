# workflow
Repo for interacting with AWS Simple Workflow Service (SWF)

## Setup Instructions

### Install Go (Windows)
Install chcolatey if you have not already done so (https://chocolatey.org/)
```
choco install golang
```

### Install Go (Mac)
https://golang.org/doc/install

### IDE
* Install an IDE.  Recommend either LiteIDE https://code.google.com/p/golangide or Atom https://atom.io
* Create a directory for your go projects, ex: `c:\source\go`

#### Setting up Intellisense in LiteIDE (Windows)
```
 go get -u -ldflags -H=windowsgui github.com/nsf/gocode
 ```
#### Setting up Intellisense in Atom (Mac/Windows)
Install `go-plus` package https://github.com/joefitzgerald/go-plus

If Intellisense isn't working, try the following
```
gocode close
gocode -s -debug
```

### Setting up Code directories
Assuming your workspace folder is `c:\source\go`, setup  environment variable `GOPATH = c:\source\go`

Create bin, pkg, src sub-folders in `$GOPATH`.

Create `$GOPATH\src\github.com\3dsim` folder.

In the `$GOPATH\src\github.com\3dsim` folder, clone the your fork of the workflow repo. When finished you should have the folder: `$GOPATH\src\github.com\3dsim\workflow`

### Packages
We use go 1.5's [new vendoring feature](https://docs.google.com/document/d/1Bz5-UB7g2uPBdOx-rw5t9MxJwkfpx90cqG9AFL0JAYo/edit) and https://github.com/Masterminds/glide to manage package dependencies.

1.  Set an environment variable:  `GO15VENDOREXPERIMENT=1`
2.  Install glide and add it to your path.  See https://github.com/Masterminds/glide#install
3.  Inside `solver-svc` directory run  `glide up`.  This downloads all package dependencies
to a `vendor/` folder.
4.  If you need to add more dependencies to the project at any point run
`glide get <package name>`, which will add the package to the `glide.yaml`.  Make
sure to update the `rev:` field of the package in `glide.yaml` to the latest revision.  
Running `glide pin` will print to standard out the latest revision of every package.

### Running tests
Golang has testing built in.  

To run tests for a single package run:
```
go test ./<package name>
```

To run all tests run:
```
go test $(glide novendor)
```

## Automated deployments via Jenkins (Preferred)

TODO

## To run locally
* From root directory of project run
```
go install
```
* To see what commands you can run:
```
workflow
```

## Coding tips

### Testing
* Generate mocks using https://github.com/vektra/mockery

```
go get github.com/vektra/mockery/.../
$GOPATH/bin/mockery -name <interface name> -recursive
```

If you need to generate mocks in the same package to avoid circular dependencies use
```
$GOPATH/bin/mockery -name <interface name> -recursive -inpkg
```

## Contributing code
Read this article and follow the steps they outline: http://scottchacon.com/2011/08/31/github-flow.html

All PRs should be signed off by a member of the team before merging.

## Team
* Tim Sublette
* Ryan Walls
* Chad Queen
* Pete Krull

## Original release
January 2016
