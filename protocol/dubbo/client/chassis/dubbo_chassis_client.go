/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package chassisclient

import (
	"context"
	"os"
	"sync"

	mesherCommon "github.com/go-chassis/mesher/common"
	dubboClient "github.com/go-chassis/mesher/protocol/dubbo/client"
	"github.com/go-chassis/mesher/protocol/dubbo/dubbo"
	"github.com/go-chassis/mesher/protocol/dubbo/proxy"
	"github.com/go-chassis/mesher/protocol/dubbo/utils"

	"github.com/ServiceComb/go-chassis/core/client"
	"github.com/ServiceComb/go-chassis/core/invocation"
	"github.com/ServiceComb/go-chassis/core/lager"
)

//Name is a constant
const Name = "dubbo"

func init() {
	client.InstallPlugin(Name, NewDubboChassisClient)
}

type dubboChassisClient struct {
	once     sync.Once
	opts     client.Options
	reqMutex sync.Mutex
}

//NewDubboChassisClient create new client
func NewDubboChassisClient(options client.Options) client.ProtocolClient {

	rc := &dubboChassisClient{
		once: sync.Once{},
		opts: options,
	}
	return client.ProtocolClient(rc)
}

func (c *dubboChassisClient) String() string {
	return "highway_client"
}

func (c *dubboChassisClient) Call(ctx context.Context, addr string, inv *invocation.Invocation, rsp interface{}) error {
	dubboReq := inv.Args.(*dubbo.Request)

	endPoint := addr
	if endPoint == dubboproxy.DubboListenAddr {
		endPoint = os.Getenv(mesherCommon.EnvSpecificAddr)
	}
	if endPoint == "" {
		return &util.BaseError{" The endpoint is empty"}
	}
	lager.Logger.Info("Dubbo invoke endPont: " + endPoint)
	dubboCli, err := dubboClient.CachedClients.GetClient(endPoint)
	if err != nil {
		lager.Logger.Error("Invalid Request addr ="+endPoint, err)
		return err
	}

	dubboRsp, errSnd := dubboCli.Send(dubboReq)
	if errSnd != nil {
		lager.Logger.Error("Dubbo server exception:", errSnd)
		return errSnd
	}
	resp := rsp.(*dubboClient.WrapResponse)
	resp.Resp = dubboRsp
	return nil
}
