/*
** Copyright [2013-2017] [Megam Systems]
**
** Licensed under the Apache License, Version 2.0 (the "License");
** you may not use this file except in compliance with the License.
** You may obtain a copy of the License at
**
** http://www.apache.org/licenses/LICENSE-2.0
**
** Unless required by applicable law or agreed to in writing, software
** distributed under the License is distributed on an "AS IS" BASIS,
** WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
** See the License for the specific language governing permissions and
** limitations under the License.
 */

package one

import (
	"errors"
	"github.com/megamsys/libgo/hc"
	"strings"
)

func init() {
	hc.AddChecker("vertice:one", healthCheck)
}

func healthCheck() (interface{}, error) {
	if !strings.Contains(mainOneProvisioner.String(), "ready") {
		return nil, hc.ErrDisabledComponent
	}
	nodes, err := mainOneProvisioner.Cluster().Nodes()
	if err != nil {
		return nil, err
	}
	if len(nodes) < 1 {
		return nil, errors.New("no nodes available for running vm")
	}
	return "one " + nodes[0].Address + " ready", nil
}
