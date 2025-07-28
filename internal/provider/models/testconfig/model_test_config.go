// Copyright © 2025 Scott Wiederhold <s.e.wiederhold@gmail.com>
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package testconfig

type ModelTestConfig struct {
	Name         string
	Model        string
	Resource     string
	Data         string
	ModelOnly    bool
	Dependencies map[string]*ModelTestConfig
	TestPre      string
}

func (c ModelTestConfig) GetModelListConfig() string {
	if c.Model == "\n" {
		return "null"
	} else {
		return "[" + c.Model + "]"
	}
}

func (c ModelTestConfig) GetModelConfig() string {
	if c.Model == "{}" {
		return "null"
	} else {
		return c.Model
	}
}

func (c ModelTestConfig) GetDataConfig() string {
	return c.GetResourceConfig() + c.Data
}

func (c ModelTestConfig) GetResourceConfig() string {
	return c.getPrequisites() + c.Resource
}

func (c *ModelTestConfig) getPrequisites() string {
	referencesTo := make(map[string]*ModelTestConfig)
	preReqs := c.TestPre
	for _, v := range referencesTo {
		if !v.ModelOnly {
			preReqs = preReqs + v.Resource
		}
	}
	return preReqs
}
