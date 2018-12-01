// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package hippo

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testSubscriptions(t *testing.T) {
	t.Parallel()

	query := Subscriptions()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testSubscriptionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Subscription{}
	if err = randomize.Struct(seed, o, subscriptionDBTypes, true, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Subscriptions().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSubscriptionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Subscription{}
	if err = randomize.Struct(seed, o, subscriptionDBTypes, true, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Subscriptions().DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Subscriptions().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSubscriptionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Subscription{}
	if err = randomize.Struct(seed, o, subscriptionDBTypes, true, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SubscriptionSlice{o}

	if rowsAff, err := slice.DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Subscriptions().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSubscriptionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Subscription{}
	if err = randomize.Struct(seed, o, subscriptionDBTypes, true, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := SubscriptionExists(tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Subscription exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SubscriptionExists to return true, but got false.")
	}
}

func testSubscriptionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Subscription{}
	if err = randomize.Struct(seed, o, subscriptionDBTypes, true, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	subscriptionFound, err := FindSubscription(tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if subscriptionFound == nil {
		t.Error("want a record, got nil")
	}
}

func testSubscriptionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Subscription{}
	if err = randomize.Struct(seed, o, subscriptionDBTypes, true, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Subscriptions().Bind(nil, tx, o); err != nil {
		t.Error(err)
	}
}

func testSubscriptionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Subscription{}
	if err = randomize.Struct(seed, o, subscriptionDBTypes, true, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Subscriptions().One(tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSubscriptionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	subscriptionOne := &Subscription{}
	subscriptionTwo := &Subscription{}
	if err = randomize.Struct(seed, subscriptionOne, subscriptionDBTypes, false, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}
	if err = randomize.Struct(seed, subscriptionTwo, subscriptionDBTypes, false, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = subscriptionOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = subscriptionTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Subscriptions().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSubscriptionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	subscriptionOne := &Subscription{}
	subscriptionTwo := &Subscription{}
	if err = randomize.Struct(seed, subscriptionOne, subscriptionDBTypes, false, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}
	if err = randomize.Struct(seed, subscriptionTwo, subscriptionDBTypes, false, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = subscriptionOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = subscriptionTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Subscriptions().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func subscriptionBeforeInsertHook(e boil.Executor, o *Subscription) error {
	*o = Subscription{}
	return nil
}

func subscriptionAfterInsertHook(e boil.Executor, o *Subscription) error {
	*o = Subscription{}
	return nil
}

func subscriptionAfterSelectHook(e boil.Executor, o *Subscription) error {
	*o = Subscription{}
	return nil
}

func subscriptionBeforeUpdateHook(e boil.Executor, o *Subscription) error {
	*o = Subscription{}
	return nil
}

func subscriptionAfterUpdateHook(e boil.Executor, o *Subscription) error {
	*o = Subscription{}
	return nil
}

func subscriptionBeforeDeleteHook(e boil.Executor, o *Subscription) error {
	*o = Subscription{}
	return nil
}

func subscriptionAfterDeleteHook(e boil.Executor, o *Subscription) error {
	*o = Subscription{}
	return nil
}

func subscriptionBeforeUpsertHook(e boil.Executor, o *Subscription) error {
	*o = Subscription{}
	return nil
}

func subscriptionAfterUpsertHook(e boil.Executor, o *Subscription) error {
	*o = Subscription{}
	return nil
}

func testSubscriptionsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Subscription{}
	o := &Subscription{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, subscriptionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Subscription object: %s", err)
	}

	AddSubscriptionHook(boil.BeforeInsertHook, subscriptionBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	subscriptionBeforeInsertHooks = []SubscriptionHook{}

	AddSubscriptionHook(boil.AfterInsertHook, subscriptionAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	subscriptionAfterInsertHooks = []SubscriptionHook{}

	AddSubscriptionHook(boil.AfterSelectHook, subscriptionAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	subscriptionAfterSelectHooks = []SubscriptionHook{}

	AddSubscriptionHook(boil.BeforeUpdateHook, subscriptionBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	subscriptionBeforeUpdateHooks = []SubscriptionHook{}

	AddSubscriptionHook(boil.AfterUpdateHook, subscriptionAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	subscriptionAfterUpdateHooks = []SubscriptionHook{}

	AddSubscriptionHook(boil.BeforeDeleteHook, subscriptionBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	subscriptionBeforeDeleteHooks = []SubscriptionHook{}

	AddSubscriptionHook(boil.AfterDeleteHook, subscriptionAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	subscriptionAfterDeleteHooks = []SubscriptionHook{}

	AddSubscriptionHook(boil.BeforeUpsertHook, subscriptionBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	subscriptionBeforeUpsertHooks = []SubscriptionHook{}

	AddSubscriptionHook(boil.AfterUpsertHook, subscriptionAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	subscriptionAfterUpsertHooks = []SubscriptionHook{}
}

func testSubscriptionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Subscription{}
	if err = randomize.Struct(seed, o, subscriptionDBTypes, true, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Subscriptions().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSubscriptionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Subscription{}
	if err = randomize.Struct(seed, o, subscriptionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Whitelist(subscriptionColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Subscriptions().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSubscriptionToManyTenants(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Subscription
	var b, c Tenant

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, subscriptionDBTypes, true, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, tenantDBTypes, false, tenantColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, tenantDBTypes, false, tenantColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.SubscriptionID = a.ID
	c.SubscriptionID = a.ID

	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	tenant, err := a.Tenants().All(tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range tenant {
		if v.SubscriptionID == b.SubscriptionID {
			bFound = true
		}
		if v.SubscriptionID == c.SubscriptionID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := SubscriptionSlice{&a}
	if err = a.L.LoadTenants(tx, false, (*[]*Subscription)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Tenants); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Tenants = nil
	if err = a.L.LoadTenants(tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Tenants); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", tenant)
	}
}

func testSubscriptionToManyAddOpTenants(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Subscription
	var b, c, d, e Tenant

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, subscriptionDBTypes, false, strmangle.SetComplement(subscriptionPrimaryKeyColumns, subscriptionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Tenant{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, tenantDBTypes, false, strmangle.SetComplement(tenantPrimaryKeyColumns, tenantColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Tenant{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddTenants(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.SubscriptionID {
			t.Error("foreign key was wrong value", a.ID, first.SubscriptionID)
		}
		if a.ID != second.SubscriptionID {
			t.Error("foreign key was wrong value", a.ID, second.SubscriptionID)
		}

		if first.R.Subscription != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Subscription != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Tenants[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Tenants[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Tenants().Count(tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testSubscriptionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Subscription{}
	if err = randomize.Struct(seed, o, subscriptionDBTypes, true, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testSubscriptionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Subscription{}
	if err = randomize.Struct(seed, o, subscriptionDBTypes, true, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SubscriptionSlice{o}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}

func testSubscriptionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Subscription{}
	if err = randomize.Struct(seed, o, subscriptionDBTypes, true, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Subscriptions().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	subscriptionDBTypes = map[string]string{`Description`: `character varying`, `ID`: `integer`, `Name`: `character varying`, `Price`: `numeric`, `SubscriptionID`: `character varying`, `TrialDuration`: `integer`}
	_                   = bytes.MinRead
)

func testSubscriptionsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(subscriptionPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(subscriptionColumns) == len(subscriptionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Subscription{}
	if err = randomize.Struct(seed, o, subscriptionDBTypes, true, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Subscriptions().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, subscriptionDBTypes, true, subscriptionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	if rowsAff, err := o.Update(tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testSubscriptionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(subscriptionColumns) == len(subscriptionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Subscription{}
	if err = randomize.Struct(seed, o, subscriptionDBTypes, true, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Subscriptions().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, subscriptionDBTypes, true, subscriptionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(subscriptionColumns, subscriptionPrimaryKeyColumns) {
		fields = subscriptionColumns
	} else {
		fields = strmangle.SetComplement(
			subscriptionColumns,
			subscriptionPrimaryKeyColumns,
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

	slice := SubscriptionSlice{o}
	if rowsAff, err := slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testSubscriptionsUpsert(t *testing.T) {
	t.Parallel()

	if len(subscriptionColumns) == len(subscriptionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Subscription{}
	if err = randomize.Struct(seed, &o, subscriptionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Subscription: %s", err)
	}

	count, err := Subscriptions().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, subscriptionDBTypes, false, subscriptionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	if err = o.Upsert(tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Subscription: %s", err)
	}

	count, err = Subscriptions().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
