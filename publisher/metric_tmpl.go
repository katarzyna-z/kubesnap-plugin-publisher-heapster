/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2016 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package publisher

const builtinMetricTemplate = `{
	"id":"!!",
	"name":"!!",
	"aliases":[
	],
	"labels":{
	},
	"subcontainers":[
	],
	"spec":{
		"creation_time":"/creation_time",
		"labels":{
		},
		"has_cpu":true,
		"cpu":{
			"limit":2,
			"max_limit":2,
			"mask":"0-1"
		},
		"has_memory":true,
		"memory":{
			"limit":18446744073709551615,
			"swap_limit":18446744073709551615
		},
		"has_network":true,
		"has_filesystem":true,
		"has_diskio":false,
		"has_custom_metrics":false,
		"image":"/image_name"
	},
	"stats":[
		{
			"timestamp":"!!",
			"cpu":{
				"usage":{
					"total":"/cgroups/cpu_stats/cpu_usage/total_usage",
					"user":"/cgroups/cpu_stats/cpu_usage/usage_in_usermode",
					"system":"/cgroups/cpu_stats/cpu_usage/usage_in_kernelmode"
				},
				"load_average":0
			},
			"diskio":{
			},
			"memory":{
				"usage":"/cgroups/memory_stats/usage/usage",
				"cache":"/cgroups/memory_stats/cache",
				"rss":"/cgroups/memory_stats/stats/rss",
				"working_set":0,
				"failcnt":0,
				"container_data":{
					"pgfault":"/cgroups/memory_stats/stats/pgfault",
					"pgmajfault":"/cgroups/memory_stats/stats/pgmajfault"
				},
				"hierarchical_data":{
					"pgfault":0,
					"pgmajfault":0
				}
			},
			"network":{
				"name":"",
				"rx_bytes":0,
				"rx_packets":0,
				"rx_errors":0,
				"rx_dropped":0,
				"tx_bytes":0,
				"tx_packets":0,
				"tx_errors":0,
				"tx_dropped":0,
				"interfaces":[
				],
				"tcp":{
					"Established":0,
					"SynSent":0,
					"SynRecv":0,
					"FinWait1":0,
					"FinWait2":0,
					"TimeWait":0,
					"Close":0,
					"CloseWait":0,
					"LastAck":0,
					"Listen":0,
					"Closing":0
				},
				"tcp6":{
					"Established":0,
					"SynSent":0,
					"SynRecv":0,
					"FinWait1":0,
					"FinWait2":0,
					"TimeWait":0,
					"Close":0,
					"CloseWait":0,
					"LastAck":0,
					"Listen":0,
					"Closing":0
				}
			},
			"filesystem":[
				{
					"device":"/dev/sdx1",
					"type":"vfs",
					"capacity":0,
					"usage":0,
					"base_usage":0,
					"available":0,
					"inodes_free":0,
					"reads_completed":0,
					"reads_merged":0,
					"sectors_read":0,
					"read_time":0,
					"writes_completed":0,
					"writes_merged":0,
					"sectors_written":0,
					"write_time":0,
					"io_in_progress":0,
					"io_time":0,
					"weighted_io_time":0
				}
			]
		}
	]
}
`
