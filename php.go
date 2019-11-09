// Copyright 2011 Xing Xing <mikespook@gmail.com> All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package php

// #include "php_embed.h"
import "C"

import (
	"errors"
	"os"
	"syscall"
)

const (
	Success = 0
	Failure = -1
)

var (
	ErrInvalidType  = errors.New("Invalide type")
	ErrInvalidValue = errors.New("Invalide value")
)

type PHP struct {
	stdout, stderr *os.File
	inifile        string
}

func New() (php *PHP) {
	php = &PHP{
		stdout: os.Stdout,
		stderr: os.Stderr,
	}
	return
}

func (php *PHP) Stdout(f *os.File) {
	php.stdout = f
}

func (php *PHP) Stderr(f *os.File) {
	php.stderr = f
}

func (php *PHP) IniFile(ini string) {
	php.inifile = ini
}

func (php *PHP) Startup() {
	// #TODO issue #8
	// We should not use syscall for this purpose,
	// it will affect whole app's output.
	syscall.Dup2(int(php.stdout.Fd()), 1)
	syscall.Dup2(int(php.stderr.Fd()), 2)
	C.php_set_ini(C.CString(php.inifile))
	C.php_startup()
}

func (php *PHP) Exec(filepath string) (err error) {
	if _, err = os.Stat(filepath); err != nil {
		return
	}

	_ = C.php_exec_file(C.CString(filepath))

	err_result := C.php_exec_error()

	if err_result != nil {
		return errors.New(C.GoString(err_result))
	}

	return
}

func (php *PHP) Close() {
	C.php_shutdown()
}
