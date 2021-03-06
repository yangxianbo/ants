// Copyright 2018 github.com/xiaoenai. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	tp "github.com/henrylee2cn/teleport"
	"github.com/henrylee2cn/teleport/socket"
)

// ProxyHooks proxy hooks
type ProxyHooks interface {
	// BeforePull is called before pulling the internal service.
	BeforePull(uri string, args interface{}, reply interface{}, setting ...socket.PacketSetting) *tp.Rerror
	// BeforePush is called before pushing the internal service.
	BeforePush(uri string, args interface{}, setting ...socket.PacketSetting) *tp.Rerror
}

// DefaultProxyHooks creates a new default ProxyHooks object.
func DefaultProxyHooks() ProxyHooks {
	return new(defProxyHooks)
}

type defProxyHooks struct{}

func (d *defProxyHooks) BeforePull(uri string, args interface{}, reply interface{}, setting ...socket.PacketSetting) *tp.Rerror {
	return nil
}

func (d *defProxyHooks) BeforePush(uri string, args interface{}, setting ...socket.PacketSetting) *tp.Rerror {
	return nil
}
