// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/aakarim/pland/ent/migrate"

	"github.com/aakarim/pland/ent/arbitrarysection"
	"github.com/aakarim/pland/ent/day"
	"github.com/aakarim/pland/ent/header"
	"github.com/aakarim/pland/ent/plan"
	"github.com/aakarim/pland/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// ArbitrarySection is the client for interacting with the ArbitrarySection builders.
	ArbitrarySection *ArbitrarySectionClient
	// Day is the client for interacting with the Day builders.
	Day *DayClient
	// Header is the client for interacting with the Header builders.
	Header *HeaderClient
	// Plan is the client for interacting with the Plan builders.
	Plan *PlanClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// additional fields for node api
	tables tables
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.ArbitrarySection = NewArbitrarySectionClient(c.config)
	c.Day = NewDayClient(c.config)
	c.Header = NewHeaderClient(c.config)
	c.Plan = NewPlanClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:              ctx,
		config:           cfg,
		ArbitrarySection: NewArbitrarySectionClient(cfg),
		Day:              NewDayClient(cfg),
		Header:           NewHeaderClient(cfg),
		Plan:             NewPlanClient(cfg),
		User:             NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:              ctx,
		config:           cfg,
		ArbitrarySection: NewArbitrarySectionClient(cfg),
		Day:              NewDayClient(cfg),
		Header:           NewHeaderClient(cfg),
		Plan:             NewPlanClient(cfg),
		User:             NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		ArbitrarySection.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.ArbitrarySection.Use(hooks...)
	c.Day.Use(hooks...)
	c.Header.Use(hooks...)
	c.Plan.Use(hooks...)
	c.User.Use(hooks...)
}

// ArbitrarySectionClient is a client for the ArbitrarySection schema.
type ArbitrarySectionClient struct {
	config
}

