/*
 *  Copyright 2021 The Trader Authors
 *
 *  Licensed under the GNU General Public License v3.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      <https:fsf.org/>
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */
package cerebro

import "testing"

func TestDefaultLogger_Info(t *testing.T) {
	cerebroLogger.Info("test")
}

func TestDefaultLogger_Infof(t *testing.T) {
	cerebroLogger.Infof("%s test", "info")
}

func TestDefaultLogger_Debug(t *testing.T) {
	cerebroLogger.Debug("debug test")
}

func TestDefaultLogger_Debugf(t *testing.T) {
	cerebroLogger.Debugf("%s test", "debug")
}

func TestDefaultLogger_Error(t *testing.T) {
	cerebroLogger.Error("error test")
}
