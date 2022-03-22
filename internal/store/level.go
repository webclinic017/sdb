package store

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
	"github.com/yemingfeng/sdb/internal/conf"
	util2 "github.com/yemingfeng/sdb/internal/util"
	"log"
)

type LevelStore struct {
	db *leveldb.DB
}

func NewLevelStore() *LevelStore {
	dbPath := conf.Conf.Store.Path + "/level"
	db, err := leveldb.OpenFile(dbPath, &opt.Options{Filter: filter.NewBloomFilter(10)})
	if err != nil {
		log.Fatalf("failed to open file: %+v", err)
	}
	log.Printf("db init %s complete", dbPath)

	return &LevelStore{db: db}
}

func (store *LevelStore) NewBatch() Batch {
	transaction, _ := store.db.OpenTransaction()
	return &LevelBatch{db: store.db, transaction: transaction}
}

func (store *LevelStore) Close() error {
	return store.db.Close()
}

type LevelBatch struct {
	db          *leveldb.DB
	transaction *leveldb.Transaction
}

func (batch *LevelBatch) Get(key []byte) ([]byte, error) {
	value, err := batch.transaction.Get(key, nil)
	if err == leveldb.ErrNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return value, err
}

func (batch *LevelBatch) Set(key []byte, value []byte) error {
	return batch.transaction.Put(key, value, &opt.WriteOptions{Sync: true})
}

func (batch *LevelBatch) Del(key []byte) error {
	return batch.transaction.Delete(key, &opt.WriteOptions{Sync: true})
}

func (batch *LevelBatch) Commit() error {
	return batch.transaction.Commit()
}

func (batch *LevelBatch) Iterate(opt *PrefixIteratorOption, handle func([]byte, []byte) error) error {
	it := batch.transaction.NewIterator(util.BytesPrefix(opt.Prefix), nil)
	defer func() {
		it.Release()
	}()

	if opt.Offset >= 0 {
		i := 0
		for it.First(); i < int(opt.Offset) && it.Valid(); it.Next() {
			i++
		}

		i = 0
		for ; it.Valid(); it.Next() {
			err := handle(util2.Copy2(it.Key()), util2.Copy2(it.Value()))
			if err != nil {
				return err
			}
			i++
			if opt.Limit > 0 && i == int(opt.Limit) {
				break
			}
		}
	} else {
		i := 0
		for it.Last(); i < int(-opt.Offset-1) && it.Valid(); it.Prev() {
			i++
		}

		i = 0
		for ; it.Valid(); it.Prev() {
			err := handle(util2.Copy2(it.Key()), util2.Copy2(it.Value()))
			if err != nil {
				return err
			}
			i++
			if opt.Limit > 0 && i == int(opt.Limit) {
				break
			}
		}
	}
	return nil
}

func (batch *LevelBatch) Close() {
	batch.transaction.Discard()
}