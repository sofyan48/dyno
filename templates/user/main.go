package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	consulapi "github.com/hashicorp/consul/api"
)

// User ...
type User struct {
	ID       uint64    `json:"id"`
	Username string    `json:"username"`
	Products []product `json:"products"`
}

// UserOrder ...
type UserOrder struct {
	ID       uint64  `json:"id"`
	Username string  `json:"username"`
	Order    []order `json:"orders"`
}

// product ...
type product struct {
	ID    uint64  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// order ...
type order struct {
	ID     uint64  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Status bool    `json:"status"`
}

// registerServiceWithConsul ...
func registerServiceWithConsul() {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatalln(err)
	}

	registration := new(consulapi.AgentServiceRegistration)

	registration.ID = "user"
	registration.Name = "user"
	address := hostname()
	registration.Address = address
	p, err := strconv.Atoi(port()[1:len(port())])
	if err != nil {
		log.Fatalln(err)
	}
	registration.Port = p
	registration.Check = new(consulapi.AgentServiceCheck)
	registration.Check.HTTP = fmt.Sprintf("http://%s:%v/healthcheck", address, p)
	registration.Check.Interval = "5s"
	registration.Check.Timeout = "3s"
	consul.Agent().ServiceRegister(registration)
}

// lookupServiceWithConsul ...
func lookupServiceWithConsul(svc string) (string, error) {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		return "", err
	}
	services, err := consul.Agent().Services()
	if err != nil {
		return "", err
	}
	srvc := services[svc]
	address := srvc.Address
	port := srvc.Port
	return fmt.Sprintf("http://%s:%v", address, port), nil
}

func main() {
	registerServiceWithConsul()
	http.HandleFunc("/healthcheck", healthcheck)
	http.HandleFunc("/user/product", UserProduct)
	http.HandleFunc("/user/orders", OrderView)
	fmt.Printf("user service is up on port: %s", port())
	http.ListenAndServe(port(), nil)
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `user service is good`)
}

// UserProduct ...
func UserProduct(w http.ResponseWriter, r *http.Request) {
	p := []product{}
	url, err := lookupServiceWithConsul("product")
	fmt.Println("URL: ", url)
	if err != nil {
		fmt.Fprintf(w, "Error. %s", err)
		return
	}
	client := &http.Client{}
	resp, err := client.Get(url + "/products")
	if err != nil {
		fmt.Fprintf(w, "Error. %s", err)
		return
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		fmt.Fprintf(w, "Error. %s", err)
		return
	}
	u := User{
		ID:       1,
		Username: "wn48@gmail.com",
	}
	u.Products = p
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&u)
}

// OrderView ...
func OrderView(w http.ResponseWriter, r *http.Request) {
	p := []order{}
	url, err := lookupServiceWithConsul("order")
	fmt.Println("URL: ", url)
	if err != nil {
		fmt.Fprintf(w, "Error. %s", err)
		return
	}
	client := &http.Client{}
	resp, err := client.Get(url + "/orders")
	if err != nil {
		fmt.Fprintf(w, "Error. %s", err)
		return
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		fmt.Fprintf(w, "Error. %s", err)
		return
	}
	u := UserOrder{
		ID:       1,
		Username: "wn48@gmail.com",
	}
	u.Order = p
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&u)
}

// port ...
func port() string {
	p := os.Getenv("USER_SERVICE_PORT")
	h := os.Getenv("USER_SERVICE_HOST")
	if len(strings.TrimSpace(p)) == 0 {
		return ":8080"
	}
	return fmt.Sprintf("%s:%s", h, p)
}

// hostname ...
func hostname() string {
	// return os.Getenv("CONSUL_HTTP_ADDR")
	hn, err := os.Hostname()
	if err != nil {
		log.Fatalln(err)
	}
	return hn
}
