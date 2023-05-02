module github.com/PasanovicHalid/XWS_Project/BookingService/APIGateway

go 1.20

replace github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC => ../SharedLibraries/gRPC

replace github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/Saga => ../SharedLibraries/Saga

require (
	github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC v0.0.0-00010101000000-000000000000
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.2
	google.golang.org/grpc v1.54.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)
