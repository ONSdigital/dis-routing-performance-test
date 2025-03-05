# dis-routing-performance-test

Performance Tests for Routing Investigation POCs

## Getting started

The service consists of a handler that can handle any requests that are used in the performance tests (`handle-anything-server`, port 30001)

```shell
go run main.go
```

To check that it's running just go to any path on port 30001 e.g. http://localhost:30001/anythinghere

Then this message should be returned:

"Hello there. You have reached the server that handles any requests from upstream services!"

To run any of the existing test plans, which are all in the jmeter-test-files directory, run a jmeter command that follows this format:

```shell
jmeter -n -t /path/to/test/plan/Name-Of-Test-Plan.jmx -l Name-Of-Output.csv
```
E.g.
jmeter -n -t /Users/.../dis-routing-performance-test/jmeter-test-files/Hundred-Definitive-Proxies-Test-Plan.jmx -l 
Hundred-Definitive-Proxies-Output.csv

To create a new test plan you can either use the jmeter UI or copy and amend an existing one. Instructions for using the jmeter UI are here:

https://github.com/ONSdigital/dp-operations/blob/main/guides/performance-testing.md#create-a-test-plan-with-a-thread-group

### Dependencies

Before running any of the test plans you will need to install jmeter e.g. using brew:

```shell
brew install jmeter
```

Also, if you would like to use any of the routes that have been added to the Frontend Router, for testing purposes, then the following services need to be running:
- zebedee
```shell
cd zebedee
./run.sh
```

- the api router (port 23200):
```shell
cd dp-api-router
make debug
```

- the frontend router (port 20000):
```shell
cd dp-frontend-router
git checkout feature/performance-testing-routes
make debug
```

### Configuration

The following routes and redirects have been created, in the feature/performance-testing-routes branch of the Frontend Router, and can be added to other routers for similar tests:

#### Routes

- http://localhost:20000/
- http://localhost:20000/beans
- http://localhost:20000/means
- http://localhost:20000/heinz

- http://localhost:20000/beans0 
http://localhost:20000/beans1 
...
http://localhost:20000/beans999

- http://localhost:20000/*anything*/beans
- http://localhost:20000/*anything*/means
- http://localhost:20000/*anything*/heinz

- http://localhost:20000/*anything*/beans0 
http://localhost:20000/*anything*/beans1 
...
http://localhost:20000/*anything*/beans999

- http://localhost:20000/beans/*anything*
- http://localhost:20000/means/*anything*
- http://localhost:20000/heinz/*anything*

- http://localhost:20000/beans0/*anything* 
http://localhost:20000/beans1/*anything* 
...
http://localhost:20000/beans999/*anything*

#### Redirects
- http://localhost:20000/cola redirects to http://localhost:20000/beans
- http://localhost:20000/capri-sun redirects to http://localhost:20000/beans
- http://localhost:20000/choc%20pudding redirects to http://localhost:20000/beans
- http://localhost:20000/cola0 http://localhost:20000/cola1 
...
http://localhost:20000/cola999 redirects to http://localhost:20000/beans

- http://localhost:20000/*anything*/cola redirects to http://localhost:20000/means
- http://localhost:20000/*anything*/capri-sun redirects to http://localhost:20000/means
- http://localhost:20000/*anything*/choc%20pudding redirects to http://localhost:20000/means
- http://localhost:20000/*anything*/cola0 http://localhost:20000/*anything*/cola1 
...
http://localhost:20000/*anything*/cola999 redirects to http://localhost:20000/means

- http://localhost:20000/cola/*anything* redirects to http://localhost:20000/heinz
- http://localhost:20000/capri-sun/*anything* redirects to http://localhost:20000/heinz
- http://localhost:20000/choc%20pudding/*anything* redirects to http://localhost:20000/heinz
- http://localhost:20000/cola0/*anything* http://localhost:20000/cola1/*anything* 
...
http://localhost:20000/cola999/*anything* redirects to http://localhost:20000/heinz

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details.

## License

Copyright © 2025, Office for National Statistics (https://www.ons.gov.uk)

Released under MIT license, see [LICENSE](LICENSE.md) for details.
