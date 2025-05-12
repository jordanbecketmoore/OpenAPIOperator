# OpenAPIOperator
A Kubernetes operator for automatic HTTP routing according to OpenAPI specifications. 

# Custom Resources 
## OpenApiSpec
The `OpenApiSpec` resource is a straightforward Kubernetes resource
representation of an OpenAPI specification document.
## OpenApiRouter
The `OpenApiRouter` resource will consist of references to a source `Gateway`, a
target `Service`, and an `OpenApiSpec`. 

The `OpenApiRouter` resource will use the `OpenApiSpec` to generate and deploy a set of
`HTTPRoute`s that correspond to the routing rules defined by the `OpenApiSpec`.
These `HTTPRoutes` will be attached to the source `Gateway`, whose traffic they
will route to the target `Service`. The `OpenApiRouter` will also deploy any
`ReferenceGrant`s necessary to satisfy the routing rules. 

The `OpenApiRouter` resource can optionally host a copy of your `OpenApiSpec` on your endpoint, accessible at your desired path and served by the operator `Pod` itself. 
