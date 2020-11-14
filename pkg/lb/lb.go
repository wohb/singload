package lb

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// LoadBalancer balances requests to a single node
type LoadBalancer struct {
	TargetAddr   string
	ListenerPort int
}

// Run the reverse-proxy load balancer
func (lb *LoadBalancer) Run() error {
	targetAddr := fmt.Sprintf("http://%s/", lb.TargetAddr)
	target, err := url.Parse(targetAddr)
	if err != nil {
		log.Println(err)
	}

	// Round-robinning here
	director := func(req *http.Request) {
		req.Header.Add("X-Forwarded-Host", req.Host)
		req.Header.Add("X-Origin-Host", target.Host)
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		log.Printf("Forwarding request to -> %s://%s%s\n", target.Scheme, target.Host, req.URL.Path)
	}
	proxy := &httputil.ReverseProxy{Director: director}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		req.Host = req.URL.Host
		w.Header().Add("X-Distributed-Systems-Are-Good", "False")
		proxy.ServeHTTP(w, req)
	})

	listenerAddr := fmt.Sprintf(":%d", lb.ListenerPort)
	log.Printf("Serving requests on %s", listenerAddr)
	err = http.ListenAndServe(listenerAddr, nil)
	if err != nil {
		panic(err)
	}

	return nil
}
