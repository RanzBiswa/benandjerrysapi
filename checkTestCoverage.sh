go test -coverprofile=coverage_models.out ./models/
go tool cover -func=coverage_models.out
go tool cover -html=coverage_models.out
go test -coverprofile=coverage_resources_icecreams.out ./resources/icecreams/

# icecreams resources
go tool cover -func=coverage_resources_icecreams.out
go tool cover -html=coverage_resources_icecreams.out	


go test -coverprofile=coverage_handlers.out ./

go tool cover -func=coverage_handlers.out
go tool cover -html=coverage_handlers.out