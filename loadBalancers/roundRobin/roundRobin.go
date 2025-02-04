package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

// simpleServer represents a backend server with reverse proxy capabilities
type simpleServer struct {
	addr  string
	proxy *httputil.ReverseProxy
}

// Server interface defines the methods required by any server in the load balancer
type Server interface {
	Address() string
	isAlive() bool
	Serve(rw http.ResponseWriter, r *http.Request)	
}

// newSimpleServer creates a new simple server with a reverse proxy to the given address
func newSimpleServer(addr string) *simpleServer {
	serverUrl, err := url.Parse(addr)
	handleErr(err)

	return &simpleServer{
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

// LoadBalancer manages multiple backend servers and handles request forwarding
type LoadBalancer struct {
	port            string
	roundRobinCount int
	servers         []Server
}

// NewLoadBalancer creates a new load balancer with a given port and list of backend servers
func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port : port,
		roundRobinCount: 0,
		servers: servers,
	}
}

// Address returns the address of the simpleServer
func (s *simpleServer) Address() string {return s.addr}

// checks if a servers is available
func (s *simpleServer) isAlive() bool {return true}

// forwards the request to the backend server via a reverse proxy
func (s *simpleServer) Serve(rw http.ResponseWriter, req *http.Request) {
	s.proxy.ServeHTTP(rw, req)
}

// select the next available server through round robin strategy
func (lb *LoadBalancer)getNextAvailableServer() Server {
	server := lb.servers[(lb.roundRobinCount%len(lb.servers))]

	for !server.isAlive() {
		server = lb.servers[lb.roundRobinCount%len(lb.servers)]
	}
	lb.roundRobinCount++
	return server
}

// forwards request to the next available server
func (lb *LoadBalancer) serveProxy(rw http.ResponseWriter, req* http.Request) {
	targetServer := lb.getNextAvailableServer()
	fmt.Printf("forwarding request to address %q\n", targetServer.Address())
	targetServer.Serve(rw, req)
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	// random servers
	servers := []Server {
		newSimpleServer("https://www.facebook.com"),
		newSimpleServer("https://www.bing.com"),
		newSimpleServer("https://www.duckduckgo.com"),
	}
		lb :=NewLoadBalancer("8000", servers)
		// to forward incoming requests to the load balancer
		handleRedirect := func(rw http.ResponseWriter, req *http.Request){
			lb.serveProxy(rw, req)
		}
		http.HandleFunc ("/",handleRedirect)

		fmt.Printf("serving at 'localhost%s'\n", lb.port)
		http.ListenAndServe(":" + lb.port, nil)
	}
