//
// =========================================================================
// Copyright (C) 2017 by Yunify, Inc...
// -------------------------------------------------------------------------
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this work except in compliance with the License.
// You may obtain a copy of the License in the LICENSE file, or at:
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// =========================================================================
//

package provider

import "fmt"

var initializerMap = make(map[string]Initializer)

// Register register new provider
func Register(name string, init Initializer) {
	initializerMap[name] = init
}

//New create new nic provider from config
func New(name string, conf map[string]interface{}) (NicProvider, error) {
	if init := initializerMap[name]; init != nil {
		return init(conf)
	}
	return nil, fmt.Errorf("Unsupported provider: %s", name)
}
