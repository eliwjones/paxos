Quick Start
===========

1. If you don't have GO, go get it [http://golang.org/doc/install]

2. Clone repo and run the server:
```
$ git clone git@github.com:eliwjones/paxos.git
$ cd paxos/one
$ go run message_hasher.go
```

3. Test it out:
```
# In another terminal

# Try to get something that is not there:
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

# Put it there:
$ curl -X POST -H "Content-Type: application/json" -d '{"message": "foo"}' http://localhost:9999/messages
{
  "digest":"2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae"
}

# Get it:
$ curl http://127.0.0.1:9999/messages/2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae
{
  "message":"foo"
}
```
