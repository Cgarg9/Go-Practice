# Load Balancer with Reverse Proxy

This Go application implements a basic **load balancer** that forwards incoming HTTP requests to a set of predefined servers. The load balancer uses a **round-robin algorithm** to distribute traffic across available servers. Each server is connected through a reverse proxy, enabling transparent forwarding of requests to the backend servers.

## Overview

1. **Simple Server**: Represents the backend servers that the load balancer will forward requests to. It uses a `ReverseProxy` to forward requests.
2. **Load Balancer**: Manages multiple backend servers and forwards incoming requests to them based on the round-robin algorithm. It ensures that traffic is evenly distributed to the available servers.

## Reverse Proxy

A **reverse proxy** is a server that sits between client devices and backend servers. It intercepts client requests and forwards them to one of the backend servers. Unlike a traditional proxy (which forwards requests from clients to the internet), a reverse proxy serves on behalf of the backend server. It acts as an intermediary between the client and the backend, making it appear as though all traffic is handled by the reverse proxy server.

### Benefits of Using a Reverse Proxy:
- **Load Balancing**: The reverse proxy can distribute incoming requests across multiple backend servers to optimize resource utilization and performance.
- **Security**: It can hide the identity and structure of the backend servers, providing a layer of protection.
- **Caching**: Reverse proxies can cache responses from the backend servers to reduce load and improve response time.
- **SSL Termination**: The reverse proxy can handle SSL encryption and decryption, reducing the load on backend servers.
  
In this project, we use the Go `httputil.ReverseProxy` to implement reverse proxy functionality. This allows the load balancer to forward client requests to backend servers like `facebook.com`, `bing.com`, or `duckduckgo.com`.

## Code Flow

### 1. **Server Setup**

A **simpleServer** is a struct representing a backend server. Each simple server holds:
- `addr`: The server's address.
- `proxy`: A reverse proxy instance used to forward requests to the actual backend server.

### 2. **Creating the Simple Server**

The `newSimpleServer` function initializes a new backend server:
- Takes an address (`addr`) as input.
- Parses the URL and creates a `ReverseProxy` using `httputil.NewSingleHostReverseProxy`.

### 3. **Load Balancer**

The `LoadBalancer` struct:
- Maintains a list of backend servers.
- Uses a round-robin counter (`roundRobinCount`) to ensure requests are distributed evenly across the servers.
- Contains methods like:
  - `getNextAvailableServer()`: Chooses the next server based on round-robin logic.
  - `serveProxy()`: Forwards the request to the chosen backend server.

### 4. **Reverse Proxying**

The `simpleServer` struct implements the `Server` interface, which requires:
- `Address()`: Returns the server's address.
- `isAlive()`: Checks if the server is healthy (always returns `true` in this example).
- `Serve()`: Handles incoming HTTP requests and forwards them using a reverse proxy.

### 5. **Main Flow**

The main function initializes the following:
- A list of backend servers (`facebook.com`, `bing.com`, `duckduckgo.com`).
- A load balancer (`lb`) is created with the specified port (8000) and the list of servers.
- The load balancer listens for incoming HTTP requests and forwards them using the round-robin algorithm.

---

## Code Breakdown

### 1. **simpleServer Struct**

```go
type simpleServer struct {
    addr  string
    proxy *httputil.ReverseProxy
}
```

- The *simpleServer* struct has two fields:
    - *addr*: The address of the backend server.
    - *proxy*: A reverse proxy that will forward requests to the server.

### 2. **Server Interface**

```go
type Server interface {
    Address() string
    isAlive() bool
    Serve(rw http.ResponseWriter, r *http.Request)  
}
```

- The *Server* interface defines the necessary methods for any backend server:
    - *Address()* returns the server's address.
    - *isAlive()* checks if the server is healthy (in this example, always returns *true*).
    - *Serve()* handles incoming HTTP requests and forwards them using a reverse proxy.

### 3. **Load Balancer Struct**

```go
type LoadBalancer struct {
    port            string
    roundRobinCount int
    servers         []Server
}
```

- The *LoadBalancer* struct holds:
    - *port*: The port on which the load balancer will listen.
    - *roundRobinCount*: A counter used to distribute requests evenly.
    - *servers*: A list of backend servers.

### 4. **newSimpleServer**

```go
func newSimpleServer(addr string) *simpleServer {
    serverUrl, err := url.Parse(addr)
    handleErr(err)

    return &simpleServer{
        addr:  addr,
        proxy: httputil.NewSingleHostReverseProxy(serverUrl),
    }
}
```

- Takes an address as input.
- Parses it into a *url.URL* object.
- Initializes a reverse proxy for that server.

### 5. **Round-Robin Server Selection**

```go
func (lb *LoadBalancer) getNextAvailableServer() Server {
    server := lb.servers[(lb.roundRobinCount % len(lb.servers))]

    for !server.isAlive() {
        server = lb.servers[lb.roundRobinCount % len(lb.servers)]
    }
    lb.roundRobinCount++
    return server
}
```

- Selects the next backend server in a round-robin fashion.
- If a server is unavailable, it skips to the next.

### 6. **Request Forwarding**

```go
func (lb *LoadBalancer) serveProxy(rw http.ResponseWriter, req *http.Request) {
    targetServer := lb.getNextAvailableServer()
    fmt.Printf("forwarding request to address %q\n", targetServer.Address())
    targetServer.Serve(rw, req)
}
```

- Determines the next backend server.
- Uses its *Serve()* method to forward the request.

## Future Improvements

- *Health Checks:* Implement periodic health checks to detect if a backend server is down and exclude it from the round-robin rotation.
- *Dynamic Server Addition/Removal:* Allow adding/removing backend servers at runtime.
- *Logging & Monitoring:* Add request logs and metrics to track performance.
- *Sticky Sessions:* Maintain session affinity for specific clients.