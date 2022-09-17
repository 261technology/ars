# guin

### RUN DEVELOPMENT ###
```sh run``` 

### RUN BUILD ####
###### docker ######
```docker compose up -d --build```

###### local ######
*LINUX*```GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build```
*OSX*  ```GOOS=darwin GOARCH=amd64  CGO_ENABLED=0 go build```
*START-APP* ```./guin```

### TEST ###
```go test -v ./api/v1/shorten``` 