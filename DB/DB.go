package DB

import (
	"bytes"
	"log"

	"go.mills.io/bitcask/v2"
)

func CheckForKey(userk []byte) bool {
	dbCheck, err := bitcask.Open("DB", bitcask.WithOpenReadonly(true))
	if err != nil {
		log.Fatal(err)
		dbCheck.Close()
	}
	dbCheck.Close()
	t := dbCheck.Has(userk)
	dbCheck.Close()
	return t
}

func ValueMatchesKey(userk []byte, userp []byte) bool {
	var checker bool = false

	dbMatch, err1 := bitcask.Open("DB", bitcask.WithOpenReadonly(true))

	if err1 != nil {
		log.Fatal(err1)
		checker = false
		dbMatch.Close()
		return checker
	}
	dbMatch.Close()
	get, err2 := db.Get(userk)
	if err2 != nil {
		log.Fatal(err2)
		checker = false
		dbMatch.Close()
		return checker
	}
	dbMatch.Close()
	checker = bytes.Equal(userp, get)
	dbMatch.Close()
	return checker
}

func NewKeyValue(userk []byte, userp []byte) {
	dbNew, err := bitcask.Open("DB", bitcask.WithOpenReadonly(false))

	if err != nil {
		log.Fatal(err)
		dbNew.Close()
	}
	dbNew.Close()
	dbNew.Put(userk, userp)
	dbNew.Close()
}
