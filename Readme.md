# Architecture

Team developed by __Hexagonal Architecture__ and __Clean Architecture__  mindset especially   __Dependency Injection__  and __S.O.L.I.D__


## microservice Structure


each project have two main part __core__ and __adapter__.

### core
core is the main part of our app and contains business rules , port, domain and dto.
 and port (dependencies model) at the 

#### Usecase
 __usecase__ is a business rules but this component need other part for working well like dao
we model other part in port  and inject implementation of port on usecase
you can find   at the __app/internal/core/usecase__

### Port
port is a model of our dependency
you can find port on __app/internal/core/port__


#### DTO
dto or data transfer object is a standard for send and recive data from usecase 


### Adapter
each adapter is implementation of one port that injected in usecase we impliment somthing like query in this part. 

![This is an alt text.](./hexa.png)




# THE GOLDEN RULE IS USECASE NOT DEPEND ON ANY THING

![This is an alt text.](./clean.jpg)




### Database Migration
you can find database migration on __app/assets__


## App lifecycle , What happens when we run cmd/main.go?

1. Read configs from env vars
2. In the second step we prepering our __driven port__ like database. at the end of this step we have a implementation of driven port
    1. when we need local data in this step we have a __database instance__.
    2. when we need a data from other microservices we have __sync or async commenication__ object like grpc or rabbitmq
3. In the third step we call the usecase and pass objects of step 2 for use as dao in usecase
4. In the last Pass usecase in the driver port implementation.
