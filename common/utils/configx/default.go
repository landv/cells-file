/*
 * Copyright (c) 2019-2022. Abstrium SAS <team (at) pydio.com>
 * This file is part of Pydio Cells.
 *
 * Pydio Cells is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Pydio Cells is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Pydio Cells.  If not, see <http://www.gnu.org/licenses/>.
 *
 * The latest code can be found at <https://pydio.com>.
 */

package configx

import (
	"fmt"
	"time"

	"github.com/spf13/cast"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	json "github.com/pydio/cells/v4/common/utils/jsonx"
)

type def struct {
	v interface{}
}

func (d *def) Default(i interface{}) Value {
	if d.v == nil {
		d.v = i
	}

	return d
}

func (d *def) Nil() bool {
	return d.v == nil
}

func (d *def) Bool() bool {
	return cast.ToBool(d.v)
}

func (d *def) Bytes() []byte {
	return []byte(cast.ToString(d.v))
}
func (d *def) Key() []string {
	// TODO ?
	fmt.Println("Are we in  here ?")
	return []string{}
}
func (d *def) Reference() Ref {
	r := &ref{}
	if err := d.Scan(r); err != nil {
		return nil
	}

	rr, ok := GetReference(r)
	if ok {
		return rr
	}

	return nil
}
func (d *def) Interface() interface{} {
	return d.v
}
func (d *def) Int() int {
	return cast.ToInt(d.v)
}
func (d *def) Int64() int64 {
	return cast.ToInt64(d.v)
}
func (d *def) Duration() time.Duration {
	return cast.ToDuration(d.v)
}
func (d *def) String() string {
	return cast.ToString(d.v)
}
func (d *def) StringMap() map[string]string {
	return cast.ToStringMapString(d.v)
}
func (d *def) StringArray() []string {
	return cast.ToStringSlice(d.v)
}
func (d *def) Slice() []interface{} {
	return cast.ToSlice(d.v)
}
func (d *def) Map() map[string]interface{} {
	r, _ := cast.ToStringMapE(d.v)
	return r
}
func (d *def) Scan(val interface{}, opts ...Option) error {
	jsonStr, err := json.Marshal(d.v)
	if err != nil {
		return err
	}

	switch v := val.(type) {
	case proto.Message:
		err = protojson.Unmarshal(jsonStr, val.(proto.Message))
	default:
		err = json.Unmarshal(jsonStr, v)
	}

	return err
}
func (d *def) Clone() Value {
	return d
}
