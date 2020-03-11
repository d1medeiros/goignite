package grpc

import (
	"log"

	"github.com/b2wdigital/goignite/pkg/config"
)

const (
	Port                 = "grpc.server.port"
	MaxConcurrentStreams = "grpc.server.maxconcurrentstreams"
	TlsEnabled = "grpc.server.tls.enabled"
	CertFile   = "grpc.server.tls.certfile"
	KeyFile    = "grpc.server.tls.keyfile"
	CaFile     = "grpc.server.tls.cafile"
)

func init() {

	log.Println("getting configurations for grpc server")

	config.Add(Port, 9090, "server grpc port")
	config.Add(MaxConcurrentStreams, 5000, "server grpc max concurrent streams")
	config.Add(TlsEnabled, false, "Use TLS - required for HTTP2.")
	config.Add(CertFile, "./cert/out/localhost.crt", "Path to the CRT/PEM file.")
	config.Add(KeyFile, "./cert/out/localhost.key", "Path to the private key file.")
	config.Add(CaFile, "./cert/out/blackbox.crt", "Path to the certificate authority (CA).")

}