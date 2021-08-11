// Gulu - Golang common utilities for everyone.
// Copyright (c) 2019-present, b3log.org
//
// Gulu is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package gulu

import (
	"compress/gzip"
	"encoding/json"
	"net/http"
)

// Result represents a common-used result struct.
type Result struct {
	Code int         `json:"code"` // return code
	Msg  string      `json:"msg"`  // message
	Data interface{} `json:"data"` // data object
}

// NewResult creates a result with Code=0, Msg="", Data=nil.
func (*GuluRet) NewResult() *Result {
	return &Result{
		Code: 0,
		Msg:  "",
		Data: nil,
	}
}

// RetResult writes HTTP response with "Content-Type, application/json".
func (*GuluRet) RetResult(w http.ResponseWriter, r *http.Request, res *Result) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(res)
	if err != nil {
		logger.Error(err)

		return
	}

	w.Write(data)
}

// RetGzResult writes HTTP response with "Content-Type, application/json" and "Content-Encoding, gzip".
func (*GuluRet) RetGzResult(w http.ResponseWriter, r *http.Request, res *Result) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Encoding", "gzip")

	gz := gzip.NewWriter(w)
	err := json.NewEncoder(gz).Encode(res)
	if nil != err {
		logger.Error(err)

		return
	}

	err = gz.Close()
	if nil != err {
		logger.Error(err)

		return
	}
}

// RetJSON writes HTTP response with "Content-Type, application/json".
func (*GuluRet) RetJSON(w http.ResponseWriter, r *http.Request, res map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(res)
	if err != nil {
		logger.Error(err)

		return
	}

	w.Write(data)
}

// RetGzJSON writes HTTP response with "Content-Type, application/json" and "Content-Encoding, gzip".
func (*GuluRet) RetGzJSON(w http.ResponseWriter, r *http.Request, res map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Encoding", "gzip")

	gz := gzip.NewWriter(w)
	err := json.NewEncoder(gz).Encode(res)
	if nil != err {
		logger.Error(err)

		return
	}

	err = gz.Close()
	if nil != err {
		logger.Error(err)

		return
	}
}
