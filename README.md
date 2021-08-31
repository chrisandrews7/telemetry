# Telemetry

Consume telemetry from ground station TCP servers, store the data in InfluxDB and visualise in Grafana.

## Start dependencies
This will start Influx, Grafana and the ground station telemetry server.   
It will require you to have Docker installed.

```
make start-deps
```

## Start app

Start the application to start consuming the telemetry.   
When stopping the app it will finish processing the telemetry consumed so far.

```
make start
```

## Visualising Telemetry
Grafana is used to visualise the data.  
Visit [http://localhost:3000](http://localhost:3000) and use username:`admin` and password:`admin`.  
Use the premade dashboard called `Telemetry`.

## Testing

### Unit tests

```
make test
```

### Integration tests
This will require Influx to be running.  

```
make integration-test
```

### Generating mocks
```
make generate-mocks
```