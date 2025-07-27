package client

import (
	"context"
	"time"

	"scd-service/proto"

	"google.golang.org/grpc"
)

type Client struct {
	conn   *grpc.ClientConn
	client proto.SCDServiceClient
}

func NewClient(address string) (*Client, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		return nil, err
	}
	client := proto.NewSCDServiceClient(conn)
	return &Client{conn: conn, client: client}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) GetLatestJobs(statusFilter string) ([]*proto.Job, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req := &proto.GetLatestJobsRequest{StatusFilter: statusFilter}
	res, err := c.client.GetLatestJobs(ctx, req)
	if err != nil {
		return nil, err
	}

	return res.Jobs, nil
}

func (c *Client) UpdateJob(id string, updates map[string]string) (*proto.Job, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req := &proto.UpdateJobRequest{
		Id:            id,
		UpdatedFields: updates,
	}

	res, err := c.client.UpdateJob(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
