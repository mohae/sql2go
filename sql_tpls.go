// Copyright 2016-17 Joel Scoble.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package dbsql2go

// TableSQL is used to describe basic components of a sql statement for a
// single table. Everything specified for the WHERE clause is assumed to be an
// AND. This is mainly meant for basic INSERT, UPDATE, SELECT, DELETE
// statements on a table.
type TableSQL struct {
	Columns []string // the columns that will be SELECTed
	Table   string   // the table from which to SELECT
	Where   []string // the where fragment, excluding ?; e.g. id ==, id !=
}
