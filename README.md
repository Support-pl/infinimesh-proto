# infinimesh API Proto

This repo containes protobuf files and Go generated code for the infinimesh APIs.

## Install

### Go

At your project(where go.mod is):

```shell
go get github.com/infinimesh/proto@latest
```

### Node (Connect ES)

> Install protoc before going further: [manual](https://grpc.io/docs/protoc-installation/)

```shell
# Clone repo
git clone git@github.com:infinimesh/proto.git
cd infinimesh-proto

# Install generators
npm i
npm run generate
```

## How to use

### Generating

1. Clone [this repo](github.com/infinimesh/proto)
2. Navigate to cloned repo directory

#### Docker

Just run:

```shell
docker run -it \
  -v $(pwd):/go/src/github.com/infinimesh/proto \
  ghcr.io/infinimesh/proto/buf:latest
```

#### Buf

Set up [`buf`](buf.build), follow Dockerfile for additional dependencies.

Run `buf generate`

## How to generate code for you language

Check according module for buf, add it to `buf.gen.yaml`, run `buf generate`.

If you're generating using Docker, you'd have to build image first.
