## Comando para gerar PB do gRPC
protoc --go_out=application/grpc/pb --go_opt=paths=source_relative --go-grpc_out=application/grpc/pb --go-grpc_opt=paths=source_relative --proto_path=application/grpc/protofiles application/grpc/protofiles/*.proto

## Comand para abrir client do gRPC
evans -r repl
**importante:** para que este client esteja disponível a imagem do docker tem que ter o evans instalado e no StartGrpcServer tem que definir o reflection do gRPC  - - reflection.Register(grpcServer)
