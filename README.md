# K U B E

Build image with dokka

```shell
$ docker build -t kube-dev .
```

Run image

```shell
# This also mounts local /src to container for live reloading, yay
$ docker run -it --rm -p 3000:3000 v $PWD/src:/go/src/kube kube-dev
```

Go into `/src` and run

```shell
$ go mod download
$ go mod vendor
$ go mod verify
```

Docker stuff in the project root
