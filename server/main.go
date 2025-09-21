package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	blodBank "github.com/dracuxan/blod-bank/proto"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	port  = flag.Int("port", 5001, "Server port")
	dbURL = flag.String("db", "postgres://dracuxan:pass@localhost:5432/blodbank?sslmode=disable", "Postgres connection URL")
)

type server struct {
	blodBank.UnimplementedBlodBankServiceServer
	db *sql.DB
}

func (s *server) RegisterConfig(_ context.Context, configItem *blodBank.ConfigItem) (*blodBank.Status, error) {
	query := `
		INSERT INTO configs (name, content, created_at, updated_at)
		VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		RETURNING id, created_at, updated_at;
	`

	var id int
	var createdAt, updatedAt time.Time
	err := s.db.QueryRow(query, configItem.Name, configItem.Content).Scan(&id, &createdAt, &updatedAt)
	if err != nil {
		log.Printf("ERROR inserting config: %v", err)
		return nil, status.Error(codes.Internal, "failed to insert config")
	}

	log.Printf("New config inserted with id %d", id)

	return &blodBank.Status{Status: fmt.Sprintf("Registered new config with id %d", id)}, nil
}

func (s *server) GetConfig(_ context.Context, configItemID *blodBank.ConfigID) (*blodBank.ConfigItem, error) {
	query := "SELECT * FROM configs WHERE id = $1;"
	var conf blodBank.ConfigItem
	row := s.db.QueryRow(query, configItemID.Id)

	var createdAt, updatedAt time.Time
	if err := row.Scan(&conf.Id, &conf.Name, &conf.Content, &createdAt, &updatedAt); err != nil {
		return nil, status.Errorf(codes.NotFound, "config not found")
	}
	conf.CreatedAt = createdAt.Format(time.RFC3339)
	conf.UpdatedAt = updatedAt.Format(time.RFC3339)

	log.Printf("Fetched config with id: %d", configItemID.Id)
	return &conf, nil
}

func (s *server) ListAllConfig(configItem *blodBank.NoParam, stream grpc.ServerStreamingServer[blodBank.ConfigItem]) error {
	log.Println("streaming list of all the configs")
	query := "SELECT * FROM configs;"

	row, err := s.db.Query(query)
	if err != nil {
		return status.Error(codes.Aborted, "bad request")
	}
	defer row.Close()

	for row.Next() {
		var id int
		var name, content string
		var createdAt, updatedAt string

		row.Scan(&id, &name, &content, &createdAt, &updatedAt)

		item := &blodBank.ConfigItem{
			Id:        int64(id),
			Name:      name,
			Content:   content,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}

		if err := stream.Send(item); err != nil {
			return status.Error(codes.Aborted, "bad request")
		}
	}

	return nil
}

func (s *server) DeleteConfig(ctx context.Context, configID *blodBank.ConfigID) (*blodBank.Status, error) {
	query := "DELETE FROM configs WHERE id = $1;"
	_, err := s.db.Exec(query, configID.Id)
	if err != nil {
		return &blodBank.Status{Status: ""}, status.Error(codes.Internal, "failed to delete config")
	}

	log.Printf("Deleted config with id %d", configID.Id)
	return &blodBank.Status{Status: fmt.Sprintf("Deleted config with id: %d", configID.Id)}, nil
}

func (s *server) UpdateConfig(ctx context.Context, configItem *blodBank.ConfigItem) (*blodBank.Status, error) {
	query := "UPDATE configs SET name = $2, content = $3, updated_at = $4 WHERE id = $1"

	_, err := s.db.Exec(query, configItem.Id, configItem.Name, configItem.Content, time.Now())
	if err != nil {
		return &blodBank.Status{Status: ""}, status.Error(codes.Internal, "failed to update config")
	}

	log.Printf("Updated config with id %d", configItem.Id)

	return &blodBank.Status{Status: fmt.Sprintf("updated config with id %d", configItem.Id)}, nil
}

func newServer(db *sql.DB) *server {
	return &server{db: db}
}

func main() {
	flag.Parse()
	db, err := sql.Open("postgres", *dbURL)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("unable to reach db: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	blodBank.RegisterBlodBankServiceServer(grpcServer, newServer(db))

	log.Printf("gRPC server listening on %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
