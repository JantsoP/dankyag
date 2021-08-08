// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

	"emperror.dev/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// RSVPSession is an object representing the database table.
type RSVPSession struct {
	MessageID       int64     `boil:"message_id" json:"message_id" toml:"message_id" yaml:"message_id"`
	GuildID         int64     `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	ChannelID       int64     `boil:"channel_id" json:"channel_id" toml:"channel_id" yaml:"channel_id"`
	LocalID         int64     `boil:"local_id" json:"local_id" toml:"local_id" yaml:"local_id"`
	AuthorID        int64     `boil:"author_id" json:"author_id" toml:"author_id" yaml:"author_id"`
	CreatedAt       time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	StartsAt        time.Time `boil:"starts_at" json:"starts_at" toml:"starts_at" yaml:"starts_at"`
	Title           string    `boil:"title" json:"title" toml:"title" yaml:"title"`
	Description     string    `boil:"description" json:"description" toml:"description" yaml:"description"`
	MaxParticipants int       `boil:"max_participants" json:"max_participants" toml:"max_participants" yaml:"max_participants"`
	SendReminders   bool      `boil:"send_reminders" json:"send_reminders" toml:"send_reminders" yaml:"send_reminders"`
	SentReminders   bool      `boil:"sent_reminders" json:"sent_reminders" toml:"sent_reminders" yaml:"sent_reminders"`

	R *rsvpSessionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L rsvpSessionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RSVPSessionColumns = struct {
	MessageID       string
	GuildID         string
	ChannelID       string
	LocalID         string
	AuthorID        string
	CreatedAt       string
	StartsAt        string
	Title           string
	Description     string
	MaxParticipants string
	SendReminders   string
	SentReminders   string
}{
	MessageID:       "message_id",
	GuildID:         "guild_id",
	ChannelID:       "channel_id",
	LocalID:         "local_id",
	AuthorID:        "author_id",
	CreatedAt:       "created_at",
	StartsAt:        "starts_at",
	Title:           "title",
	Description:     "description",
	MaxParticipants: "max_participants",
	SendReminders:   "send_reminders",
	SentReminders:   "sent_reminders",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var RSVPSessionWhere = struct {
	MessageID       whereHelperint64
	GuildID         whereHelperint64
	ChannelID       whereHelperint64
	LocalID         whereHelperint64
	AuthorID        whereHelperint64
	CreatedAt       whereHelpertime_Time
	StartsAt        whereHelpertime_Time
	Title           whereHelperstring
	Description     whereHelperstring
	MaxParticipants whereHelperint
	SendReminders   whereHelperbool
	SentReminders   whereHelperbool
}{
	MessageID:       whereHelperint64{field: "\"rsvp_sessions\".\"message_id\""},
	GuildID:         whereHelperint64{field: "\"rsvp_sessions\".\"guild_id\""},
	ChannelID:       whereHelperint64{field: "\"rsvp_sessions\".\"channel_id\""},
	LocalID:         whereHelperint64{field: "\"rsvp_sessions\".\"local_id\""},
	AuthorID:        whereHelperint64{field: "\"rsvp_sessions\".\"author_id\""},
	CreatedAt:       whereHelpertime_Time{field: "\"rsvp_sessions\".\"created_at\""},
	StartsAt:        whereHelpertime_Time{field: "\"rsvp_sessions\".\"starts_at\""},
	Title:           whereHelperstring{field: "\"rsvp_sessions\".\"title\""},
	Description:     whereHelperstring{field: "\"rsvp_sessions\".\"description\""},
	MaxParticipants: whereHelperint{field: "\"rsvp_sessions\".\"max_participants\""},
	SendReminders:   whereHelperbool{field: "\"rsvp_sessions\".\"send_reminders\""},
	SentReminders:   whereHelperbool{field: "\"rsvp_sessions\".\"sent_reminders\""},
}

// RSVPSessionRels is where relationship names are stored.
var RSVPSessionRels = struct {
	RSVPSessionsMessageRSVPParticipants string
}{
	RSVPSessionsMessageRSVPParticipants: "RSVPSessionsMessageRSVPParticipants",
}

// rsvpSessionR is where relationships are stored.
type rsvpSessionR struct {
	RSVPSessionsMessageRSVPParticipants RSVPParticipantSlice
}

// NewStruct creates a new relationship struct
func (*rsvpSessionR) NewStruct() *rsvpSessionR {
	return &rsvpSessionR{}
}

// rsvpSessionL is where Load methods for each relationship are stored.
type rsvpSessionL struct{}

var (
	rsvpSessionAllColumns            = []string{"message_id", "guild_id", "channel_id", "local_id", "author_id", "created_at", "starts_at", "title", "description", "max_participants", "send_reminders", "sent_reminders"}
	rsvpSessionColumnsWithoutDefault = []string{"message_id", "guild_id", "channel_id", "local_id", "author_id", "created_at", "starts_at", "title", "description", "max_participants", "send_reminders", "sent_reminders"}
	rsvpSessionColumnsWithDefault    = []string{}
	rsvpSessionPrimaryKeyColumns     = []string{"message_id"}
)

type (
	// RSVPSessionSlice is an alias for a slice of pointers to RSVPSession.
	// This should generally be used opposed to []RSVPSession.
	RSVPSessionSlice []*RSVPSession

	rsvpSessionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	rsvpSessionType                 = reflect.TypeOf(&RSVPSession{})
	rsvpSessionMapping              = queries.MakeStructMapping(rsvpSessionType)
	rsvpSessionPrimaryKeyMapping, _ = queries.BindMapping(rsvpSessionType, rsvpSessionMapping, rsvpSessionPrimaryKeyColumns)
	rsvpSessionInsertCacheMut       sync.RWMutex
	rsvpSessionInsertCache          = make(map[string]insertCache)
	rsvpSessionUpdateCacheMut       sync.RWMutex
	rsvpSessionUpdateCache          = make(map[string]updateCache)
	rsvpSessionUpsertCacheMut       sync.RWMutex
	rsvpSessionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single rsvpSession record from the query using the global executor.
func (q rsvpSessionQuery) OneG(ctx context.Context) (*RSVPSession, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single rsvpSession record from the query.
func (q rsvpSessionQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RSVPSession, error) {
	o := &RSVPSession{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.WrapIf(err, "models: failed to execute a one query for rsvp_sessions")
	}

	return o, nil
}

// AllG returns all RSVPSession records from the query using the global executor.
func (q rsvpSessionQuery) AllG(ctx context.Context) (RSVPSessionSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all RSVPSession records from the query.
func (q rsvpSessionQuery) All(ctx context.Context, exec boil.ContextExecutor) (RSVPSessionSlice, error) {
	var o []*RSVPSession

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.WrapIf(err, "models: failed to assign all query results to RSVPSession slice")
	}

	return o, nil
}

// CountG returns the count of all RSVPSession records in the query, and panics on error.
func (q rsvpSessionQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all RSVPSession records in the query.
func (q rsvpSessionQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.WrapIf(err, "models: failed to count rsvp_sessions rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q rsvpSessionQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q rsvpSessionQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.WrapIf(err, "models: failed to check if rsvp_sessions exists")
	}

	return count > 0, nil
}

// RSVPSessionsMessageRSVPParticipants retrieves all the rsvp_participant's RSVPParticipants with an executor via rsvp_sessions_message_id column.
func (o *RSVPSession) RSVPSessionsMessageRSVPParticipants(mods ...qm.QueryMod) rsvpParticipantQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"rsvp_participants\".\"rsvp_sessions_message_id\"=?", o.MessageID),
	)

	query := RSVPParticipants(queryMods...)
	queries.SetFrom(query.Query, "\"rsvp_participants\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"rsvp_participants\".*"})
	}

	return query
}

// LoadRSVPSessionsMessageRSVPParticipants allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (rsvpSessionL) LoadRSVPSessionsMessageRSVPParticipants(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRSVPSession interface{}, mods queries.Applicator) error {
	var slice []*RSVPSession
	var object *RSVPSession

	if singular {
		object = maybeRSVPSession.(*RSVPSession)
	} else {
		slice = *maybeRSVPSession.(*[]*RSVPSession)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &rsvpSessionR{}
		}
		args = append(args, object.MessageID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &rsvpSessionR{}
			}

			for _, a := range args {
				if a == obj.MessageID {
					continue Outer
				}
			}

			args = append(args, obj.MessageID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`rsvp_participants`), qm.WhereIn(`rsvp_sessions_message_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.WrapIf(err, "failed to eager load rsvp_participants")
	}

	var resultSlice []*RSVPParticipant
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.WrapIf(err, "failed to bind eager loaded slice rsvp_participants")
	}

	if err = results.Close(); err != nil {
		return errors.WrapIf(err, "failed to close results in eager load on rsvp_participants")
	}
	if err = results.Err(); err != nil {
		return errors.WrapIf(err, "error occurred during iteration of eager loaded relations for rsvp_participants")
	}

	if singular {
		object.R.RSVPSessionsMessageRSVPParticipants = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &rsvpParticipantR{}
			}
			foreign.R.RSVPSessionsMessage = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.MessageID == foreign.RSVPSessionsMessageID {
				local.R.RSVPSessionsMessageRSVPParticipants = append(local.R.RSVPSessionsMessageRSVPParticipants, foreign)
				if foreign.R == nil {
					foreign.R = &rsvpParticipantR{}
				}
				foreign.R.RSVPSessionsMessage = local
				break
			}
		}
	}

	return nil
}

