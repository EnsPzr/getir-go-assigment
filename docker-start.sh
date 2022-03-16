docker rm getir-go-assigment
docker rmi getir-go-assigment
docker build -t getir-go-assigment .
docker run -it -p 8080:8080 --name getir-go-assigment getir-go-assigment