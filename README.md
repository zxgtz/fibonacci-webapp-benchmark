# Fibonacci Web Application Benchmark

`fibonacci-webapp-benchmark` contains dockerized web-apps written in different languages and frameworks that get a number as url parameter and calculate the corresponding fibonacci number. 

All these implementations are used for benchmarking.

## Requirements

- [docker](https://www.docker.com/)
- [ab](https://httpd.apache.org/docs/2.4/programs/ab.html)

## Usage

Run `./docker-compose up -d` against this repository:

```
$ ./docker-compose up -d
```

Then execute the command:

```
$ docker ps
CONTAINER ID        IMAGE                             COMMAND                  CREATED             STATUS              PORTS                    NAMES
14e0d2388dca        fibonacciwebappbenchmark_node     "npm start"              6 seconds ago       Up 5 seconds        0.0.0.0:8080->8080/tcp   fibonacciwebappbenchmark_node_1
8b1bdd070f83        fibonacciwebappbenchmark_ruby     "bundle exec ruby sin"   23 seconds ago      Up 22 seconds       0.0.0.0:4567->4567/tcp   fibonacciwebappbenchmark_ruby_1
333360123b56        fibonacciwebappbenchmark_go       "go run martini.go"      34 seconds ago      Up 32 seconds       0.0.0.0:3000->3000/tcp   fibonacciwebappbenchmark_go_1
df50829f511b        fibonacciwebappbenchmark_python   "python app.py"          42 seconds ago      Up 41 seconds       0.0.0.0:5000->5000/tcp   fibonacciwebappbenchmark_python_1
```

## Screenshots

Now you can open your favorite Web Browser and visit http://localhost:{port}/{number}, for example:

Open http://localhost:5000/10, and you can see:

```
Python + Flask

fib(10): 55
```

## Benchmark

```
$ ab -n 100000 -c 100 http://localhost:5000/10
This is ApacheBench, Version 2.3 <$Revision: 1430300 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:        Werkzeug/0.11.4
Server Hostname:        localhost
Server Port:            5000

Document Path:          /10
Document Length:        29 bytes

Concurrency Level:      100
Time taken for tests:   171.011 seconds
Complete requests:      100000
Failed requests:        0
Write errors:           0
Total transferred:      18400000 bytes
HTML transferred:       2900000 bytes
Requests per second:    584.76 [#/sec] (mean)
Time per request:       171.011 [ms] (mean)
Time per request:       1.710 [ms] (mean, across all concurrent requests)
Transfer rate:          105.07 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.3      0      10
Processing:    21  171  16.2    168     284
Waiting:        6  170  15.9    167     276
Total:         21  171  16.2    169     284

Percentage of the requests served within a certain time (ms)
  50%    169
  66%    174
  75%    178
  80%    180
  90%    189
  95%    200
  98%    218
  99%    233
 100%    284 (longest request)
```

For more detailed results, please refer to my blog: 

http://startover.github.io/


## Links

Related articles:

- https://medium.com/@tschundeee/express-vs-flask-vs-go-acc0879c2122#.wfmtvnnxq 
- https://realpython.com/blog/python/python-ruby-and-golang-a-web-Service-application-comparison/

