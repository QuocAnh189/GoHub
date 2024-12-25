protobuf-review:
	protoc --go_out=./proto/gen/pb_reviews --go-grpc_out=./proto/gen/pb_reviews ./proto/review/review.proto