// AddRSVPSessionsMessageRSVPParticipantsG adds the given related objects to the existing relationships
// of the rsvp_session, optionally inserting them as new records.
// Appends related to o.R.RSVPSessionsMessageRSVPParticipants.
// Sets related.R.RSVPSessionsMessage appropriately.
// Uses the global database handle.
func (o *RSVPSession) AddRSVPSessionsMessageRSVPParticipantsG(ctx context.Context, insert bool, related ...*RSVPParticipant) error {
	return o.AddRSVPSessionsMessageRSVPParticipants(ctx, boil.GetContextDB(), insert, related...)
}

// AddRSVPSessionsMessageRSVPParticipants adds the given related objects to the existing relationships
// of the rsvp_session, optionally inserting them as new records.
// Appends related to o.R.RSVPSessionsMessageRSVPParticipants.
// Sets related.R.RSVPSessionsMessage appropriately.
func (o *RSVPSession) AddRSVPSessionsMessageRSVPParticipants(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*RSVPParticipant) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.RSVPSessionsMessageID = o.MessageID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.WrapIf(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"rsvp_participants\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"rsvp_sessions_message_id"}),
				strmangle.WhereClause("\"", "\"", 2, rsvpParticipantPrimaryKeyColumns),
			)
			values := []interface{}{o.MessageID, rel.RSVPSessionsMessageID, rel.UserID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.WrapIf(err, "failed to update foreign table")
			}

			rel.RSVPSessionsMessageID = o.MessageID
		}
	}

	if o.R == nil {
		o.R = &rsvpSessionR{
			RSVPSessionsMessageRSVPParticipants: related,
		}
	} else {
		o.R.RSVPSessionsMessageRSVPParticipants = append(o.R.RSVPSessionsMessageRSVPParticipants, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &rsvpParticipantR{
				RSVPSessionsMessage: o,
			}
		} else {
			rel.R.RSVPSessionsMessage = o
		}
	}
	return nil
}

