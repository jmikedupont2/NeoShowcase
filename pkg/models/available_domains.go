// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// AvailableDomain is an object representing the database table.
type AvailableDomain struct {
	ID        string `boil:"id" json:"id" toml:"id" yaml:"id"`
	Domain    string `boil:"domain" json:"domain" toml:"domain" yaml:"domain"`
	Subdomain bool   `boil:"subdomain" json:"subdomain" toml:"subdomain" yaml:"subdomain"`

	R *availableDomainR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L availableDomainL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AvailableDomainColumns = struct {
	ID        string
	Domain    string
	Subdomain string
}{
	ID:        "id",
	Domain:    "domain",
	Subdomain: "subdomain",
}

// Generated where

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var AvailableDomainWhere = struct {
	ID        whereHelperstring
	Domain    whereHelperstring
	Subdomain whereHelperbool
}{
	ID:        whereHelperstring{field: "`available_domains`.`id`"},
	Domain:    whereHelperstring{field: "`available_domains`.`domain`"},
	Subdomain: whereHelperbool{field: "`available_domains`.`subdomain`"},
}

// AvailableDomainRels is where relationship names are stored.
var AvailableDomainRels = struct {
}{}

// availableDomainR is where relationships are stored.
type availableDomainR struct {
}

// NewStruct creates a new relationship struct
func (*availableDomainR) NewStruct() *availableDomainR {
	return &availableDomainR{}
}

// availableDomainL is where Load methods for each relationship are stored.
type availableDomainL struct{}

var (
	availableDomainAllColumns            = []string{"id", "domain", "subdomain"}
	availableDomainColumnsWithoutDefault = []string{"id", "domain"}
	availableDomainColumnsWithDefault    = []string{"subdomain"}
	availableDomainPrimaryKeyColumns     = []string{"id"}
)

type (
	// AvailableDomainSlice is an alias for a slice of pointers to AvailableDomain.
	// This should generally be used opposed to []AvailableDomain.
	AvailableDomainSlice []*AvailableDomain
	// AvailableDomainHook is the signature for custom AvailableDomain hook methods
	AvailableDomainHook func(context.Context, boil.ContextExecutor, *AvailableDomain) error

	availableDomainQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	availableDomainType                 = reflect.TypeOf(&AvailableDomain{})
	availableDomainMapping              = queries.MakeStructMapping(availableDomainType)
	availableDomainPrimaryKeyMapping, _ = queries.BindMapping(availableDomainType, availableDomainMapping, availableDomainPrimaryKeyColumns)
	availableDomainInsertCacheMut       sync.RWMutex
	availableDomainInsertCache          = make(map[string]insertCache)
	availableDomainUpdateCacheMut       sync.RWMutex
	availableDomainUpdateCache          = make(map[string]updateCache)
	availableDomainUpsertCacheMut       sync.RWMutex
	availableDomainUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var availableDomainBeforeInsertHooks []AvailableDomainHook
var availableDomainBeforeUpdateHooks []AvailableDomainHook
var availableDomainBeforeDeleteHooks []AvailableDomainHook
var availableDomainBeforeUpsertHooks []AvailableDomainHook

var availableDomainAfterInsertHooks []AvailableDomainHook
var availableDomainAfterSelectHooks []AvailableDomainHook
var availableDomainAfterUpdateHooks []AvailableDomainHook
var availableDomainAfterDeleteHooks []AvailableDomainHook
var availableDomainAfterUpsertHooks []AvailableDomainHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AvailableDomain) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range availableDomainBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AvailableDomain) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range availableDomainBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AvailableDomain) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range availableDomainBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AvailableDomain) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range availableDomainBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AvailableDomain) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range availableDomainAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AvailableDomain) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range availableDomainAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AvailableDomain) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range availableDomainAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AvailableDomain) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range availableDomainAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AvailableDomain) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range availableDomainAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAvailableDomainHook registers your hook function for all future operations.
func AddAvailableDomainHook(hookPoint boil.HookPoint, availableDomainHook AvailableDomainHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		availableDomainBeforeInsertHooks = append(availableDomainBeforeInsertHooks, availableDomainHook)
	case boil.BeforeUpdateHook:
		availableDomainBeforeUpdateHooks = append(availableDomainBeforeUpdateHooks, availableDomainHook)
	case boil.BeforeDeleteHook:
		availableDomainBeforeDeleteHooks = append(availableDomainBeforeDeleteHooks, availableDomainHook)
	case boil.BeforeUpsertHook:
		availableDomainBeforeUpsertHooks = append(availableDomainBeforeUpsertHooks, availableDomainHook)
	case boil.AfterInsertHook:
		availableDomainAfterInsertHooks = append(availableDomainAfterInsertHooks, availableDomainHook)
	case boil.AfterSelectHook:
		availableDomainAfterSelectHooks = append(availableDomainAfterSelectHooks, availableDomainHook)
	case boil.AfterUpdateHook:
		availableDomainAfterUpdateHooks = append(availableDomainAfterUpdateHooks, availableDomainHook)
	case boil.AfterDeleteHook:
		availableDomainAfterDeleteHooks = append(availableDomainAfterDeleteHooks, availableDomainHook)
	case boil.AfterUpsertHook:
		availableDomainAfterUpsertHooks = append(availableDomainAfterUpsertHooks, availableDomainHook)
	}
}

