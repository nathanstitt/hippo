// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package hippo

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Tenant is an object representing the database table.
type Tenant struct {
	ID             string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Identifier     string      `boil:"identifier" json:"identifier" toml:"identifier" yaml:"identifier"`
	Name           string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	SubscriptionID int         `boil:"subscription_id" json:"subscription_id" toml:"subscription_id" yaml:"subscription_id"`
	LogoURL        null.String `boil:"logo_url" json:"logo_url,omitempty" toml:"logo_url" yaml:"logo_url,omitempty"`
	HomepageURL    null.String `boil:"homepage_url" json:"homepage_url,omitempty" toml:"homepage_url" yaml:"homepage_url,omitempty"`
	Email          string      `boil:"email" json:"email" toml:"email" yaml:"email"`
	CreatedAt      time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt      time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *tenantR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L tenantL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TenantColumns = struct {
	ID             string
	Identifier     string
	Name           string
	SubscriptionID string
	LogoURL        string
	HomepageURL    string
	Email          string
	CreatedAt      string
	UpdatedAt      string
}{
	ID:             "id",
	Identifier:     "identifier",
	Name:           "name",
	SubscriptionID: "subscription_id",
	LogoURL:        "logo_url",
	HomepageURL:    "homepage_url",
	Email:          "email",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// TenantRels is where relationship names are stored.
var TenantRels = struct {
	Subscription string
	Users        string
}{
	Subscription: "Subscription",
	Users:        "Users",
}

// tenantR is where relationships are stored.
type tenantR struct {
	Subscription *Subscription
	Users        UserSlice
}

// NewStruct creates a new relationship struct
func (*tenantR) NewStruct() *tenantR {
	return &tenantR{}
}

// tenantL is where Load methods for each relationship are stored.
type tenantL struct{}

var (
	tenantColumns               = []string{"id", "identifier", "name", "subscription_id", "logo_url", "homepage_url", "email", "created_at", "updated_at"}
	tenantColumnsWithoutDefault = []string{"identifier", "name", "logo_url", "homepage_url", "email"}
	tenantColumnsWithDefault    = []string{"id", "subscription_id", "created_at", "updated_at"}
	tenantPrimaryKeyColumns     = []string{"id"}
)

type (
	// TenantSlice is an alias for a slice of pointers to Tenant.
	// This should generally be used opposed to []Tenant.
	TenantSlice []*Tenant
	// TenantHook is the signature for custom Tenant hook methods
	TenantHook func(boil.Executor, *Tenant) error

	tenantQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	tenantType                 = reflect.TypeOf(&Tenant{})
	tenantMapping              = queries.MakeStructMapping(tenantType)
	tenantPrimaryKeyMapping, _ = queries.BindMapping(tenantType, tenantMapping, tenantPrimaryKeyColumns)
	tenantInsertCacheMut       sync.RWMutex
	tenantInsertCache          = make(map[string]insertCache)
	tenantUpdateCacheMut       sync.RWMutex
	tenantUpdateCache          = make(map[string]updateCache)
	tenantUpsertCacheMut       sync.RWMutex
	tenantUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var tenantBeforeInsertHooks []TenantHook
var tenantBeforeUpdateHooks []TenantHook
var tenantBeforeDeleteHooks []TenantHook
var tenantBeforeUpsertHooks []TenantHook

var tenantAfterInsertHooks []TenantHook
var tenantAfterSelectHooks []TenantHook
var tenantAfterUpdateHooks []TenantHook
var tenantAfterDeleteHooks []TenantHook
var tenantAfterUpsertHooks []TenantHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Tenant) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range tenantBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Tenant) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range tenantBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Tenant) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range tenantBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Tenant) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range tenantBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Tenant) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range tenantAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Tenant) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range tenantAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Tenant) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range tenantAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Tenant) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range tenantAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Tenant) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range tenantAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTenantHook registers your hook function for all future operations.
func AddTenantHook(hookPoint boil.HookPoint, tenantHook TenantHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		tenantBeforeInsertHooks = append(tenantBeforeInsertHooks, tenantHook)
	case boil.BeforeUpdateHook:
		tenantBeforeUpdateHooks = append(tenantBeforeUpdateHooks, tenantHook)
	case boil.BeforeDeleteHook:
		tenantBeforeDeleteHooks = append(tenantBeforeDeleteHooks, tenantHook)
	case boil.BeforeUpsertHook:
		tenantBeforeUpsertHooks = append(tenantBeforeUpsertHooks, tenantHook)
	case boil.AfterInsertHook:
		tenantAfterInsertHooks = append(tenantAfterInsertHooks, tenantHook)
	case boil.AfterSelectHook:
		tenantAfterSelectHooks = append(tenantAfterSelectHooks, tenantHook)
	case boil.AfterUpdateHook:
		tenantAfterUpdateHooks = append(tenantAfterUpdateHooks, tenantHook)
	case boil.AfterDeleteHook:
		tenantAfterDeleteHooks = append(tenantAfterDeleteHooks, tenantHook)
	case boil.AfterUpsertHook:
		tenantAfterUpsertHooks = append(tenantAfterUpsertHooks, tenantHook)
	}
}

