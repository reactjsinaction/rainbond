// Copyright (C) 2014-2018 Goodrain Co., Ltd.
// RAINBOND, Application Management Platform

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package controller

import (
	"github.com/goodrain/rainbond/pkg/api/proxy"
	"net/http"

	"github.com/Sirupsen/logrus"
)

//EntranceStruct entrance struct
type EntranceStruct struct {
	HTTPProxy proxy.Proxy
}

//Entrance entrance
func (e *EntranceStruct) Entrance(w http.ResponseWriter, r *http.Request) {
	logrus.Debug("proxy entrance api %s", r.RequestURI)
	e.HTTPProxy.Proxy(w, r)
}
