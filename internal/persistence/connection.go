package persistence

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	"github.com/goakshit/gandalf/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

// Singleton
var db *gorm.DB
var once sync.Once

type DBIface interface {
	DB() (*sql.DB, error)
	Where(query interface{}, args ...interface{}) DBIface
	Clauses(conds ...clause.Expression) DBIface
	Or(query interface{}, args ...interface{}) DBIface
	Not(query interface{}, args ...interface{}) DBIface
	Limit(value int) DBIface
	Offset(value int) DBIface
	Order(value interface{}) DBIface
	Select(query interface{}, args ...interface{}) DBIface
	Omit(columns ...string) DBIface
	Group(query string) DBIface
	Having(query string, values ...interface{}) DBIface
	Joins(query string, args ...interface{}) DBIface
	Scopes(funcs ...func(*gorm.DB) *gorm.DB) DBIface
	Unscoped() DBIface
	Attrs(attrs ...interface{}) DBIface
	Assign(attrs ...interface{}) DBIface
	First(out interface{}, where ...interface{}) DBIface
	Last(out interface{}, where ...interface{}) DBIface
	Find(out interface{}, where ...interface{}) DBIface
	Scan(dest interface{}) DBIface
	Row() *sql.Row
	Rows() (*sql.Rows, error)
	ScanRows(rows *sql.Rows, result interface{}) error
	Pluck(column string, value interface{}) DBIface
	Count(value *int64) DBIface
	FirstOrInit(out interface{}, where ...interface{}) DBIface
	FirstOrCreate(out interface{}, where ...interface{}) DBIface
	Update(column string, value interface{}) DBIface
	Updates(values interface{}) DBIface
	UpdateColumn(column string, value interface{}) DBIface
	UpdateColumns(values interface{}) DBIface
	Save(value interface{}) DBIface
	Create(value interface{}) DBIface
	Delete(value interface{}, where ...interface{}) DBIface
	Raw(sql string, values ...interface{}) DBIface
	Exec(sql string, values ...interface{}) DBIface
	Model(value interface{}) DBIface
	Table(name string) DBIface
	Debug() DBIface
	Begin() DBIface
	Commit() DBIface
	Rollback() DBIface
	Association(column string) *gorm.Association
	Preload(column string, conditions ...interface{}) DBIface
	Set(name string, value interface{}) DBIface
	Get(name string) (value interface{}, ok bool)
	WithContext(ctx context.Context) DBIface

	// extra
	Error() error
	RowsAffected() int64
}

type gormHandler struct {
	w *gorm.DB
}

// Wrap wraps gorm.DB in an interface
func Wrap(db *gorm.DB) DBIface {
	return &gormHandler{db}
}

func (handler *gormHandler) DB() (*sql.DB, error) {
	return handler.w.DB()
}

func (handler *gormHandler) Where(query interface{}, args ...interface{}) DBIface {
	return Wrap(handler.w.Where(query, args...))
}

func (handler *gormHandler) Clauses(conds ...clause.Expression) DBIface {
	return Wrap(handler.w.Clauses(conds...))
}

func (handler *gormHandler) Or(query interface{}, args ...interface{}) DBIface {
	return Wrap(handler.w.Or(query, args...))
}

func (handler *gormHandler) Not(query interface{}, args ...interface{}) DBIface {
	return Wrap(handler.w.Not(query, args...))
}

func (handler *gormHandler) Limit(value int) DBIface {
	return Wrap(handler.w.Limit(value))
}

func (handler *gormHandler) Offset(value int) DBIface {
	return Wrap(handler.w.Offset(value))
}

func (handler *gormHandler) Order(value interface{}) DBIface {
	return Wrap(handler.w.Order(value))
}

func (handler *gormHandler) Select(query interface{}, args ...interface{}) DBIface {
	return Wrap(handler.w.Select(query, args...))
}

func (handler *gormHandler) Omit(columns ...string) DBIface {
	return Wrap(handler.w.Omit(columns...))
}

func (handler *gormHandler) Group(query string) DBIface {
	return Wrap(handler.w.Group(query))
}

func (handler *gormHandler) Having(query string, values ...interface{}) DBIface {
	return Wrap(handler.w.Having(query, values...))
}

func (handler *gormHandler) Joins(query string, args ...interface{}) DBIface {
	return Wrap(handler.w.Joins(query, args...))
}

func (handler *gormHandler) Scopes(funcs ...func(*gorm.DB) *gorm.DB) DBIface {
	return Wrap(handler.w.Scopes(funcs...))
}

func (handler *gormHandler) Unscoped() DBIface {
	return Wrap(handler.w.Unscoped())
}

func (handler *gormHandler) Attrs(attrs ...interface{}) DBIface {
	return Wrap(handler.w.Attrs(attrs...))
}

func (handler *gormHandler) Assign(attrs ...interface{}) DBIface {
	return Wrap(handler.w.Assign(attrs...))
}

