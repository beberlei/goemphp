// Copyright 2011 Xing Xing <mikespook@gmail.com> All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package php

import (
	"testing"
)

const (
	FileName     = "test.php"
	TestFileName = "test.txt"
)

var (
	php *PHP
)

func init() {
	php = New()
	// #TODO issue #8
	/*
		devNull, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		php.Stdout(devNull)
		php.Stderr(devNull)
	*/
	php.Startup()
}

func TestExec(t *testing.T) {
	if err := php.Exec("fixtures/exec1.php"); err != nil {
		t.Errorf("php.Exec: %s", err)
	}
}

func TestExecErr(t *testing.T) {
	if err := php.Exec("fixtures/exec2.php"); err != nil {
		t.Errorf("php.Exec: %s", err)
	}
}

func TestExecNotFound(t *testing.T) {
	if err := php.Exec("fixtures/not-found.php"); err == nil {
		t.Errorf("php.Exec should have a panic.")
	}
}
