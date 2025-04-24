package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"

	Log "github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/divyank-lakhera/HiringPlatform/Backend/api"
	db "github.com/divyank-lakhera/HiringPlatform/Backend/db/sqlc"
	"github.com/divyank-lakhera/HiringPlatform/Backend/gapi"
	pb "github.com/divyank-lakhera/HiringPlatform/Backend/pb"
	"github.com/divyank-lakhera/HiringPlatform/Backend/util"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:hiring_platform@localhost:5432/hiring_platform?sslmode=disable"
	serverAddress = "localhost:8080"
)

func main() {

	serverType := flag.String("server", "", "Which server")
	flag.Parse()

	config, err := util.LoadConfig(".")
	fmt.Println("Access toke n duration", config.AccessTokenDuration)
	if err != nil {
		Log.Fatal().Err(err).Msg("cannot load config")
	}

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	// server, err := api.NewServer(config, store)
	// if err != nil {
	// 	log.Fatal("Cannot create server:", err)
	// }

	// err = server.Start(serverAddress)
	// if err != nil {
	// 	log.Fatal("Cannot start server:", err)
	// }
	switch *serverType {
	case "http":
		runGinServer(config, store)
	case "grpc":
		runGRPCServer(config, store)
	default:
		fmt.Println("Restart server, incorrect server type: Options - http & grpc")
	}
	// runGinServer(config, store)

}

func runGinServer(config util.Config, store *db.SQLStore) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create server", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}

func runGRPCServer(config util.Config, store *db.SQLStore) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create server", err)
	}

	gRPCServer := grpc.NewServer()
	pb.RegisterHiringPlatformServer(gRPCServer, server)
	reflection.Register(gRPCServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener")
	}

	log.Printf("gRPC server started : %s", listener.Addr().String())
	err = gRPCServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server", err)
	}

}