func (handler *gormHandler) First(out interface{}, where ...interface{}) DBIface {
	return Wrap(handler.w.First(out, where...))
}

func (handler *gormHandler) Last(out interface{}, where ...interface{}) DBIface {
	return Wrap(handler.w.Last(out, where...))
}

func (handler *gormHandler) Find(out interface{}, where ...interface{}) DBIface {
	return Wrap(handler.w.Find(out, where...))
}

func (handler *gormHandler) Scan(dest interface{}) DBIface {
	return Wrap(handler.w.Scan(dest))
}

func (handler *gormHandler) Row() *sql.Row {
	return handler.w.Row()
}

func (handler *gormHandler) Rows() (*sql.Rows, error) {
	return handler.w.Rows()
}

func (handler *gormHandler) ScanRows(rows *sql.Rows, result interface{}) error {
	return handler.w.ScanRows(rows, result)
}

func (handler *gormHandler) Pluck(column string, value interface{}) DBIface {
	return Wrap(handler.w.Pluck(column, value))
}

func (handler *gormHandler) Count(value *int64) DBIface {
	return Wrap(handler.w.Count(value))
}

func (handler *gormHandler) FirstOrInit(out interface{}, where ...interface{}) DBIface {
	return Wrap(handler.w.FirstOrInit(out, where...))
}

func (handler *gormHandler) FirstOrCreate(out interface{}, where ...interface{}) DBIface {
	return Wrap(handler.w.FirstOrCreate(out, where...))
}

func (handler *gormHandler) Update(column string, value interface{}) DBIface {
	return Wrap(handler.w.Update(column, value))
}

func (handler *gormHandler) Updates(values interface{}) DBIface {
	return Wrap(handler.w.Updates(values))
}

func (handler *gormHandler) UpdateColumn(column string, value interface{}) DBIface {
	return Wrap(handler.w.UpdateColumn(column, value))
}

func (handler *gormHandler) UpdateColumns(values interface{}) DBIface {
	return Wrap(handler.w.UpdateColumns(values))
}

func (handler *gormHandler) Save(value interface{}) DBIface {
	return Wrap(handler.w.Save(value))
}

func (handler *gormHandler) Create(value interface{}) DBIface {
	return Wrap(handler.w.Create(value))
}

func (handler *gormHandler) Delete(value interface{}, where ...interface{}) DBIface {
	return Wrap(handler.w.Delete(value, where...))
}

func (handler *gormHandler) Raw(sql string, values ...interface{}) DBIface {
	return Wrap(handler.w.Raw(sql, values...))
}

func (handler *gormHandler) Exec(sql string, values ...interface{}) DBIface {
	return Wrap(handler.w.Exec(sql, values...))
}

func (handler *gormHandler) Model(value interface{}) DBIface {
	return Wrap(handler.w.Model(value))
}

func (handler *gormHandler) Table(name string) DBIface {
	return Wrap(handler.w.Table(name))
}

func (handler *gormHandler) Debug() DBIface {
	return Wrap(handler.w.Debug())
}

func (handler *gormHandler) Begin() DBIface {
	return Wrap(handler.w.Begin())
}

func (handler *gormHandler) Commit() DBIface {
	return Wrap(handler.w.Commit())
}

func (handler *gormHandler) Rollback() DBIface {
	return Wrap(handler.w.Rollback())
}

func (handler *gormHandler) Association(column string) *gorm.Association {
	return handler.w.Association(column)
}

func (handler *gormHandler) Preload(column string, conditions ...interface{}) DBIface {
	return Wrap(handler.w.Preload(column, conditions...))
}

func (handler *gormHandler) Set(name string, value interface{}) DBIface {
	return Wrap(handler.w.Set(name, value))
}

func (handler *gormHandler) Get(name string) (value interface{}, ok bool) {
	return handler.w.Get(name)
}

func (handler *gormHandler) RowsAffected() int64 {
	return handler.w.RowsAffected
}

func (handler *gormHandler) Error() error {
	return handler.w.Error
}

func (handler *gormHandler) WithContext(ctx context.Context) DBIface {
	return Wrap(handler.w.WithContext(ctx))
}

func getGORMConfig() *gorm.Config {
	return &gorm.Config{
		// Ignore default transaction started by GORM. Improves performance by upto 30%
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// Doesn't pluralize the table names
			// Eg: 'user' table won't be pluralized to 'users' table
			SingularTable: true,
		},
	}
}

// Return postgres connection string
func getPostgresConnString() string {
	config := config.New()
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		config.Database.User, config.Database.Passwd, config.Database.Host, config.Database.Port, config.Database.DatabaseName)
}

// GetGormClient - Returns db client
func GetGormClient() DBIface {
	once.Do(func() {
		var err error
		db, err = gorm.Open(postgres.Open(getPostgresConnString()), getGORMConfig())
		if err != nil {
			panic("Failed to open postgres connection\n" + err.Error())
		}
	})
	return Wrap(db)
}
