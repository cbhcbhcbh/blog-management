package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"blog/internal/pkg/log"
	pb "blog/pkg/proto/blog/v1"
)

var (
	addr  = flag.String("addr", "localhost:9090", "The address to connect to.")
	limit = flag.Int64("limit", 10, "Limit to list users.")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalw("Did not connect", "err", err)
	}
	defer conn.Close()
	c := pb.NewBlogClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.ListUser(ctx, &pb.ListUserRequest{Offset: 0, Limit: *limit})
	if err != nil {
		log.Fatalw("could not greet: %v", err)
	}

	fmt.Println("TotalCount:", r.TotalCount)
	for _, u := range r.Users {
		d, _ := json.Marshal(u)
		fmt.Println(string(d))
	}
}
