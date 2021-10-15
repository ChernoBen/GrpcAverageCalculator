#!/bin/bash
protoc src/average_proto/average.proto --go_out=. --go-grpc_out=.