package client

import (
	"context"
	"fmt"
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

func (c *Client) GetPaymentLineItem(uid string) (*proto.PaymentLineItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req := &proto.GetLatestPaymentLineItemsRequest{}
	res, err := c.client.GetLatestPaymentLineItems(ctx, req)
	if err != nil {
		return nil, err
	}

	for _, item := range res.Items {
		if item.Uid == uid {
			return item, nil
		}
	}

	return nil, fmt.Errorf("PaymentLineItem with UID %s not found", uid)
}

func (c *Client) GetTimelog(uid string) (*proto.Timelog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req := &proto.GetLatestTimelogsRequest{}
	res, err := c.client.GetLatestTimelogs(ctx, req)
	if err != nil {
		return nil, err
	}

	for _, tl := range res.Timelogs {
		if tl.Uid == uid {
			return tl, nil
		}
	}

	return nil, fmt.Errorf("Timelog with UID %s not found", uid)
}

func (c *Client) UpdatePaymentLineItem(uid string, updates map[string]string) (*proto.PaymentLineItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req := &proto.UpdatePaymentLineItemRequest{
		Id:            uid,
		UpdatedFields: updates,
	}

	res, err := c.client.UpdatePaymentLineItem(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) UpdateTimelog(uid string, updates map[string]string) (*proto.Timelog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req := &proto.UpdateTimelogRequest{
		Id:            uid,
		UpdatedFields: updates,
	}

	res, err := c.client.UpdateTimelog(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
