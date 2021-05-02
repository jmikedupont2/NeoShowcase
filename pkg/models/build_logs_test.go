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

func testBuildLogs(t *testing.T) {
	t.Parallel()

	query := BuildLogs()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBuildLogsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildLog{}
	if err = randomize.Struct(seed, o, buildLogDBTypes, true, buildLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildLog struct: %s", err)
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

	count, err := BuildLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBuildLogsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildLog{}
	if err = randomize.Struct(seed, o, buildLogDBTypes, true, buildLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := BuildLogs().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BuildLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBuildLogsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildLog{}
	if err = randomize.Struct(seed, o, buildLogDBTypes, true, buildLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BuildLogSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BuildLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBuildLogsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildLog{}
	if err = randomize.Struct(seed, o, buildLogDBTypes, true, buildLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BuildLogExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if BuildLog exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BuildLogExists to return true, but got false.")
	}
}

func testBuildLogsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildLog{}
	if err = randomize.Struct(seed, o, buildLogDBTypes, true, buildLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	buildLogFound, err := FindBuildLog(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if buildLogFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBuildLogsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildLog{}
	if err = randomize.Struct(seed, o, buildLogDBTypes, true, buildLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = BuildLogs().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBuildLogsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildLog{}
	if err = randomize.Struct(seed, o, buildLogDBTypes, true, buildLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := BuildLogs().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBuildLogsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	buildLogOne := &BuildLog{}
	buildLogTwo := &BuildLog{}
	if err = randomize.Struct(seed, buildLogOne, buildLogDBTypes, false, buildLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildLog struct: %s", err)
	}
	if err = randomize.Struct(seed, buildLogTwo, buildLogDBTypes, false, buildLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = buildLogOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = buildLogTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BuildLogs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBuildLogsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	buildLogOne := &BuildLog{}
	buildLogTwo := &BuildLog{}
	if err = randomize.Struct(seed, buildLogOne, buildLogDBTypes, false, buildLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildLog struct: %s", err)
	}
	if err = randomize.Struct(seed, buildLogTwo, buildLogDBTypes, false, buildLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = buildLogOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = buildLogTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BuildLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func buildLogBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *BuildLog) error {
	*o = BuildLog{}
	return nil
}

func buildLogAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *BuildLog) error {
	*o = BuildLog{}
	return nil
}

func buildLogAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *BuildLog) error {
	*o = BuildLog{}
	return nil
}

func buildLogBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *BuildLog) error {
	*o = BuildLog{}
	return nil
}

func buildLogAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *BuildLog) error {
	*o = BuildLog{}
	return nil
}

func buildLogBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *BuildLog) error {
	*o = BuildLog{}
	return nil
}

func buildLogAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *BuildLog) error {
	*o = BuildLog{}
	return nil
}

func buildLogBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *BuildLog) error {
	*o = BuildLog{}
	return nil
}

func buildLogAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *BuildLog) error {
	*o = BuildLog{}
	return nil
}

func testBuildLogsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &BuildLog{}
	o := &BuildLog{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, buildLogDBTypes, false); err != nil {
		t.Errorf("Unable to randomize BuildLog object: %s", err)
	}

	AddBuildLogHook(boil.BeforeInsertHook, buildLogBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	buildLogBeforeInsertHooks = []BuildLogHook{}

	AddBuildLogHook(boil.AfterInsertHook, buildLogAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	buildLogAfterInsertHooks = []BuildLogHook{}

	AddBuildLogHook(boil.AfterSelectHook, buildLogAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	buildLogAfterSelectHooks = []BuildLogHook{}

	AddBuildLogHook(boil.BeforeUpdateHook, buildLogBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	buildLogBeforeUpdateHooks = []BuildLogHook{}

	AddBuildLogHook(boil.AfterUpdateHook, buildLogAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	buildLogAfterUpdateHooks = []BuildLogHook{}

	AddBuildLogHook(boil.BeforeDeleteHook, buildLogBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	buildLogBeforeDeleteHooks = []BuildLogHook{}

	AddBuildLogHook(boil.AfterDeleteHook, buildLogAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	buildLogAfterDeleteHooks = []BuildLogHook{}

	AddBuildLogHook(boil.BeforeUpsertHook, buildLogBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	buildLogBeforeUpsertHooks = []BuildLogHook{}

	AddBuildLogHook(boil.AfterUpsertHook, buildLogAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	buildLogAfterUpsertHooks = []BuildLogHook{}
}

func testBuildLogsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildLog{}
	if err = randomize.Struct(seed, o, buildLogDBTypes, true, buildLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BuildLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBuildLogsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildLog{}
	if err = randomize.Struct(seed, o, buildLogDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BuildLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(buildLogColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := BuildLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBuildLogOneToOneArtifactUsingArtifact(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var foreign Artifact
	var local BuildLog

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, artifactDBTypes, true, artifactColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Artifact struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, buildLogDBTypes, true, buildLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildLog struct: %s", err)
	}

	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreign.BuildLogID = local.ID
	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Artifact().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.BuildLogID != foreign.BuildLogID {
		t.Errorf("want: %v, got %v", foreign.BuildLogID, check.BuildLogID)
	}

	slice := BuildLogSlice{&local}
	if err = local.L.LoadArtifact(ctx, tx, false, (*[]*BuildLog)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Artifact == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Artifact = nil
	if err = local.L.LoadArtifact(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Artifact == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testBuildLogOneToOneSetOpArtifactUsingArtifact(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BuildLog
	var b, c Artifact

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, buildLogDBTypes, false, strmangle.SetComplement(buildLogPrimaryKeyColumns, buildLogColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, artifactDBTypes, false, strmangle.SetComplement(artifactPrimaryKeyColumns, artifactColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, artifactDBTypes, false, strmangle.SetComplement(artifactPrimaryKeyColumns, artifactColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Artifact{&b, &c} {
		err = a.SetArtifact(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Artifact != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.BuildLog != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.ID != x.BuildLogID {
			t.Error("foreign key was wrong value", a.ID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.BuildLogID))
		reflect.Indirect(reflect.ValueOf(&x.BuildLogID)).Set(zero)

		if err = x.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ID != x.BuildLogID {
			t.Error("foreign key was wrong value", a.ID, x.BuildLogID)
		}

		if _, err = x.Delete(ctx, tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}

func testBuildLogToManyBuildEnvironments(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BuildLog
	var b, c Environment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, buildLogDBTypes, true, buildLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildLog struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, environmentDBTypes, false, environmentColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, environmentDBTypes, false, environmentColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.BuildID, a.ID)
	queries.Assign(&c.BuildID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.BuildEnvironments().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.BuildID, b.BuildID) {
			bFound = true
		}
		if queries.Equal(v.BuildID, c.BuildID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BuildLogSlice{&a}
	if err = a.L.LoadBuildEnvironments(ctx, tx, false, (*[]*BuildLog)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.BuildEnvironments); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.BuildEnvironments = nil
	if err = a.L.LoadBuildEnvironments(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.BuildEnvironments); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testBuildLogToManyAddOpBuildEnvironments(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BuildLog
	var b, c, d, e Environment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, buildLogDBTypes, false, strmangle.SetComplement(buildLogPrimaryKeyColumns, buildLogColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Environment{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, environmentDBTypes, false, strmangle.SetComplement(environmentPrimaryKeyColumns, environmentColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Environment{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddBuildEnvironments(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.BuildID) {
			t.Error("foreign key was wrong value", a.ID, first.BuildID)
		}
		if !queries.Equal(a.ID, second.BuildID) {
			t.Error("foreign key was wrong value", a.ID, second.BuildID)
		}

		if first.R.Build != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Build != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.BuildEnvironments[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.BuildEnvironments[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.BuildEnvironments().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testBuildLogToManySetOpBuildEnvironments(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BuildLog
	var b, c, d, e Environment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, buildLogDBTypes, false, strmangle.SetComplement(buildLogPrimaryKeyColumns, buildLogColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Environment{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, environmentDBTypes, false, strmangle.SetComplement(environmentPrimaryKeyColumns, environmentColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetBuildEnvironments(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.BuildEnvironments().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetBuildEnvironments(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.BuildEnvironments().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.BuildID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.BuildID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.BuildID) {
		t.Error("foreign key was wrong value", a.ID, d.BuildID)
	}
	if !queries.Equal(a.ID, e.BuildID) {
		t.Error("foreign key was wrong value", a.ID, e.BuildID)
	}

	if b.R.Build != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Build != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Build != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Build != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.BuildEnvironments[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.BuildEnvironments[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testBuildLogToManyRemoveOpBuildEnvironments(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BuildLog
	var b, c, d, e Environment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, buildLogDBTypes, false, strmangle.SetComplement(buildLogPrimaryKeyColumns, buildLogColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Environment{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, environmentDBTypes, false, strmangle.SetComplement(environmentPrimaryKeyColumns, environmentColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddBuildEnvironments(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.BuildEnvironments().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveBuildEnvironments(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.BuildEnvironments().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.BuildID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.BuildID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Build != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Build != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Build != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Build != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.BuildEnvironments) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.BuildEnvironments[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.BuildEnvironments[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testBuildLogToOneEnvironmentUsingEnvironment(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local BuildLog
	var foreign Environment

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, buildLogDBTypes, true, buildLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildLog struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, environmentDBTypes, false, environmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Environment struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.EnvironmentID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Environment().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := BuildLogSlice{&local}
	if err = local.L.LoadEnvironment(ctx, tx, false, (*[]*BuildLog)(&slice), nil); err != nil {
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

func testBuildLogToOneSetOpEnvironmentUsingEnvironment(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BuildLog
	var b, c Environment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, buildLogDBTypes, false, strmangle.SetComplement(buildLogPrimaryKeyColumns, buildLogColumnsWithoutDefault)...); err != nil {
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

		if x.R.BuildLogs[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.EnvironmentID, x.ID) {
			t.Error("foreign key was wrong value", a.EnvironmentID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.EnvironmentID))
		reflect.Indirect(reflect.ValueOf(&a.EnvironmentID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.EnvironmentID, x.ID) {
			t.Error("foreign key was wrong value", a.EnvironmentID, x.ID)
		}
	}
}

func testBuildLogToOneRemoveOpEnvironmentUsingEnvironment(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BuildLog
	var b Environment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, buildLogDBTypes, false, strmangle.SetComplement(buildLogPrimaryKeyColumns, buildLogColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, environmentDBTypes, false, strmangle.SetComplement(environmentPrimaryKeyColumns, environmentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetEnvironment(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveEnvironment(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Environment().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Environment != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.EnvironmentID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.BuildLogs) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testBuildLogsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildLog{}
	if err = randomize.Struct(seed, o, buildLogDBTypes, true, buildLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildLog struct: %s", err)
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

func testBuildLogsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildLog{}
	if err = randomize.Struct(seed, o, buildLogDBTypes, true, buildLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BuildLogSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBuildLogsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildLog{}
	if err = randomize.Struct(seed, o, buildLogDBTypes, true, buildLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BuildLogs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	buildLogDBTypes = map[string]string{`ID`: `varchar`, `Result`: `enum('BUILDING','SUCCEEDED','FAILED','CANCELED')`, `StartedAt`: `datetime`, `FinishedAt`: `datetime`, `EnvironmentID`: `varchar`}
	_               = bytes.MinRead
)

func testBuildLogsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(buildLogPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(buildLogAllColumns) == len(buildLogPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BuildLog{}
	if err = randomize.Struct(seed, o, buildLogDBTypes, true, buildLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BuildLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, buildLogDBTypes, true, buildLogPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BuildLog struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBuildLogsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(buildLogAllColumns) == len(buildLogPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BuildLog{}
	if err = randomize.Struct(seed, o, buildLogDBTypes, true, buildLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BuildLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, buildLogDBTypes, true, buildLogPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BuildLog struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(buildLogAllColumns, buildLogPrimaryKeyColumns) {
		fields = buildLogAllColumns
	} else {
		fields = strmangle.SetComplement(
			buildLogAllColumns,
			buildLogPrimaryKeyColumns,
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

	slice := BuildLogSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBuildLogsUpsert(t *testing.T) {
	t.Parallel()

	if len(buildLogAllColumns) == len(buildLogPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLBuildLogUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := BuildLog{}
	if err = randomize.Struct(seed, &o, buildLogDBTypes, false); err != nil {
		t.Errorf("Unable to randomize BuildLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BuildLog: %s", err)
	}

	count, err := BuildLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, buildLogDBTypes, false, buildLogPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BuildLog struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BuildLog: %s", err)
	}

	count, err = BuildLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
