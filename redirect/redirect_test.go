package main

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func tryConnection(t *testing.T) {
	t.Logf("trying connection at %s...", time.Now().Format("2006/01/02 15:04:05.000"))

	url := "http://localhost:3000"

	myClient := http.Client{
		Timeout: time.Second * 1,
	}

	response, getErr := myClient.Get(url)

	if getErr != nil {
		t.Error("problem with response")
	} else {
		buf, bodyErr := ioutil.ReadAll(response.Body)

		if bodyErr != nil {
			t.Error("problem with body")
		} else {
			exp := "this is fred"
			if string(buf) != exp {
				t.Errorf("the body is wrong: expected %s, got %s", exp, buf)
			}
		}
	}
}

func TestRedirect(t *testing.T) {
	t.Logf("starting test of Redirect at %s...", time.Now().Format("2006/01/02 15:04:05.000"))
	go main()
	//time.Sleep(time.Second * 1)
	tryConnection(t)
	t.Logf("...done test of Redirect at %s...", time.Now().Format("2006/01/02 15:04:05.000"))
}
