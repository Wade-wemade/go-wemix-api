// rocksdb_database.go
//go:build !rocksdb
// +build !rocksdb

package rocksdb

import "github.com/wemixarchive/go-wemix/ethdb/leveldb"

func New(file string, cache int, handles int, namespace string, readonly bool) (*leveldb.Database, error) {
	return leveldb.New(file, cache, handles, namespace, readonly)
}
