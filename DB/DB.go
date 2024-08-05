package DB

import (
	"bytes"
	"log"

	"go.mills.io/bitcask/v2"
)

func CheckForKey(userk []byte) (bool, error) {
	var checker bool
	
	db, err := bitcask.Open("DB", bitcask.WithOpenReadonly(true))
	defer db.Close()
	if err != nil {
		log.Fatal(err)
		checker = false
		return checker, err
	}
	checker = db.Has(userk)
	return checker, nil
}

func ValueMatchesKey(userk []byte, userp []byte) (bool, error) {
	var checker bool = false

	db, err := bitcask.Open("DB", bitcask.WithOpenReadonly(true))
	defer db.Close()
	if err != nil {
		log.Fatal(err)
		checker = false
		return checker, err
	}
	
	get, err := db.Get(userk)
	if err != nil {
		log.Fatal(err)
		checker = false
		return checker, err
	}
	checker = bytes.Equal(userp, get)
	return checker, nil
}

func NewKeyValue(userk []byte, userp []byte) error {
	db, err := bitcask.Open("DB", bitcask.WithOpenReadonly(false))
	defer db.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}
	db.Put(userk, userp)
	return nil
}
