# Exploring grpc and channels

This is a learning repo to explore how to structure programs that work with gRPC bidirectional streams, specifically when we want the stream to be running concurrent to a set of other functions.

Secondly, we are looking at how to incorpporate godog into this scenario, so that you can setup a stream that runs alongside a set of godog functions, and ensure that your tests run correctly.

# The example
This repo has an external gRPC server running our maths schema.  The main function dials into this server and sets up a stream.  The stream listens in on a numbers channel and sends any numbers it gets to this server, which returns the square or cube depneding on service you're using.  
When the number channel closes, the stream closes as well.

We have two tests setup at `features/maths.feature`.  The scnenarios re-use steps, but each one is testing a different gRPC rpc, similar to testing CDS or LDS.

We implement the tests by setting up a runner struct that has a sending channel, a receiving channel, and a number cache.  Before each scenario, we reset our cache and start up new channel instances.  The given step initiates a stream that runs as a goroutine, meaning it will stay running even as we move to other functions. 

We have several functions that affect the stream, namely by sending more numbers along our sending channel. 
The final `Then` step in each scenario closes the sending channel, waits for responses on the receiving channel, then validates the responses.
This pattern, at least in this simple context, allows us to run the stream concurrently, have it be affected by the test, and guarantees the final step has all the results it should.  It will be a similar setup when testing xDS, but with discovery request and discovery response channels.

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

You can also verbosely run through the tests with

``` sh
go run main.go --debug
```

# current status
This seems to be working!  Would love a review of the go code, though.  Another useful step might be to set up the scenarios with an example table, to make the steps even more reusable.
