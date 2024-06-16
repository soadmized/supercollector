gen-grpc:
	cd api/grpc && protoc --go_out=../../internal/ --go-grpc_out=../../internal/ --grpc-gateway_out=../../internal/ --grpc-gateway_opt generate_unbound_methods=true --openapiv2_out ../swagger --openapiv2_opt generate_unbound_methods=true supercollector.proto
