// Copyright 2014, The Serviced Authors. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package host

import (
	"github.com/zenoss/glog"
	"github.com/zenoss/serviced/datastore/elastic"

	"testing"
)

func Test_AddFileMapping(t *testing.T) {

	esDriver := elastic.New("localhost", 9200, "twitter")
	esDriver.AddMappingFilie("host", "./host_mapping.json")
	err := esDriver.Initialize()
	glog.Infof("initialized: %v", err)
	if err != nil {
		t.Errorf("Error initializing db driver %v", err)
	}

}
