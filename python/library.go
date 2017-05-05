// Copyright 2017 Google Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package python

// This file contains the module types for building Python library.

import (
	"github.com/google/blueprint"

	"android/soong/android"
)

func init() {
	android.RegisterModuleType("python_library_host", PythonLibraryHostFactory)
}

type PythonLibrary struct {
	pythonBaseModule
}

var _ PythonSubModule = (*PythonLibrary)(nil)

func PythonLibraryHostFactory() (blueprint.Module, []interface{}) {
	module := &PythonLibrary{}

	return InitPythonBaseModule(&module.pythonBaseModule, module, android.HostSupportedNoCross)
}

func (p *PythonLibrary) GeneratePythonAndroidMk() (ret android.AndroidMkData, err error) {
	return
}
