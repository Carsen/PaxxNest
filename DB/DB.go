package DB

import (
	"bytes"
	"log"

	"go.mills.io/bitcask/v2"
)

func CheckForKey(usrk []byte) bool {
	db, err := bitcask.Open("~/PaxxNestDB")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
		db.Close()
	}
	t := db.Has(userk)
	db.Close()
	return t
}

func ValueMatchesKey(userk []byte, userp []byte) bool {
	var checker bool = false
	
	db, err1 := bitcask.Open("DB")
	if err1 != nil {
		log.Fatal(err1)
		checker = false
		db.Close()
		return checker
	}
	
	get, err2 := db.Get(userk)
	if err2 != nil {
		log.Fatal(err2)
		checker = false
		db.Close()
		return checker
	}
	
	checker = bytes.Equal(userp, get)
	db.Close()
	return checker
}

func NewKeyValue(userk []byte, userp []byte) {
	db, err := bitcask.Open("DB")

	if err != nil {
		log.Fatal(err)
		db.Close()
	}
	
	db.Put(userk, userp)
	db.Close()
}
