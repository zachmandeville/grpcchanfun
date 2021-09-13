package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"io"
	"log"
	"sync"
	"time"

	mathspb "github.com/zachmandeville/grpcchanfun/api/maths"
	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("server_addr", "localhost:10000", "The server address")
)

func runSquares (client mathspb.MathsClient, nums chan int32, ch chan string, done chan bool) {
	// run a bidirectional stream that will sends the server response to the ch channel.
	// the sending stream will only close when it receive no more messages across the numschannel
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.Squares(ctx)
	if err != nil {
		log.Fatalf("%v.Squares(_) = _, %v", client, err)
	}
	var wg sync.WaitGroup
	go func() {
		for {
		wg.Add(1)
			in, err := stream.Recv()
			if err == io.EOF {
				close(ch)
				done <- true
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive number to square: %v\n", err)
			}
			result := fmt.Sprintf("The Square of %v is %v", in.Number, in.Square)
			ch <- result
		}
	}()

	// this for loop will run until the num channel closes.
	// this function doesn't control whether the num channel closes,
	// that is handled outside this function.
	for num := range nums {
		fmt.Printf("received %v from channel\n", num)
		request := &mathspb.SquaresRequest{
			Number: num,
		}
		if err := stream.Send(request); err != nil {
			log.Fatalf("error sending number: %v", num)
		}
	}
	stream.CloseSend()
	wg.Wait() // keep the for loops running until we are fully done sending all our numbers to the server.
}

func main () {
	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("unable to connect to server: %v", err)
	}
	defer conn.Close()

	client := mathspb.NewMathsClient(conn)
	ch := make(chan string, 2)
	done := make(chan bool, 2)
	nums := make(chan int32, 2)

	go runSquares(client, nums, ch, done) // this is how we interact with the external server, like a function that intereacts with an xDS server

	go func () {
		// here we have a concurrent routine that is random, and so we can't anticipate when the number channel will close.
		// we want to make sure we have the program setup well enough that we never try to send across a closed channel
		// and we get all our responses...even if we odn't know how many these would be.
		for {
		  rand.Seed(time.Now().UnixNano())
		  n := rand.Intn(12 - 1) + 1
			if n == 7 {
				fmt.Println("sending lucky number 7 across channel")
				nums <- int32(n)
				close(nums)
				return

			} else {
				fmt.Printf("sending %v across channel\n",n)
				nums <- int32(n)
			}
		}
	}()

	// keep the for loop running until we get a done signal from our runSquares program.  Without this, the program exits too quickly.
	// the messages received will be pulled from our server.
	for {
		select {
		case msg := <- ch:
			if len(msg) > 0 {
				fmt.Println(msg)
			}
		case <-done:
			fmt.Println("All done!!")
			return
		}
	}
}
