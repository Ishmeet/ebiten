// Copyright 2013 The Ebiten Authors
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

// Code generated by gen.go using 'go generate'. DO NOT EDIT.

package keyboard

import (
	"image"
)

var keyboardKeyRects = map[string]image.Rectangle{}

func init() {
	keyboardKeyRects["'"] = image.Rect(208, 36, 224, 54)
	keyboardKeyRects[","] = image.Rect(168, 54, 184, 72)
	keyboardKeyRects["-"] = image.Rect(192, 0, 208, 18)
	keyboardKeyRects["."] = image.Rect(184, 54, 200, 72)
	keyboardKeyRects["/"] = image.Rect(200, 54, 216, 72)
	keyboardKeyRects["0"] = image.Rect(176, 0, 192, 18)
	keyboardKeyRects["1"] = image.Rect(32, 0, 48, 18)
	keyboardKeyRects["2"] = image.Rect(48, 0, 64, 18)
	keyboardKeyRects["3"] = image.Rect(64, 0, 80, 18)
	keyboardKeyRects["4"] = image.Rect(80, 0, 96, 18)
	keyboardKeyRects["5"] = image.Rect(96, 0, 112, 18)
	keyboardKeyRects["6"] = image.Rect(112, 0, 128, 18)
	keyboardKeyRects["7"] = image.Rect(128, 0, 144, 18)
	keyboardKeyRects["8"] = image.Rect(144, 0, 160, 18)
	keyboardKeyRects["9"] = image.Rect(160, 0, 176, 18)
	keyboardKeyRects[";"] = image.Rect(192, 36, 208, 54)
	keyboardKeyRects["="] = image.Rect(208, 0, 224, 18)
	keyboardKeyRects["A"] = image.Rect(48, 36, 64, 54)
	keyboardKeyRects["Alt"] = image.Rect(64, 72, 96, 90)
	keyboardKeyRects["B"] = image.Rect(120, 54, 136, 72)
	keyboardKeyRects["BS"] = image.Rect(232, 18, 272, 36)
	keyboardKeyRects["C"] = image.Rect(88, 54, 104, 72)
	keyboardKeyRects["Ctrl"] = image.Rect(0, 36, 48, 54)
	keyboardKeyRects["D"] = image.Rect(80, 36, 96, 54)
	keyboardKeyRects["Down"] = image.Rect(48, 126, 96, 144)
	keyboardKeyRects["E"] = image.Rect(72, 18, 88, 36)
	keyboardKeyRects["Enter"] = image.Rect(224, 36, 272, 54)
	keyboardKeyRects["Esc"] = image.Rect(0, 0, 32, 18)
	keyboardKeyRects["F"] = image.Rect(96, 36, 112, 54)
	keyboardKeyRects["G"] = image.Rect(112, 36, 128, 54)
	keyboardKeyRects["H"] = image.Rect(128, 36, 144, 54)
	keyboardKeyRects["I"] = image.Rect(152, 18, 168, 36)
	keyboardKeyRects["J"] = image.Rect(144, 36, 160, 54)
	keyboardKeyRects["K"] = image.Rect(160, 36, 176, 54)
	keyboardKeyRects["L"] = image.Rect(176, 36, 192, 54)
	keyboardKeyRects["Left"] = image.Rect(0, 126, 48, 144)
	keyboardKeyRects["M"] = image.Rect(152, 54, 168, 72)
	keyboardKeyRects["N"] = image.Rect(136, 54, 152, 72)
	keyboardKeyRects["O"] = image.Rect(168, 18, 184, 36)
	keyboardKeyRects["P"] = image.Rect(184, 18, 200, 36)
	keyboardKeyRects["Q"] = image.Rect(40, 18, 56, 36)
	keyboardKeyRects["R"] = image.Rect(88, 18, 104, 36)
	keyboardKeyRects["Right"] = image.Rect(96, 126, 144, 144)
	keyboardKeyRects["S"] = image.Rect(64, 36, 80, 54)
	keyboardKeyRects["Shift"] = image.Rect(0, 54, 56, 72)
	keyboardKeyRects["Space"] = image.Rect(96, 72, 176, 90)
	keyboardKeyRects["T"] = image.Rect(104, 18, 120, 36)
	keyboardKeyRects["Tab"] = image.Rect(0, 18, 40, 36)
	keyboardKeyRects["U"] = image.Rect(136, 18, 152, 36)
	keyboardKeyRects["Up"] = image.Rect(48, 108, 96, 126)
	keyboardKeyRects["V"] = image.Rect(104, 54, 120, 72)
	keyboardKeyRects["W"] = image.Rect(56, 18, 72, 36)
	keyboardKeyRects["X"] = image.Rect(72, 54, 88, 72)
	keyboardKeyRects["Y"] = image.Rect(120, 18, 136, 36)
	keyboardKeyRects["Z"] = image.Rect(56, 54, 72, 72)
	keyboardKeyRects["["] = image.Rect(200, 18, 216, 36)
	keyboardKeyRects["\\"] = image.Rect(224, 0, 240, 18)
	keyboardKeyRects["]"] = image.Rect(216, 18, 232, 36)
	keyboardKeyRects["`"] = image.Rect(240, 0, 256, 18)
}

func KeyRect(name string) (image.Rectangle, bool) {
	r, ok := keyboardKeyRects[name]
	return r, ok
}
