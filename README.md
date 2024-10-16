Golang microservices example with docker & consul support.

### Microservices presented here are: ###

1. Video metadata service - `metadata` package
2. Video rating service - `rating` package
3. General video service - `video` package, retranslating requests to the above microservices.

Communication is accomplished through grpc framework (the proto file is `api/video.proto`).
