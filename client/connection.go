package client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"github.com/jumba-nl/go-sdk/credential"
	"crypto/tls"
	"google.golang.org/grpc/metadata"
	"context"
)

var defaultBackendHost = "api.jumba.nl:443"

var client *Client

func CreateContext(token string) context.Context {
	md := metadata.Pairs("token", token)
	return metadata.NewOutgoingContext(context.Background(), md)
}

func SetClient(cl *Client) {
	client = cl
}

func GetClient() *Client {
	return client
}

type Client struct {
	Host  string
	Token string
	Dev   bool
	Conn  *grpc.ClientConn
}

func NewDefaultClient() *Client {
	return &Client{
		Host: defaultBackendHost,
	}
}

func NewClient(host, token string, dev bool) *Client {
	client := NewDefaultClient()
	if host != "" {
		client.Host = host
	}
	client.Token = token
	client.Dev = dev

	return client
}

func (c *Client) Open() error {
	var err error

	dialOptions := []grpc.DialOption{}
	if c.Dev {
		dialOptions = append(dialOptions, grpc.WithInsecure())
	}

	if !c.Dev {
		creds := credentials.NewServerTLSFromCert(&tls.Certificate{})
		dialOptions = append(dialOptions, grpc.WithTransportCredentials(creds))
	}

	dialOptions = append(dialOptions, grpc.WithPerRPCCredentials(&credential.TokenCreds{Token: c.Token}))

	c.Conn, err = grpc.Dial(
		c.Host,
		dialOptions...,
	)

	return err
}

func (c *Client) Close() error {
	if c.Conn != nil {
		err := c.Conn.Close()
		if err != nil {
			return err
		}
	}

	return nil
}
