package db

import (
	"encoding/binary"
	"encoding/json"
	"time"

	"github.com/boltdb/bolt"
)

var statsBucket = []byte("stats")
var db *bolt.DB

type Stat struct {
	Key    int
	Player string
	Added  int
	Note   string
	Value  map[string]interface{}
}

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(statsBucket)
		return err
	})
}

func CreateStat(player, note string, stats map[string]interface{}) error {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(statsBucket)
		id64, err := b.NextSequence()
		if err != nil {
			return err
		}
		id = int(id64)
		key := itob(id)
		stat := Stat{Key: id, Player: player, Note: note, Added: int(time.Now().Unix()), Value: stats["stats"].(map[string]interface{})}
		encoded, errtwo := json.Marshal(stat)
		if errtwo != nil {
			return errtwo
		}
		return b.Put(key, []byte(encoded))
	})
	if err != nil {
		return err
	}
	return nil
}

func AllStats() ([]Stat, error) {
	var stats []Stat
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(statsBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var stat Stat
			json.Unmarshal(v, &stat)
			stats = append(stats, stat)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return stats, nil
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
