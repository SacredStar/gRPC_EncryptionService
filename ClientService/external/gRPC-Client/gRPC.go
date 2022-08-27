package gRPC_Client

import (
	pb "ClientService/external/gRPCproto"
	"ClientService/internal/logging"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strings"
	"time"
)

func StartConnectionAndGetStorage(port string, token string, timeout int64, logger *logging.Logger) {
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal().Str("Port:", port).Msg("Cant Dial connection with server")
	}
	//TODO: error processing
	defer conn.Close()
	c := pb.NewPasswordsStorageClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	r, err := c.GetStorage(ctx, &pb.Token{
		//TODO: understand why encodings is so ...
		Token: strings.ToValidUTF8(token, ""),
	})
	if err != nil {
		logger.Fatal().Msgf("could not Get Storage: %v", err)
	}
	logger.Info().Msgf("Greeting: %s", r.GetLogin())
}
