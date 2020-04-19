// shipping-service-consignment/main.go
package main

import (
	"context"
	"log"
	"sync"

	consignmentpb "GoMicroservice/shipping-service-consignment/proto/consignment"

	"github.com/micro/go-micro"
)

type Repository interface {
	Create(*consignmentpb.Consignment) (*consignmentpb.Consignment, error)
	GetAll() ([]*consignmentpb.Consignment, error)
}

type repository struct {
	mu           *sync.RWMutex
	consignments []*consignmentpb.Consignment
}

type service struct {
	repo Repository
}

func (repo *repository) Create(consignment *consignmentpb.Consignment) (*consignmentpb.Consignment, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.consignments = append(repo.consignments, consignment)
	return consignment, nil
}

func (repo *repository) GetAll() ([]*consignmentpb.Consignment, error) {
	return repo.consignments, nil
}

func (s *service) CreateConsignment(ctx context.Context, consignment *consignmentpb.Consignment, res *consignmentpb.Response) error {
	respConsignment, err := s.repo.Create(consignment)
	if err != nil {
		return err
	}
	res.Created = true
	res.Consignment = respConsignment
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *consignmentpb.GetConsignmentRequest, res *consignmentpb.ConsignmentResponse) error {
	consignmentList, err := s.repo.GetAll()
	if err != nil {
		log.Fatalf("Error in getting consignmentList: %v", err)
	}
	res.Consignments = consignmentList
	return nil
}

func main() {
	srv := micro.NewService(
		// This name must match the package name given in your protobuf definition
		micro.Name("shipping.service.consignment"),
	)

	// Init will parse the command line flags.
	srv.Init()

	serviceObj := service{
		repo: &repository{
			mu:           &sync.RWMutex{},
			consignments: make([]*consignmentpb.Consignment, 0),
		},
	}

	consignmentpb.RegisterShippingServiceHandler(srv.Server(), &serviceObj)

	log.Println("Starting server")
	if err := srv.Run(); err != nil {
		log.Fatalf("failed to RUN: %v", err)
	}
}
