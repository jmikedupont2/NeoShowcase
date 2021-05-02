// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testWebsites(t *testing.T) {
	t.Parallel()

	query := Websites()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testWebsitesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Website{}
	if err = randomize.Struct(seed, o, websiteDBTypes, true, websiteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Website struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Websites().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testWebsitesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Website{}
	if err = randomize.Struct(seed, o, websiteDBTypes, true, websiteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Website struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Websites().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Websites().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testWebsitesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Website{}
	if err = randomize.Struct(seed, o, websiteDBTypes, true, websiteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Website struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := WebsiteSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Websites().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testWebsitesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Website{}
	if err = randomize.Struct(seed, o, websiteDBTypes, true, websiteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Website struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := WebsiteExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Website exists: %s", err)
	}
	if !e {
		t.Errorf("Expected WebsiteExists to return true, but got false.")
	}
}

func testWebsitesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Website{}
	if err = randomize.Struct(seed, o, websiteDBTypes, true, websiteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Website struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	websiteFound, err := FindWebsite(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if websiteFound == nil {
		t.Error("want a record, got nil")
	}
}

func testWebsitesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Website{}
	if err = randomize.Struct(seed, o, websiteDBTypes, true, websiteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Website struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Websites().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testWebsitesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Website{}
	if err = randomize.Struct(seed, o, websiteDBTypes, true, websiteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Website struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Websites().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testWebsitesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	websiteOne := &Website{}
	websiteTwo := &Website{}
	if err = randomize.Struct(seed, websiteOne, websiteDBTypes, false, websiteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Website struct: %s", err)
	}
	if err = randomize.Struct(seed, websiteTwo, websiteDBTypes, false, websiteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Website struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = websiteOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = websiteTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Websites().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testWebsitesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	websiteOne := &Website{}
	websiteTwo := &Website{}
	if err = randomize.Struct(seed, websiteOne, websiteDBTypes, false, websiteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Website struct: %s", err)
	}
	if err = randomize.Struct(seed, websiteTwo, websiteDBTypes, false, websiteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Website struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = websiteOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = websiteTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Websites().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func websiteBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Website) error {
	*o = Website{}
	return nil
}

func websiteAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Website) error {
	*o = Website{}
	return nil
}

func websiteAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Website) error {
	*o = Website{}
	return nil
}

func websiteBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Website) error {
	*o = Website{}
	return nil
}

func websiteAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Website) error {
	*o = Website{}
	return nil
}

func websiteBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Website) error {
	*o = Website{}
	return nil
}

func websiteAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Website) error {
	*o = Website{}
	return nil
}

func websiteBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Website) error {
	*o = Website{}
	return nil
}

func websiteAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Website) error {
	*o = Website{}
	return nil
}

func testWebsitesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Website{}
	o := &Website{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, websiteDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Website object: %s", err)
	}

	AddWebsiteHook(boil.BeforeInsertHook, websiteBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	websiteBeforeInsertHooks = []WebsiteHook{}

	AddWebsiteHook(boil.AfterInsertHook, websiteAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	websiteAfterInsertHooks = []WebsiteHook{}

	AddWebsiteHook(boil.AfterSelectHook, websiteAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	websiteAfterSelectHooks = []WebsiteHook{}

	AddWebsiteHook(boil.BeforeUpdateHook, websiteBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	websiteBeforeUpdateHooks = []WebsiteHook{}

	AddWebsiteHook(boil.AfterUpdateHook, websiteAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	websiteAfterUpdateHooks = []WebsiteHook{}

	AddWebsiteHook(boil.BeforeDeleteHook, websiteBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	websiteBeforeDeleteHooks = []WebsiteHook{}

	AddWebsiteHook(boil.AfterDeleteHook, websiteAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	websiteAfterDeleteHooks = []WebsiteHook{}

	AddWebsiteHook(boil.BeforeUpsertHook, websiteBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	websiteBeforeUpsertHooks = []WebsiteHook{}

	AddWebsiteHook(boil.AfterUpsertHook, websiteAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	websiteAfterUpsertHooks = []WebsiteHook{}
}

func testWebsitesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Website{}
	if err = randomize.Struct(seed, o, websiteDBTypes, true, websiteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Website struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Websites().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testWebsitesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Website{}
	if err = randomize.Struct(seed, o, websiteDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Website struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(websiteColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Websites().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testWebsiteToOneEnvironmentUsingEnvironment(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Website
	var foreign Environment

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, websiteDBTypes, false, websiteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Website struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, environmentDBTypes, false, environmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Environment struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.EnvironmentID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Environment().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := WebsiteSlice{&local}
	if err = local.L.LoadEnvironment(ctx, tx, false, (*[]*Website)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Environment == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Environment = nil
	if err = local.L.LoadEnvironment(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Environment == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testWebsiteToOneSetOpEnvironmentUsingEnvironment(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Website
	var b, c Environment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, websiteDBTypes, false, strmangle.SetComplement(websitePrimaryKeyColumns, websiteColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, environmentDBTypes, false, strmangle.SetComplement(environmentPrimaryKeyColumns, environmentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, environmentDBTypes, false, strmangle.SetComplement(environmentPrimaryKeyColumns, environmentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Environment{&b, &c} {
		err = a.SetEnvironment(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Environment != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Website != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.EnvironmentID != x.ID {
			t.Error("foreign key was wrong value", a.EnvironmentID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.EnvironmentID))
		reflect.Indirect(reflect.ValueOf(&a.EnvironmentID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.EnvironmentID != x.ID {
			t.Error("foreign key was wrong value", a.EnvironmentID, x.ID)
		}
	}
}

func testWebsitesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Website{}
	if err = randomize.Struct(seed, o, websiteDBTypes, true, websiteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Website struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testWebsitesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Website{}
	if err = randomize.Struct(seed, o, websiteDBTypes, true, websiteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Website struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := WebsiteSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testWebsitesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Website{}
	if err = randomize.Struct(seed, o, websiteDBTypes, true, websiteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Website struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Websites().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	websiteDBTypes = map[string]string{`ID`: `varchar`, `FQDN`: `varchar`, `HTTPPort`: `int`, `CreatedAt`: `datetime`, `UpdatedAt`: `datetime`, `EnvironmentID`: `varchar`}
	_              = bytes.MinRead
)

func testWebsitesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(websitePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(websiteAllColumns) == len(websitePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Website{}
	if err = randomize.Struct(seed, o, websiteDBTypes, true, websiteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Website struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Websites().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, websiteDBTypes, true, websitePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Website struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testWebsitesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(websiteAllColumns) == len(websitePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Website{}
	if err = randomize.Struct(seed, o, websiteDBTypes, true, websiteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Website struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Websites().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, websiteDBTypes, true, websitePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Website struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(websiteAllColumns, websitePrimaryKeyColumns) {
		fields = websiteAllColumns
	} else {
		fields = strmangle.SetComplement(
			websiteAllColumns,
			websitePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := WebsiteSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testWebsitesUpsert(t *testing.T) {
	t.Parallel()

	if len(websiteAllColumns) == len(websitePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLWebsiteUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Website{}
	if err = randomize.Struct(seed, &o, websiteDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Website struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Website: %s", err)
	}

	count, err := Websites().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, websiteDBTypes, false, websitePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Website struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Website: %s", err)
	}

	count, err = Websites().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
