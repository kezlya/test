package main

import (
	"net/http"
	"time"
	"io/ioutil"
	"fmt"
)

func main() {
	var httpClient = http.Client{Timeout: 1 * time.Second}

	hosts := []string{"http://njs:7070/","http://ph2:7070/","http://ph3:7070/"}

	for {
		for _,host := range hosts{
			req, e := http.NewRequest("POST", host, nil)
			if e != nil {
				fmt.Println("Fail to create Request", e)
				continue
			}

			response, e := httpClient.Do(req)
			if e != nil {
				fmt.Println("Fail to do a call", e)
				continue
			}
			defer response.Body.Close()

			responseData, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(responseData))
		}
		time.Sleep(3 * time.Second)
	}
}
