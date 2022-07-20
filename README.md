# Octopus energy tracker
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


# Properties






Those can be found under Developer settings in https://octopus.energy/dashboard/new/accounts/personal-details 

account_number=xxxxxx
api_key=xxxxxx
gas_account_number=xxxxxxx
gas_serial_number=xxxxxxx
electricity_account_number=xxxxxxxx
electricity_serial_number=xxxxxxxxx

---
Those are the endpoints that we need to call in order to get the consumption.

gas_endpoint=https://api.octopus.energy/v1/gas-meter-points/{account_number}/meters/{serial_number}/consumption/
electricity_endpoint=https://api.octopus.energy/v1/electricity-meter-points/{account_number}/meters/{serial_number}/consumption/

---
The following details can be found in the latest bill if you have fixed price

price_electricity=21.32
price_gas=6.47
standing_electricity=23.68
standing_gas=24.86
