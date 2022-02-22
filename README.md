# 97vaibhav-go-api-tweets-with-clean-architecture-

# project structure 

```

cmd 
----main
      ----main.go                                                                        // starting point of project
      
pkg
----config
       ----- app.go                                                                      //use for databse connection returs db variable to be used in different package
       
----controllers
       ----- tweet-controller.go                                                         // actual logic controllers 
       
----models
       ----- tweet.go                                                                    //struct tweet defined 
       
----routes
       ----- routes.go                                                                   // used for routes 
       
----utils
       ----- utils.go                                                                   // fuction to unmarshel data {json--->byte)
       
```       

# Request End Points


![endpoints drawio](https://user-images.githubusercontent.com/67567763/154916981-c83b2f34-0537-4e1c-9bc9-5480fb8a6dfa.png)

- Daigram made from draw.io




# Self Notes for clean architecture by vaibhav
![CleanArchitecture](https://user-images.githubusercontent.com/67567763/155112076-9f1fd916-da7c-4098-afdf-959e528dd155.jpg)

- Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.
- Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
- Independent of UI. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.
- Independent of Database. You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.
- Independent of any external agency. In fact your business rules simply don’t know anything at all about the outside world.


 ## The Dependency Rule
-The concentric circles represent different areas of software. In general, the further in you go, the higher level the software becomes. The outer circles are mechanisms. The inner circles are policies.
-The Dependency Rule. This rule says that source code dependencies can only point inwards. Nothing in an inner circle can know anything at all about something in an outer circle. In particular, the name of something declared in an outer circle must not be mentioned by the code in the an inner circle. 
 ### in simple lauguage dependency rules says that we can not use function ,variale of outside circle inside inner circle

## Entities.
-Entities encapsulate Enterprise wide business rules. An entity can be an object with methods, or it can be a set of data structures and functions. It doesn’t matter so long as the entities could be used by many different applications in the enterprise.
- If you don’t have an enterprise, and are just writing a single application, then these entities are the business objects of the application. They encapsulate the most general and high-level rules. They are the least likely to change when something external changes. For example, you would not expect these objects to be affected by a change to page navigation, or security. No operational change to any particular application should affect the entity layer.

## Use Cases
- The software in this layer contains ### application specific business rules.
- It encapsulates and implements all of the use cases of the system. These use cases orchestrate the flow of data to and from the entities, and direct those entities to use their enterprise wide business rules to achieve the goals of the use case.


## Interface Adapters

The software in this layer is a set of adapters that convert data from the format most convenient for the use cases and entities, to the format most convenient for some external agency such as the Database or the Web. It is this layer, for example, that will wholly contain the MVC architecture of a GUI. The Presenters, Views, and Controllers all belong in here. The models are likely just data structures that are passed from the controllers to the use cases, and then back from the use cases to the presenters and views.

## Frameworks and Drivers.
outermost layer consist of database ,ui and frameworks


## Crossing boundaries.
- At the lower right of the diagram is an example of how we cross the circle boundaries. It shows the Controllers and Presenters communicating with the Use Cases in the next layer. Note the flow of control. It begins in the controller, moves through the use case, and then winds up executing in the presenter. Note also the source code dependencies. Each one of them points inwards towards the use cases.

- We usually resolve this apparent contradiction by using the Dependency Inversion Principle. In a language like Java, for example, we would arrange interfaces and inheritance relationships such that the source code dependencies oppose the flow of control at just the right points across the boundary.

- For example, consider that the use case needs to call the presenter. However, this call must not be direct because that would violate The Dependency Rule: No name in an outer circle can be mentioned by an inner circle. So we have the use case call an interface (Shown here as Use Case Output Port) in the inner circle, and have the presenter in the outer circle implement it.

- The same technique is used to cross all the boundaries in the architectures. We take advantage of dynamic polymorphism to create source code dependencies that oppose the flow of control so that we can conform to The Dependency Rule no matter what direction the flow of control is going in.


#### What data crosses the boundaries.
Typically the data that crosses the boundaries is simple data structures.

- https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html
- https://manakuro.medium.com/clean-architecture-with-go-bce409427d31 (best resource)


# References
- https://github.com/bmf-san/go-clean-architecture-web-application-boilerplate

- http://c.xkdl666.xyz/takashabe/go-ddd-sample
