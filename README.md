# ADS-B-Simple
A Simple ADS-B query and printing tool. 

## Description
This program will use the command line arguments or application defaults to query [ADS-B Excange](https://www.adsbexchange.com) for any airplanes near by. 

The program defaults searches for planes within 20NM of Dickinson, TX. This is a small town located south of the Houston Hobby Airport (HOU). 

## Examples
### Defaults
```
go run adsb.go
```

### Search for airplanes within 5NM of LAX
```
go run adsb.go -lat 33.9416 -long -118.9416 -dist 5
```

## Go modules used
- encoding/json
- net/http
- io/ioutil

## More information
Go [here](https://www.adsbexchange.com) to find more information on the ADS-B Exchange.

Go [here](https://www.geodose.com/2018/11/create-simple-live-flight-tracking-python.html) to see a python implementation with a image display.
