package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	consulapi "github.com/hashicorp/consul/api"
)

type product struct {
	ID    uint64  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ProductConfiguration struct {
	Categories []string `json:"categories"`
}


func Configuration(w http.ResponseWriter, r *http.Request) {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		fmt.Fprintf(w, "Error. %s", err)
		return
	}
	kvpair, _, err := consul.KV().Get("product-configuration", nil)
	if err != nil {
		fmt.Fprintf(w, "Error. %s", err)
		return
	}
	if kvpair.Value == nil {
		fmt.Fprintf(w, "Configuration empty")
		return
	}
	val := string(kvpair.Value)
	fmt.Fprintf(w, "%s", val)

}

func Products(w http.ResponseWriter, r *http.Request) {
	products := []product{
		{
			ID:    1,
			Name:  "Macbook lagi",
			Price: 2000000.00,
		},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&products)
}

func main() {
	// registerServiceWithConsul()
	http.HandleFunc("/healthcheck", healthcheck)
	http.HandleFunc("/products", Products)
	http.HandleFunc("/product/config", Configuration)
	fmt.Printf("product service is up on port: %s", port())
	http.ListenAndServe(port(), nil)
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `product service is good`)
}

func port() string {
	p := os.Getenv("PRODUCT_SERVICE_PORT")
	h := os.Getenv("PRODUCT_SERVICE_HOST")
	if len(strings.TrimSpace(p)) == 0 {
		return ":8300"
	}
	return fmt.Sprintf("%s:%s", h, p)
}

func hostname() string {
	// return os.Getenv("CONSUL_HTTP_ADDR")
	hn, err := os.Hostname()
	if err != nil {
		log.Fatalln(err)
	}
	return hn
}
