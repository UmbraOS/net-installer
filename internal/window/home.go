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

import (
	"fmt"
	"github.com/rivo/tview"
)

type HomeWindow struct {
	name  Name
	modal *tview.Modal
}

func (w HomeWindow) Name() Name {
	return w.name
}

func (w HomeWindow) Modal() *tview.Modal {
	return w.modal
}

func (w HomeWindow) New(app *tview.Application, pages *tview.Pages) Window {
	return HomeWindow{
		name:  HomeWindowName,
		modal: w.Build(app, pages),
	}
}

func (w HomeWindow) Build(app *tview.Application, pages *tview.Pages) *tview.Modal {
	return tview.NewModal().
		SetText(fmt.Sprintf("Welcome to UmbraOS net installer!")).
		AddButtons([]string{"Start", "Cancel"}).
		SetDoneFunc(func(_ int, label string) {
			if label == "Start" {
				pages.SwitchToPage(string(PartitionWindowName))
			} else {
				app.Stop()
			}
		})
}