// OneP returns a single tenant record from the query, and panics on error.
func (q tenantQuery) OneP(exec boil.Executor) *Tenant {
	o, err := q.One(exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single tenant record from the query.
func (q tenantQuery) One(exec boil.Executor) (*Tenant, error) {
	o := &Tenant{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "hippo: failed to execute a one query for tenants")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Tenant records from the query, and panics on error.
func (q tenantQuery) AllP(exec boil.Executor) TenantSlice {
	o, err := q.All(exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Tenant records from the query.
func (q tenantQuery) All(exec boil.Executor) (TenantSlice, error) {
	var o []*Tenant

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "hippo: failed to assign all query results to Tenant slice")
	}

	if len(tenantAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Tenant records in the query, and panics on error.
func (q tenantQuery) CountP(exec boil.Executor) int64 {
	c, err := q.Count(exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Tenant records in the query.
func (q tenantQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "hippo: failed to count tenants rows")
	}

	return count, nil
}

// ExistsP checks if the row exists in the table, and panics on error.
func (q tenantQuery) ExistsP(exec boil.Executor) bool {
	e, err := q.Exists(exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q tenantQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "hippo: failed to check if tenants exists")
	}

	return count > 0, nil
}

// Subscription pointed to by the foreign key.
func (o *Tenant) Subscription(mods ...qm.QueryMod) subscriptionQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.SubscriptionID),
	}

	queryMods = append(queryMods, mods...)

	query := Subscriptions(queryMods...)
	queries.SetFrom(query.Query, "\"subscriptions\"")

	return query
}

// Users retrieves all the user's Users with an executor.
func (o *Tenant) Users(mods ...qm.QueryMod) userQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"users\".\"tenant_id\"=?", o.ID),
	)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"users\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"users\".*"})
	}

	return query
}

// LoadSubscription allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (tenantL) LoadSubscription(e boil.Executor, singular bool, maybeTenant interface{}, mods queries.Applicator) error {
	var slice []*Tenant
	var object *Tenant

	if singular {
		object = maybeTenant.(*Tenant)
	} else {
		slice = *maybeTenant.(*[]*Tenant)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &tenantR{}
		}
		args = append(args, object.SubscriptionID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &tenantR{}
			}

			for _, a := range args {
				if a == obj.SubscriptionID {
					continue Outer
				}
			}

			args = append(args, obj.SubscriptionID)

		}
	}

	query := NewQuery(qm.From(`subscriptions`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Subscription")
	}

	var resultSlice []*Subscription
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Subscription")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for subscriptions")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for subscriptions")
	}

	if len(tenantAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Subscription = foreign
		if foreign.R == nil {
			foreign.R = &subscriptionR{}
		}
		foreign.R.Tenants = append(foreign.R.Tenants, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.SubscriptionID == foreign.ID {
				local.R.Subscription = foreign
				if foreign.R == nil {
					foreign.R = &subscriptionR{}
				}
				foreign.R.Tenants = append(foreign.R.Tenants, local)
				break
			}
		}
	}

	return nil
}

