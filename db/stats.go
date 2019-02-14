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
	Added  string
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
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)
		row := make([]interface{}, 3)
		row[0] = player
		row[1] = note
		row[2] = stats
		encoded, errtwo := json.Marshal(row)
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

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
