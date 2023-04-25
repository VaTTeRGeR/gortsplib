package rtpmpeg2audio

import (
	"testing"

	"github.com/pion/rtp"
	"github.com/stretchr/testify/require"
)

var cases = []struct {
	name   string
	frames [][]byte
	pkt    *rtp.Packet
}{
	{
		"single",
		[][]byte{
			{
				0xff, 0xfb, 0x14, 0x64, 0x00, 0x0f, 0xf0, 0x00,
				0x00, 0x69, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00,
				0x0d, 0x20, 0x00, 0x00, 0x01, 0x00, 0x00, 0x01,
				0xa4, 0x00, 0x00, 0x00, 0x20, 0x00, 0x00, 0x34,
				0x80, 0x00, 0x00, 0x04, 0x4c, 0x41, 0x4d, 0x45,
				0x33, 0x2e, 0x31, 0x30, 0x30, 0x55, 0x55, 0x55,
				0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55,
				0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55,
				0x55, 0xc0, 0x65, 0xf4, 0xa0, 0x31, 0x8f, 0xce,
				0x8d, 0x46, 0xfc, 0x8c, 0x73, 0xb9, 0x34, 0x3e,
				0xb5, 0x03, 0x39, 0xc0, 0x04, 0x01, 0x98, 0x44,
				0x38, 0xe0, 0x98, 0x10, 0x9b, 0xa8, 0x0f, 0xa8,
			},
		},
		&rtp.Packet{
			Header: rtp.Header{
				Version:        2,
				Marker:         true,
				PayloadType:    14,
				SequenceNumber: 17645,
				Timestamp:      2289526357,
				SSRC:           0x9dbb7812,
			},
			Payload: []byte{
				0x00, 0x00, 0x00, 0x00,
				0xff, 0xfb, 0x14, 0x64, 0x00, 0x0f, 0xf0, 0x00,
				0x00, 0x69, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00,
				0x0d, 0x20, 0x00, 0x00, 0x01, 0x00, 0x00, 0x01,
				0xa4, 0x00, 0x00, 0x00, 0x20, 0x00, 0x00, 0x34,
				0x80, 0x00, 0x00, 0x04, 0x4c, 0x41, 0x4d, 0x45,
				0x33, 0x2e, 0x31, 0x30, 0x30, 0x55, 0x55, 0x55,
				0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55,
				0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55,
				0x55, 0xc0, 0x65, 0xf4, 0xa0, 0x31, 0x8f, 0xce,
				0x8d, 0x46, 0xfc, 0x8c, 0x73, 0xb9, 0x34, 0x3e,
				0xb5, 0x03, 0x39, 0xc0, 0x04, 0x01, 0x98, 0x44,
				0x38, 0xe0, 0x98, 0x10, 0x9b, 0xa8, 0x0f, 0xa8,
			},
		},
	},
	{
		"aggregated",
		[][]byte{
			{
				0xff, 0xfb, 0x14, 0x64, 0x00, 0x0f, 0xf0, 0x00,
				0x00, 0x69, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00,
				0x0d, 0x20, 0x00, 0x00, 0x01, 0x00, 0x00, 0x01,
				0xa4, 0x00, 0x00, 0x00, 0x20, 0x00, 0x00, 0x34,
				0x80, 0x00, 0x00, 0x04, 0x4c, 0x41, 0x4d, 0x45,
				0x33, 0x2e, 0x31, 0x30, 0x30, 0x55, 0x55, 0x55,
				0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55,
				0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55,
				0x55, 0xc0, 0x65, 0xf4, 0xa0, 0x31, 0x8f, 0xce,
				0x8d, 0x46, 0xfc, 0x8c, 0x73, 0xb9, 0x34, 0x3e,
				0xb5, 0x03, 0x39, 0xc0, 0x04, 0x01, 0x98, 0x44,
				0x38, 0xe0, 0x98, 0x10, 0x9b, 0xa8, 0x0f, 0xa8,
			},
			{
				0xff, 0xfb, 0x14, 0x64, 0x1e, 0x0f, 0xf0, 0x00,
				0x00, 0x69, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00,
				0x0d, 0x20, 0x00, 0x00, 0x01, 0x00, 0x00, 0x01,
				0xa4, 0x00, 0x00, 0x00, 0x20, 0x00, 0x00, 0x34,
				0x80, 0x00, 0x00, 0x04, 0xe6, 0x50, 0x10, 0x01,
				0xca, 0x13, 0x94, 0x27, 0x4a, 0x4a, 0x64, 0xce,
				0x07, 0xc2, 0x2f, 0x59, 0xc0, 0x19, 0x04, 0x05,
				0xdf, 0xe7, 0xce, 0x65, 0x24, 0xed, 0xa4, 0xe3,
				0xff, 0xc9, 0x00, 0x00, 0x05, 0x5f, 0x4a, 0x04,
				0x0e, 0xc4, 0x24, 0xfd, 0x5e, 0x4a, 0x35, 0x72,
				0x21, 0x27, 0x31, 0x08, 0x47, 0x18, 0x00, 0x06,
				0xc4, 0x02, 0x72, 0x81, 0x89, 0xc3, 0xe4, 0x0a,
			},
		},
		&rtp.Packet{
			Header: rtp.Header{
				Version:        2,
				Marker:         true,
				PayloadType:    14,
				SequenceNumber: 17645,
				Timestamp:      2289526357,
				SSRC:           0x9dbb7812,
			},
			Payload: []byte{
				0x00, 0x00, 0x00, 0x00,
				0xff, 0xfb, 0x14, 0x64, 0x00, 0x0f, 0xf0, 0x00,
				0x00, 0x69, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00,
				0x0d, 0x20, 0x00, 0x00, 0x01, 0x00, 0x00, 0x01,
				0xa4, 0x00, 0x00, 0x00, 0x20, 0x00, 0x00, 0x34,
				0x80, 0x00, 0x00, 0x04, 0x4c, 0x41, 0x4d, 0x45,
				0x33, 0x2e, 0x31, 0x30, 0x30, 0x55, 0x55, 0x55,
				0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55,
				0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55,
				0x55, 0xc0, 0x65, 0xf4, 0xa0, 0x31, 0x8f, 0xce,
				0x8d, 0x46, 0xfc, 0x8c, 0x73, 0xb9, 0x34, 0x3e,
				0xb5, 0x03, 0x39, 0xc0, 0x04, 0x01, 0x98, 0x44,
				0x38, 0xe0, 0x98, 0x10, 0x9b, 0xa8, 0x0f, 0xa8,
				0xff, 0xfb, 0x14, 0x64, 0x1e, 0x0f, 0xf0, 0x00,
				0x00, 0x69, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00,
				0x0d, 0x20, 0x00, 0x00, 0x01, 0x00, 0x00, 0x01,
				0xa4, 0x00, 0x00, 0x00, 0x20, 0x00, 0x00, 0x34,
				0x80, 0x00, 0x00, 0x04, 0xe6, 0x50, 0x10, 0x01,
				0xca, 0x13, 0x94, 0x27, 0x4a, 0x4a, 0x64, 0xce,
				0x07, 0xc2, 0x2f, 0x59, 0xc0, 0x19, 0x04, 0x05,
				0xdf, 0xe7, 0xce, 0x65, 0x24, 0xed, 0xa4, 0xe3,
				0xff, 0xc9, 0x00, 0x00, 0x05, 0x5f, 0x4a, 0x04,
				0x0e, 0xc4, 0x24, 0xfd, 0x5e, 0x4a, 0x35, 0x72,
				0x21, 0x27, 0x31, 0x08, 0x47, 0x18, 0x00, 0x06,
				0xc4, 0x02, 0x72, 0x81, 0x89, 0xc3, 0xe4, 0x0a,
			},
		},
	},
}

func TestEncode(t *testing.T) {
	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			e := &Encoder{
				SSRC: func() *uint32 {
					v := uint32(0x9dbb7812)
					return &v
				}(),
				InitialSequenceNumber: func() *uint16 {
					v := uint16(0x44ed)
					return &v
				}(),
				InitialTimestamp: func() *uint32 {
					v := uint32(0x88776655)
					return &v
				}(),
			}
			e.Init()

			pkts, err := e.Encode(ca.frames, 0)
			require.NoError(t, err)
			require.Equal(t, ca.pkt, pkts[0])
		})
	}
}

func TestEncodeRandomInitialState(t *testing.T) {
	e := &Encoder{}
	e.Init()
	require.NotEqual(t, nil, e.SSRC)
	require.NotEqual(t, nil, e.InitialSequenceNumber)
	require.NotEqual(t, nil, e.InitialTimestamp)
}