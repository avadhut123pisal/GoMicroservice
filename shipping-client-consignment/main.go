// shipping-client-consignment/main.go

package main

import (
	consignmentpb "GoMicroservice/shipping-service-consignment/proto/consignment"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	micro "github.com/micro/go-micro"
)

const (
	ADDRESS           = "localhost:50051"
	DEFAULT_FILE_NAME = "consignment.json"
)

func ParseFile(file string) (*consignmentpb.Consignment, error) {
	consignment := consignmentpb.Consignment{}
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("failed to ReadFile: %v", err)
	}
	err = json.Unmarshal(bytes, &consignment)
	if err != nil {
		log.Fatalf("failed to Unmarshal: %v", err)
	}
	return &consignment, err
}
func main() {
	service := micro.NewService(micro.Name("shipping.cli.consignment"))
	service.Init()

	client := consignmentpb.NewShippingServiceClient("shipping.service.consignment", service.Client())

	consignment, err := ParseFile(DEFAULT_FILE_NAME)
	if err != nil {
		log.Fatalf("failed to ParseFile: %v", err)
	}
	response, err := client.CreateConsignment(context.Background(), consignment)

	if err != nil {
		log.Fatalf("ERROR IN CreateConsignment: %v", err)
	}
	log.Println("CONSIGNMENT CREATED: ", response)

	req := &consignmentpb.GetConsignmentRequest{}
	consignmentResponse, err := client.GetConsignments(context.Background(), req)
	if err != nil {
		log.Fatalf("ERROR IN GetConsignments: %v", err)
	}
	log.Println("GetConsignments Response: ", consignmentResponse)
}
