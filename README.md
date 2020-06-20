# The Test

<details close>
<summary>Problem statement</summary>
<br>
We have some customer records in a text file (customers.txt) -- one customer per line, JSON lines formatted. We want to invite any customer within 100km of our Dublin office for some food and drinks on us. Write a program that will read the full list of customers and output the names and user ids of matching customers (within 100km), sorted by User ID (ascending).
<br>
You must use the first formula from [this Wikipedia article](https://en.wikipedia.org/wiki/Great-circle_distance) to calculate distance. Don't forget, you'll need to convert degrees to radians.
<br>
The GPS coordinates for our Dublin office are **53.339428, -6.257664**
<br>
You can find the Customer list [here](https://s3.amazonaws.com/intercom-take-home-test/customers.txt).
</details>


### Usage
Binaries are published in the format: https://github.com/OkBeacon/intercom/releases/download/$GIT_TAG/beer-and-cheer-go-$SYSTEM-$PLAFORM
- On macOS, you would need to run the following and then add to your path:
```
curl -LO "https://github.com/OkBeacon/intercom/releases/download/v0.2/autocannon-go-darwin-10.6-amd64" && chmod +x autocannon-go-darwin-10.6-amd64
```

- On Windows you could install from this link or if you have curl installed then run the following and then add to your path:
```
curl -LO "https://github.com/OkBeacon/intercom/releases/download/v0.2/autocannon-go-windows-4.0-amd64.exe"
```


### Code Structure
```
 .
├──  customer
│  ├──  customerMethods.go
│  ├──  customerMethods_test.go
│  └──  customerTypes.go
├──  go.mod
├──  go.sum
├──  helper
│  ├──  utilMethods.go
│  ├──  utilMethods_test.go
│  └──  utilTypes.go
├──  main.go
└──  README.md
```
