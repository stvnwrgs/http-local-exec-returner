#Locofo#

Locofo is a small tool that takes http GET requests, executes a command and gives back the response.


In simple pseudo code it is just doing:

```
GET /health

output, err = command.execute(command)

if err {
  response = 500
} else {
  response = 200
}

print output

```

## Intention ##

Get GCE Health checks working with a TLS secured etcd 2 cluster.

The tool is working as a "proxy" to enable googles checks on http while etcd2 is server on https with pki.

```
# Not tested the command

{
  "BindAddress": "0.0.0.0:2350",
  "Commands": [
    {
      "Path": "/health",
      "Command": "curl",
      "Args": "--cacert /path/to/ca.pem --key /path/key.pem --cert /path/cert.pem \"https://127.0.0.1/v2/stats/self\" "
    }
  ]
}

```



But you could use it for everything else too ;)

## build ##

```
# get a proper go installation
go get github.com/stvnwrgs/locofo
cd $GOPATH/src/github.com/stvnwrgs/locofo
go build
go install

# for multi os compiling i suggest gox.

go get github.com/mitchellh/gox
cd $GOPATH/src/github.com/mitchellh/gox
go build && go install

cd $GOPATH/src/github.com/stvnwrgs/locofo
gox -output="builds/{{.OS}}/{{.Dir}}_{{.Arch}}"
```

Or you just download the release package.
https://github.com/stvnwrgs/locofo/releases
