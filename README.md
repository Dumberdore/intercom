# The Test

<details close>
<summary>Problem statement</summary>
<br>
We have some customer records in a text file (customers.txt) -- one customer per line, JSON lines formatted. We want to invite any customer within 100km of our Dublin office for some food and drinks on us. Write a program that will read the full list of customers and output the names and user ids of matching customers (within 100km), sorted by User ID (ascending).
<br>
You must use the first formula from <a href="https://en.wikipedia.org/wiki/Great-circle_distance">this Wikipedia article<\a> to calculate distance. Don't forget, you'll need to convert degrees to radians.
<br>
The GPS coordinates for our Dublin office are **53.339428, -6.257664**
<br>
You can find the Customer list <a href="https://s3.amazonaws.com/intercom-take-home-test/customers.txt">here</a>.
</details>


### Usage
Binaries are published in the format: `https://github.com/OkBeacon/intercom/releases/download/$GIT_TAG/beer-and-cheer-go-$SYSTEM-$PLAFORM`
- On macOS, you would need to run the following and then add to your path:
```
curl -LO "https://github.com/OkBeacon/intercom/releases/download/v0.2/autocannon-go-darwin-10.6-amd64" && chmod +x autocannon-go-darwin-10.6-amd64
```

- On Windows you could install from this link or if you have curl installed then run the following and then add to your path:
```
curl -LO "https://github.com/OkBeacon/intercom/releases/download/v0.2/autocannon-go-windows-4.0-amd64.exe"
```


### Developer guide

- Code Structure
```
.
├── customer - This package holds types and methods required to implement Customers
│  ├──  customerMethods.go
│  ├──  customerMethods_test.go
│  └──  customerTypes.go

├── helper - This package holds types and methods required to implement utilities
│  ├──  utilMethods.go
│  ├──  utilMethods_test.go
│  └──  utilTypes.go

├──  main.go
└──  README.md

- Program execution starts in main.go
	1. Using utility functions in helper package, it fetches text file from Url and returns the list of strings representings lines in file
	2. For each line,
		a. Cutomer is created with by parsing the json
		b. If the customer is whithin 100km radius, it's userID and Name is appended to the list-of-customers-invited-for-Beer
	3. Sort the list-of-customers-invited-for-Beer by userID
	4. Print list-of-customers-invited-for-Beer by userID

```

- To make changes clone this repo and make changes to required package along with required test cases
- Tag the code if their is any functional code change
	For example -
```
$ git tag -a v1.0 -m "Prod Ready Version"
$ git push origin v1.0
```
	This will trigger a [github-action pipeline](https://github.com/OkBeacon/intercom/actions?query=workflow%3ARelease) which will create new release with tag **v1.0**

### ![Release](https://github.com/OkBeacon/intercom/workflows/Release/badge.svg)
