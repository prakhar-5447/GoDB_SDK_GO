package godb

import (
	"context"
	"fmt"
	"strconv"

	"github.com/prakhar-5447/GoDB_SDK_GO/proto"
	"google.golang.org/grpc"
)

// GoDBClient wraps the gRPC client and connection.
type GoDBClient struct {
	client proto.DatabaseServiceClient
	conn   *grpc.ClientConn
}

// NewGoDBClient creates a new instance of GoDBClient.
// The address parameter should be the IP and port of your Docker container running the gRPC server,
// e.g., "172.17.0.2:50051" or a DNS name if using Docker networking.
func NewGoDBClient(address string) (*GoDBClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to GoDB: %v", err)
	}
	client := proto.NewDatabaseServiceClient(conn)
	return &GoDBClient{client: client, conn: conn}, nil
}

// Close closes the underlying gRPC connection.
func (c *GoDBClient) Close() error {
	return c.conn.Close()
}

// CreateDatabase creates a new database for a user.
func (c *GoDBClient) CreateDatabase(ctx context.Context) (string, error) {
	req := &proto.CreateDatabaseRequest{}
	resp, err := c.client.CreateDatabase(ctx, req)
	if err != nil {
		return "", err
	}
	return resp.Message, nil
}

// CreateTable creates a new table in the specified user database.
func (c *GoDBClient) CreateTable(ctx context.Context, tableName string, columns map[string]string, connectionString string) (string, error) {
	req := &proto.CreateTableRequest{
		TableName:        tableName,
		Columns:          columns,
		ConnectionString: connectionString,
	}
	resp, err := c.client.CreateTable(ctx, req)
	if err != nil {
		return "", err
	}
	return resp.Message, nil
}

// InsertRecord inserts a record into a table.
func (c *GoDBClient) InsertRecord(ctx context.Context, tableName string, record map[string]string, connectionString string) (string, error) {
	req := &proto.InsertRecordRequest{
		TableName:        tableName,
		Record:           record,
		ConnectionString: connectionString,
	}
	resp, err := c.client.InsertRecord(ctx, req)
	if err != nil {
		return "", err
	}
	return resp.Message, nil
}

// QueryData queries data from a table. You can pass condition, limit, and offset.
func (c *GoDBClient) QueryData(ctx context.Context, tableName, columns, condition string, limit, offset int, connectionString string) (*proto.QueryDataResponse, error) {
	// Build query condition string with LIMIT and OFFSET if provided.
	queryCondition := condition
	if limit > 0 {
		queryCondition += " LIMIT " + strconv.Itoa(limit)
	}
	if offset > 0 {
		queryCondition += " OFFSET " + strconv.Itoa(offset)
	}
	req := &proto.QueryDataRequest{
		TableName:        tableName,
		Columns:          columns,
		Condition:        queryCondition,
		ConnectionString: connectionString,
	}
	return c.client.QueryData(ctx, req)
}

// UpdateTable updates a table by adding a new column.
func (c *GoDBClient) UpdateTable(ctx context.Context, tableName, columnName, columnType, connectionString string) (string, error) {
	req := &proto.UpdateTableRequest{
		TableName:        tableName,
		ColumnName:       columnName,
		ColumnType:       columnType,
		ConnectionString: connectionString,
	}
	resp, err := c.client.UpdateTable(ctx, req)
	if err != nil {
		return "", err
	}
	return resp.Message, nil
}

// UpdateRecord updates records in a table based on a condition.
func (c *GoDBClient) UpdateRecord(ctx context.Context, tableName, condition string, updates map[string]string, connectionString string) (string, error) {
	req := &proto.UpdateRecordRequest{
		TableName:        tableName,
		Updates:          updates,
		Condition:        condition,
		ConnectionString: connectionString,
	}
	resp, err := c.client.UpdateRecord(ctx, req)
	if err != nil {
		return "", err
	}
	return resp.Message, nil
}

// AddIndex creates an index on a table.
func (c *GoDBClient) AddIndex(ctx context.Context, tableName, indexName string, columns []string, connectionString string) (string, error) {
	req := &proto.AddIndexRequest{
		TableName:        tableName,
		IndexName:        indexName,
		Columns:          columns,
		ConnectionString: connectionString,
	}
	resp, err := c.client.AddIndex(ctx, req)
	if err != nil {
		return "", err
	}
	return resp.Message, nil
}

// DeleteIndex deletes an index from a table.
func (c *GoDBClient) DeleteIndex(ctx context.Context, indexName, connectionString string) (string, error) {
	req := &proto.DeleteIndexRequest{
		IndexName:        indexName,
		ConnectionString: connectionString,
	}
	resp, err := c.client.DeleteIndex(ctx, req)
	if err != nil {
		return "", err
	}
	return resp.Message, nil
}

// ListIndexes lists all indexes for a given user's database.
func (c *GoDBClient) ListIndexes(ctx context.Context, connectionString string) (*proto.ListIndexesResponse, error) {
	req := &proto.ListIndexesRequest{
		ConnectionString: connectionString,
	}
	return c.client.ListIndexes(ctx, req)
}
