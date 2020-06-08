package tmptry

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func DBDemo() {
	db, err := bolt.Open("bit.db", 0600, nil)
	defer db.Close()
	if err != nil {
		log.Panic(err)
	}
	n := "bucket-1"
	db.Update(func(tx *bolt.Tx) error {

		bucket := tx.Bucket([]byte(n))
		if bucket == nil {
			if bucket, err = tx.CreateBucket([]byte(n)); err != nil {
				log.Panic(err)
			}
		}

		bucket.Put([]byte("foo"), []byte("boo~~!@~!@"))
		return nil
	})

	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(n))
		if bucket == nil {
			log.Panic(err)
		}

		v1 := bucket.Get([]byte("foo"))
		v2 := bucket.Get([]byte("foo1"))

		fmt.Printf("V1:%s\n", v1)
		fmt.Printf("V2:%s\n", v2)
		return nil
	})

}

func Break3() {

	a := []int{1, 4, 7}
	b := []int{2, 5, 8}
	c := []int{3, 51, 17}

OVER:
	for _, i := range a {

		for _, j := range b {
			if j == 8 {
				fmt.Println("~~~~~~~~~~~~~~")
				break OVER
			}
			for _, k := range c {
				if k == 51 {
					break OVER
				}
				fmt.Println(i, j, k)
			}

		}
	}
	fmt.Println("ooooo")
}

func MapDemo() {
	m := make(map[string][]int)

	//m["yao"] = []int{1}
	m["yao"] = append(m["yao"], 2)

	fmt.Println(m)
}
