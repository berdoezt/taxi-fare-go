# Taxi Fare in Golang

A program to calculate taxi fare based on time and distance. Interaction only via terminal. Please find [input](#input) and [output](#output) section for the example. 

## Help
Run `make` to find helpful guidelines for available commands
```
$ make
Management commands for TAXI-FARE-GO:

Usage:
    make build                              Compile the project.
    make run                                Run binary.
    make test                               Run tests on a compiled project.
    make generate                           Find all go:generate command command(s) and execute it.
    make clean                              Clean.
```

### [Input](#input)
Consist of 2 part, `elapsed time<space>mileage`. Format should be like below. `-1` used for terminate the input and triggering the calculation. 
```
00:00:00.000 0.0
00:01:00.123 480.9
00:02:00.125 1141.2
00:03:00.100 6000.8
-1
```

### [Output](#output)
Fare as the output
```
900
```