# DYNO
Service Discovery Consul and Golang
## Getting Started
### Environment
Setup environtmen 
```
touch ~/.dyno
```
Add environtmen
```
nano ~/.dyno

######### .dyno file #########
DISCOVERY_BROKER=consul
CONSUL_HOST=127.0.0.1:8500
```
### Starting Consul
```
consul agent -dev -ui
```

### Started
see templates/examples for dyno templates and templates for other service

#### Building Dyno

Mac Os
```
$GOOS=darwin go build -o dyno src/main.go
```
Linux
```
$GOOS=linux go build -o dyno src/main.go
```
Windows
```
$GOOS=windows go build -o dyno src/main.go
```

#### Add service
```
./dyno service add -f templates/examples/dyno.yml
```

#### Delete service
```
./dyno service delete -f templates/examples/dyno.yml
```

#### Lookup service
```
./dyno service add -f templates/examples/dyno.yml
```