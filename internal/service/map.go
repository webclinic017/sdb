package service

import (
	"github.com/yemingfeng/sdb/internal/store"
	pb "github.com/yemingfeng/sdb/pkg/protobuf"
	"math"
)

var mapCollection = store.NewCollection(pb.DataType_MAP)

func MPush(key []byte, pairs []*pb.Pair) error {
	lock(LMap, key)
	defer unlock(LMap, key)

	batch := store.NewBatch()
	defer batch.Close()

	for i := range pairs {
		if err := mapCollection.UpsertRow(&store.Row{
			Key:   key,
			Id:    pairs[i].Key,
			Value: pairs[i].Value,
		}, batch); err != nil {
			return err
		}
		if err := PAdd(pb.DataType_MAP, key, batch); err != nil {
			return err
		}
	}
	return batch.Commit()
}

func MPop(key []byte, keys [][]byte) error {
	lock(LMap, key)
	defer unlock(LMap, key)

	batch := store.NewBatch()
	defer batch.Close()

	for i := range keys {
		if err := mapCollection.DelRowById(key, keys[i], batch); err != nil {
			return err
		}
	}

	// delete if not element at key
	rows, err := mapCollection.PageWithBatch(key, 0, 1, batch)
	if err != nil {
		return err
	}
	if len(rows) == 0 {
		if err := PDel(pb.DataType_MAP, key, batch); err != nil {
			return err
		}
	}

	return batch.Commit()
}

func MExist(key []byte, keys [][]byte) ([]bool, error) {
	res := make([]bool, len(keys))
	for i := range keys {
		exist, err := mapCollection.ExistRowById(key, keys[i])
		if err != nil {
			return nil, err
		}
		res[i] = exist
	}
	return res, nil
}

func MDel(key []byte) error {
	lock(LMap, key)
	defer unlock(LMap, key)

	batch := store.NewBatch()
	defer batch.Close()

	if err := mapCollection.DelAll(key, batch); err != nil {
		return err
	}
	if err := PDel(pb.DataType_MAP, key, batch); err != nil {
		return err
	}
	return batch.Commit()
}

func MCount(key []byte) (uint32, error) {
	return mapCollection.Count(key)
}

func MMembers(key []byte) ([]*pb.Pair, error) {
	rows, err := mapCollection.Page(key, 0, math.MaxUint32)
	if err != nil {
		return nil, err
	}

	res := make([]*pb.Pair, len(rows))
	for i := range rows {
		res[i] = &pb.Pair{Key: rows[i].Id, Value: rows[i].Value}
	}

	return res, nil
}
