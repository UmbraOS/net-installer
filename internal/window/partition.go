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
)

type PartitionWindow struct {
	name  Name
	modal *tview.Modal
}

func (w PartitionWindow) Name() Name {
	return w.name
}

func (w PartitionWindow) Primitive() tview.Primitive {
	return w.modal
}

func (w PartitionWindow) New(app *tview.Application, pages *tview.Pages) Window {
	return PartitionWindow{
		name:  PartitionWindowName,
		modal: w.Build(app, pages).(*tview.Modal),
	}
}

func (w PartitionWindow) Build(_ *tview.Application, pages *tview.Pages) tview.Primitive {
	return tview.NewModal().
		SetText("Select a type of installation").
		AddButtons([]string{"Automatic (for VMs)", "Manual"}).
		SetDoneFunc(func(_ int, label string) {
			if label == "Automatic (for VMs)" {
				pages.SwitchToPage(string(SwapWindowName))
			}
		})
}
