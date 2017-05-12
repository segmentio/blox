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
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GeneralAPIsTestSuite struct {
	suite.Suite
	generalAPIs          GeneralAPIs
	responseHeaderJSON   http.Header
	responseHeaderStream http.Header

	// We need a router because some of the apis use mux.Vars() which uses the URL
	// parameters parsed and stored in a global map in the global context by the router.
	router *mux.Router
}

func (suite *GeneralAPIsTestSuite) SetupTest() {
	suite.generalAPIs = NewGeneralAPIs()
	suite.router = suite.getRouter()
}

// TODO - Add the following test cases once implementation is in place
// * arn validation fails on getTask
// * streaming api

func TestGeneralAPIsTestSuite(t *testing.T) {
	suite.Run(t, new(GeneralAPIsTestSuite))
}

func (suite *GeneralAPIsTestSuite) TestPing() {
	request, err := http.NewRequest("GET", "/v1/ping", nil)
	assert.Nil(suite.T(), err, "Unexpected error creating ping request")

	responseRecorder := httptest.NewRecorder()
	suite.router.ServeHTTP(responseRecorder, request)

	header := responseRecorder.Header()
	assert.NotNil(suite.T(), header, "Unexpected empty header in ping response")
	expectedHeader := http.Header{"Content-Type": []string{"application/json; charset=UTF-8"}}
	assert.Equal(suite.T(), expectedHeader, header, "Http header in ping response is  invalid")
	assert.Equal(suite.T(), http.StatusOK, responseRecorder.Code, "Http response status in ping response is invalid")
}

func (suite *GeneralAPIsTestSuite) getRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	s := r.Path("/v1").Subrouter()

	s.Path(getPingPath).
		Methods("GET").
		HandlerFunc(suite.generalAPIs.Ping)

	return s
}
