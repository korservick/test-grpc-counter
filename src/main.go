package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/caarlos0/env"
	pbt "github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "grpc_counter/pb"
	"log"
	"net"
)

type envConfig struct {
	Port       int    `env:"PORT" envDefault:"50051"`
	DbHost     string `env:"DB_HOST" envDefault:"localhost"`
	DBPort     int    `env:"DB_PORT" envDefault:"5432"`
	DbName     string `env:"DB_NAME" envDefault:"counter_db"`
	DbUser     string `env:"DB_USER" envDefault:"counter"`
	DbPassword string `env:"DB_PASSWORD" envDefault:"password"`
}

type Counter struct {
	gorm.Model
	CurrentNumber uint32
	MaximumNumber uint32
	Increment     uint32
}

type IncrementorServer struct {
	db *gorm.DB
}

func (i *IncrementorServer) GetNumber(ctx context.Context, empty *pbt.Empty) (*pb.CurrentNumber, error) {
	var counter Counter
	//err := i.db.Select("current_number").First(&counter).Error
	err := i.db.First(&counter).Error
	return &pb.CurrentNumber{Number: counter.CurrentNumber}, err
}

func (i *IncrementorServer) IncrementNumber(ctx context.Context, empty *pbt.Empty) (*pbt.Empty, error) {
	var counter Counter
	tx := i.db.Begin()
	if tx.Error != nil {
		return &pbt.Empty{}, tx.Error
	}

	tx.Set("gorm:query_option", "FOR UPDATE").First(&counter)
	counter.CurrentNumber += counter.Increment
	if counter.CurrentNumber >= counter.MaximumNumber {
		counter.CurrentNumber = 0
	}
	//if err := tx.Model(&counter).UpdateColumn("current_number", counter.CurrentNumber).Error; err != nil {
	if err := tx.Save(&counter).Error; err != nil {
		fmt.Println("err")
		tx.Rollback()
		return &pbt.Empty{}, err
	}
	tx.Commit()
	return &pbt.Empty{}, nil
}

func CheckSettings(settings *pb.Settings) error {
	if settings.MaximumNumber <= 0 {
		return errors.New("MaximumNumber <= 0")
	}
	if settings.Increment <= 0 {
		return errors.New("Increment <= 0")
	}
	if settings.Increment >= settings.MaximumNumber {
		return errors.New("Increment >= MaximumNumber")
	}
	return nil
}

func (i *IncrementorServer) SetSettings(ctx context.Context, settings *pb.Settings) (*pbt.Empty, error) {
	var counter Counter

	err := CheckSettings(settings)
	if err != nil {
		return &pbt.Empty{}, err
	}

	tx := i.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return &pbt.Empty{}, tx.Error
	}

	tx.Set("gorm:query_option", "FOR UPDATE").First(&counter)
	counter.Increment = settings.Increment
	counter.MaximumNumber = settings.MaximumNumber
	if counter.CurrentNumber >= settings.MaximumNumber {
		counter.CurrentNumber = 0
	}
	tx.Save(&counter)
	tx.Commit()
	return &pbt.Empty{}, nil
}

func main() {
	var err error

	log.Println("starting")
	cfg := envConfig{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	} else {
		log.Fatalf("failed to parse env variables: %v", err)
	}

	db, err := gorm.Open("postgres", fmt.Sprintf(
		"host=%s port=%d sslmode=disable dbname=%s user=%s password=%s",
		cfg.DbHost, cfg.DBPort, cfg.DbName, cfg.DbUser, cfg.DbPassword))
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatalf("failed to close database: %v", err)
		}
	}()
	db.AutoMigrate(&Counter{})
	db.FirstOrCreate(&Counter{CurrentNumber: 0, MaximumNumber: 100, Increment: 1})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterIncrementorServer(server, &IncrementorServer{db: db})
	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
