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

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"net/url"

	"github.com/go-sql-driver/mysql"
	"github.com/matpool/oberth"
)

// ConvTableRename url query param
const ConvTableRename = "conv_table_name"

// MySQLDriver matpool mysql driver
type MySQLDriver struct {
	mysql.MySQLDriver
}

// mySQLConn matpool mysql conn
type mysqlConn struct {
	driver.Conn
	tableName oberth.ConvFunc
}

func (mc *mysqlConn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	query, err := oberth.ConvTable(query, mc.tableName)
	if err != nil {
		return nil, err
	}
	return mc.Conn.(driver.QueryerContext).QueryContext(ctx, query, args)
}

func (mc *mysqlConn) Query(query string, args []driver.Value) (driver.Rows, error) {
	query, err := oberth.ConvTable(query, mc.tableName)
	if err != nil {
		return nil, err
	}
	return mc.Conn.(driver.Queryer).Query(query, args)
}

func (mc *mysqlConn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	query, err := oberth.ConvTable(query, mc.tableName)
	if err != nil {
		return nil, err
	}
	return mc.Conn.(driver.ExecerContext).ExecContext(ctx, query, args)
}

func (mc *mysqlConn) Exec(query string, args []driver.Value) (driver.Result, error) {
	query, err := oberth.ConvTable(query, mc.tableName)
	if err != nil {
		return nil, err
	}
	return mc.Conn.(driver.Execer).Exec(query, args)
}

// Open new Connection
// see https://github.com/go-sql-driver/mysql/blob/v1.4.1/driver.go#L59
// add URL query param `conv_table_name`
func (d *MySQLDriver) Open(dataSourceName string) (driver.Conn, error) {
	u, err := url.Parse(dataSourceName)
	if err != nil {
		return nil, err
	}
	vs, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil, err
	}
	var conv oberth.ConvFunc
	if s := vs.Get(ConvTableRename); s != "" {
		conv = (&caesarSalt{s}).conv
	}

	vs.Del(ConvTableRename)
	u.RawQuery = vs.Encode()

	db, err := d.MySQLDriver.Open(u.String())
	if err != nil {
		return nil, err
	}

	return &mysqlConn{
		Conn:      db,
		tableName: conv,
	}, nil
}

func init() {
	sql.Register("oberth_mysql", &MySQLDriver{})
}
