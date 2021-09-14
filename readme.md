# Exploring grpc and channels

This is a learning repo to explore how to structure programs that work with gRPC bidirectional streams, specifically when we want the stream to be running concurrent to a set of other functions.

# The example
This repo has an external gRPC server running our maths schema. The main
function dials into this server and sets up a stream. The stream listens in on a
numbers channel and sends any numbers it gets to the squares server, and stops
sending when the numbers channel closes.

The main function is set up to randomly send numbers until we reach number 7,
which should be the last number squared. IT listens in for messaages coming from
our squares server, not quitting until the last message(which should always be
the square of seven) is received.

If it works correclty, then we should see all the numbers being sent in order
and receiving their square. All numbers sent should get a response, which means
the last response should always be 7.

# Running the example
from this repo, in one terminal screen

```sh
go run server/server.go
```

you should see `server started at localhost:1000`

in another terminal, run

``` sh
go run main.go
```

You will see a number of sends and responses. The key thing to look for is that
the last response is the square of 7.

`
# current status
This basic example works. Next is to add complexity by integrating some godog
tests, and having the tests steps decide when to close our stream and validate
the results.
