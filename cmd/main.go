package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

type item struct {
	Product string `json:"product"`
	Price   int    `json:"price"`
	Rating  int    `json:"rating"`
}

var (
	mostExpensive item
	highestRating item
)

func main() {
	go func() {
		var memStats runtime.MemStats
		ticker := time.Tick(time.Millisecond)

		for {
			select {
			case <-ticker:
				runtime.ReadMemStats(&memStats)
				log.Printf("TotalAlloc: %v, Alloc: %v\n", memStats.TotalAlloc, memStats.Alloc)
			}
		}
	}()

	var err error
	if len(os.Args) != 2 {
		err = fmt.Errorf("use: product_analyzer <file>")
		log.Print(err)
		return
	}

	filePath := os.Args[1]
	fileExtension := filepath.Ext(filePath)

	switch fileExtension {
	case ".csv":
		err = processingCSV(filePath)
	case ".json":
		err = processingJSON(filePath)
	default:
		err = fmt.Errorf("input file must have the extension csv/json")
	}
	if err != nil {
		log.Print(err)
		return
	}

	fmt.Printf("Most expensive product: %s. Top rated product: %s.\n",
		mostExpensive.Product, highestRating.Product)
	log.Print("success")
}

func processingCSV(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = file.Close()
	}()

	reader := csv.NewReader(file)

	//title, err := reader.Read()
	_, err = reader.Read()
	if err != nil {
		log.Fatal(err)
	}
	//log.Print(title)

	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		item, err := CSVToItem(record)
		if err != nil {
			log.Fatal(err)
		}
		//log.Print(item)
		isMostExpensive(item)
		isHighestRating(item)
	}

	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func processingJSON(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = file.Close()
	}()

	decoder := json.NewDecoder(file)
	//t, err := decoder.Token()
	_, err = decoder.Token()
	if err != nil {
		log.Fatal(err)
	}
	//log.Printf("%T: %v\n", t, t)

	for decoder.More() {
		var i item
		if err := decoder.Decode(&i); err != nil {
			log.Fatal(err)
		}
		//log.Print(i)
		isMostExpensive(i)
		isHighestRating(i)
	}

	//t, err = decoder.Token()
	_, err = decoder.Token()
	if err != nil {
		return err
	}
	//log.Printf("%T: %v\n", t, t)

	return nil
}

func CSVToItem(recordCSV []string) (item, error) {
	price, err := strconv.Atoi(recordCSV[1])
	if err != nil {
		return item{}, err
	}
	rating, err := strconv.Atoi(recordCSV[2])
	if err != nil {
		return item{}, err
	}

	return item{
		Product: recordCSV[0],
		Price:   price,
		Rating:  rating,
	}, nil
}

func isMostExpensive(i item) {
	if i.Price > mostExpensive.Price {
		mostExpensive = i
	}
}

func isHighestRating(i item) {
	if i.Rating > highestRating.Rating {
		highestRating = i
	}
}
