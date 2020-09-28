dep:
	sudo apt-get install mysql-server;
	mysql_secure_installation;
	echo "EDIT cmd/main.go FOR DATABASE AND CREATE ONE;"

build:
	go build -o builds/api_gateway api_gateway/cmd/main.go
	go build -o builds/file_service services/files/cmd/main.go
	go build -o builds/user_service services/users/cmd/main.go