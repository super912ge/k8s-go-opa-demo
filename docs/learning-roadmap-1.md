# Learning Roadmap 1 - Dev Route

The Goal of the first phase is to create a small Go application and run it in a local Kuberntes Cluster. There are many thing to learn along the way. This roodmap details all the knowledge you need to acwquire to complete this phase as well as question you should be able to answer. Those questions are ideal to test your knowledge and prepare for interviews.

In this initial phase build an application that run an HTTP server and exposes endpoints:
* GET `/v1/users` - return a list of users hard coded in the application
* GET `/v1/teams` - return a list of teams hard-coded in the application

Since you plan to deploy that application to Kubernetes you must build a container image to run in your local Kubernetes cluster.
* Write a Dockerfile to build your application
* Write a Kuberntes manifest defining a `Deployment` for you application

Deploy your application to you local Kubernetes cluster and create a `port-forward` to test it work as expected

## Golang

* [ ] How do you represent data structure in Go?
* [ ] Is Go object-oriented?
* [ ] How do you define varaibles in Go?
* [ ] How do you use pointers Go?
* [ ] What is a `struct`?
* [ ] What is an `interface` in Go?
* [ ] How to define methods in Go?
* [ ] How do you manage the scope of methods and attribute in Go?
* [ ] How can you create a loop in Go?
* [ ] How can you handle condition in Go?

### Application lifecycle

* [ ] How to structure you code repository?
* [ ] How to run your application locally?
* [ ] How to write unit test in golang?
* [ ] How to run unit tests?
* [ ] How to build an application binary?
* [ ] What is a package in Go?
* [ ] What is a module in Go?
* [ ] How do you install a new module?

### HTTP API

* [ ] How would you create an http server in Go?
* [ ] What libray would you use to simplify the HTTPs server creation?
* [ ] How do you convert JSON request and response to Go struct?

## Docker

### Basics

* [ ] What is Docker?
* [ ] What is the difference between a Docker Image and Docker Container?
* [ ] What is a Dockerfile?
* [ ] What are the most commun instruction in a Dockerfle?
* [ ] How do you build a container image with Docker?
* [ ] How do you run a container?
* [ ] How can you make you application running in the container reachable from your host machine?

### Advanced

* [ ] What is a multi-stage build?
* [ ] What is the advantage of a multistage build?
* [ ] How does build stages interact with each others?
* [ ] What is the result of a multi-stage Docker build?
* [ ] What is a layer in a Docker Image?
* [ ] Why is it considere best practices to minimize the number of layers?
* [ ] How can you minimise the number of layers in your Docker images?

### General Knowledge

* [ ] What are the most popular use of Docker?
* [ ] How do you run docker on MacOS and Windows?
* [ ] Where can you find existing Docker images?

### Paying Attention to details

* [ ] What is the difference between `ADD` and `COPY` in a Dockerfile?

## Kubernetes

### General Knowledges

* [ ] What is Kubernetes?
* [ ] How is Kubernetes different from Docker?
* [ ] Does Kubernetes uses Docker?
* [ ] What is the general architecture of Kubernetes?
* [ ] How can you create a local Kubernetes cluster?
* [ ] How do you deploy application to Kubernetes?

### Using Kubernetes

* [ ] What is `kubectl`?
* [ ] What is a Pod?
* [ ] What is a Deployement?
* [ ] How can you run a container in Kubernetes?
