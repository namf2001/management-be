// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"management-be/internal/repository/ent/match"
	"management-be/internal/repository/ent/matchplayer"
	"management-be/internal/repository/ent/player"
	"management-be/internal/repository/ent/predicate"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MatchPlayerQuery is the builder for querying MatchPlayer entities.
type MatchPlayerQuery struct {
	config
	ctx        *QueryContext
	order      []matchplayer.OrderOption
	inters     []Interceptor
	predicates []predicate.MatchPlayer
	withMatch  *MatchQuery
	withPlayer *PlayerQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MatchPlayerQuery builder.
func (mpq *MatchPlayerQuery) Where(ps ...predicate.MatchPlayer) *MatchPlayerQuery {
	mpq.predicates = append(mpq.predicates, ps...)
	return mpq
}

// Limit the number of records to be returned by this query.
func (mpq *MatchPlayerQuery) Limit(limit int) *MatchPlayerQuery {
	mpq.ctx.Limit = &limit
	return mpq
}

// Offset to start from.
func (mpq *MatchPlayerQuery) Offset(offset int) *MatchPlayerQuery {
	mpq.ctx.Offset = &offset
	return mpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mpq *MatchPlayerQuery) Unique(unique bool) *MatchPlayerQuery {
	mpq.ctx.Unique = &unique
	return mpq
}

// Order specifies how the records should be ordered.
func (mpq *MatchPlayerQuery) Order(o ...matchplayer.OrderOption) *MatchPlayerQuery {
	mpq.order = append(mpq.order, o...)
	return mpq
}

// QueryMatch chains the current query on the "match" edge.
func (mpq *MatchPlayerQuery) QueryMatch() *MatchQuery {
	query := (&MatchClient{config: mpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(matchplayer.Table, matchplayer.FieldID, selector),
			sqlgraph.To(match.Table, match.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, matchplayer.MatchTable, matchplayer.MatchColumn),
		)
		fromU = sqlgraph.SetNeighbors(mpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPlayer chains the current query on the "player" edge.
func (mpq *MatchPlayerQuery) QueryPlayer() *PlayerQuery {
	query := (&PlayerClient{config: mpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(matchplayer.Table, matchplayer.FieldID, selector),
			sqlgraph.To(player.Table, player.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, matchplayer.PlayerTable, matchplayer.PlayerColumn),
		)
		fromU = sqlgraph.SetNeighbors(mpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MatchPlayer entity from the query.
// Returns a *NotFoundError when no MatchPlayer was found.
func (mpq *MatchPlayerQuery) First(ctx context.Context) (*MatchPlayer, error) {
	nodes, err := mpq.Limit(1).All(setContextOp(ctx, mpq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{matchplayer.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mpq *MatchPlayerQuery) FirstX(ctx context.Context) *MatchPlayer {
	node, err := mpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MatchPlayer ID from the query.
// Returns a *NotFoundError when no MatchPlayer ID was found.
func (mpq *MatchPlayerQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mpq.Limit(1).IDs(setContextOp(ctx, mpq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{matchplayer.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mpq *MatchPlayerQuery) FirstIDX(ctx context.Context) int {
	id, err := mpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MatchPlayer entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MatchPlayer entity is found.
// Returns a *NotFoundError when no MatchPlayer entities are found.
func (mpq *MatchPlayerQuery) Only(ctx context.Context) (*MatchPlayer, error) {
	nodes, err := mpq.Limit(2).All(setContextOp(ctx, mpq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{matchplayer.Label}
	default:
		return nil, &NotSingularError{matchplayer.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mpq *MatchPlayerQuery) OnlyX(ctx context.Context) *MatchPlayer {
	node, err := mpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MatchPlayer ID in the query.
// Returns a *NotSingularError when more than one MatchPlayer ID is found.
// Returns a *NotFoundError when no entities are found.
func (mpq *MatchPlayerQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mpq.Limit(2).IDs(setContextOp(ctx, mpq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{matchplayer.Label}
	default:
		err = &NotSingularError{matchplayer.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mpq *MatchPlayerQuery) OnlyIDX(ctx context.Context) int {
	id, err := mpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MatchPlayers.
func (mpq *MatchPlayerQuery) All(ctx context.Context) ([]*MatchPlayer, error) {
	ctx = setContextOp(ctx, mpq.ctx, ent.OpQueryAll)
	if err := mpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*MatchPlayer, *MatchPlayerQuery]()
	return withInterceptors[[]*MatchPlayer](ctx, mpq, qr, mpq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mpq *MatchPlayerQuery) AllX(ctx context.Context) []*MatchPlayer {
	nodes, err := mpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MatchPlayer IDs.
func (mpq *MatchPlayerQuery) IDs(ctx context.Context) (ids []int, err error) {
	if mpq.ctx.Unique == nil && mpq.path != nil {
		mpq.Unique(true)
	}
	ctx = setContextOp(ctx, mpq.ctx, ent.OpQueryIDs)
	if err = mpq.Select(matchplayer.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mpq *MatchPlayerQuery) IDsX(ctx context.Context) []int {
	ids, err := mpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mpq *MatchPlayerQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mpq.ctx, ent.OpQueryCount)
	if err := mpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mpq, querierCount[*MatchPlayerQuery](), mpq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mpq *MatchPlayerQuery) CountX(ctx context.Context) int {
	count, err := mpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mpq *MatchPlayerQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mpq.ctx, ent.OpQueryExist)
	switch _, err := mpq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mpq *MatchPlayerQuery) ExistX(ctx context.Context) bool {
	exist, err := mpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MatchPlayerQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mpq *MatchPlayerQuery) Clone() *MatchPlayerQuery {
	if mpq == nil {
		return nil
	}
	return &MatchPlayerQuery{
		config:     mpq.config,
		ctx:        mpq.ctx.Clone(),
		order:      append([]matchplayer.OrderOption{}, mpq.order...),
		inters:     append([]Interceptor{}, mpq.inters...),
		predicates: append([]predicate.MatchPlayer{}, mpq.predicates...),
		withMatch:  mpq.withMatch.Clone(),
		withPlayer: mpq.withPlayer.Clone(),
		// clone intermediate query.
		sql:  mpq.sql.Clone(),
		path: mpq.path,
	}
}

// WithMatch tells the query-builder to eager-load the nodes that are connected to
// the "match" edge. The optional arguments are used to configure the query builder of the edge.
func (mpq *MatchPlayerQuery) WithMatch(opts ...func(*MatchQuery)) *MatchPlayerQuery {
	query := (&MatchClient{config: mpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mpq.withMatch = query
	return mpq
}

// WithPlayer tells the query-builder to eager-load the nodes that are connected to
// the "player" edge. The optional arguments are used to configure the query builder of the edge.
func (mpq *MatchPlayerQuery) WithPlayer(opts ...func(*PlayerQuery)) *MatchPlayerQuery {
	query := (&PlayerClient{config: mpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mpq.withPlayer = query
	return mpq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		MatchID int `json:"match_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MatchPlayer.Query().
//		GroupBy(matchplayer.FieldMatchID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mpq *MatchPlayerQuery) GroupBy(field string, fields ...string) *MatchPlayerGroupBy {
	mpq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MatchPlayerGroupBy{build: mpq}
	grbuild.flds = &mpq.ctx.Fields
	grbuild.label = matchplayer.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		MatchID int `json:"match_id,omitempty"`
//	}
//
//	client.MatchPlayer.Query().
//		Select(matchplayer.FieldMatchID).
//		Scan(ctx, &v)
func (mpq *MatchPlayerQuery) Select(fields ...string) *MatchPlayerSelect {
	mpq.ctx.Fields = append(mpq.ctx.Fields, fields...)
	sbuild := &MatchPlayerSelect{MatchPlayerQuery: mpq}
	sbuild.label = matchplayer.Label
	sbuild.flds, sbuild.scan = &mpq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MatchPlayerSelect configured with the given aggregations.
func (mpq *MatchPlayerQuery) Aggregate(fns ...AggregateFunc) *MatchPlayerSelect {
	return mpq.Select().Aggregate(fns...)
}

func (mpq *MatchPlayerQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mpq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mpq); err != nil {
				return err
			}
		}
	}
	for _, f := range mpq.ctx.Fields {
		if !matchplayer.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mpq.path != nil {
		prev, err := mpq.path(ctx)
		if err != nil {
			return err
		}
		mpq.sql = prev
	}
	return nil
}

func (mpq *MatchPlayerQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MatchPlayer, error) {
	var (
		nodes       = []*MatchPlayer{}
		_spec       = mpq.querySpec()
		loadedTypes = [2]bool{
			mpq.withMatch != nil,
			mpq.withPlayer != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MatchPlayer).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MatchPlayer{config: mpq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mpq.withMatch; query != nil {
		if err := mpq.loadMatch(ctx, query, nodes, nil,
			func(n *MatchPlayer, e *Match) { n.Edges.Match = e }); err != nil {
			return nil, err
		}
	}
	if query := mpq.withPlayer; query != nil {
		if err := mpq.loadPlayer(ctx, query, nodes, nil,
			func(n *MatchPlayer, e *Player) { n.Edges.Player = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mpq *MatchPlayerQuery) loadMatch(ctx context.Context, query *MatchQuery, nodes []*MatchPlayer, init func(*MatchPlayer), assign func(*MatchPlayer, *Match)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*MatchPlayer)
	for i := range nodes {
		fk := nodes[i].MatchID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(match.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "match_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (mpq *MatchPlayerQuery) loadPlayer(ctx context.Context, query *PlayerQuery, nodes []*MatchPlayer, init func(*MatchPlayer), assign func(*MatchPlayer, *Player)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*MatchPlayer)
	for i := range nodes {
		fk := nodes[i].PlayerID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(player.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "player_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (mpq *MatchPlayerQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mpq.querySpec()
	_spec.Node.Columns = mpq.ctx.Fields
	if len(mpq.ctx.Fields) > 0 {
		_spec.Unique = mpq.ctx.Unique != nil && *mpq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mpq.driver, _spec)
}

func (mpq *MatchPlayerQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(matchplayer.Table, matchplayer.Columns, sqlgraph.NewFieldSpec(matchplayer.FieldID, field.TypeInt))
	_spec.From = mpq.sql
	if unique := mpq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mpq.path != nil {
		_spec.Unique = true
	}
	if fields := mpq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, matchplayer.FieldID)
		for i := range fields {
			if fields[i] != matchplayer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if mpq.withMatch != nil {
			_spec.Node.AddColumnOnce(matchplayer.FieldMatchID)
		}
		if mpq.withPlayer != nil {
			_spec.Node.AddColumnOnce(matchplayer.FieldPlayerID)
		}
	}
	if ps := mpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mpq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mpq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mpq *MatchPlayerQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mpq.driver.Dialect())
	t1 := builder.Table(matchplayer.Table)
	columns := mpq.ctx.Fields
	if len(columns) == 0 {
		columns = matchplayer.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mpq.sql != nil {
		selector = mpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mpq.ctx.Unique != nil && *mpq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range mpq.predicates {
		p(selector)
	}
	for _, p := range mpq.order {
		p(selector)
	}
	if offset := mpq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mpq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MatchPlayerGroupBy is the group-by builder for MatchPlayer entities.
type MatchPlayerGroupBy struct {
	selector
	build *MatchPlayerQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mpgb *MatchPlayerGroupBy) Aggregate(fns ...AggregateFunc) *MatchPlayerGroupBy {
	mpgb.fns = append(mpgb.fns, fns...)
	return mpgb
}

// Scan applies the selector query and scans the result into the given value.
func (mpgb *MatchPlayerGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mpgb.build.ctx, ent.OpQueryGroupBy)
	if err := mpgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MatchPlayerQuery, *MatchPlayerGroupBy](ctx, mpgb.build, mpgb, mpgb.build.inters, v)
}

func (mpgb *MatchPlayerGroupBy) sqlScan(ctx context.Context, root *MatchPlayerQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mpgb.fns))
	for _, fn := range mpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mpgb.flds)+len(mpgb.fns))
		for _, f := range *mpgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mpgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mpgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MatchPlayerSelect is the builder for selecting fields of MatchPlayer entities.
type MatchPlayerSelect struct {
	*MatchPlayerQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mps *MatchPlayerSelect) Aggregate(fns ...AggregateFunc) *MatchPlayerSelect {
	mps.fns = append(mps.fns, fns...)
	return mps
}

// Scan applies the selector query and scans the result into the given value.
func (mps *MatchPlayerSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mps.ctx, ent.OpQuerySelect)
	if err := mps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MatchPlayerQuery, *MatchPlayerSelect](ctx, mps.MatchPlayerQuery, mps, mps.inters, v)
}

func (mps *MatchPlayerSelect) sqlScan(ctx context.Context, root *MatchPlayerQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(mps.fns))
	for _, fn := range mps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*mps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
