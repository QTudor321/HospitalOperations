.DEFAULT_GOAL:=build
fmt:
	go fmt Launcher.go
	go fmt logic/doctor/ClientDoc.go
	go fmt network/Server.go
	go fmt database/DoctorsDatabase.go
.PHONY:fmt
lint:fmt
	golint Launcher.go
	golint logic/doctor/ClientDoc.go
	golint network/Server.go
	golint database/DoctorsDatabase.go
.PHONY:lint
vet:fmt
	go vet Launcher.go
	go vet logic/doctor/ClientDoc.go
	go vet network/Server.go
	go vet database/DoctorsDatabase.go
.PHONY:vet
build:vet
	go build -o Launcher.exe
.PHONY:build