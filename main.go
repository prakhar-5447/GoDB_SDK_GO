package godb

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/prakhar-5447/GoDB_SDK_GO/proto"

	"google.golang.org/grpc"
)

// GoDBClient wraps the gRPC client and connection.
type GoDBClient struct {
	client           proto.DatabaseServiceClient
	conn             *grpc.ClientConn
	connectionString string
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

// SetConnectionString stores the connection string for subsequent operations.
func (c *GoDBClient) SetConnectionString(connStr string) {
	c.connectionString = connStr
}

// CreateUser calls the gRPC CreateUser method to register a new user and returns
// both a message and a connection string with a placeholder for the database name.
func (c *GoDBClient) CreateUser(ctx context.Context, username, password string) (string, string, error) {
	req := &proto.CreateUserRequest{
		Username: username,
		Password: password,
	}
	resp, err := c.client.CreateUser(ctx, req)
	if err != nil {
		return "", "", fmt.Errorf("failed to create user: %w", err)
	}
	return resp.Message, resp.ConnectionString, nil
}

// CreateDatabase creates a new database for a user.
func (c *GoDBClient) CreateDatabase(ctx context.Context, connectionString string) (string, error) {
	req := &proto.CreateDatabaseRequest{ConnectionString: connectionString}
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

// ---------- Fluent Builder Types and Methods ----------

// UpdateTableBuilder provides a fluent interface for updating table structure.
type UpdateTableBuilder struct {
	client     *GoDBClient
	ctx        context.Context
	tableName  string
	columnName string
	columnType string
}

// NewUpdateTable creates a new UpdateTableBuilder using the client's stored connection string.
func (client *GoDBClient) UpdateTable(ctx context.Context) *UpdateTableBuilder {
	return &UpdateTableBuilder{
		client: client,
		ctx:    ctx,
	}
}

// Table sets the table name.
func (utb *UpdateTableBuilder) Table(table string) *UpdateTableBuilder {
	utb.tableName = table
	return utb
}

// AddColumn sets the new column name and type.
func (utb *UpdateTableBuilder) AddColumn(name, colType string) *UpdateTableBuilder {
	utb.columnName = name
	utb.columnType = colType
	return utb
}

// Exec executes the update table operation.
func (utb *UpdateTableBuilder) Exec() (string, error) {
	if utb.tableName == "" {
		return "", fmt.Errorf("table name is required")
	}
	if utb.columnName == "" || utb.columnType == "" {
		return "", fmt.Errorf("column name and type are required")
	}
	req := &proto.UpdateTableRequest{
		TableName:        utb.tableName,
		ColumnName:       utb.columnName,
		ColumnType:       utb.columnType,
		ConnectionString: utb.client.connectionString,
	}
	resp, err := utb.client.client.UpdateTable(utb.ctx, req)
	if err != nil {
		return "", err
	}
	return resp.Message, nil
}

// InsertBuilder provides a fluent interface for building an insert operation.
type InsertBuilder struct {
	client    *GoDBClient
	ctx       context.Context
	tableName string
	record    map[string]string
}

// Insert returns a new InsertBuilder using the client's stored connection string.
func (client *GoDBClient) Insert(ctx context.Context) *InsertBuilder {
	return &InsertBuilder{
		client: client,
		ctx:    ctx,
		record: make(map[string]string),
	}
}

// Table sets the table name for the InsertBuilder.
func (ib *InsertBuilder) Table(table string) *InsertBuilder {
	ib.tableName = table
	return ib
}

// Values sets the record values.
func (ib *InsertBuilder) Values(record map[string]string) *InsertBuilder {
	ib.record = record
	return ib
}

// Exec executes the insert operation.
func (ib *InsertBuilder) Exec() (string, error) {
	if ib.tableName == "" {
		return "", fmt.Errorf("table name is required")
	}
	if ib.record == nil || len(ib.record) == 0 {
		return "", fmt.Errorf("no record provided")
	}
	// Construct the request directly.
	req := &proto.InsertRecordRequest{
		TableName:        ib.tableName,
		Record:           ib.record,
		ConnectionString: ib.client.connectionString,
	}
	// Directly call the gRPC method on the underlying client.
	resp, err := ib.client.client.InsertRecord(ib.ctx, req)
	if err != nil {
		return "", err
	}
	return resp.Message, nil
}

// InsertMultipleBuilder provides a fluent interface for inserting multiple records.
type InsertMultipleBuilder struct {
	client    *GoDBClient
	ctx       context.Context
	tableName string
	records   []*proto.Record
}

// NewInsertMultiple returns a new InsertMultipleBuilder using the client's stored connection string.
func (client *GoDBClient) InsertMultiple(ctx context.Context) *InsertMultipleBuilder {
	return &InsertMultipleBuilder{
		client:  client,
		ctx:     ctx,
		records: make([]*proto.Record, 0),
	}
}

// Table sets the table name.
func (imb *InsertMultipleBuilder) Table(table string) *InsertMultipleBuilder {
	imb.tableName = table
	return imb
}

// Records sets multiple records at once.
func (imb *InsertMultipleBuilder) Records(records []map[string]string) *InsertMultipleBuilder {
	for _, rec := range records {
		imb.records = append(imb.records, &proto.Record{Data: rec})
	}
	return imb
}

// Exec executes the insert operation by directly calling the gRPC InsertMultipleRecords API.
func (imb *InsertMultipleBuilder) Exec() (string, error) {
	if imb.tableName == "" {
		return "", fmt.Errorf("table name is required")
	}
	if len(imb.records) == 0 {
		return "", fmt.Errorf("no records provided")
	}

	req := &proto.InsertMultipleRecordsRequest{
		TableName:        imb.tableName,
		Records:          imb.records,
		ConnectionString: imb.client.connectionString,
	}
	resp, err := imb.client.client.InsertMultipleRecords(imb.ctx, req)
	if err != nil {
		return "", err
	}
	return resp.Message, nil
}

// UpdateRecordBuilder provides a fluent interface for updating records.
type UpdateRecordBuilder struct {
	client           *GoDBClient
	ctx              context.Context
	tableName        string
	updates          map[string]string
	condition        string
	connectionString string
}

// NewUpdateRecord creates a new UpdateRecordBuilder using the client's stored connection string.
func (client *GoDBClient) UpdateRecord(ctx context.Context) *UpdateRecordBuilder {
	return &UpdateRecordBuilder{
		client:           client,
		ctx:              ctx,
		updates:          make(map[string]string),
	}
}

// Table sets the table name.
func (urb *UpdateRecordBuilder) Table(table string) *UpdateRecordBuilder {
	urb.tableName = table
	return urb
}

// SetUpdate sets a key-value update.
func (urb *UpdateRecordBuilder) SetUpdate(field string, value interface{}) *UpdateRecordBuilder {
	urb.updates[field] = fmt.Sprintf("%v", value)
	return urb
}

// Updates sets multiple updates at once.
func (urb *UpdateRecordBuilder) Updates(upds map[string]interface{}) *UpdateRecordBuilder {
	for k, v := range upds {
		urb.updates[k] = fmt.Sprintf("%v", v)
	}
	return urb
}

// Condition sets a custom WHERE condition.
func (urb *UpdateRecordBuilder) Condition(cond string) *UpdateRecordBuilder {
	urb.condition = cond
	return urb
}

// Equal adds an equality condition.
func (urb *UpdateRecordBuilder) Equal(field string, value interface{}) *UpdateRecordBuilder {
	cond := formatCondition(field, "=", value)
	urb.addCondition(cond)
	return urb
}

// Greater adds a greater-than condition.
func (urb *UpdateRecordBuilder) Greater(field string, value interface{}) *UpdateRecordBuilder {
	cond := formatCondition(field, ">", value)
	urb.addCondition(cond)
	return urb
}

// Less adds a less-than condition.
func (urb *UpdateRecordBuilder) Less(field string, value interface{}) *UpdateRecordBuilder {
	cond := formatCondition(field, "<", value)
	urb.addCondition(cond)
	return urb
}

// addCondition appends a condition to the builder.
func (urb *UpdateRecordBuilder) addCondition(cond string) {
	if urb.condition != "" {
		urb.condition += " AND " + cond
	} else {
		urb.condition = cond
	}
}

// Exec executes the update record operation.
func (urb *UpdateRecordBuilder) Exec() (string, error) {
	if urb.tableName == "" {
		return "", fmt.Errorf("table name is required")
	}
	if len(urb.updates) == 0 {
		return "", fmt.Errorf("no updates provided")
	}
	req := &proto.UpdateRecordRequest{
		TableName:        urb.tableName,
		Updates:          urb.updates,
		Condition:        urb.condition,
		ConnectionString: urb.client.connectionString,
	}
	resp, err := urb.client.client.UpdateRecord(urb.ctx, req)
	if err != nil {
		return "", err
	}
	return resp.Message, nil
}

// QueryBuilder provides a fluent interface for building queries.
type QueryBuilder struct {
	client    *GoDBClient
	ctx       context.Context
	tableName string
	columns   string
	condition string
	orderBy   string
	limit     int
	offset    int
	cursor    string
}

// Query creates a new QueryBuilder using the client's stored connection string.
func (client *GoDBClient) Query(ctx context.Context) *QueryBuilder {
	return &QueryBuilder{
		client: client,
		ctx:    ctx,
		limit:  0,
		offset: 0,
	}
}

// Table sets the table name.
func (qb *QueryBuilder) Table(table string) *QueryBuilder {
	qb.tableName = table
	return qb
}

// Columns sets the columns to select.
func (qb *QueryBuilder) Columns(cols string) *QueryBuilder {
	qb.columns = cols
	return qb
}

// Condition sets a custom WHERE condition.
func (qb *QueryBuilder) Condition(cond string) *QueryBuilder {
	qb.condition = cond
	return qb
}

// Equal adds an equality condition (e.g., field = value).
func (qb *QueryBuilder) Equal(field string, value interface{}) *QueryBuilder {
	condition := formatCondition(field, "=", value)
	qb.addCondition(condition)
	return qb
}

// Greater adds a greater-than condition (e.g., field > value).
func (qb *QueryBuilder) Greater(field string, value interface{}) *QueryBuilder {
	condition := formatCondition(field, ">", value)
	qb.addCondition(condition)
	return qb
}

// Less adds a less-than condition (e.g., field < value).
func (qb *QueryBuilder) Less(field string, value interface{}) *QueryBuilder {
	condition := formatCondition(field, "<", value)
	qb.addCondition(condition)
	return qb
}

// LessEqual adds a less-than-or-equal condition (e.g., field <= value).
func (qb *QueryBuilder) LessEqual(field string, value interface{}) *QueryBuilder {
	condition := formatCondition(field, "<=", value)
	qb.addCondition(condition)
	return qb
}

// addCondition appends a new condition to the builder.
func (qb *QueryBuilder) addCondition(cond string) {
	if qb.condition != "" {
		qb.condition += " AND " + cond
	} else {
		qb.condition = cond
	}
}

// Cursor sets a cursor for pagination. It will add a condition like "id > {cursor}".
func (qb *QueryBuilder) Cursor(cursor string) *QueryBuilder {
	qb.cursor = cursor
	return qb
}

// OrderBy sets the ORDER BY clause.
func (qb *QueryBuilder) OrderBy(order string) *QueryBuilder {
	qb.orderBy = order
	return qb
}

// Limit sets the maximum number of rows.
func (qb *QueryBuilder) Limit(limit int) *QueryBuilder {
	qb.limit = limit
	return qb
}

// Offset sets the offset.
func (qb *QueryBuilder) Offset(offset int) *QueryBuilder {
	qb.offset = offset
	return qb
}

// Exec constructs the QueryDataRequest and directly calls the gRPC QueryData API.
func (qb *QueryBuilder) Exec() (*proto.QueryDataResponse, error) {
	// Build conditions.
	var conditions []string
	if qb.condition != "" {
		conditions = append(conditions, qb.condition)
	}
	// If cursor is provided, add a condition for pagination.
	if qb.cursor != "" {
		conditions = append(conditions, fmt.Sprintf("id > %s", qb.cursor))
	}

	finalCondition := ""
	if len(conditions) > 0 {
		finalCondition = strings.Join(conditions, " AND ")
	}

	// Append ORDER BY clause if provided.
	if qb.orderBy != "" {
		finalCondition += " ORDER BY " + qb.orderBy
	}
	// Append LIMIT and OFFSET.
	if qb.limit > 0 {
		finalCondition += " LIMIT " + strconv.Itoa(qb.limit)
	}
	if qb.offset > 0 {
		finalCondition += " OFFSET " + strconv.Itoa(qb.offset)
	}

	req := &proto.QueryDataRequest{
		ConnectionString: qb.client.connectionString,
		TableName:        qb.tableName,
		Columns:          qb.columns,
		Condition:        finalCondition,
	}
	return qb.client.client.QueryData(qb.ctx, req)
}

// formatCondition formats the condition based on the operator and value.
// If the value is a string, it adds quotes around it.
func formatCondition(field, operator string, value interface{}) string {
	switch v := value.(type) {
	case string:
		escaped := strings.ReplaceAll(v, "'", "''")
		return fmt.Sprintf("%s %s '%s'", field, operator, escaped)
	default:
		return fmt.Sprintf("%s %s %v", field, operator, v)
	}
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
