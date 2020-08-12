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
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"os/exec"
)

type SwapSizeWindow struct {
	name Name
	form *tview.InputField
}

func (w SwapSizeWindow) Name() Name {
	return w.name
}

func (w SwapSizeWindow) Primitive() tview.Primitive {
	return w.form
}

func (w SwapSizeWindow) New(app *tview.Application, pages *tview.Pages) Window {
	return SwapSizeWindow{
		name: SwapSizeWindowName,
		form: w.Build(app, pages).(*tview.InputField),
	}
}

func (w SwapSizeWindow) Build(app *tview.Application, pages *tview.Pages) tview.Primitive {
	field := tview.NewInputField()

	return field.SetLabel("Enter the size in MiB: ").
		SetPlaceholder("E.g. 2048").
		SetFieldWidth(9).
		SetAcceptanceFunc(tview.InputFieldInteger).
		SetDoneFunc(func(key tcell.Key) {
			if key == tcell.KeyEnter {
				pages.SwitchToPage(string(LoadingWindowName))
				cmd := exec.Command("sh", "/umbra/scripts/auto-install-swap.sh", field.GetText())

				go func() {
					defer app.Stop()
					_ = cmd.Start()
					_ = cmd.Wait()
				}()
			}
		})
}
