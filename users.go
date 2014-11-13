package main

import "net/http"
import "os"
import "log"
import "fmt"
import "bufio"
import "io/ioutil"
import "errors"

func main() {
	url := os.Getenv("NAVEED_USERS_URL")
	username := os.Getenv("NAVEED_USERS_USERNAME")
	password := os.Getenv("NAVEED_USERS_PASSWORD")

	if url == "" || username == "" || password == "" { // TODO: optional auth
		log.Printf("ERROR missing settings")
		return
	}

	download("users.json", url, username, password)
}

func download(filePath, url, username, password string) {
	body, err := retrieve(url, username, password)
	if err != nil {
		log.Printf("ERROR retrieving %s", url)
		return
	}

	err = store(body, filePath)
	if err != nil {
		log.Printf("ERROR storing %s", filePath)
	}

	log.Printf("stored %s", filePath)
}

func retrieve(url, username, password string) (body []byte, err error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("request failed for %s", url))
	}
	req.SetBasicAuth(username, password)

	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to request %s", url))
	}
	if res.StatusCode < 200 || res.StatusCode > 299 {
		return nil, errors.New(fmt.Sprintf("unexpected response for %s: %s",
			url, res.Status))
	}

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to retrieve %s", url))
	}

	return body, nil
}

func store(contents []byte, filePath string) (err error) {
	fh, err := os.Create(filePath)
	defer fh.Close()
	if err != nil {
		return errors.New(fmt.Sprintf("failed to open file %s", filePath))
	}

	buffer := bufio.NewWriter(fh)
	defer buffer.Flush()
	_, err = buffer.Write(contents)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to create file %s", filePath))
	}

	return
}
