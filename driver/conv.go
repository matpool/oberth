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

var reservedWord = map[string]bool{
	"dual": true, // https://cloud.tencent.com/developer/article/1374246
}

type caesarSalt struct {
	salt string
}

func (c *caesarSalt) conv(s string) string {
	if _, ok := reservedWord[s]; ok {
		return s
	}
	return caesarEn(c.salt + s)
}
