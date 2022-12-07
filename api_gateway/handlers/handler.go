package handlers

import (
	"yorqinbek/microservices/Blogpost/api_gateway/clients"
	"yorqinbek/microservices/Blogpost/api_gateway/config"
)

// Handler ...
type handler struct {
	cfg         config.Config
	grpcClients *clients.GrpcClients
}

// NewHandler ...
func NewHandler(cfg config.Config, grpcClients *clients.GrpcClients) handler {
	return handler{
		cfg:         cfg,
		grpcClients: grpcClients,
	}
}
