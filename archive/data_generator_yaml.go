package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Entry struct {
	Fruit
}
type Fruit struct {
	Name        string
	Weight      float32
	Measure     string
	Price       float32
	GrossWeight float32 `yaml:"gross_weight"`
}

func main() {
	entries := []Fruit{}
	dat, err := ioutil.ReadFile("./data_template.yaml")
	if err != nil {
		log.Println("File Read error: ", err)
	}
	err = yaml.Unmarshal(dat, &entries)
	if err != nil {
		log.Println("File Read error: ", err)
	}
	log.Println("Entry count: ", len(entries))

	fmt.Println(entries[2])
	/*
		entries = append(entries, entry)

		fruit := Fruit{}
		fruit.Name = "Fred"
		fruit.Weight = 32
		fruit.Measure = "Weight"

		entry.Fruit = fruit

		entries = append(entries, entry)
		d, err := yaml.Marshal(&entries)
		if err != nil {
			log.Println("Error: ", err)
		}

		fmt.Println(string(d))
	*/
}