// RSVPSessions retrieves all the records using an executor.
func RSVPSessions(mods ...qm.QueryMod) rsvpSessionQuery {
	mods = append(mods, qm.From("\"rsvp_sessions\""))
	return rsvpSessionQuery{NewQuery(mods...)}
}

// FindRSVPSessionG retrieves a single record by ID.
func FindRSVPSessionG(ctx context.Context, messageID int64, selectCols ...string) (*RSVPSession, error) {
	return FindRSVPSession(ctx, boil.GetContextDB(), messageID, selectCols...)
}

// FindRSVPSession retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRSVPSession(ctx context.Context, exec boil.ContextExecutor, messageID int64, selectCols ...string) (*RSVPSession, error) {
	rsvpSessionObj := &RSVPSession{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"rsvp_sessions\" where \"message_id\"=$1", sel,
	)

	q := queries.Raw(query, messageID)

	err := q.Bind(ctx, exec, rsvpSessionObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.WrapIf(err, "models: unable to select from rsvp_sessions")
	}

	return rsvpSessionObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *RSVPSession) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *RSVPSession) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no rsvp_sessions provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(rsvpSessionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	rsvpSessionInsertCacheMut.RLock()
	cache, cached := rsvpSessionInsertCache[key]
	rsvpSessionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			rsvpSessionAllColumns,
			rsvpSessionColumnsWithDefault,
			rsvpSessionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(rsvpSessionType, rsvpSessionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(rsvpSessionType, rsvpSessionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"rsvp_sessions\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"rsvp_sessions\" %sDEFAULT VALUES%s"
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
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.WrapIf(err, "models: unable to insert into rsvp_sessions")
	}

	if !cached {
		rsvpSessionInsertCacheMut.Lock()
		rsvpSessionInsertCache[key] = cache
		rsvpSessionInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single RSVPSession record using the global executor.
// See Update for more documentation.
func (o *RSVPSession) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the RSVPSession.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *RSVPSession) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	rsvpSessionUpdateCacheMut.RLock()
	cache, cached := rsvpSessionUpdateCache[key]
	rsvpSessionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			rsvpSessionAllColumns,
			rsvpSessionPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update rsvp_sessions, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"rsvp_sessions\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, rsvpSessionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(rsvpSessionType, rsvpSessionMapping, append(wl, rsvpSessionPrimaryKeyColumns...))
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
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to update rsvp_sessions row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.WrapIf(err, "models: failed to get rows affected by update for rsvp_sessions")
	}

	if !cached {
		rsvpSessionUpdateCacheMut.Lock()
		rsvpSessionUpdateCache[key] = cache
		rsvpSessionUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q rsvpSessionQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q rsvpSessionQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to update all for rsvp_sessions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to retrieve rows affected for rsvp_sessions")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o RSVPSessionSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RSVPSessionSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), rsvpSessionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"rsvp_sessions\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, rsvpSessionPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to update all in rsvpSession slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to retrieve rows affected all in update all rsvpSession")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *RSVPSession) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *RSVPSession) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no rsvp_sessions provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(rsvpSessionColumnsWithDefault, o)

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

	rsvpSessionUpsertCacheMut.RLock()
	cache, cached := rsvpSessionUpsertCache[key]
	rsvpSessionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			rsvpSessionAllColumns,
			rsvpSessionColumnsWithDefault,
			rsvpSessionColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			rsvpSessionAllColumns,
			rsvpSessionPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert rsvp_sessions, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(rsvpSessionPrimaryKeyColumns))
			copy(conflict, rsvpSessionPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"rsvp_sessions\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(rsvpSessionType, rsvpSessionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(rsvpSessionType, rsvpSessionMapping, ret)
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
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.WrapIf(err, "models: unable to upsert rsvp_sessions")
	}

	if !cached {
		rsvpSessionUpsertCacheMut.Lock()
		rsvpSessionUpsertCache[key] = cache
		rsvpSessionUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single RSVPSession record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *RSVPSession) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single RSVPSession record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *RSVPSession) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no RSVPSession provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), rsvpSessionPrimaryKeyMapping)
	sql := "DELETE FROM \"rsvp_sessions\" WHERE \"message_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to delete from rsvp_sessions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.WrapIf(err, "models: failed to get rows affected by delete for rsvp_sessions")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q rsvpSessionQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no rsvpSessionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to delete all from rsvp_sessions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.WrapIf(err, "models: failed to get rows affected by deleteall for rsvp_sessions")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o RSVPSessionSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RSVPSessionSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), rsvpSessionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"rsvp_sessions\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, rsvpSessionPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to delete all from rsvpSession slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.WrapIf(err, "models: failed to get rows affected by deleteall for rsvp_sessions")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *RSVPSession) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no RSVPSession provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *RSVPSession) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRSVPSession(ctx, exec, o.MessageID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RSVPSessionSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty RSVPSessionSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RSVPSessionSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RSVPSessionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), rsvpSessionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"rsvp_sessions\".* FROM \"rsvp_sessions\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, rsvpSessionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.WrapIf(err, "models: unable to reload all in RSVPSessionSlice")
	}

	*o = slice

	return nil
}

// RSVPSessionExistsG checks if the RSVPSession row exists.
func RSVPSessionExistsG(ctx context.Context, messageID int64) (bool, error) {
	return RSVPSessionExists(ctx, boil.GetContextDB(), messageID)
}

// RSVPSessionExists checks if the RSVPSession row exists.
func RSVPSessionExists(ctx context.Context, exec boil.ContextExecutor, messageID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"rsvp_sessions\" where \"message_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, messageID)
	}

	row := exec.QueryRowContext(ctx, sql, messageID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.WrapIf(err, "models: unable to check if rsvp_sessions exists")
	}

	return exists, nil
}
