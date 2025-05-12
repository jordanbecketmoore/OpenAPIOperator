# OpenAPIOperator
A Kubernetes operator for automatic HTTP routing according to OpenAPI specifications. 

# Custom Resources 
## OpenApiSpec
The `OpenApiSpec` resource is a straightforward Kubernetes resource
representation of an OpenAPI specification document, managed via the [libopenapi](https://github.com/pb33f/libopenapi) Go library. 
## OpenApi
The `OpenApi` resource will consist of references to a source `Gateway`, a
target `Service`, and an `OpenApiSpec`. 

The `OpenApi` resource will use the `OpenApiSpec` to generate a set of
`HTTPRoute`s that correspond to the routing rules defined by the `OpenApiSpec`.
These `HTTPRoutes` will be attached to the source `Gateway`, whose traffic they
will route to the target `Service`. 

The `OpenApi` resource can optionally host a copy of your `OpenApiSpec` on your desired endpoint 