// LoadUsers allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (tenantL) LoadUsers(e boil.Executor, singular bool, maybeTenant interface{}, mods queries.Applicator) error {
	var slice []*Tenant
	var object *Tenant

	if singular {
		object = maybeTenant.(*Tenant)
	} else {
		slice = *maybeTenant.(*[]*Tenant)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &tenantR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &tenantR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	query := NewQuery(qm.From(`users`), qm.WhereIn(`tenant_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load users")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice users")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(userAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Users = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &userR{}
			}
			foreign.R.Tenant = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.TenantID {
				local.R.Users = append(local.R.Users, foreign)
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.Tenant = local
				break
			}
		}
	}

	return nil
}

// SetSubscriptionP of the tenant to the related item.
// Sets o.R.Subscription to related.
// Adds o to related.R.Tenants.
// Panics on error.
func (o *Tenant) SetSubscriptionP(exec boil.Executor, insert bool, related *Subscription) {
	if err := o.SetSubscription(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetSubscription of the tenant to the related item.
// Sets o.R.Subscription to related.
// Adds o to related.R.Tenants.
func (o *Tenant) SetSubscription(exec boil.Executor, insert bool, related *Subscription) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"tenants\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"subscription_id"}),
		strmangle.WhereClause("\"", "\"", 2, tenantPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.SubscriptionID = related.ID
	if o.R == nil {
		o.R = &tenantR{
			Subscription: related,
		}
	} else {
		o.R.Subscription = related
	}

	if related.R == nil {
		related.R = &subscriptionR{
			Tenants: TenantSlice{o},
		}
	} else {
		related.R.Tenants = append(related.R.Tenants, o)
	}

	return nil
}

// AddUsersP adds the given related objects to the existing relationships
// of the tenant, optionally inserting them as new records.
// Appends related to o.R.Users.
// Sets related.R.Tenant appropriately.
// Panics on error.
func (o *Tenant) AddUsersP(exec boil.Executor, insert bool, related ...*User) {
	if err := o.AddUsers(exec, insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddUsers adds the given related objects to the existing relationships
// of the tenant, optionally inserting them as new records.
// Appends related to o.R.Users.
// Sets related.R.Tenant appropriately.
func (o *Tenant) AddUsers(exec boil.Executor, insert bool, related ...*User) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.TenantID = o.ID
			if err = rel.Insert(exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"users\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"tenant_id"}),
				strmangle.WhereClause("\"", "\"", 2, userPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.TenantID = o.ID
		}
	}

	if o.R == nil {
		o.R = &tenantR{
			Users: related,
		}
	} else {
		o.R.Users = append(o.R.Users, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &userR{
				Tenant: o,
			}
		} else {
			rel.R.Tenant = o
		}
	}
	return nil
}

// Tenants retrieves all the records using an executor.
func Tenants(mods ...qm.QueryMod) tenantQuery {
	mods = append(mods, qm.From("\"tenants\""))
	return tenantQuery{NewQuery(mods...)}
}

// FindTenantP retrieves a single record by ID with an executor, and panics on error.
func FindTenantP(exec boil.Executor, iD string, selectCols ...string) *Tenant {
	retobj, err := FindTenant(exec, iD, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindTenant retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTenant(exec boil.Executor, iD string, selectCols ...string) (*Tenant, error) {
	tenantObj := &Tenant{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"tenants\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, tenantObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "hippo: unable to select from tenants")
	}

	return tenantObj, nil
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Tenant) InsertP(exec boil.Executor, columns boil.Columns) {
	if err := o.Insert(exec, columns); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Tenant) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("hippo: no tenants provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	if o.UpdatedAt.IsZero() {
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tenantColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	tenantInsertCacheMut.RLock()
	cache, cached := tenantInsertCache[key]
	tenantInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			tenantColumns,
			tenantColumnsWithDefault,
			tenantColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(tenantType, tenantMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(tenantType, tenantMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"tenants\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"tenants\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "hippo: unable to insert into tenants")
	}

	if !cached {
		tenantInsertCacheMut.Lock()
		tenantInsertCache[key] = cache
		tenantInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateP uses an executor to update the Tenant, and panics on error.
// See Update for more documentation.
func (o *Tenant) UpdateP(exec boil.Executor, columns boil.Columns) int64 {
	rowsAff, err := o.Update(exec, columns)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// Update uses an executor to update the Tenant.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Tenant) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	tenantUpdateCacheMut.RLock()
	cache, cached := tenantUpdateCache[key]
	tenantUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			tenantColumns,
			tenantPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("hippo: unable to update tenants, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"tenants\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, tenantPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(tenantType, tenantMapping, append(wl, tenantPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.Exec(cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "hippo: unable to update tenants row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "hippo: failed to get rows affected by update for tenants")
	}

	if !cached {
		tenantUpdateCacheMut.Lock()
		tenantUpdateCache[key] = cache
		tenantUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q tenantQuery) UpdateAllP(exec boil.Executor, cols M) int64 {
	rowsAff, err := q.UpdateAll(exec, cols)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// UpdateAll updates all rows with the specified column values.
func (q tenantQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "hippo: unable to update all for tenants")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "hippo: unable to retrieve rows affected for tenants")
	}

	return rowsAff, nil
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o TenantSlice) UpdateAllP(exec boil.Executor, cols M) int64 {
	rowsAff, err := o.UpdateAll(exec, cols)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TenantSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("hippo: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tenantPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"tenants\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, tenantPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "hippo: unable to update all in tenant slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "hippo: unable to retrieve rows affected all in update all tenant")
	}
	return rowsAff, nil
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Tenant) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, insertColumns); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Tenant) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("hippo: no tenants provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tenantColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
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
	key := buf.String()
	strmangle.PutBuffer(buf)

	tenantUpsertCacheMut.RLock()
	cache, cached := tenantUpsertCache[key]
	tenantUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			tenantColumns,
			tenantColumnsWithDefault,
			tenantColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			tenantColumns,
			tenantPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("hippo: unable to upsert tenants, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(tenantPrimaryKeyColumns))
			copy(conflict, tenantPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"tenants\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(tenantType, tenantMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(tenantType, tenantMapping, ret)
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

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "hippo: unable to upsert tenants")
	}

	if !cached {
		tenantUpsertCacheMut.Lock()
		tenantUpsertCache[key] = cache
		tenantUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Tenant record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Tenant) DeleteP(exec boil.Executor) int64 {
	rowsAff, err := o.Delete(exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// Delete deletes a single Tenant record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Tenant) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("hippo: no Tenant provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), tenantPrimaryKeyMapping)
	sql := "DELETE FROM \"tenants\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "hippo: unable to delete from tenants")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "hippo: failed to get rows affected by delete for tenants")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q tenantQuery) DeleteAllP(exec boil.Executor) int64 {
	rowsAff, err := q.DeleteAll(exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// DeleteAll deletes all matching rows.
func (q tenantQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("hippo: no tenantQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "hippo: unable to delete all from tenants")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "hippo: failed to get rows affected by deleteall for tenants")
	}

	return rowsAff, nil
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o TenantSlice) DeleteAllP(exec boil.Executor) int64 {
	rowsAff, err := o.DeleteAll(exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TenantSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("hippo: no Tenant slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(tenantBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tenantPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"tenants\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, tenantPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "hippo: unable to delete all from tenant slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "hippo: failed to get rows affected by deleteall for tenants")
	}

	if len(tenantAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Tenant) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Tenant) Reload(exec boil.Executor) error {
	ret, err := FindTenant(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *TenantSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TenantSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TenantSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tenantPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"tenants\".* FROM \"tenants\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, tenantPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "hippo: unable to reload all in TenantSlice")
	}

	*o = slice

	return nil
}

// TenantExistsP checks if the Tenant row exists. Panics on error.
func TenantExistsP(exec boil.Executor, iD string) bool {
	e, err := TenantExists(exec, iD)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// TenantExists checks if the Tenant row exists.
func TenantExists(exec boil.Executor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"tenants\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "hippo: unable to check if tenants exists")
	}

	return exists, nil
}
