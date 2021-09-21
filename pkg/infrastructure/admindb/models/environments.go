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

// Environment is an object representing the database table.
type Environment struct {
	ID       string `boil:"id" json:"id" toml:"id" yaml:"id"`
	BranchID string `boil:"branch_id" json:"branch_id" toml:"branch_id" yaml:"branch_id"`
	Key      string `boil:"key" json:"key" toml:"key" yaml:"key"`
	Value    string `boil:"value" json:"value" toml:"value" yaml:"value"`

	R *environmentR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L environmentL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EnvironmentColumns = struct {
	ID       string
	BranchID string
	Key      string
	Value    string
}{
	ID:       "id",
	BranchID: "branch_id",
	Key:      "key",
	Value:    "value",
}

// Generated where

var EnvironmentWhere = struct {
	ID       whereHelperstring
	BranchID whereHelperstring
	Key      whereHelperstring
	Value    whereHelperstring
}{
	ID:       whereHelperstring{field: "`environments`.`id`"},
	BranchID: whereHelperstring{field: "`environments`.`branch_id`"},
	Key:      whereHelperstring{field: "`environments`.`key`"},
	Value:    whereHelperstring{field: "`environments`.`value`"},
}

// EnvironmentRels is where relationship names are stored.
var EnvironmentRels = struct {
	Branch string
}{
	Branch: "Branch",
}

// environmentR is where relationships are stored.
type environmentR struct {
	Branch *Branch `boil:"Branch" json:"Branch" toml:"Branch" yaml:"Branch"`
}

// NewStruct creates a new relationship struct
func (*environmentR) NewStruct() *environmentR {
	return &environmentR{}
}

// environmentL is where Load methods for each relationship are stored.
type environmentL struct{}

var (
	environmentAllColumns            = []string{"id", "branch_id", "key", "value"}
	environmentColumnsWithoutDefault = []string{"id", "branch_id", "key", "value"}
	environmentColumnsWithDefault    = []string{}
	environmentPrimaryKeyColumns     = []string{"id"}
)

type (
	// EnvironmentSlice is an alias for a slice of pointers to Environment.
	// This should generally be used opposed to []Environment.
	EnvironmentSlice []*Environment
	// EnvironmentHook is the signature for custom Environment hook methods
	EnvironmentHook func(context.Context, boil.ContextExecutor, *Environment) error

	environmentQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	environmentType                 = reflect.TypeOf(&Environment{})
	environmentMapping              = queries.MakeStructMapping(environmentType)
	environmentPrimaryKeyMapping, _ = queries.BindMapping(environmentType, environmentMapping, environmentPrimaryKeyColumns)
	environmentInsertCacheMut       sync.RWMutex
	environmentInsertCache          = make(map[string]insertCache)
	environmentUpdateCacheMut       sync.RWMutex
	environmentUpdateCache          = make(map[string]updateCache)
	environmentUpsertCacheMut       sync.RWMutex
	environmentUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var environmentBeforeInsertHooks []EnvironmentHook
var environmentBeforeUpdateHooks []EnvironmentHook
var environmentBeforeDeleteHooks []EnvironmentHook
var environmentBeforeUpsertHooks []EnvironmentHook

var environmentAfterInsertHooks []EnvironmentHook
var environmentAfterSelectHooks []EnvironmentHook
var environmentAfterUpdateHooks []EnvironmentHook
var environmentAfterDeleteHooks []EnvironmentHook
var environmentAfterUpsertHooks []EnvironmentHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Environment) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range environmentBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Environment) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range environmentBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Environment) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range environmentBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Environment) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range environmentBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Environment) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range environmentAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Environment) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range environmentAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Environment) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range environmentAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Environment) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range environmentAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Environment) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range environmentAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddEnvironmentHook registers your hook function for all future operations.
func AddEnvironmentHook(hookPoint boil.HookPoint, environmentHook EnvironmentHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		environmentBeforeInsertHooks = append(environmentBeforeInsertHooks, environmentHook)
	case boil.BeforeUpdateHook:
		environmentBeforeUpdateHooks = append(environmentBeforeUpdateHooks, environmentHook)
	case boil.BeforeDeleteHook:
		environmentBeforeDeleteHooks = append(environmentBeforeDeleteHooks, environmentHook)
	case boil.BeforeUpsertHook:
		environmentBeforeUpsertHooks = append(environmentBeforeUpsertHooks, environmentHook)
	case boil.AfterInsertHook:
		environmentAfterInsertHooks = append(environmentAfterInsertHooks, environmentHook)
	case boil.AfterSelectHook:
		environmentAfterSelectHooks = append(environmentAfterSelectHooks, environmentHook)
	case boil.AfterUpdateHook:
		environmentAfterUpdateHooks = append(environmentAfterUpdateHooks, environmentHook)
	case boil.AfterDeleteHook:
		environmentAfterDeleteHooks = append(environmentAfterDeleteHooks, environmentHook)
	case boil.AfterUpsertHook:
		environmentAfterUpsertHooks = append(environmentAfterUpsertHooks, environmentHook)
	}
}

