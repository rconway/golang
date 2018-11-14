package main

import (
	"testing"
	"time"
)

func TestRedirect(t *testing.T) {
	t.Logf("starting test of Redirect at %s...", time.Now().Format("2006/01/02 15:04:05.000"))
	go main()
	t.Logf("...done test of Redirect at %s...", time.Now().Format("2006/01/02 15:04:05.000"))
}
