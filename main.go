package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	mathspb "github.com/zachmandeville/grpcchanfun/api/maths"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type Runner struct {
	Conn *grpc.ClientConn
	Numbers []int32
	Numchan chan int32 // numbers we send to a stream
	Reschan chan int32 // results from that stream
}

var (
	runner *Runner
	serverAddr = flag.String("server_addr", "localhost:10000", "The server address")
	debug      = flag.Bool("debug", false, "sets log level to debug")
	godogOpts  = godog.Options{
		Output:      colors.Colored(os.Stdout),
	}
)

func newRunner() *Runner {
	return &Runner{
		Numbers: []int32{},
		Numchan: make(chan int32),
		Reschan: make(chan int32),
	}
}

func InitializeTestSuite(tc *godog.TestSuiteContext) {
	tc.BeforeSuite(func() {
		runner = newRunner()
		log.Debug().
			Msg("Runner initalized")
		conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
		if err != nil {
			log.Fatal().
				Err(err).
				Msg("unable to connect to gRPC server")
		}
		runner.Conn = conn
		log.Debug().
			Msgf("Runner connected to %v", *serverAddr)
	})
	tc.AfterSuite(func () {
		runner.Conn.Close()
		log.Debug().
			Msgf("Closed connection to %v", *serverAddr)
	})
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		runner.Numbers = []int32{}
		return ctx, nil
	})
	ctx.Step(`^an initiated squares stream$`, runner.AnInitiatedSquaresStream)
    ctx.Step(`^an initiated cubes stream$`, runner.AnInitiatedCubesStream)
	ctx.Step(`^I add (\d+) to the stream$`, runner.IAddToTheStream)
	ctx.Step(`^no cached numbers$`, runner.NoCachedNumbers)
	ctx.Step(`^the sum of cached numbers should be (\d+)$`, runner.TheSumOfCachedNumbersShouldBe)
}

func (r *Runner) AnInitiatedSquaresStream() error {
	runner.Numchan = make(chan int32)
	runner.Reschan = make(chan int32)
	client := mathspb.NewMathsClient(r.Conn)
	go runSquares(client, r.Numchan, r.Reschan)
	return nil
}

func (r *Runner) AnInitiatedCubesStream() error {
	runner.Numchan = make(chan int32)
	runner.Reschan = make(chan int32)
	client := mathspb.NewMathsClient(r.Conn)
	go runCubes(client, r.Numchan, r.Reschan)
	return nil
}

func (r *Runner) NoCachedNumbers() error {
	log.Debug().
		Msgf("Our Numbers Cache: %v", r.Numbers)
	if len(r.Numbers) > 0 {
		err := fmt.Errorf("There are cached numbers: %v", r.Numbers)
		return err
	}
	return nil
}

func (r *Runner) IAddToTheStream(arg1 int) error {
	log.Debug().
		Msgf("Sending %v across numChan", arg1)
	r.Numchan <- int32(arg1)
	return nil
}

func (r * Runner) TheSumOfCachedNumbersShouldBe(arg1 int) error {
	close(r.Numchan)
	for square := range r.Reschan {
		r.Numbers = append(r.Numbers, square)
		log.Debug().
			Msg(fmt.Sprintf("received on res channel: %v", square))
	}
	sum := int32(0)
	for _, n := range r.Numbers {
		sum = sum + n
	}

	if sum != int32(arg1) {
		err := fmt.Errorf("Incorrect sum of squares.\n Expected: %v\nActual: %v", arg1, sum)
		return err
	}
	return nil
}

func runCubes(client mathspb.MathsClient, nums chan int32, res chan int32) {
	// run a bidirectional stream that will sends the server response to the ch channel.
	// the sending stream will only close when it receive no more messages across the numschannel
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.Cubes(ctx)
	if err != nil {
		log.Fatal().
			Err(err).
			Msgf("%v.Cubes(_) = _, %v", client, err)
	}
	log.Debug().
		Msg("Starting Cubes Stream")

	// receiving stream for our gRPC server
	// will loop until the sending channel is closed.
	//
	var wg sync.WaitGroup
	go func() {
		for {
			wg.Add(1)
			in, err := stream.Recv()
			if err == io.EOF {
				log.Debug().
					Msg("No more numbers from cubes stream")
				close(res)
				return
			}
			if err != nil {
				log.Fatal().
					Err(err).
					Msg("Failed to receive a number to cube")
			}
			log.Debug().
				Msgf("received from server: cube of %v is %v", in.Number, in.Cube)
			res <- in.Cube
		}
	}()

	// Sending stream.
	// this for loop will run until the num channel closes.
	// that is handled outside this function.
	for num := range nums {
		log.Debug().
			Msgf("received %v from channel\n", num)
		request := &mathspb.CubesRequest{
			Number: num,
		}
		if err := stream.Send(request); err != nil {
			log.Fatal().
				Err(err).
				Msgf("error sending number: %v", num)
		}
	}
	log.Debug().
		Msg("number loop ended")
	stream.CloseSend()
	wg.Wait() // keep the for loops running until we've received all squares from our server.
}

func runSquares(client mathspb.MathsClient, nums chan int32, res chan int32) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.Squares(ctx)
	if err != nil {
		log.Fatal().
			Err(err).
			Msgf("%v.Squares(_) = _, %v", client, err)
	}
	log.Debug().
		Msg("Starting Squares Stream")

	var wg sync.WaitGroup
	go func() {
		for {
			wg.Add(1)
			in, err := stream.Recv()
			if err == io.EOF {
				log.Debug().
					Msg("No more numbers from squares stream")
				close(res)
				return
			}
			if err != nil {
				log.Fatal().
					Err(err).
					Msg("Failed to receive a number to square")
			}
			log.Debug().
				Msgf("received from server: square of %v is %v", in.Number, in.Square)
			res <- in.Square
		}
	}()

	for num := range nums {
		log.Debug().
			Msgf("received %v from channel\n", num)
		request := &mathspb.SquaresRequest{
			Number: num,
		}
		if err := stream.Send(request); err != nil {
			log.Fatal().
				Err(err).
				Msgf("error sending number: %v", num)
		}
	}
	log.Debug().
		Msg("number loop ended")
	stream.CloseSend()
	wg.Wait()
}

func init() {
	flag.Parse()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	godog.BindCommandLineFlags("godog.", &godogOpts)
}

func main() {
	status := godog.TestSuite{
		Name:                 "xDS Test Suite",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
		Options:              &godogOpts,
	}.Run()
	os.Exit(status)
}
