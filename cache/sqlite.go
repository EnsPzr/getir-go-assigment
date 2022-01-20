package cache

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
)

var (
	ErrRequiredKey = errors.New("key is required")
)

var db *sql.DB

func DB() *sql.DB {
	return db
}

// InitSqliteCache
// This function initialize sqlite cache in memory.
// Create a table by name "cache".
// Table has two column. 1- Key. 2- Value.
// Function return one variable. Error.
// If any error exits, returns filled error. But if no error, return nil.
func InitSqliteCache() error {
	var err error
	db, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		return errors.New("cannot open an SQLite memory database => " + err.Error())
	}
	_, err = db.Exec("CREATE TABLE cache (key string, value string);")
	if err != nil {
		return errors.New("cannot create schema => " + err.Error())
	}
	return nil
}

type SqlCacheService struct {
	Db *sql.DB
}

func NewSqlCacheService() *SqlCacheService {
	return &SqlCacheService{Db: db}
}

// Get
// This method returns value according to key in sqlite cache table.
// If no key, return empty value and error. But if any error, returns error.
func (s *SqlCacheService) Get(key string) (string, error) {
	value, err := s.getFromCache(key)
	if err != nil && err == sql.ErrNoRows {
		return "", nil
	}
	if err != nil {
		return "", errors.New("cannot scan addition => " + err.Error())
	}
	return value, nil
}

// Set
// This method create a row in SQLite cache table.
// If exist key in cache table, overwrite value.
// If any error, returns error.
func (s *SqlCacheService) Set(key, value string) error {
	if key == "" {
		return ErrRequiredKey
	}
	_, err := s.getFromCache(key)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if err != nil && err == sql.ErrNoRows {
		_, err = s.Db.Exec("insert into cache (key, value) values (?,?)", key, value)
		return err
	}

	_, err = s.Db.Exec("update cache set value = ? where key=?", value, key)
	return err
}

// This common method return value according to key in sqlite cache table for two methods.
func (s *SqlCacheService) getFromCache(key string) (string, error) {
	// Exec query.
	row := s.Db.QueryRow("SELECT value FROM cache where key = ?", key)
	// If any error, returns error.
	if row.Err() != nil {
		return "", row.Err()
	}
	value := ""
	// Bind row into value.
	err := row.Scan(&value)
	return value, err
}
