# tarantula
Build the web

## What is tarantula?

Tarantula is a set of tools for practical and robust deployments of web application backends; the main tool groups consist of:

tarantula/log - a logging interface.

tarantula/router - http routing packages.

tarantula/graph - a graph database package.

tarantula/silk - a package which converts a database schema into an API model.

tarantula/markup - html wrapper for SSR webpages.


## Package silk

Package silk will generate your web application API model from the database schema.

A typical system will consist of an instance of *graph.Database, a tarantula/router root node, an object that implements  http.Server, a database client that implements graph.Client interface, and a logger that implements tarantula/log Logger interface.

Using an instance of a *silk.System, vertex classes can be defined in the graph package's database.

```

animals := system.AddClass("animal")

animals.AddProperty("species", validation.StringSet("cat", "dog"))
animals.AddProperty("age", validation.Int())

```

The above class declaration would create these endpoints:

```

/animal/:uid

/animal/:uid/properties/:propKeys

```

Defining classes as users is achieved with the .IsUser() method. Adding this attribute will allow silk to generate the registration and login routes.

```

users := system.AddClass("user").IsUser()

users.AddProperty("name", validation.String(1, 30))
users.AddProperty("age", validation.Int())

```

Relationships between classes can be modelled using the classes' .Link(...) method. The silk package will use these relationships to model corresponding endpoints exposing pre-configured queries to the graph database.


```

users.Link("owning", animals)

```
The above relationship would create:

```

/user/in/owning/count

/user/in/owning/list

/user/in/owning/export

/user/in/owning/animal/count

/user/in/owning/animal/list

/user/in/owning/animal/export

```
