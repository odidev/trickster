/*
 * Copyright 2018 Comcast Cable Communications Management, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package options

import (
	"testing"

	"github.com/BurntSushi/toml"
)

func TestNewOptions(t *testing.T) {
	o := NewOptions()
	o.CollectorUser = "trickster"
	o2 := o.Clone()
	if o2.CollectorUser != "trickster" {
		t.Error("clone failed")
	}
}

func TestProcessTracingConfigs(t *testing.T) {

	ProcessTracingOptions(nil, nil)

	o := NewOptions()
	o.SampleRate = 0

	mo := map[string]*Options{
		"test": o,
	}

	ProcessTracingOptions(mo, &toml.MetaData{})

	if int(o.SampleRate) != 1 {
		t.Errorf("expected 1 got %d", int(o.SampleRate))
	}

}
