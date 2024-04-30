// Code generated by type.go.tmpl. DO NOT EDIT.

// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package math

import (
	"github.com/apache/arrow/go/v17/arrow/array"
)

type Uint64Funcs struct {
	sum func(a *array.Uint64) uint64
}

var (
	Uint64 Uint64Funcs
)

// Sum returns the summation of all elements in a.
func (f Uint64Funcs) Sum(a *array.Uint64) uint64 {
	if a.Len() == 0 {
		return uint64(0)
	}
	return f.sum(a)
}

func sum_uint64_go(a *array.Uint64) uint64 {
	acc := uint64(0)
	for _, v := range a.Uint64Values() {
		acc += v
	}
	return acc
}
