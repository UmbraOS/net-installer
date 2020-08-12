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
	"github.com/umbraos/net-installer/internal/configuration"
	"os/exec"
)

type ConfigurationWindow struct {
	name     Name
	dropDown *tview.Form
}

func (w ConfigurationWindow) Name() Name {
	return w.name
}

func (w ConfigurationWindow) Primitive() tview.Primitive {
	return w.dropDown
}

func (w ConfigurationWindow) New(app *tview.Application, pages *tview.Pages) Window {
	return ConfigurationWindow{
		name:     ConfigurationWindowName,
		dropDown: w.Build(app, pages).(*tview.Form),
	}
}

func (w ConfigurationWindow) Build(app *tview.Application, _ *tview.Pages) tview.Primitive {
	form := tview.NewForm()
	form.SetBorder(true).SetTitle("Configuration (use TAB to navigate)")

	return form.
		AddDropDown("Language", configuration.GetLanguages(), 0, nil).
		AddDropDown("Keyboard layout", configuration.GetKeyboardLayouts(), 0, nil).
		AddInputField("Time zone (Zone/City)", "e.g. Europe/Lisbon", 18, nil, nil).
		AddButton("Save", func() {
			args := []string{"/umbra/scripts/setup-config.sh"}
			nItem := 3

			for i := 0; i < nItem - 1; i++ {
				_, value := form.GetFormItem(i).(*tview.DropDown).GetCurrentOption()
				args = append(args, value)
			}

			args = append(args, form.GetFormItem(nItem - 1).(*tview.InputField).GetText())

			_ = exec.Command("sh", args...).Start()
			app.Stop()
		})
}
