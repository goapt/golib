package mathutil

import (
	"fmt"
	"strconv"
)

// Copyright 2014 com authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// PowInt is int type of math.Pow function.
func PowInt(x int, y int) int {
	if y <= 0 {
		return 1
	} else {
		if y%2 == 0 {
			sqrt := PowInt(x, y/2)
			return sqrt * sqrt
		} else {
			return PowInt(x, y-1) * x
		}
	}
}

func Round(f float64) (int, error) {
	return strconv.Atoi(fmt.Sprintf("%.f", f))
}

func MustRound(f float64) (int) {
	n, _ := strconv.Atoi(fmt.Sprintf("%.f", f))
	return n
}

func AbsInt(x int) int {
	if x > 0 {
		return x
	}
	return -1 * x
}
