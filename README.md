# talk-observability

The content found in this repo are for a talk on observability that will be 
presented at nidevconf 2018.


### requirements

 - >= docker-compose version 1.17.1
 - >= docker version 1.17.1
 - free ports 9090 and 8080 (this can be changed in `docker-compose.yml`)

### run

You can run this demonstration using docker-compose by running the following:

```sh
docker-compose up
```

This should run all of the required services. You can view prometheus on 
[localhost:9090](http://localhost:9090) and you can view nginx (load balancing 
over 3 replica /api/foo services) at 
[localhost:8080/api/foo](http://localhost:8080/api/foo).


### test

The endpoint /api/foo is configured to produce some metrics, visible from the
/metrics endpoint. By loading /api/foo several times, prometheus will start to 
scrape these metrics from each of the underlying docker containers

In prometheus you will see the metrics listed as foo_seconds_bucket, 
foo_seconds_count and foo_seconds_sum.