// One returns a single environment record from the query.
func (q environmentQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Environment, error) {
	o := &Environment{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for environments")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Environment records from the query.
func (q environmentQuery) All(ctx context.Context, exec boil.ContextExecutor) (EnvironmentSlice, error) {
	var o []*Environment

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Environment slice")
	}

	if len(environmentAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Environment records in the query.
func (q environmentQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count environments rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q environmentQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if environments exists")
	}

	return count > 0, nil
}

// Branch pointed to by the foreign key.
func (o *Environment) Branch(mods ...qm.QueryMod) branchQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.BranchID),
	}

	queryMods = append(queryMods, mods...)

	query := Branches(queryMods...)
	queries.SetFrom(query.Query, "`branches`")

	return query
}

// LoadBranch allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (environmentL) LoadBranch(ctx context.Context, e boil.ContextExecutor, singular bool, maybeEnvironment interface{}, mods queries.Applicator) error {
	var slice []*Environment
	var object *Environment

	if singular {
		object = maybeEnvironment.(*Environment)
	} else {
		slice = *maybeEnvironment.(*[]*Environment)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &environmentR{}
		}
		args = append(args, object.BranchID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &environmentR{}
			}

			for _, a := range args {
				if a == obj.BranchID {
					continue Outer
				}
			}

			args = append(args, obj.BranchID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`branches`),
		qm.WhereIn(`branches.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Branch")
	}

	var resultSlice []*Branch
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Branch")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for branches")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for branches")
	}

	if len(environmentAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Branch = foreign
		if foreign.R == nil {
			foreign.R = &branchR{}
		}
		foreign.R.Environments = append(foreign.R.Environments, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.BranchID == foreign.ID {
				local.R.Branch = foreign
				if foreign.R == nil {
					foreign.R = &branchR{}
				}
				foreign.R.Environments = append(foreign.R.Environments, local)
				break
			}
		}
	}

	return nil
}

// SetBranch of the environment to the related item.
// Sets o.R.Branch to related.
// Adds o to related.R.Environments.
func (o *Environment) SetBranch(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Branch) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `environments` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"branch_id"}),
		strmangle.WhereClause("`", "`", 0, environmentPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.BranchID = related.ID
	if o.R == nil {
		o.R = &environmentR{
			Branch: related,
		}
	} else {
		o.R.Branch = related
	}

	if related.R == nil {
		related.R = &branchR{
			Environments: EnvironmentSlice{o},
		}
	} else {
		related.R.Environments = append(related.R.Environments, o)
	}

	return nil
}

// Environments retrieves all the records using an executor.
func Environments(mods ...qm.QueryMod) environmentQuery {
	mods = append(mods, qm.From("`environments`"))
	return environmentQuery{NewQuery(mods...)}
}

// FindEnvironment retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEnvironment(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Environment, error) {
	environmentObj := &Environment{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `environments` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, environmentObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from environments")
	}

	return environmentObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Environment) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no environments provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(environmentColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	environmentInsertCacheMut.RLock()
	cache, cached := environmentInsertCache[key]
	environmentInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			environmentAllColumns,
			environmentColumnsWithDefault,
			environmentColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(environmentType, environmentMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(environmentType, environmentMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `environments` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `environments` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `environments` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, environmentPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into environments")
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
		return errors.Wrap(err, "models: unable to populate default values for environments")
	}

CacheNoHooks:
	if !cached {
		environmentInsertCacheMut.Lock()
		environmentInsertCache[key] = cache
		environmentInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Environment.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Environment) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	environmentUpdateCacheMut.RLock()
	cache, cached := environmentUpdateCache[key]
	environmentUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			environmentAllColumns,
			environmentPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update environments, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `environments` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, environmentPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(environmentType, environmentMapping, append(wl, environmentPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update environments row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for environments")
	}

	if !cached {
		environmentUpdateCacheMut.Lock()
		environmentUpdateCache[key] = cache
		environmentUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q environmentQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for environments")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for environments")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EnvironmentSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), environmentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `environments` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, environmentPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in environment slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all environment")
	}
	return rowsAff, nil
}

var mySQLEnvironmentUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Environment) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no environments provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(environmentColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLEnvironmentUniqueColumns, o)

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

	environmentUpsertCacheMut.RLock()
	cache, cached := environmentUpsertCache[key]
	environmentUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			environmentAllColumns,
			environmentColumnsWithDefault,
			environmentColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			environmentAllColumns,
			environmentPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert environments, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`environments`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `environments` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(environmentType, environmentMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(environmentType, environmentMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for environments")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(environmentType, environmentMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for environments")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for environments")
	}

CacheNoHooks:
	if !cached {
		environmentUpsertCacheMut.Lock()
		environmentUpsertCache[key] = cache
		environmentUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Environment record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Environment) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Environment provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), environmentPrimaryKeyMapping)
	sql := "DELETE FROM `environments` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from environments")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for environments")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q environmentQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no environmentQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from environments")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for environments")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EnvironmentSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(environmentBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), environmentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `environments` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, environmentPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from environment slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for environments")
	}

	if len(environmentAfterDeleteHooks) != 0 {
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
func (o *Environment) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEnvironment(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EnvironmentSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EnvironmentSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), environmentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `environments`.* FROM `environments` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, environmentPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in EnvironmentSlice")
	}

	*o = slice

	return nil
}

// EnvironmentExists checks if the Environment row exists.
func EnvironmentExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `environments` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if environments exists")
	}

	return exists, nil
}
