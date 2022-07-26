package main

// import (
// 	"context"
// 	"net"

// 	"github.com/alecthomas/kong"
// 	"github.com/hantonelli/ghipster/product/internal/repository/entgen"
// 	"github.com/hantonelli/ghipster/product/internal/repository/entgen/migrate"
// 	svc "github.com/hantonelli/ghipster/product/internal/service"
// 	pb "github.com/hantonelli/ghipster/proto/product"
// 	"go.uber.org/zap"
// 	"google.golang.org/grpc"
// )

// func main() {
// 	var cli struct {
// 		Addr  string `name:"address" default:":50051" help:"Address to listen on."`
// 		Debug bool   `name:"debug" help:"Enable debugging mode."`
// 	}
// 	kong.Parse(&cli)

// 	log, _ := zap.NewProduction()

// 	client, err := entgen.Open("sqlite3", "file:entproduct?mode=memory&cache=shared&_fk=1")
// 	if err != nil {
// 		log.Fatal("opening ent client", zap.Error(err))
// 	}
// 	defer client.Close()

// 	if err := client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true)); err != nil {
// 		log.Fatal("running schema migration", zap.Error(err))
// 	}

// 	productSvc := svc.NewProductsService(log, client)

// 	grpcServer := grpc.NewServer()

// 	pb.RegisterProductsServiceServer(grpcServer, productSvc)

// 	lis, err := net.Listen("tcp", cli.Addr)
// 	if err != nil {
// 		log.Fatal("failed to listen: %v", zap.Error(err))
// 	}
// 	// defer lis.Close()

// 	if err := grpcServer.Serve(lis); err != nil {
// 		log.Fatal("failed to serve", zap.Error(err))
// 	}
// }
