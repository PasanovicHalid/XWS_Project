module github.com/PasanovicHalid/XWS_Project/BookingService/AccommodationService

go 1.20

replace github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC => ../SharedLibraries/gRPC

replace github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/Saga => ../SharedLibraries/Saga

require (
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	go.mongodb.org/mongo-driver v1.11.4
	google.golang.org/grpc v1.54.0
)

require github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.2 // indirect

require (
	github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC v0.0.0-20230512223728-3bd7451f1839
	github.com/golang/protobuf v1.5.3
	github.com/golang/snappy v0.0.1 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.1 // indirect
	github.com/xdg-go/stringprep v1.0.3 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)
