// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
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

package messages

import (
	"testing"
	"time"

	"github.com/smartystreets/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func timePtr(time time.Time) *time.Time { return &time }

func TestFromDownlinkMessage(t *testing.T) {
	for _, tc := range []struct {
		Name                    string
		DownlinkMessage         ttnpb.DownlinkMessage
		GatewayIDs              ttnpb.GatewayIdentifiers
		ExpectedDownlinkMessage DownlinkMessage
	}{
		{
			Name: "SampleDownlink",
			DownlinkMessage: ttnpb.DownlinkMessage{
				RawPayload: []byte("Ymxhamthc25kJ3M=="),
				EndDeviceIDs: &ttnpb.EndDeviceIdentifiers{
					DeviceID: "testdevice",
				},
				Settings: &ttnpb.DownlinkMessage_Scheduled{
					Scheduled: &ttnpb.TxSettings{
						DataRateIndex: 2,
						Frequency:     868500000,
						Downlink: &ttnpb.TxSettings_Downlink{
							AntennaIndex: 2,
						},
						Timestamp: 1553300787,
					},
				},
			},
			GatewayIDs: ttnpb.GatewayIdentifiers{GatewayID: "test-gateway"},
			ExpectedDownlinkMessage: DownlinkMessage{
				DevEUI:      "00-00-00-00-00-00-00-00",
				DeviceClass: 0,
				Pdu:         "596d7868616d74686332356b4a334d3d3d",
				RxDelay:     1,
				Rx1DR:       2,
				Rx1Freq:     868500000,
				RCtx:        2,
				Priority:    25,
				MuxTime:     1554300787.123456,
			},
		},
		{
			Name: "WithAbsoluteTime",
			DownlinkMessage: ttnpb.DownlinkMessage{
				RawPayload: []byte("Ymxhamthc25kJ3M=="),
				EndDeviceIDs: &ttnpb.EndDeviceIdentifiers{
					DeviceID: "testdevice",
				},
				Settings: &ttnpb.DownlinkMessage_Scheduled{
					Scheduled: &ttnpb.TxSettings{
						DataRateIndex: 2,
						Frequency:     869525000,
						Downlink: &ttnpb.TxSettings_Downlink{
							AntennaIndex: 2,
						},
					},
				},
			},
			GatewayIDs: ttnpb.GatewayIdentifiers{GatewayID: "test-gateway"},
			ExpectedDownlinkMessage: DownlinkMessage{
				DevEUI:      "00-00-00-00-00-00-00-00",
				DeviceClass: 0,
				Pdu:         "596d7868616d74686332356b4a334d3d3d",
				RxDelay:     1,
				Rx1DR:       2,
				Rx1Freq:     869525000,
				RCtx:        2,
				Priority:    25,
				MuxTime:     1554300787.123456,
			},
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			a := assertions.New(t)
			dnmsg := FromDownlinkMessage(tc.GatewayIDs, tc.DownlinkMessage.GetRawPayload(), tc.DownlinkMessage.GetScheduled(), 0, time.Unix(1554300787, 123456000), 0x00)
			dnmsg.XTime = 0
			if !a.So(dnmsg, should.Resemble, tc.ExpectedDownlinkMessage) {
				t.Fatalf("Invalid DownlinkMessage: %v", dnmsg)
			}
		})
	}
}

func TestToDownlinkMessage(t *testing.T) {
	for _, tc := range []struct {
		Name                    string
		DownlinkMessage         DownlinkMessage
		ExpectedDownlinkMessage ttnpb.DownlinkMessage
	}{
		{
			Name: "SampleDownlink",
			DownlinkMessage: DownlinkMessage{
				DeviceClass: 0,
				Pdu:         "Ymxhamthc25kJ3M==",
				RxDelay:     1,
				Rx1DR:       2,
				Rx1Freq:     868500000,
				RCtx:        2,
				Priority:    25,
				XTime:       1554300785,
			},
			ExpectedDownlinkMessage: ttnpb.DownlinkMessage{
				RawPayload: []byte("Ymxhamthc25kJ3M=="),
				Settings: &ttnpb.DownlinkMessage_Scheduled{
					Scheduled: &ttnpb.TxSettings{
						DataRateIndex: 2,
						Frequency:     868500000,
						Downlink: &ttnpb.TxSettings_Downlink{
							AntennaIndex: 2,
						},
						Timestamp: 1554300785,
					},
				},
			},
		},
		{
			Name: "WithAbsoluteTime",
			DownlinkMessage: DownlinkMessage{
				DeviceClass: 1,
				Pdu:         "Ymxhamthc25kJ3M==",
				RxDelay:     1,
				Rx1DR:       2,
				Rx1Freq:     869525000,
				RCtx:        2,
				Priority:    25,
			},
			ExpectedDownlinkMessage: ttnpb.DownlinkMessage{
				RawPayload: []byte("Ymxhamthc25kJ3M=="),
				Settings: &ttnpb.DownlinkMessage_Scheduled{
					Scheduled: &ttnpb.TxSettings{
						DataRateIndex: 2,
						Frequency:     869525000,
						Downlink: &ttnpb.TxSettings_Downlink{
							AntennaIndex: 2,
						},
					},
				},
			},
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			a := assertions.New(t)
			dlMesg := tc.DownlinkMessage.ToDownlinkMessage()
			if !a.So(dlMesg, should.Resemble, tc.ExpectedDownlinkMessage) {
				t.Fatalf("Invalid DownlinkMessage: %v", dlMesg)
			}
		})
	}
}
