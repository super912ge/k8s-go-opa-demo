Following this tutorial [An easy and practical approach to structuring Golang applications](https://levelup.gitconnected.com/a-practical-approach-to-structuring-go-applications-7f77d7f9c189), 
I learned that in main.go, setting up web server, defining routes(could also be defined in its own file), setting up db connections, initialize dependencies, starting the server. 

I also learned some useful command, such as go tidy, go mod init, go get "dependency url", go run cmd/server/main.go, go build etc

Go has 2 types of ? : struct and interface

To implement a function to a struct, it works as follow: 

```
func (s *StructName) funcName() returnType {
    return returnType
}
```

---

Docker initdb script: 

By mount a volume to docker postgres image as follow, 
```
    volumes:
      - ./initdb:/docker-entrypoint-initdb.d
```
one can add init-db.sh to initialize db when db is first created.


