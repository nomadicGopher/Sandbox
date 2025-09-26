// https://gist.github.com/rlappert/6ad7620fa98dd58b10e39799dfff0ad7

package main

import (
	"fmt"
	"log"
	"time"
)

type Product struct {
	ID    string  // relevant to test
	Price float64 // not relevant to test
}

var (
	pages          = 8
	productStorage = map[string]Product{}
)

func main() {
	t := time.Now()
	go func() {
		// TODO: Understand select {case...} syntax
		select {
		case <-time.After(7 * time.Second): // must complete in 7 seconds
			log.Fatal("timeout exceeded. test cannot take longer than 7 seconds")
		}
	}()

	getAndStoreAllProducts()

	if err := verifyAllProductIDsAreStored(); err != nil {
		log.Fatal("error verifying products:", err)
	}
	log.Println("All products stored successfully. Took", time.Since(t))
}

// ---------- MAIN FUNCTIONS ----------

func getAndStoreAllProducts() {
	/* Problem:
	Requesting product pages and storing them takes too long.
	- requestProducts call takes 1 second (8 seconds total for 8 pages)
	- storeAllProducts call takes 5 seconds
	We have to request and store the products in 7 seconds.*/
	// TODO: Edit getAndStoreAllProducts() to use goroutines to speed up the process.
	// TODO: Research how sync.Mutex() can be is used

	var products []Product
	for i := range pages {
		for _, product := range requestProducts(i) { // each request takes 1 second
			products = append(products, product)
		}
	}
	storeAllProducts(products) // takes 5 seconds. call this only once
}

func verifyAllProductIDsAreStored() error {
	for i := 0; i < pages; i++ {
		for j := 0; j < 10; j++ {
			id := fmt.Sprintf("pg%d_pid%d", i, j)
			if _, ok := productStorage[id]; !ok {
				return fmt.Errorf("product ID %s not stored", id)
			}
		}
	}
	return nil
}

// ---------- HELPERS ----------

// request returns 10 products per page.
func requestProducts(page int) []Product {
	time.Sleep(time.Second)
	return requestInner(page)
}

// storeProduct stores the product in the database.
func storeAllProducts(products []Product) {
	log.Println("Storing products:", len(products))
	time.Sleep(time.Second * 5)
	for _, p := range products {
		storeInner(p)
	}
	log.Println("Products stored:", len(products))
}

func requestInner(page int) []Product {
	if page < 0 {
		panic("invalid page number")
	}
	if page > pages {
		return []Product{}
	}
	products := make([]Product, 0)
	for i := range 10 {
		p := Product{
			ID:    fmt.Sprintf("pg%d_pid%d", page, i),
			Price: float64(page*10 + i),
		}
		products = append(products, p)
	}
	return products
}

func storeInner(p Product) {
	productStorage[p.ID] = p
}
