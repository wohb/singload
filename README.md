# singload
Reduce load-balancing complexity by routing HTTP requests to a single node.

## Description
Routing traffic to multiple nodes using a load-balancer increases your system's complexity, and singload solves exactly that problem.  
Simply enter a single IP address and singload will route traffic to it using the round-robin strategy.
Singload avoids complexity by not evaluating the target node's health, not routing HTTPS requests, and routing on port 80 only.

## Installation
-- coming soon --

## Usage
Run singload and specify an IP address
```shell
singload --ip '8.8.8.8'
```

## Contributing
1. Run `dig +short github.com` and copy the IP
2. Open chrome and paste the IP with the suffix as follows: `ip_addr_you_copied_here/wohb/singload/issues`
3. Create an issue describing why you love system complexity
