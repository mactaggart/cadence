// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package sql

import (
	"fmt"
	"io/ioutil"

	"github.com/iancoleman/strcase"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

const driverName = "mysql"

// TODO: driverName parameter
func newConnection(host string, port int, username, password, dbName string) (*sqlx.DB, error) {
	var db, err = sqlx.Connect(driverName,
		fmt.Sprintf(dataSourceName, username, password, host, port, dbName))
	if err != nil {
		return nil, err
	}
	// Maps struct names in CamelCase to snake without need for db struct tags.
	db.MapperFunc(strcase.ToSnake)
	return db, nil
}

func createDatabase(host string, port int, username, password, dbName string, overwrite bool) error {
	var db, err = sqlx.Connect(driverName,
		fmt.Sprintf(dataSourceName, username, password, host, port, ""))
	if err != nil {
		return fmt.Errorf("failure connecting to mysql database: %v", err)
	}

	if overwrite {
		dropDatabase(db, dbName)
	}
	_, err = db.Exec(`CREATE DATABASE ` + dbName)
	if err != nil {
		return fmt.Errorf("failure creating database %v: %v", dbName, err)
	}
	log.WithField(`database-name`, dbName).Debug(`created database`)
	return nil
}

// DropCassandraKeyspace drops the given keyspace, if it exists
func dropDatabase(db *sqlx.DB, dbName string) (err error) {
	_, err = db.Exec("DROP DATABASE " + dbName)
	if err != nil {
		return err
	}
	log.WithField(`database-name`, dbName).Info(`dropped database`)
	return nil
}

// LoadCassandraSchema loads the schema from the given .sql files on this database
func loadDatabaseSchema(dir string, fileNames []string, db *sqlx.DB, override bool) (err error) {

	for _, file := range fileNames {
		content, err := ioutil.ReadFile(dir + "/" + file)
		if err != nil {
			return fmt.Errorf("error reading contents of file %v:%v", file, err.Error())
		}
		_, err = db.Exec(string(content))
		if err != nil {
			err = fmt.Errorf("error loading schema from %v: %v", file, err.Error())
		}
	}
	return nil
}
