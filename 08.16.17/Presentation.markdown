# Building microservices with Go-Micro

Kansas City Go-lang Meetup

08/16/17

---

## Admin stuff

- This presentation is in markdown, and it is uploaded to the [KC_GOLAMG repo](https://github.com/NDari/KC_GOLANG)
- Built with [markdown-to-slides](https://github.com/partageit/markdown-to-slides)
- Will make the repository under its own name-space, and share access with the admins.
- If you have any old presentations that you want to put here, please let me know to make a PR!
- We are looking for presentations for September and October. Let us know ASAP if you want to present!

---

## The "mature point of view"

- What is a microservice anyway?
  - Taking ques from OOP
  - [Look at this talk](https://www.youtube.com/watch?v=JXEjAwNWays&t)
- What are the advantages of microservices?
  - From that video: Solves org problems, causes tech problems.
  - The interface problem
  - Independently developed parts speaking a common format (e.g JSON)
- Drawbacks?
  - Increased surface area
  - You need an ops team or an ops culture.
  - From that video: 40 things you have to worry about for each and every MS
    among which is: language, architecture, transport protocol, on call rotation,
    logging, license issues, CI/CD pipelines, API documentation, etc.
- When should you use a microservice?
  - Maybe monolith first is a good approach...
- Have a sit. have a good logical discussion about your needs, and decide!

---

## Meh. Logical reasons are overrated.

![Alt text](burnitall.png)

---

## Enter Go-Micro

- A set of packages that can be used together to build microservices
- Packages are defined by interfaces. List of packages:
  - Registry  : Client side service discovery (abstraction over Consul, Etcd, Zookeeper, Kubernetes, etc)
  - Transport : Synchronous communication (abstraction over HTTP, RabbitMQ, NATS, etc)
  - Broker    : Asynchronous communication (abstraction over NATS, RabbitMQ, kafka)
  - Selector  : Node filtering and load balancing (abstraction over node selection strategies)
  - Codec     : Message encoding/decoding (rpc => json, bson => rpc, etc)
  - Server    : RPC server building on the above
  - Client    : RPC client building on the above 
- Many concrete implementations provided for each (called Plugins)
- Can be swapped in and even replaced entirely by your own implementation.

---

## The Service interfaces

Everything starts with the Service interface:

```go
type Service interface {
    Init(...Option)
    Options() Options
    Client() client.Client
    Server() server.Server
    Run() error
    String() string
}
```

This can be initialized as:

```go
service := micro.NewService() 
// or with options
service := micro.NewService(
	micro.Name("greeter"),
	micro.Version("latest"),
    micro.Flags(
		cli.StringFlag{
			Name:  "environment",
			Usage: "The environment",
		},
	),
)
```

---

## Options

The Options is a struct that wraps all the components of a service:

```go
type Options struct {
	Broker    broker.Broker
	Cmd       cmd.Cmd
	Client    client.Client
	Server    server.Server
	Registry  registry.Registry
	Transport transport.Transport

	RegisterInterval time.Duration

	BeforeStart []func() error
	BeforeStop  []func() error
	AfterStart  []func() error
	AfterStop   []func() error

	Context context.Context
}
```

Each component in Options (broker, cmd, client, server, registry, and transport) is itself
and interface which means that they can be swapped as you see fit.

---

## option 

Finally, each of the components can be customized to your liking using the
[functional options pattern](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis)

```go 
type Option func(*Options)
```

The multi-level option and options can be a bit much to take in right away, so lets break it down.
Before we do that, we need one last piece of the puzzle.

---

## Registry

```go
type Registry interface {
	Register(*Service, ...RegisterOption) error
	Deregister(*Service) error
	GetService(string) ([]*Service, error)
	ListServices() ([]*Service, error)
	Watch() (Watcher, error)
	String() string
}
```

Just like the service above, a registry also has options:

```go 
type Options struct {
	Addrs     []string
	Timeout   time.Duration
	Secure    bool
	TLSConfig *tls.Config

	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}
```

---

## How does it look?

We want a new service with all the default options, except we want the registry
to have a timeout of 10 seconds

```go
// first, configure the registry

timeout := time.Duration(10*time.Second)

r := registry.NewRegistry(registry.Timeout(timeout))

// use it to create the service

s := micro.NewService(
    micro.Registry(r)
)
```

---

## Plugins

Each of the components can be swapped out very easily. Go-Micro has a large list of
plugins, for each of the components.

You can also roll your own!

---

## Demo Time 





