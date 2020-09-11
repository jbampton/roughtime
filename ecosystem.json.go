// Code generated by "go run ./ecosystem_json_go_builder". DO NOT EDIT.
// source: ecosystem.json
package roughtime

import (
	"roughtime.googlesource.com/roughtime.git/go/config"
)

var Ecosystem = []config.Server{
	{
		Name:          "Chainpoint-Roughtime",
		PublicKeyType: "ed25519",
		PublicKey:     []byte{109, 180, 254, 68, 244, 187, 204, 165, 250, 195, 189, 108, 176, 248, 155, 206, 108, 22, 169, 79, 95, 125, 21, 121, 162, 61, 142, 173, 235, 18, 154, 17},
		Addresses: []config.ServerAddress{
			{
				Protocol: "udp",
				Address:  "roughtime.chainpoint.org:2002",
			},
		},
	},
	{
		Name:          "Cloudflare-Roughtime",
		PublicKeyType: "ed25519",
		PublicKey:     []byte{128, 62, 183, 133, 40, 247, 73, 196, 190, 194, 227, 158, 26, 187, 155, 94, 90, 183, 228, 221, 92, 228, 182, 242, 253, 47, 147, 236, 195, 83, 143, 26},
		Addresses: []config.ServerAddress{
			{
				Protocol: "udp",
				Address:  "roughtime.cloudflare.com:2002",
			},
		},
	},
	{
		Name:          "Google-Sandbox-Roughtime",
		PublicKeyType: "ed25519",
		PublicKey:     []byte{122, 211, 218, 104, 140, 92, 4, 198, 53, 161, 71, 134, 167, 11, 207, 48, 34, 76, 194, 84, 85, 55, 27, 249, 212, 162, 191, 182, 75, 104, 37, 52},
		Addresses: []config.ServerAddress{
			{
				Protocol: "udp",
				Address:  "roughtime.sandbox.google.com:2002",
			},
		},
	},
	{
		Name:          "int08h-Roughtime",
		PublicKeyType: "ed25519",
		PublicKey:     []byte{1, 110, 110, 2, 132, 210, 76, 55, 198, 228, 215, 216, 213, 180, 225, 211, 193, 148, 156, 234, 165, 69, 191, 135, 86, 22, 201, 220, 224, 201, 190, 193},
		Addresses: []config.ServerAddress{
			{
				Protocol: "udp",
				Address:  "roughtime.int08h.com:2002",
			},
		},
	},
	{
		Name:          "ticktock",
		PublicKeyType: "ed25519",
		PublicKey:     []byte{114, 63, 6, 178, 35, 101, 70, 74, 162, 12, 73, 64, 120, 211, 18, 4, 19, 48, 172, 9, 117, 230, 22, 15, 129, 126, 116, 248, 101, 151, 254, 80},
		Addresses: []config.ServerAddress{
			{
				Protocol: "udp",
				Address:  "ticktock.mixmin.net:5333",
			},
		},
	},
	{
		Name:          "time.0xt.ca",
		PublicKeyType: "ed25519",
		PublicKey:     []byte{136, 21, 99, 198, 15, 245, 143, 188, 181, 250, 68, 20, 76, 22, 29, 77, 166, 241, 10, 154, 94, 177, 79, 244, 236, 62, 15, 48, 50, 100, 217, 96},
		Addresses: []config.ServerAddress{
			{
				Protocol: "udp",
				Address:  "time.0xt.ca:2002",
			},
		},
	},
}
