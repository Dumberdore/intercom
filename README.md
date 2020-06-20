# The Test

<details close>
<summary>Problem statement</summary>
<br>
<p>
We have some customer records in a text file (customers.txt) -- one customer per line, JSON lines formatted. We want to invite any customer within 100km of our Dublin office for some food and drinks on us. Write a program that will read the full list of customers and output the names and user ids of matching customers (within 100km), sorted by User ID (ascending).
</p>
<p>
<ul>
<li>You must use the first formula from <a href="https://en.wikipedia.org/wiki/Great-circle_distance">this Wikipedia article</a> to calculate distance. Don't forget, you'll need to convert degrees to radians.</li>
<li>The GPS coordinates for our Dublin office are <b>53.339428, -6.257664</b></li>
<li>You can find the Customer list <a href="https://s3.amazonaws.com/intercom-take-home-test/customers.txt">here</a>.</li>
</ul>
</p>
</details>


### Usage
Binaries are published in the format: `https://github.com/OkBeacon/intercom/releases/download/$GIT_TAG/beer-and-cheer-go-$SYSTEM-$PLAFORM`
- On macOS, you would need to run the following and then add to your path:
```
curl -LO "https://github.com/OkBeacon/intercom/releases/download/v1.0/beer-and-cheer-darwin-10.6-amd64" && chmod +x beer-and-cheer-darwin-10.6-amd64
```

- On Windows you could install from this link or if you have curl installed then run the following and then add to your path:
```
curl -LO "https://github.com/OkBeacon/intercom/releases/download/v1.0/beer-and-cheer-windows-4.0-amd64.exe"
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
  4. Print list-of-customers-invited-for-Beer

```

- **To make changes**, clone this repo and make changes to required package along with required test cases
- **To test locally**
```
$ make test
go test ./...
?       intercom        [no test files]
ok      intercom/customer       (cached)
ok      intercom/helper (cached)
```
- **To run locally**
```
$ make run
go run main.go
  4  |  Ian Kehoe
  5  |  Nora Dempsey
  6  |  Theresa Enright
  8  |  Eoin Ahearn
 11  |  Richard Finnegan
 12  |  Christina McArdle
 13  |  Olive Ahearn
 15  |  Michael Ahearn
 17  |  Patricia Cahill
 23  |  Eoin Gallagher
 24  |  Rose Enright
 26  |  Stephen McArdle
 29  |  Oliver Ahearn
 30  |  Nick Enright
 31  |  Alan Behan
 39  |  Lisa Ahearn
```
- **To publish**, tag the code to trigger a [github-action pipeline](https://github.com/OkBeacon/intercom/actions?query=workflow%3ARelease) which will:
	- Execute tests
	- Build binaries for targets: `windows/amd64, linux/amd64, darwin/amd64`
	- Create new release with given tag
	  For example - to create [beer-and-cheer v1.0](https://github.com/OkBeacon/intercom/releases/tag/v1.0)
```
$ git tag -a v1.0 -m "Prod Ready Version"
$ git push origin v1.0
```
### ![Build/Release](https://github.com/OkBeacon/intercom/workflows/Release/badge.svg)
