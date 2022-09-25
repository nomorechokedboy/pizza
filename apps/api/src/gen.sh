npx grpc_tools_node_protoc \
--js_out=import_style=commonjs,binary:./ \
--grpc_out=./src/grpc/ \
--ts_out=./src/grpc-ts/ \
./src/house/house.proto