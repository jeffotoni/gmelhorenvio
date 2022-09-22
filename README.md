# Go Melhor Envio
Esta api faz busca na plataforma [rastreio melhor envio](melhorenvio.com.br) para calculos de frete em diversas outra plataformas dentre elas o correio.
 

## Docker

Build docker image:

```
make docker-build
```

Run docker image:

```
make docker-run
```

## Metrics

#### api fiber
```
$ k6 run -d 90s -u 200 --rps 100000 ./k6/package.stress.js

data_received..................: 2.7 GB  30 MB/s
data_sent......................: 448 MB  5.0 MB/s
http_req_blocked...............: avg=5.5µs    min=779ns   med=1.6µs   max=46.9ms  p(90)=2.28µs   p(95)=2.92µs 
http_req_connecting............: avg=457ns    min=0s      med=0s      max=11.06ms p(90)=0s       p(95)=0s     
http_req_duration..............: avg=13.54ms  min=72.43µs med=10.46ms max=8.53s   p(90)=27.59ms  p(95)=34.03ms
{ expected_response:true }...: avg=13.54ms  min=72.43µs med=10.46ms max=8.53s   p(90)=27.59ms  p(95)=34.03ms
http_req_failed................: 0.00%   ✓ 0            ✗ 1286209
http_req_receiving.............: avg=813.04µs min=12.2µs  med=22.38µs max=110.5ms p(90)=285.66µs p(95)=4.55ms 
http_req_sending...............: avg=56.04µs  min=6.77µs  med=10µs    max=90.79ms p(90)=19.55µs  p(95)=62.99µs
http_req_tls_handshaking.......: avg=0s       min=0s      med=0s      max=0s      p(90)=0s       p(95)=0s     
http_req_waiting...............: avg=12.67ms  min=45µs    med=9.75ms  max=8.53s   p(90)=25.82ms  p(95)=31.81ms
http_reqs......................: 1286209 14290.465091/s
iteration_duration.............: avg=13.94ms  min=128µs   med=10.77ms max=8.53s   p(90)=28.23ms  p(95)=34.83ms
iterations.....................: 1286209 14290.465091/s
vus............................: 200     min=200        max=200  
vus_max........................: 200     min=200        max=200  
```


#### api gin
```
$ k6 run -d 90s -u 200 --rps 100000 ./k6/package.stress.js

data_received..................: 2.4 GB  27 MB/s
data_sent......................: 402 MB  4.5 MB/s
http_req_blocked...............: avg=6.4µs    min=778ns    med=1.62µs  max=52.76ms  p(90)=2.29µs   p(95)=2.94µs 
http_req_connecting............: avg=243ns    min=0s       med=0s      max=34.17ms  p(90)=0s       p(95)=0s     
http_req_duration..............: avg=15.14ms  min=85.99µs  med=12.29ms max=4.52s    p(90)=29.36ms  p(95)=36.31ms
{ expected_response:true }...: avg=15.14ms  min=85.99µs  med=12.29ms max=4.52s    p(90)=29.36ms  p(95)=36.31ms
http_req_failed................: 0.00%   ✓ 0          ✗ 1155028
http_req_receiving.............: avg=657.03µs min=11.19µs  med=21.11µs max=150.75ms p(90)=196.52µs p(95)=3ms    
http_req_sending...............: avg=52.67µs  min=6.84µs   med=10.12µs max=107.99ms p(90)=19.74µs  p(95)=55.62µs
http_req_tls_handshaking.......: avg=0s       min=0s       med=0s      max=0s       p(90)=0s       p(95)=0s     
http_req_waiting...............: avg=14.43ms  min=61.33µs  med=11.71ms max=4.52s    p(90)=27.9ms   p(95)=34.49ms
http_reqs......................: 1155028 12833.1061/s
iteration_duration.............: avg=15.53ms  min=140.39µs med=12.6ms  max=4.52s    p(90)=30.02ms  p(95)=37.21ms
iterations.....................: 1155028 12833.1061/s
vus............................: 200     min=200      max=200  
vus_max........................: 200     min=200      max=200
```

## Endpoints

### POST /v1/frete/calc

```
$ go test -v -run ^TestPostCalc$
```

#### products

```
curl --location --request POST 'http://localhost:8080/v1/frete/calc' \
--header 'Accept: application/json' \
--header 'Authorization: <YOUR_TOKEN>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "from": {
        "postal_code": "96020360"
    },
    "to": {
        "postal_code": "01018020"
    },
    "products": [
        {
            "id": "x",
            "width": 11,
            "height": 17,
            "length": 11,
            "weight": 0.3,
            "insurance_value": 10.1,
            "quantity": 1
        }
    ],
    "options": {
        "receipt": false,
        "own_hand": false
    },
    "services": "1,2,18"
}'
```

```
$ k6 run -d 90s -u 200 --rps 100000 ./k6/products.stress.js
```

#### package

```
curl --location --request POST 'http://localhost:8080/v1/frete/calc' \
--header 'Accept: application/json' \
--header 'Authorization: <YOUR_TOKEN>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "from": {
        "postal_code": "90570020"
    },
    "to": {
        "postal_code": "90570020"
    },
    "package": {
        "height": 4,
        "width": 12,
        "length": 17,
        "weight": 0.3
    }
}'
```

```
$ k6 run -d 90s -u 200 --rps 100000 ./k6/package.stress.js
```

### DELETE /v1/cache

```
curl --location --request DELETE 'http://localhost:8080/v1/cache' \
--header 'Authorization: <API_STATIC_TOKEN>'
```

### DELETE /v1/cache/:key

```
curl --location --request DELETE 'http://localhost:8080/v1/cache/:key' \
--header 'Authorization: <API_STATIC_TOKEN>'

```
