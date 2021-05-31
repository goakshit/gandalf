package persistence

import (
	"context"
	"database/sql"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type gormRepoMock struct {
	mock.Mock
}

// GetGormMock - Returns mock gorm object
func GetGormMock() *gormRepoMock {
	return &gormRepoMock{}
}

func (m *gormRepoMock) DB() (*sql.DB, error) {
	args := m.Called()
	return args.Get(0).(*sql.DB), args.Error(1)
}

func (m *gormRepoMock) Create(value interface{}) DBIface {
	args := m.Called(value)
	return args.Get(0).(DBIface)
}

func (m *gormRepoMock) Clauses(conds ...clause.Expression) DBIface {
	args := m.Called(conds)
	return args.Get(0).(DBIface)
}

func (m *gormRepoMock) Delete(value interface{}, where ...interface{}) DBIface {
	args := m.Called(value, where)
	return args.Get(0).(DBIface)
}

func (m *gormRepoMock) Model(value interface{}) DBIface {
	args := m.Called(value)
	return args.Get(0).(DBIface)
}

func (m *gormRepoMock) Debug() DBIface {
	args := m.Called()
	return args.Get(0).(DBIface)
}

func (m *gormRepoMock) Exec(sql string, values ...interface{}) DBIface {
	args := m.Called(sql, values)
	return args.Get(0).(DBIface)
}

func (m *gormRepoMock) Where(query interface{}, args ...interface{}) DBIface {
	arguments := m.Called(query, args)
	return arguments.Get(0).(DBIface)
}

func (m *gormRepoMock) Or(query interface{}, args ...interface{}) DBIface {
	arguments := m.Called(query, args)
	return arguments.Get(0).(DBIface)
}

func (m *gormRepoMock) Not(query interface{}, args ...interface{}) DBIface {
	arguments := m.Called(query, args)
	return arguments.Get(0).(DBIface)
}

func (m *gormRepoMock) WithContext(ctx context.Context) DBIface {
	args := m.Called()
	return args.Get(0).(DBIface)
}

func (m *gormRepoMock) Limit(value int) DBIface {
	args := m.Called(value)
	return args.Get(0).(DBIface)
}

func (m *gormRepoMock) Offset(value int) DBIface {
	args := m.Called(value)
	return args.Get(0).(DBIface)
}

func (m *gormRepoMock) Order(value interface{}) DBIface {
	args := m.Called(value)
	return args.Get(0).(DBIface)
}

func (m *gormRepoMock) Omit(columns ...string) DBIface {
	args := m.Called(columns)
	return args.Get(0).(DBIface)
}

func (m *gormRepoMock) Group(query string) DBIface {
	args := m.Called(query)
	return args.Get(0).(DBIface)
}

func (m *gormRepoMock) Select(query interface{}, args ...interface{}) DBIface {
	arguments := m.Called(query, args)
	return arguments.Get(0).(DBIface)
}

func (m *gormRepoMock) Having(query string, values ...interface{}) DBIface {
	args := m.Called(query, values)
	return args.Get(0).(DBIface)
}

func (m *gormRepoMock) Scopes(funcs ...func(*gorm.DB) *gorm.DB) DBIface {
	arguments := m.Called(funcs)
	return arguments.Get(0).(DBIface)
}

func (m *gormRepoMock) Unscoped() DBIface {
	arguments := m.Called()
	return arguments.Get(0).(DBIface)
}

func (m *gormRepoMock) Attrs(attrs ...interface{}) DBIface {
	arguments := m.Called(attrs...)
	return arguments.Get(0).(DBIface)
}

func (m *gormRepoMock) Assign(attrs ...interface{}) DBIface {
	arguments := m.Called(attrs...)
	return arguments.Get(0).(DBIface)
}

func (m *gormRepoMock) First(out interface{}, where ...interface{}) DBIface {
	arguments := m.Called(out, where)
	return arguments.Get(0).(DBIface)
}

func (m *gormRepoMock) Joins(query string, args ...interface{}) DBIface {
	arguments := m.Called(query, args)
	return arguments.Get(0).(DBIface)
}

func (m *gormRepoMock) Last(out interface{}, where ...interface{}) DBIface {
	arguments := m.Called(out, where)
	return arguments.Get(0).(DBIface)
}

func (m *gormRepoMock) Find(out interface{}, where ...interface{}) DBIface {
	arguments := m.Called(out, where)
	return arguments.Get(0).(DBIface)
}

func (m *gormRepoMock) Scan(dest interface{}) DBIface {
	arguments := m.Called(dest)
	return arguments.Get(0).(DBIface)
}

func (m *gormRepoMock) Row() *sql.Row {
	arguments := m.Called()
	return arguments.Get(0).(*sql.Row)
}

func (m *gormRepoMock) Rows() (*sql.Rows, error) {
	arguments := m.Called()
	return arguments.Get(0).(*sql.Rows), arguments.Error(1)
}

func (m *gormRepoMock) ScanRows(rows *sql.Rows, result interface{}) error {
	arguments := m.Called(rows, result)
	return arguments.Error(1)
}

func (m *gormRepoMock) Pluck(column string, value interface{}) DBIface {
	arguments := m.Called(column, value)
	return arguments.Get(0).(DBIface)
}

func (m *gormRepoMock) Count(value *int64) DBIface {
	arguments := m.Called(value)
	return arguments.Get(0).(DBIface)
}

func (m *gormRepoMock) FirstOrInit(out interface{}, where ...interface{}) DBIface {
	arguments := m.Called(out, where)
	return arguments.Get(0).(DBIface)
}

func (m *gormRepoMock) FirstOrCreate(out interface{}, where ...interface{}) DBIface {
	arguments := m.Called(out, where)
	return arguments.Get(0).(DBIface)
}

func (m *gormRepoMock) Update(column string, value interface{}) DBIface {
	arguments := m.Called(column, value)
	return arguments.Get(0).(DBIface)
}

func (m *gormRepoMock) Updates(values interface{}) DBIface {
	arguments := m.Called(values)
	return arguments.Get(0).(DBIface)
}

func (m *gormRepoMock) UpdateColumn(column string, value interface{}) DBIface {
	arguments := m.Called(column, value)
	return arguments.Get(0).(DBIface)
}

func (m *gormRepoMock) UpdateColumns(values interface{}) DBIface {
	arguments := m.Called(values)
	return arguments.Get(0).(DBIface)
}

func (m *gormRepoMock) Save(value interface{}) DBIface {
	arguments := m.Called(value)
	return arguments.Get(0).(DBIface)
}

func (m *gormRepoMock) Raw(sql string, values ...interface{}) DBIface {
	arguments := m.Called(sql, values)
	return arguments.Get(0).(DBIface)
}

func (m *gormRepoMock) Table(name string) DBIface {
	arguments := m.Called(name)
	return arguments.Get(0).(DBIface)
}

func (m *gormRepoMock) Begin() DBIface {
	arguments := m.Called()
	return arguments.Get(0).(DBIface)
}

func (m *gormRepoMock) Commit() DBIface {
	arguments := m.Called()
	return arguments.Get(0).(DBIface)
}

func (m *gormRepoMock) Rollback() DBIface {
	arguments := m.Called()
	return arguments.Get(0).(DBIface)
}

func (m *gormRepoMock) Association(column string) *gorm.Association {
	arguments := m.Called(column)
	return arguments.Get(0).(*gorm.Association)
}

func (m *gormRepoMock) Preload(column string, conditions ...interface{}) DBIface {
	arguments := m.Called(column, conditions)
	return arguments.Get(0).(DBIface)
}

func (m *gormRepoMock) Set(name string, value interface{}) DBIface {
	arguments := m.Called(name, value)
	return arguments.Get(0).(DBIface)
}

func (m *gormRepoMock) Get(name string) (value interface{}, ok bool) {
	arguments := m.Called(name)
	return arguments.Get(0).(interface{}), arguments.Bool(1)
}

func (m *gormRepoMock) Error() error {
	arguments := m.Called()
	return arguments.Error(0)
}

func (m *gormRepoMock) RowsAffected() int64 {
	arguments := m.Called()
	return arguments.Get(0).(int64)
}
