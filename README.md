# singload
Reduce load-balancing complexity by routing HTTP requests to a single node.

## Description
Routing traffic to multiple nodes using a load-balancer increases your system's complexity, and singload solves exactly that problem.  
Simply enter a single node address and singload will route traffic to it using the round-robin strategy.
Singload avoids complexity by not evaluating the target node's health, not routing HTTPS requests, and routing on port 80 only.

## Installation
```shell
wget https://github.com/wohb/singload/releases/latest/download/singload
mv singload /path/of/choice/singload
chmod +x /path/of/choice/singload
```

## Usage
The singload load-balancer listens on port 80, and by default routes requests to '127.0.0.1' to reduce complexity even more
```shell
singload
```
Run singload and specify an address
```shell
singload --address 'some.address.com'
```
Run singload and specify an IP address
```shell
singload --ip '10.10.55.55'
```

## Contributing
1. Run `dig +short github.com` and copy the IP
2. Open chrome and paste the IP with the suffix as follows: `addr_you_copied_here/wohb/singload/issues`
3. Create an issue describing why you love system complexity
