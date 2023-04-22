module github.com/PasanovicHalid/XWS_Project/BookingService/APIGateway

go 1.20

replace github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC => ../SharedLibraries/gRPC
replace github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/Saga => ../SharedLibraries/Saga

require github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.2

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/genproto v0.0.0-20230223222841-637eb2293923 // indirect
	google.golang.org/grpc v1.53.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)
