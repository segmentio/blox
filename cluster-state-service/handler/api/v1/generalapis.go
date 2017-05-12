// Copyright 2016-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package v1

import (
	"net/http"

	"github.com/blox/blox/cluster-state-service/handler/store"
)

// TaskAPIs encapsulates the backend datastore with which the task APIs interact
type GeneralAPIs struct {
	taskStore store.TaskStore
}

// NewTaskAPIs initializes the TaskAPIs struct
func NewGeneralAPIs() GeneralAPIs {
	return GeneralAPIs{}
}

// Ping is used to perform server health checks
func (api GeneralAPIs) Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentTypeKey, contentTypeJSON)
	w.WriteHeader(http.StatusOK)
}
