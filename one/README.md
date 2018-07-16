Quick Start
===========

1. If you don't have Go 1.9+ (required for sync.Map support), go get it [http://golang.org/doc/install]

2. Clone repo and run the server:
```
$ git clone git@github.com:eliwjones/paxos.git
$ cd paxos/one
$ PORT=9999 go run message_hasher.go
```

3. Try to get something that is not there:
```
$ curl -v http://127.0.0.1:9999/messages/2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae
* Hostname was NOT found in DNS cache
*   Trying 127.0.0.1...
* Connected to 127.0.0.1 (127.0.0.1) port 9999 (#0)
> GET /messages/2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae HTTP/1.1
> User-Agent: curl/7.35.0
> Host: 127.0.0.1:9999
> Accept: */*
>
< HTTP/1.1 404 Not Found
< Content-Type: application/json
< Date: Fri, 13 Jul 2018 17:58:11 GMT
< Content-Length: 31
<
{
  "error":"Message not found."
}
```

4. Put it there:
```
$ curl -X POST -H "Content-Type: application/json" -d '{"message": "foo"}' http://localhost:9999/messages
{
  "digest":"2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae"
}
```

5. Get it:
```
$ curl http://127.0.0.1:9999/messages/2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae
{
  "message":"foo"
}
```

Test Deployed Code
==================
```
$ curl https://message-hasher.herokuapp.com/messages/2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae
{
  "error": "Message not found."
}

$ curl -X POST -H "Content-Type: application/json" -d '{"message": "foo"}' https://message-hasher.herokuapp.com/messages
{
  "digest": "2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae"
}

$ curl https://message-hasher.herokuapp.com/messages/2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae
{
  "message": "foo"
}
```

Run Local Tests
===============
```
$ go test -v ./...
=== RUN   TestGetHandlerNotFound
--- PASS: TestGetHandlerNotFound (0.00s)
=== RUN   TestPostHandler
--- PASS: TestPostHandler (0.00s)
PASS
ok  	_/home/mrz/eliwjones/paxos/one	0.007s
```

NOTES
=====
1. I've used Go for this problem since my Ubuntu Trusty install is several years old and only has Python 3.4 built-in.  This prevents me from easily using AIOHTTP 3.x (which requires Python 3.5+), and I find older versions of AIOHTTP to be fairly ugly, syntax-wise.  Also, Go is fairly ideal for something like this.

2. The db is just a global sync.Map since adding Redis or whatnot to this simple setup would still be a toy solution.

BOTTLENECKS:
============
1. CPU and RAM available to the compiled binary and also the number of logical CPUs we let the runtime use with GOMAXPROCS.  One would want to write an intelligent load test and see if there was a sweet spot for GOMAXPROCS.

2. Once an external, persistent store was added, we'd be limited by how many open file handles we can have.  So, one would want to find some nicely written, performant package that pooled connections to our datastore of choice.  Then, crank the pool size as high as we can safely and ensure we can open that many filehandles.

3. Without knowing the particulars of the role this service performs in the overall architecture, I'd just say focus on allocating the most CPU/RAM one can that fits the budget.  If it can't go down, then spread the resources across 3 (or more) instances and load balance in your preferred manner.  In the end, this is just a fancy proxy sitting in front of our datastore.
