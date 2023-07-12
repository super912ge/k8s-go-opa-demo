# Learning Roadmap 1 - Dev Route

The Goal of the first phase is to create a small Go application and run it in a local Kuberntes Cluster. There are many thing to learn along the way. This roodmap details all the knowledge you need to acwquire to complete this phase as well as question you should be able to answer. Those questions are ideal to test your knowledge and prepare for interviews.

In this initial phase build an application that run an HTTP server and exposes endpoints:
* GET `/v1/users` - return a list of users hard coded in the application
* GET `/v1/teams` - return a list of teams hard-coded in the application

Since you plan to deploy that application to Kubernetes you must build a container image to run in your local Kubernetes cluster.
* Write a Dockerfile to build your application

## Golang

* [ ] How do you represent data structure in Go?
* [ ] Is Go object-oriented?
* [ ] How do you define varaibles in Go?
* [ ] How do you use pointers Go?
* [ ] What is a `struct`?
* [ ] What is an `interface` in Go?
* [ ] How to define methods in Go?

### Application lifecycle

* [ ] How to structure you code repository?
* [ ] How to run your application locally?
* [ ] How to write unit test in golang?
* [ ] How to run unit tests?
* [ ] How to build an application binary?

### HTTP API

* [ ] How would you create an http server in Go?
* [ ] What libray would you use to simplify the HTTPs server creation?
* [ ] How do you convert JSON request and response to Go struct?
* [ ] 

## Docker

* [ ] What is Docker?
* [ ] What is the difference between a Docker Image and Docker Container?
* [ ] What is a Dockerfile?
* [ ] What are the most commun instruction in a Dockerfle?
* [ ] What is a multi-stage