// One returns a single availableDomain record from the query.
func (q availableDomainQuery) One(ctx context.Context, exec boil.ContextExecutor) (*AvailableDomain, error) {
	o := &AvailableDomain{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for available_domains")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all AvailableDomain records from the query.
func (q availableDomainQuery) All(ctx context.Context, exec boil.ContextExecutor) (AvailableDomainSlice, error) {
	var o []*AvailableDomain

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AvailableDomain slice")
	}

	if len(availableDomainAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all AvailableDomain records in the query.
func (q availableDomainQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count available_domains rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q availableDomainQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if available_domains exists")
	}

	return count > 0, nil
}

// AvailableDomains retrieves all the records using an executor.
func AvailableDomains(mods ...qm.QueryMod) availableDomainQuery {
	mods = append(mods, qm.From("`available_domains`"))
	return availableDomainQuery{NewQuery(mods...)}
}

// FindAvailableDomain retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAvailableDomain(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*AvailableDomain, error) {
	availableDomainObj := &AvailableDomain{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `available_domains` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, availableDomainObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from available_domains")
	}

	return availableDomainObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AvailableDomain) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no available_domains provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(availableDomainColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	availableDomainInsertCacheMut.RLock()
	cache, cached := availableDomainInsertCache[key]
	availableDomainInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			availableDomainAllColumns,
			availableDomainColumnsWithDefault,
			availableDomainColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(availableDomainType, availableDomainMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(availableDomainType, availableDomainMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `available_domains` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `available_domains` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `available_domains` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, availableDomainPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into available_domains")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for available_domains")
	}

CacheNoHooks:
	if !cached {
		availableDomainInsertCacheMut.Lock()
		availableDomainInsertCache[key] = cache
		availableDomainInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the AvailableDomain.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AvailableDomain) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	availableDomainUpdateCacheMut.RLock()
	cache, cached := availableDomainUpdateCache[key]
	availableDomainUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			availableDomainAllColumns,
			availableDomainPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update available_domains, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `available_domains` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, availableDomainPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(availableDomainType, availableDomainMapping, append(wl, availableDomainPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update available_domains row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for available_domains")
	}

	if !cached {
		availableDomainUpdateCacheMut.Lock()
		availableDomainUpdateCache[key] = cache
		availableDomainUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q availableDomainQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for available_domains")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for available_domains")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AvailableDomainSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), availableDomainPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `available_domains` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, availableDomainPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in availableDomain slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all availableDomain")
	}
	return rowsAff, nil
}

var mySQLAvailableDomainUniqueColumns = []string{
	"id",
	"domain",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *AvailableDomain) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no available_domains provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(availableDomainColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLAvailableDomainUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	availableDomainUpsertCacheMut.RLock()
	cache, cached := availableDomainUpsertCache[key]
	availableDomainUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			availableDomainAllColumns,
			availableDomainColumnsWithDefault,
			availableDomainColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			availableDomainAllColumns,
			availableDomainPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert available_domains, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`available_domains`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `available_domains` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(availableDomainType, availableDomainMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(availableDomainType, availableDomainMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for available_domains")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(availableDomainType, availableDomainMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for available_domains")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for available_domains")
	}

CacheNoHooks:
	if !cached {
		availableDomainUpsertCacheMut.Lock()
		availableDomainUpsertCache[key] = cache
		availableDomainUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single AvailableDomain record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AvailableDomain) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no AvailableDomain provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), availableDomainPrimaryKeyMapping)
	sql := "DELETE FROM `available_domains` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from available_domains")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for available_domains")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q availableDomainQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no availableDomainQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from available_domains")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for available_domains")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AvailableDomainSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(availableDomainBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), availableDomainPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `available_domains` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, availableDomainPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from availableDomain slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for available_domains")
	}

	if len(availableDomainAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *AvailableDomain) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAvailableDomain(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AvailableDomainSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AvailableDomainSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), availableDomainPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `available_domains`.* FROM `available_domains` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, availableDomainPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AvailableDomainSlice")
	}

	*o = slice

	return nil
}

// AvailableDomainExists checks if the AvailableDomain row exists.
func AvailableDomainExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `available_domains` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if available_domains exists")
	}

	return exists, nil
}