// NewArbitrarySectionClient returns a client for the ArbitrarySection from the given config.
func NewArbitrarySectionClient(c config) *ArbitrarySectionClient {
	return &ArbitrarySectionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `arbitrarysection.Hooks(f(g(h())))`.
func (c *ArbitrarySectionClient) Use(hooks ...Hook) {
	c.hooks.ArbitrarySection = append(c.hooks.ArbitrarySection, hooks...)
}

// Create returns a builder for creating a ArbitrarySection entity.
func (c *ArbitrarySectionClient) Create() *ArbitrarySectionCreate {
	mutation := newArbitrarySectionMutation(c.config, OpCreate)
	return &ArbitrarySectionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ArbitrarySection entities.
func (c *ArbitrarySectionClient) CreateBulk(builders ...*ArbitrarySectionCreate) *ArbitrarySectionCreateBulk {
	return &ArbitrarySectionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ArbitrarySection.
func (c *ArbitrarySectionClient) Update() *ArbitrarySectionUpdate {
	mutation := newArbitrarySectionMutation(c.config, OpUpdate)
	return &ArbitrarySectionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ArbitrarySectionClient) UpdateOne(as *ArbitrarySection) *ArbitrarySectionUpdateOne {
	mutation := newArbitrarySectionMutation(c.config, OpUpdateOne, withArbitrarySection(as))
	return &ArbitrarySectionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ArbitrarySectionClient) UpdateOneID(id int) *ArbitrarySectionUpdateOne {
	mutation := newArbitrarySectionMutation(c.config, OpUpdateOne, withArbitrarySectionID(id))
	return &ArbitrarySectionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ArbitrarySection.
func (c *ArbitrarySectionClient) Delete() *ArbitrarySectionDelete {
	mutation := newArbitrarySectionMutation(c.config, OpDelete)
	return &ArbitrarySectionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ArbitrarySectionClient) DeleteOne(as *ArbitrarySection) *ArbitrarySectionDeleteOne {
	return c.DeleteOneID(as.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *ArbitrarySectionClient) DeleteOneID(id int) *ArbitrarySectionDeleteOne {
	builder := c.Delete().Where(arbitrarysection.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ArbitrarySectionDeleteOne{builder}
}

// Query returns a query builder for ArbitrarySection.
func (c *ArbitrarySectionClient) Query() *ArbitrarySectionQuery {
	return &ArbitrarySectionQuery{
		config: c.config,
	}
}

// Get returns a ArbitrarySection entity by its id.
func (c *ArbitrarySectionClient) Get(ctx context.Context, id int) (*ArbitrarySection, error) {
	return c.Query().Where(arbitrarysection.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ArbitrarySectionClient) GetX(ctx context.Context, id int) *ArbitrarySection {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPlan queries the plan edge of a ArbitrarySection.
func (c *ArbitrarySectionClient) QueryPlan(as *ArbitrarySection) *PlanQuery {
	query := &PlanQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := as.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(arbitrarysection.Table, arbitrarysection.FieldID, id),
			sqlgraph.To(plan.Table, plan.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, arbitrarysection.PlanTable, arbitrarysection.PlanPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(as.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ArbitrarySectionClient) Hooks() []Hook {
	return c.hooks.ArbitrarySection
}

// DayClient is a client for the Day schema.
type DayClient struct {
	config
}

// NewDayClient returns a client for the Day from the given config.
func NewDayClient(c config) *DayClient {
	return &DayClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `day.Hooks(f(g(h())))`.
func (c *DayClient) Use(hooks ...Hook) {
	c.hooks.Day = append(c.hooks.Day, hooks...)
}

// Create returns a builder for creating a Day entity.
func (c *DayClient) Create() *DayCreate {
	mutation := newDayMutation(c.config, OpCreate)
	return &DayCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Day entities.
func (c *DayClient) CreateBulk(builders ...*DayCreate) *DayCreateBulk {
	return &DayCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Day.
func (c *DayClient) Update() *DayUpdate {
	mutation := newDayMutation(c.config, OpUpdate)
	return &DayUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DayClient) UpdateOne(d *Day) *DayUpdateOne {
	mutation := newDayMutation(c.config, OpUpdateOne, withDay(d))
	return &DayUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DayClient) UpdateOneID(id int) *DayUpdateOne {
	mutation := newDayMutation(c.config, OpUpdateOne, withDayID(id))
	return &DayUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Day.
func (c *DayClient) Delete() *DayDelete {
	mutation := newDayMutation(c.config, OpDelete)
	return &DayDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *DayClient) DeleteOne(d *Day) *DayDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *DayClient) DeleteOneID(id int) *DayDeleteOne {
	builder := c.Delete().Where(day.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DayDeleteOne{builder}
}

// Query returns a query builder for Day.
func (c *DayClient) Query() *DayQuery {
	return &DayQuery{
		config: c.config,
	}
}

// Get returns a Day entity by its id.
func (c *DayClient) Get(ctx context.Context, id int) (*Day, error) {
	return c.Query().Where(day.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DayClient) GetX(ctx context.Context, id int) *Day {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPlan queries the plan edge of a Day.
func (c *DayClient) QueryPlan(d *Day) *PlanQuery {
	query := &PlanQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(day.Table, day.FieldID, id),
			sqlgraph.To(plan.Table, plan.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, day.PlanTable, day.PlanPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DayClient) Hooks() []Hook {
	return c.hooks.Day
}

// HeaderClient is a client for the Header schema.
type HeaderClient struct {
	config
}

// NewHeaderClient returns a client for the Header from the given config.
func NewHeaderClient(c config) *HeaderClient {
	return &HeaderClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `header.Hooks(f(g(h())))`.
func (c *HeaderClient) Use(hooks ...Hook) {
	c.hooks.Header = append(c.hooks.Header, hooks...)
}

// Create returns a builder for creating a Header entity.
func (c *HeaderClient) Create() *HeaderCreate {
	mutation := newHeaderMutation(c.config, OpCreate)
	return &HeaderCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Header entities.
func (c *HeaderClient) CreateBulk(builders ...*HeaderCreate) *HeaderCreateBulk {
	return &HeaderCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Header.
func (c *HeaderClient) Update() *HeaderUpdate {
	mutation := newHeaderMutation(c.config, OpUpdate)
	return &HeaderUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *HeaderClient) UpdateOne(h *Header) *HeaderUpdateOne {
	mutation := newHeaderMutation(c.config, OpUpdateOne, withHeader(h))
	return &HeaderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *HeaderClient) UpdateOneID(id int) *HeaderUpdateOne {
	mutation := newHeaderMutation(c.config, OpUpdateOne, withHeaderID(id))
	return &HeaderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Header.
func (c *HeaderClient) Delete() *HeaderDelete {
	mutation := newHeaderMutation(c.config, OpDelete)
	return &HeaderDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *HeaderClient) DeleteOne(h *Header) *HeaderDeleteOne {
	return c.DeleteOneID(h.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *HeaderClient) DeleteOneID(id int) *HeaderDeleteOne {
	builder := c.Delete().Where(header.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &HeaderDeleteOne{builder}
}

// Query returns a query builder for Header.
func (c *HeaderClient) Query() *HeaderQuery {
	return &HeaderQuery{
		config: c.config,
	}
}

// Get returns a Header entity by its id.
func (c *HeaderClient) Get(ctx context.Context, id int) (*Header, error) {
	return c.Query().Where(header.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *HeaderClient) GetX(ctx context.Context, id int) *Header {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPlan queries the plan edge of a Header.
func (c *HeaderClient) QueryPlan(h *Header) *PlanQuery {
	query := &PlanQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := h.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(header.Table, header.FieldID, id),
			sqlgraph.To(plan.Table, plan.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, header.PlanTable, header.PlanColumn),
		)
		fromV = sqlgraph.Neighbors(h.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *HeaderClient) Hooks() []Hook {
	return c.hooks.Header
}

// PlanClient is a client for the Plan schema.
type PlanClient struct {
	config
}

// NewPlanClient returns a client for the Plan from the given config.
func NewPlanClient(c config) *PlanClient {
	return &PlanClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `plan.Hooks(f(g(h())))`.
func (c *PlanClient) Use(hooks ...Hook) {
	c.hooks.Plan = append(c.hooks.Plan, hooks...)
}

// Create returns a builder for creating a Plan entity.
func (c *PlanClient) Create() *PlanCreate {
	mutation := newPlanMutation(c.config, OpCreate)
	return &PlanCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Plan entities.
func (c *PlanClient) CreateBulk(builders ...*PlanCreate) *PlanCreateBulk {
	return &PlanCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Plan.
func (c *PlanClient) Update() *PlanUpdate {
	mutation := newPlanMutation(c.config, OpUpdate)
	return &PlanUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PlanClient) UpdateOne(pl *Plan) *PlanUpdateOne {
	mutation := newPlanMutation(c.config, OpUpdateOne, withPlan(pl))
	return &PlanUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PlanClient) UpdateOneID(id int) *PlanUpdateOne {
	mutation := newPlanMutation(c.config, OpUpdateOne, withPlanID(id))
	return &PlanUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Plan.
func (c *PlanClient) Delete() *PlanDelete {
	mutation := newPlanMutation(c.config, OpDelete)
	return &PlanDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PlanClient) DeleteOne(pl *Plan) *PlanDeleteOne {
	return c.DeleteOneID(pl.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *PlanClient) DeleteOneID(id int) *PlanDeleteOne {
	builder := c.Delete().Where(plan.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PlanDeleteOne{builder}
}

// Query returns a query builder for Plan.
func (c *PlanClient) Query() *PlanQuery {
	return &PlanQuery{
		config: c.config,
	}
}

// Get returns a Plan entity by its id.
func (c *PlanClient) Get(ctx context.Context, id int) (*Plan, error) {
	return c.Query().Where(plan.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PlanClient) GetX(ctx context.Context, id int) *Plan {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAuthor queries the author edge of a Plan.
func (c *PlanClient) QueryAuthor(pl *Plan) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(plan.Table, plan.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, plan.AuthorTable, plan.AuthorColumn),
		)
		fromV = sqlgraph.Neighbors(pl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDays queries the days edge of a Plan.
func (c *PlanClient) QueryDays(pl *Plan) *DayQuery {
	query := &DayQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(plan.Table, plan.FieldID, id),
			sqlgraph.To(day.Table, day.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, plan.DaysTable, plan.DaysPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(pl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryArbitrarySections queries the arbitrarySections edge of a Plan.
func (c *PlanClient) QueryArbitrarySections(pl *Plan) *ArbitrarySectionQuery {
	query := &ArbitrarySectionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(plan.Table, plan.FieldID, id),
			sqlgraph.To(arbitrarysection.Table, arbitrarysection.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, plan.ArbitrarySectionsTable, plan.ArbitrarySectionsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(pl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryHeader queries the header edge of a Plan.
func (c *PlanClient) QueryHeader(pl *Plan) *HeaderQuery {
	query := &HeaderQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(plan.Table, plan.FieldID, id),
			sqlgraph.To(header.Table, header.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, plan.HeaderTable, plan.HeaderColumn),
		)
		fromV = sqlgraph.Neighbors(pl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPrev queries the prev edge of a Plan.
func (c *PlanClient) QueryPrev(pl *Plan) *PlanQuery {
	query := &PlanQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(plan.Table, plan.FieldID, id),
			sqlgraph.To(plan.Table, plan.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, plan.PrevTable, plan.PrevColumn),
		)
		fromV = sqlgraph.Neighbors(pl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryNext queries the next edge of a Plan.
func (c *PlanClient) QueryNext(pl *Plan) *PlanQuery {
	query := &PlanQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(plan.Table, plan.FieldID, id),
			sqlgraph.To(plan.Table, plan.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, plan.NextTable, plan.NextColumn),
		)
		fromV = sqlgraph.Neighbors(pl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PlanClient) Hooks() []Hook {
	return c.hooks.Plan
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPlans queries the plans edge of a User.
func (c *UserClient) QueryPlans(u *User) *PlanQuery {
	query := &PlanQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(plan.Table, plan.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.PlansTable, user.PlansColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
