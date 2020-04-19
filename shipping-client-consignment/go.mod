module shipping-client-consignment

require (
	google.golang.org/grpc v1.28.1
	shipping-service-consignment v0.0.0
)

replace shipping-service-consignment => ../shipping-service-consignment
