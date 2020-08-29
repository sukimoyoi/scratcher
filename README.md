# scratcher
create base image with go

# Usage

### Create Script

```sh
$ cd script
$ go build
```

cross compile example

```sh
$ GOOS=linux GOARCH=amd64 go build
```


### Move script file

```sh
$ cp script ../work/
```

### Create tar.gz file

```sh
$ cd .../work
$ tar cvzf ../work.tar.gz *
```

### Create Docker Image

```sh
$ cd ..  # pwd: scratch_test
$ go run main.go
```
