# dis-routing-performance-test

Performance Tests for Routing Investigation POCs

This repository contains the things that are required for running jmeter test scripts against either of the following:
1. The feature/performance-testing-routes branch of the Frontend Router, running locally, and a local `handle-anything-server` that it can proxy requests to.
2. A vanilla NginX server, and an http echo server, which it can proxy requests to, both running in local docker containers.

The jmeter test scripts contain GET requests of the following 3 types:
- definitive paths e.g. /beans123
- paths containing a wildcard prefix e.g. /*/beans123
- paths containing a wildcard suffix e.g. /beans123/*

The Frontend Router (on port 20000) and the NginX server (on port 8080) are both configured to handle the following routes and redirects:

Routes:
```shell
- /beans e.g. http://localhost:20000/beans
- /means e.g. http://localhost:8080/means
- /beans0 /beans1 ... /beans999 e.g. http://localhost:8080/beans118
- /*/beans e.g. http://localhost:20000/some/beans
- /*/beans0 /*/beans1 ... /*/beans999 e.g. http://localhost:8080/some/beans118 
- beans/* e.g.  http://localhost:20000/beans/chips
- beans0/* beans1/* ... beans999/* e.g. http://localhost:8080/beans12/chips
```

Redirects:
```shell
- /cola e.g. http://localhost:20000/cola --> http://localhost:20000/beans
- /cola0 /cola1 ... /cola999 e.g. http://localhost:8080/cola118 --> http://localhost:8080/beans
- cola/* e.g. http://localhost:20000/cola/chips --> http://localhost:20000/heinz
- cola0/* cola1/* ... cola999/* e.g. http://localhost:8080/cola12/chips --> http://localhost:8080/heinz
```

But the NginX server cannot handle wildcard prefix redirects, so these ones are only handled by the Frontend Router:
```shell
- /*/cola e.g. http://localhost:20000/some/cola --> http://localhost:20000/means
- /*/cola0 /*/cola1 ... /*/cola999 e.g. http://localhost:8080/some/cola118 --> http://localhost:8080/means
```

## Getting started

### Testing with the Frontend Router and 'Handle Anything' Server

Firstly, from the root directory, run the server that can handle any upstream requests from the Frontend Router:

```shell
go run main.go
```

To check that it's running just go to any path on port 30001 e.g. http://localhost:30001/anythinghere

Then this message should be returned:

"Hello there. You have reached the server that handles any requests from upstream services!"

Next run the services that the Frontend Router depends on - see [Dependencies](#dependencies)

Next run the Frontend Router itself:

- the frontend router (port 20000):
```shell
cd dp-frontend-router
git checkout feature/performance-testing-routes
make debug
```

Next make sure that you have installed JMeter - see [Dependencies](#dependencies)

Then, to run any of the existing test plans, which are all in the jmeter-test-files directory, run a jmeter command that follows this format:

```shell
jmeter -n -t /path/to/test/plan/Name-Of-Test-Plan.jmx -l Name-Of-Output.csv
```
E.g.
jmeter -n -t /Users/.../dis-routing-performance-test/jmeter-test-files/Hundred-Definitive-Proxies-Test-Plan.jmx -l 
Hundred-Definitive-Proxies-Output.csv

To create a new test plan you can either use the jmeter UI or copy and amend an existing one. Instructions for using the jmeter UI are here:

https://github.com/ONSdigital/dp-operations/blob/main/guides/performance-testing.md#create-a-test-plan-with-a-thread-group

### Dependencies

Services that the Frontend Router depends on:

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

Before running any of the test plans you will need to install jmeter e.g. using brew:

```shell
brew install jmeter
```

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details.

## License

Copyright © 2025, Office for National Statistics (https://www.ons.gov.uk)

Released under MIT license, see [LICENSE](LICENSE.md) for details.
