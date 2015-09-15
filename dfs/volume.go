// Copyright 2014 The Serviced Authors.
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

package dfs

import (
	"github.com/control-center/serviced/volume"
	"github.com/zenoss/glog"
)

func (dfs *DistributedFilesystem) GetVolume(serviceID string) (volume.Volume, error) {
	return dfs.getServiceVolume(dfs.fsType, dfs.volumesPath, serviceID)
}

func serviceVolumeGet(fsType volume.DriverType, root, serviceID string) (volume.Volume, error) {
	drv, err := volume.GetDriver(root)
	if err != nil {
		glog.Errorf("Could not find driver at root %s: %s", root, err)
		return nil, err
	}
	return drv.Get(serviceID)
}
