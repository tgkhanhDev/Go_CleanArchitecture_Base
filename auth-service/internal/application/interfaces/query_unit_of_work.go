package interfaces

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// IQueryUnitOfWork interface
type IQueryUnitOfWork interface {
	Repository[T any]() GenericRepository[T]
	Dispose()
}

// ICommandUnitOfWork interface
type ICommandUnitOfWork interface {
	Repository[T any]() GenericRepository[T]
	SaveChanges(ctx context.Context) error
	BeginTransaction(ctx context.Context) error
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	Dispose()
}

// QueryUnitOfWork implementation
type QueryUnitOfWork struct {
	db *gorm.DB
}

func NewQueryUnitOfWork(db *gorm.DB) *QueryUnitOfWork {
	return &QueryUnitOfWork{db: db}
}

func (q *QueryUnitOfWork) Repository[T any]() GenericRepository[T] {
	return NewGenericRepository[T](q.db)
}

func (q *QueryUnitOfWork) Dispose() {
	sqlDB, _ := q.db.DB()
	sqlDB.Close()
}

// CommandUnitOfWork implementation
type CommandUnitOfWork struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewCommandUnitOfWork(db *gorm.DB) *CommandUnitOfWork {
	return &CommandUnitOfWork{db: db}
}

func (c *CommandUnitOfWork) Repository[T any]() GenericRepository[T] {
	if c.tx != nil {
		return NewGenericRepository[T](c.tx)
	}
	return NewGenericRepository[T](c.db)
}

func (c *CommandUnitOfWork) SaveChanges(ctx context.Context) error {
	if c.tx != nil {
		return c.tx.WithContext(ctx).Commit().Error
	}
	return c.db.WithContext(ctx).Commit().Error
}

func (c *CommandUnitOfWork) BeginTransaction(ctx context.Context) error {
	c.tx = c.db.WithContext(ctx).Begin()
	return c.tx.Error
}

func (c *CommandUnitOfWork) Commit(ctx context.Context) error {
	if c.tx != nil {
		err := c.tx.Commit().Error
		c.tx = nil
		return err
	}
	return nil
}

func (c *CommandUnitOfWork) Rollback(ctx context.Context) error {
	if c.tx != nil {
		err := c.tx.Rollback().Error
		c.tx = nil
		return err
	}
	return nil
}

func (c *CommandUnitOfWork) Dispose() {
	sqlDB, _ := c.db.DB()
	sqlDB.Close()
}
