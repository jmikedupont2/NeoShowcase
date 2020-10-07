// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testDynamicSiteDetails(t *testing.T) {
	t.Parallel()

	query := DynamicSiteDetails()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDynamicSiteDetailsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicSiteDetail{}
	if err = randomize.Struct(seed, o, dynamicSiteDetailDBTypes, true, dynamicSiteDetailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicSiteDetail struct: %s", err)
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

	count, err := DynamicSiteDetails().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDynamicSiteDetailsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicSiteDetail{}
	if err = randomize.Struct(seed, o, dynamicSiteDetailDBTypes, true, dynamicSiteDetailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicSiteDetail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := DynamicSiteDetails().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DynamicSiteDetails().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDynamicSiteDetailsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicSiteDetail{}
	if err = randomize.Struct(seed, o, dynamicSiteDetailDBTypes, true, dynamicSiteDetailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicSiteDetail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DynamicSiteDetailSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DynamicSiteDetails().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDynamicSiteDetailsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicSiteDetail{}
	if err = randomize.Struct(seed, o, dynamicSiteDetailDBTypes, true, dynamicSiteDetailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicSiteDetail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DynamicSiteDetailExists(ctx, tx, o.SiteID)
	if err != nil {
		t.Errorf("Unable to check if DynamicSiteDetail exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DynamicSiteDetailExists to return true, but got false.")
	}
}

func testDynamicSiteDetailsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicSiteDetail{}
	if err = randomize.Struct(seed, o, dynamicSiteDetailDBTypes, true, dynamicSiteDetailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicSiteDetail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	dynamicSiteDetailFound, err := FindDynamicSiteDetail(ctx, tx, o.SiteID)
	if err != nil {
		t.Error(err)
	}

	if dynamicSiteDetailFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDynamicSiteDetailsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicSiteDetail{}
	if err = randomize.Struct(seed, o, dynamicSiteDetailDBTypes, true, dynamicSiteDetailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicSiteDetail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = DynamicSiteDetails().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDynamicSiteDetailsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicSiteDetail{}
	if err = randomize.Struct(seed, o, dynamicSiteDetailDBTypes, true, dynamicSiteDetailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicSiteDetail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := DynamicSiteDetails().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDynamicSiteDetailsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dynamicSiteDetailOne := &DynamicSiteDetail{}
	dynamicSiteDetailTwo := &DynamicSiteDetail{}
	if err = randomize.Struct(seed, dynamicSiteDetailOne, dynamicSiteDetailDBTypes, false, dynamicSiteDetailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicSiteDetail struct: %s", err)
	}
	if err = randomize.Struct(seed, dynamicSiteDetailTwo, dynamicSiteDetailDBTypes, false, dynamicSiteDetailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicSiteDetail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = dynamicSiteDetailOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = dynamicSiteDetailTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DynamicSiteDetails().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDynamicSiteDetailsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	dynamicSiteDetailOne := &DynamicSiteDetail{}
	dynamicSiteDetailTwo := &DynamicSiteDetail{}
	if err = randomize.Struct(seed, dynamicSiteDetailOne, dynamicSiteDetailDBTypes, false, dynamicSiteDetailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicSiteDetail struct: %s", err)
	}
	if err = randomize.Struct(seed, dynamicSiteDetailTwo, dynamicSiteDetailDBTypes, false, dynamicSiteDetailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicSiteDetail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = dynamicSiteDetailOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = dynamicSiteDetailTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DynamicSiteDetails().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func dynamicSiteDetailBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *DynamicSiteDetail) error {
	*o = DynamicSiteDetail{}
	return nil
}

func dynamicSiteDetailAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *DynamicSiteDetail) error {
	*o = DynamicSiteDetail{}
	return nil
}

func dynamicSiteDetailAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *DynamicSiteDetail) error {
	*o = DynamicSiteDetail{}
	return nil
}

func dynamicSiteDetailBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DynamicSiteDetail) error {
	*o = DynamicSiteDetail{}
	return nil
}

func dynamicSiteDetailAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DynamicSiteDetail) error {
	*o = DynamicSiteDetail{}
	return nil
}

func dynamicSiteDetailBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DynamicSiteDetail) error {
	*o = DynamicSiteDetail{}
	return nil
}

func dynamicSiteDetailAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DynamicSiteDetail) error {
	*o = DynamicSiteDetail{}
	return nil
}

func dynamicSiteDetailBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DynamicSiteDetail) error {
	*o = DynamicSiteDetail{}
	return nil
}

func dynamicSiteDetailAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DynamicSiteDetail) error {
	*o = DynamicSiteDetail{}
	return nil
}

func testDynamicSiteDetailsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &DynamicSiteDetail{}
	o := &DynamicSiteDetail{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, dynamicSiteDetailDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DynamicSiteDetail object: %s", err)
	}

	AddDynamicSiteDetailHook(boil.BeforeInsertHook, dynamicSiteDetailBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	dynamicSiteDetailBeforeInsertHooks = []DynamicSiteDetailHook{}

	AddDynamicSiteDetailHook(boil.AfterInsertHook, dynamicSiteDetailAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	dynamicSiteDetailAfterInsertHooks = []DynamicSiteDetailHook{}

	AddDynamicSiteDetailHook(boil.AfterSelectHook, dynamicSiteDetailAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	dynamicSiteDetailAfterSelectHooks = []DynamicSiteDetailHook{}

	AddDynamicSiteDetailHook(boil.BeforeUpdateHook, dynamicSiteDetailBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	dynamicSiteDetailBeforeUpdateHooks = []DynamicSiteDetailHook{}

	AddDynamicSiteDetailHook(boil.AfterUpdateHook, dynamicSiteDetailAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	dynamicSiteDetailAfterUpdateHooks = []DynamicSiteDetailHook{}

	AddDynamicSiteDetailHook(boil.BeforeDeleteHook, dynamicSiteDetailBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	dynamicSiteDetailBeforeDeleteHooks = []DynamicSiteDetailHook{}

	AddDynamicSiteDetailHook(boil.AfterDeleteHook, dynamicSiteDetailAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	dynamicSiteDetailAfterDeleteHooks = []DynamicSiteDetailHook{}

	AddDynamicSiteDetailHook(boil.BeforeUpsertHook, dynamicSiteDetailBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	dynamicSiteDetailBeforeUpsertHooks = []DynamicSiteDetailHook{}

	AddDynamicSiteDetailHook(boil.AfterUpsertHook, dynamicSiteDetailAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	dynamicSiteDetailAfterUpsertHooks = []DynamicSiteDetailHook{}
}

func testDynamicSiteDetailsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicSiteDetail{}
	if err = randomize.Struct(seed, o, dynamicSiteDetailDBTypes, true, dynamicSiteDetailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicSiteDetail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DynamicSiteDetails().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDynamicSiteDetailsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicSiteDetail{}
	if err = randomize.Struct(seed, o, dynamicSiteDetailDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DynamicSiteDetail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(dynamicSiteDetailColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := DynamicSiteDetails().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDynamicSiteDetailToOneSiteUsingSite(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local DynamicSiteDetail
	var foreign Site

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, dynamicSiteDetailDBTypes, false, dynamicSiteDetailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicSiteDetail struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, siteDBTypes, false, siteColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Site struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.SiteID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Site().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := DynamicSiteDetailSlice{&local}
	if err = local.L.LoadSite(ctx, tx, false, (*[]*DynamicSiteDetail)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Site == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Site = nil
	if err = local.L.LoadSite(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Site == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testDynamicSiteDetailToOneSetOpSiteUsingSite(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DynamicSiteDetail
	var b, c Site

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dynamicSiteDetailDBTypes, false, strmangle.SetComplement(dynamicSiteDetailPrimaryKeyColumns, dynamicSiteDetailColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, siteDBTypes, false, strmangle.SetComplement(sitePrimaryKeyColumns, siteColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, siteDBTypes, false, strmangle.SetComplement(sitePrimaryKeyColumns, siteColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Site{&b, &c} {
		err = a.SetSite(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Site != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.DynamicSiteDetail != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.SiteID != x.ID {
			t.Error("foreign key was wrong value", a.SiteID)
		}

		if exists, err := DynamicSiteDetailExists(ctx, tx, a.SiteID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testDynamicSiteDetailsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicSiteDetail{}
	if err = randomize.Struct(seed, o, dynamicSiteDetailDBTypes, true, dynamicSiteDetailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicSiteDetail struct: %s", err)
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

func testDynamicSiteDetailsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicSiteDetail{}
	if err = randomize.Struct(seed, o, dynamicSiteDetailDBTypes, true, dynamicSiteDetailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicSiteDetail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DynamicSiteDetailSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDynamicSiteDetailsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicSiteDetail{}
	if err = randomize.Struct(seed, o, dynamicSiteDetailDBTypes, true, dynamicSiteDetailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicSiteDetail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DynamicSiteDetails().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	dynamicSiteDetailDBTypes = map[string]string{`SiteID`: `varchar`}
	_                        = bytes.MinRead
)

func testDynamicSiteDetailsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(dynamicSiteDetailPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(dynamicSiteDetailAllColumns) == len(dynamicSiteDetailPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DynamicSiteDetail{}
	if err = randomize.Struct(seed, o, dynamicSiteDetailDBTypes, true, dynamicSiteDetailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicSiteDetail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DynamicSiteDetails().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, dynamicSiteDetailDBTypes, true, dynamicSiteDetailPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DynamicSiteDetail struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDynamicSiteDetailsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(dynamicSiteDetailAllColumns) == len(dynamicSiteDetailPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DynamicSiteDetail{}
	if err = randomize.Struct(seed, o, dynamicSiteDetailDBTypes, true, dynamicSiteDetailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicSiteDetail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DynamicSiteDetails().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, dynamicSiteDetailDBTypes, true, dynamicSiteDetailPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DynamicSiteDetail struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(dynamicSiteDetailAllColumns, dynamicSiteDetailPrimaryKeyColumns) {
		fields = dynamicSiteDetailAllColumns
	} else {
		fields = strmangle.SetComplement(
			dynamicSiteDetailAllColumns,
			dynamicSiteDetailPrimaryKeyColumns,
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

	slice := DynamicSiteDetailSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDynamicSiteDetailsUpsert(t *testing.T) {
	t.Parallel()

	if len(dynamicSiteDetailAllColumns) == len(dynamicSiteDetailPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLDynamicSiteDetailUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := DynamicSiteDetail{}
	if err = randomize.Struct(seed, &o, dynamicSiteDetailDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DynamicSiteDetail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DynamicSiteDetail: %s", err)
	}

	count, err := DynamicSiteDetails().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, dynamicSiteDetailDBTypes, false, dynamicSiteDetailPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DynamicSiteDetail struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DynamicSiteDetail: %s", err)
	}

	count, err = DynamicSiteDetails().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
