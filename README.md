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

# References
- https://github.com/bmf-san/go-clean-architecture-web-application-boilerplate

- http://c.xkdl666.xyz/takashabe/go-ddd-sample
