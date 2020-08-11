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
	"github.com/rivo/tview"
	"os/exec"
)

type SwapWindow struct {
	name Name
	form *tview.Modal
}

func (w SwapWindow) Name() Name {
	return w.name
}

func (w SwapWindow) Primitive() tview.Primitive {
	return w.form
}

func (w SwapWindow) New(app *tview.Application, pages *tview.Pages) Window {
	return SwapWindow{
		name: SwapWindowName,
		form: w.Build(app, pages).(*tview.Modal),
	}
}

func (w SwapWindow) Build(_ *tview.Application, pages *tview.Pages) tview.Primitive {
	return tview.NewModal().
		SetText("Do you want a swap partition?").
		AddButtons([]string{"Yes", "No"}).
		SetDoneFunc(func(_ int, label string) {
			if label == "Yes" {
				pages.SwitchToPage(string(SwapSizeWindowName))
			} else {
				exec.Command("sh", "/umbra/scripts/auto-install.sh")
			}
		})
}
