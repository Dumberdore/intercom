package helper

import (
	"bufio"
	"encoding/json"
	"io"
	"net/http"
)

// GetJSONCustomer reads json and returns JSONCutomer
func GetJSONCustomer(line string) (JSONCustomer, error) {
	var j JSONCustomer
	err := json.Unmarshal([]byte(line), &j)
	if err != nil {
		return JSONCustomer{}, err
	}
	return j, nil
}

// GetLinesFromURL returns the list of lines as read from input URL
func GetLinesFromURL(s string) []string {
	resp, err := http.Get(s)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var lines []string
	if resp.StatusCode == http.StatusOK {
		lines = httpGetToLines(resp.Body)
	}
	return lines
}

func httpGetToLines(r io.Reader) []string {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}
