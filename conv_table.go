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

package oberth

import (
	"github.com/matpool/vitess-sqlparser/go/vt/sqlparser"
)

// ConvFunc conv func
type ConvFunc func(string) string

// ConvTable conv SQL table name
func ConvTable(query string, f ConvFunc) (string, error) {
	if f == nil {
		return query, nil
	}

	stmt, err := sqlparser.Parse(query)
	if err != nil {
		return "", err
	}

	post := func(cursor *sqlparser.Cursor) bool {
		switch n := cursor.Node().(type) {
		case sqlparser.TableIdent:
			if n.String() != "" {
				cursor.Replace(sqlparser.NewTableIdent(f(n.String())))
			}
		case sqlparser.TableName:
			if n.Name.String() != "" {
				cursor.Replace(sqlparser.TableName{
					Name:      sqlparser.NewTableIdent(f(n.Name.String())),
					Qualifier: n.Qualifier,
				})
			}
		}
		return true
	}
	s := sqlparser.Rewrite(stmt, nil, post)

	buf := sqlparser.NewTrackedBuffer(nil)
	s.Format(buf)
	return buf.String(), nil
}
