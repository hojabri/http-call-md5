package app

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/hojabri/http-call-md5/filelog"
	"github.com/hojabri/http-call-md5/models"
	"github.com/hojabri/http-call-md5/waitgroup"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

var err error

// StartApplication function gets http addresses and handles requests with goroutines
func StartApplication(addresses []string, parallel int) {

	defer filelog.Close()

	// Create a new custom waitgroup with parallel limit
	wg := waitgroup.NewWaitGroup(parallel)

	chRequests := make(chan string, len(addresses))

	// push all required addresses to the request channel
	pushAddressesToChannel(addresses, chRequests)

	// http call each addresses in the channel
	callRequests(chRequests, wg)

	// Wait until all urls be called
	wg.Wait()

}

// pushAddressesToChannel function pushes requested http urls into chRequestsChannel
func pushAddressesToChannel(addresses []string, chRequests chan string) {
	defer close(chRequests)
	for _, address := range addresses {
		err = checkUrlFormat(&address)
		if err != nil {
			filelog.Logger.Print(fmt.Sprintf("format error for url (%s) err:%s", address, err))
			continue
		}
		chRequests <- address
	}
}

func checkUrlFormat(address *string) error {

	if strings.TrimSpace(*address) =="" {
		return errors.New("blank address")
	}

	parsedURL, err := url.Parse(*address)
	if err != nil {
		return err
	}

	// In case there is no scheme specified for the given url, add http:// before httpcall
	if parsedURL.Scheme == "" {
		*address = "http://" + parsedURL.String()
	}

	return nil
}

// callRequests http calls each addresses in the channel
func callRequests(chRequests chan string, wg *waitgroup.WaitGroup) {

	// Channel for getting httpcall results
	messages := make(chan string)
	defer close(messages)

	// waitgroup for printing all messages
	var wgPrint sync.WaitGroup

	go printMessage(messages, &wgPrint)

	for req := range chRequests {
		// run httpCall with goroutine in parallel
		wg.Add()
		wgPrint.Add(1)
		address :=req
		go func() {
			md5hash, err := httpCall(address)
			wg.Done()
			if err !=nil {
				// TODO: we can then push all failed addresses to a separate channel to try again later
				filelog.Logger.Print(fmt.Sprintf("Invalid response: %s\n", err.Error()))
				wgPrint.Done()
			} else {
				messages <-md5hash
			}
		}()
	}

	wgPrint.Wait()
}

func httpCall(url string) (string, error) {

	httpResponse, err := http.Get(url)

	if err != nil {
		return "", err
	} else if responseByte, err := ioutil.ReadAll(httpResponse.Body); err != nil {
		defer httpResponse.Body.Close()
		return "", err
	} else {

		responseString := string(responseByte)
		responseMD5Bytes := MD5(responseString)
		responseMD5String := hex.EncodeToString(responseMD5Bytes)

		//create a new Response struct and fill with the result
		var response models.Response

		response.Address = url
		response.Hash = responseMD5String
		return response.String(), nil

	}

}

//MD5 function hashes the input string and returns the hash bytes
func MD5(inputText string) []byte {
	h := md5.New()
	io.WriteString(h, inputText)
	return h.Sum(nil)
}

func printMessage(messages <-chan string, wgPrint *sync.WaitGroup) {
	for {
		select {
		case message := <-messages:
			fmt.Println(message)
			wgPrint.Done()
		}
	}
}
