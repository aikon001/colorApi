package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/aikon001/color-api/colorapi"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "color-api_test"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewColorsClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.AddColorFromHexOrRgb(ctx, &pb.Color{Name: "testcolor", Hex: "1B3C9B"})

	log.Println("****************First call*****************")
	log.Println(res)
	log.Println(err)

	res, err = c.AddColorFromHexOrRgb(ctx, &pb.Color{Name: "testcolor", RGB: 1784987})

	log.Println("****************Second call*****************")
	log.Println(res)
	log.Println(err)

	stream, err := c.PickAllColors(ctx, &pb.StartIndex{Index: 0})
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Println("****************Third call*****************")
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("cannot receive response: ", err)
		}

		c := res.GetColorResp()

		log.Println(c.GetName() + "(" + c.GetHex() + ")")
	}

}
