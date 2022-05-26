package main

import (
	app "gRPCClientServerForEncryption/ClientService/internal/application"
	"gRPCClientServerForEncryption/ClientService/internal/logging"
	"gRPCClientServerForEncryption/ClientService/internal/settings"
	_ "google.golang.org/protobuf/proto"
	"log"
)

/*var (
	addr = fmt.Sprintf("%s", "localhost:9999")
)*/

func main() {
	log.Println("Starting client application...")
	log.Println("Collecting config...")
	cfg := Settings.GetConfig()
	log.Println("Getting logger of client application...")
	logger := logging.StartLog(cfg.LoggerConfig.LogLevel, cfg.LoggerConfig.LogFile)
	logger.Info().Msg("Starting application...")
	_, err := app.NewApplication(cfg, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("Application cant start")
	}
	//logger.Info().Str("foo", "bar").Msg("Hello world")
	//StartgRPC()
}

/*
func StartgRPC() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	err = fmt.Errorf("err")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPasswordsStorageClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetStorage(ctx, &pb.Token{
		Token: "123",
	})
	if err != nil {
		log.Fatalf("could not Get Strage: %v", err)
	}
	log.Printf("Greeting: %s", r.GetLogin())
}
*/
