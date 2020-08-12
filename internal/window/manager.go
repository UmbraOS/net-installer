/*
 * Copyright © 2020 github.com/UmbraOS
 *
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package window

import "github.com/rivo/tview"

type Manager struct {
	pages *tview.Pages
}

func NewManager(app *tview.Application) *Manager {
	pages := tview.NewPages()
	windows := []Window{
		HomeWindow{}.New(app, pages),
		PartitionWindow{}.New(app, pages),
		SwapWindow{}.New(app, pages),
		SwapSizeWindow{}.New(app, pages),
		LoadingWindow{}.New(app, pages),
		ConfigurationWindow{}.New(app, pages),
	}

	for i, w := range windows {
		pages.AddPage(string(w.Name()), w.Primitive(), true, i == 0)
	}

	return &Manager{
		pages,
	}
}

func (w *Manager) Pages() *tview.Pages {
	return w.pages
}
