package main

import (
	//"fmt"
	// "gopkg.in/yaml.v2"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v1"
)

// Facts is the basis
type Facts struct {
	About        string
	Name         string
	DateScale    string `yaml:"date_scale"`
	QuantityLow  int    `yaml:"quantity_low"`
	QuantityHigh int    `yaml:"quantity_high"`
}

// Config sets the stage for building data
type Config struct {
	Facts
	Dimensions []string
}

type entry struct {
	fruit
}

type fruit struct {
	Name        string
	Weight      float32
	Measure     string
	Price       float32
	GrossWeight float32 `yaml:"gross_weight"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	f, err := os.Open("./Apples_and_Oranges/fruit.csv")
	if err != nil {
		panic(err)
	}

	var config Config
	log.Print("About to load config")
	configFile, err := ioutil.ReadFile("./Apples_and_Oranges/config.yaml")
	check(err)

	err = yaml.Unmarshal(configFile, &config)
	check(err)

	fmt.Println(config.DateScale)
	fmt.Println(config.QuantityLow)
	fmt.Println(config.QuantityHigh)

	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Println("CSV Read Error: ", err)
	}
	log.Println("Entry count: ", len(lines))
	// fmt.Printf(lines[0].Name)
}
