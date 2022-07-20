# energy-tracker
A simple CLI application that parses data from the Octopus API in order to give me an overview of my consumption.


A sample output at the moment looks like this

```shell
Data from the last  10  days
electricity £ 18
gas £ 2
```

# Arguments


|Key | Description| Optional |
| -------------- | :--------- | ----------: | 
|f|The properties file that will contain the information that are required to make the calls| Y|
|days| the number of days that we want to go back ( default 10 ) | N |


