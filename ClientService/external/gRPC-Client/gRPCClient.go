package gRPC_Client

import (
	Settings "ClientService/internal/config"
	"ClientService/internal/logging"
)

type Application struct {
	cfg    *Settings.ClientConfig
	logger *logging.Logger
}

/*func StartgRPC(app Application) {
	conn, err := grpc.Dial(app.cfg.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		app.logger.Fatal().Str("Port:", app.cfg.Port).Msg("Cant Dial connection with server")
	}
	defer conn.Close()
	c := pb.NewPasswordsStorageClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(app.cfg.TimeoutContext)*time.Second)
	defer cancel()
	r, err := c.GetStorage(ctx, &pb.Token{
		Token: "123",
	})
	if err != nil {
		app.logger.Fatal().Msgf("could not Get Storage: %v", err)
	}
	app.logger.Info().Msgf("Greeting: %s", r.GetLogin())
}*/
