# Basic Redis K6 InfluxDB Grafana

Start with docker-compose.yml
-
  ```
  docker-compose up or docker-compose up -d
  ```

Setting Grafana
-
 ```
  http://localhost:3000
  ```
- Configuration -> data source
  - Add data source and select InfluxDB
  - Import via grafana.com -> Load and Import
   ```
    HTTP URL
    http://influxdb:8086

    Database
    k6

    ID
    2587
   ```
K6 test
-
  - Change config to load test in **test.js**
```
docker-compose run --rm k6 run /scripts/test.js
```
  
    
