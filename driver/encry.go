/*
Copyright 2020 The Matpool Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package driver

// e.g. a -> l  A -> L
func caesarEn(s string) string {
	ns := []rune(s)
	for i, c := range ns {
		switch {
		case c >= 'a' && c <= 'z':
			ns[i] = (c-'a'+11)%26 + 'a'

		case c >= 'A' && c <= 'Z':
			ns[i] = (c-'A'+11)%26 + 'A'
		}
	}
	return string(ns)
}
