// Copyright 2019 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"errors"
)

// common errors
// use a language builtin error type without error trace stack
var (
	ErrWriteTsConflict               = errors.New("write ts conflict")
	ErrChangeFeedNotExists           = errors.New("changefeed not exists")
	ErrSubChangeFeedInfoNotExists    = errors.New("subchangefeedinfo not exists")
	ErrWriteSubChangeFeedInfoConlict = errors.New("write subchangefeedinfo conflict")
	ErrFindPLockNotCommit            = errors.New("subchangefeedinfo has p-lock not commited")
)
