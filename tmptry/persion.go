package tmptry

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type Person struct {
	Name string
	Age  uint
}

func (p *Person) trySer() []byte {

	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(p)
	if err != nil {
		log.Panic(err)
	}

	return buffer.Bytes()
}
func unSer(b []byte) Person {

	var p Person
	decoder := gob.NewDecoder(bytes.NewReader(b))

	err := decoder.Decode(&p)
	if err != nil {
		log.Panic(err)
	}

	p.Age = 10
	return p
}

func SerDemo() {

	p := Person{
		Name: "yao",
		Age:  9,
	}
	fmt.Println(p)
	bytes := p.trySer()

	w := unSer(bytes)
	fmt.Println(w)
}
