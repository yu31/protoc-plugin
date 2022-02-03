package tests

import (
	"encoding/json"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/yu31/protoc-plugin/xgo/tests/gojsontest"
)

var jsonStringModel2 = []byte(`
{"type_double1":64.11,"type_double2":64.22,"type_double3":64.33,"type_double4":64.44,"type_double5":64.55,"type_float":32.11,"type_int32":321,"type_int64":641,"type_uint32":322,"type_uint64":642,"type_sint32":323,"type_sint64":643,"type_fixed32":324,"type_fixed64":644,"type_sfixed32":325,"type_sfixed64":646,"type_bool1":true,"type_bool2":true,"type_string1":"TypeString1","type_string2":"TypeString2","type_string3":"TypeString3","type_string4":"TypeString4","type_string5":"TypeString5","type_bytes":"VHlwZUJ5dGVz","type_embed_message":{"age1":"201","age2":"202","age3":"203"},"type_stand_message":{"name1":"q21","name2":"q22","name3":"q23"},"type_embed_enum":1,"type_stand_enum":1,"type_external_enum":4,"type_external_message":{"ip1":"127.0.1.1","ip2":"127.0.1.2","ip3":"127.0.1.3"},"array_double":[64.11,64.12,64.13,64.15,64.15],"array_float":[32.11,32.12,32.13,32.15,32.15],"array_int32":[3211,3212,3213,3214,3215],"array_int64":[6411,6412,6413,6414,6415],"array_uint32":[3221,3222,3223,3224,3225],"array_uint64":[6421,6422,6423,6424,6425],"array_sint32":[3231,3232,3233,3234,3235],"array_sint64":[6431,6432,6433,6434,6435],"array_fixed32":[3241,3242,3243,3244,3245],"array_fixed64":[6441,6442,6443,6444,6445],"array_sfixed32":[3251,3252,3253,3254,3255],"array_sfixed64":[6451,6452,6453,6454,6455],"array_bool":[true,false,true,false,true],"array_string":["s_arr_1","s_arr_2","s_arr_3","s_arr_4","s_arr_5"],"array_bytes":["Yl9hcnJfMQ==","Yl9hcnJfMg==","Yl9hcnJfMw==","Yl9hcnJfNA==","Yl9hcnJfNQ=="],"array_embed_message":[{"age1":"311","age2":"312","age3":"313"}],"array_stand_message":[{"name1":"q11","name2":"q12","name3":"q13"}],"array_external_message":[{"ip1":"127.0.1.1","ip2":"127.0.1.2","ip3":"127.0.1.3"}],"array_embed_enum":[0,1,2,3,4,5],"array_stand_enum":[0,1,2,3,4,5],"array_external_enum":[0,1,2,3,4,5,6],"map_int32_double":{"3211":64.11},"map_int32_float":{"3221":32.11},"map_int32_int32":{"3231":31},"map_int32_int64":{"3241":41},"map_int32_uint32":{"3251":51},"map_int32_uint64":{"3261":61},"map_int32_sint32":{"3271":71},"map_int32_sint64":{"3281":81},"map_int32_fixed32":{"3291":91},"map_int32_fixed64":{"32101":101},"map_int32_sfixed32":{"32111":111},"map_int32_sfixed64":{"32121":121},"map_int32_bool":{"32131":true},"map_int32_string":{"32141":"mvs11"},"map_int32_bytes":{"32151":"bXZiMTE="},"map_int32_embed_message":{"32161":{"age1":"311","age2":"312","age3":"313"}},"map_int32_stand_message":{"32171":{"name1":"q11","name2":"q12","name3":"q13"}},"map_int32_embed_enum":{"32181":0},"map_int32_stand_enum":{"32191":0},"map_int64_int32":{"6411":11},"map_uint32_int32":{"3221":21},"map_uint64_int32":{"6431":31},"map_sint32_int32":{"3241":41},"map_sint64_int32":{"6451":51},"map_fixed32_int32":{"3261":61},"map_fixed64_int32":{"6471":71},"map_sfixed32_int32":{"3281":81},"map_sfixed64_int32":{"6491":91},"map_string_int32":{"mks11":100},"map_string_string":{"mks21":"mvs21"},"map_string_embed_message":{"mks31":{"age1":"411","age2":"412","age3":"413"}},"map_string_stand_message":{"mks41":{"name1":"q111","name2":"q112","name3":"q113"}},"map_string_external_message":{"mks51":{"ip1":"127.0.1.1","ip2":"127.0.1.2","ip3":"127.0.1.3"}},"map_string_embed_enum":{"mks61":0},"map_string_stand_enum":{"mks71":0},"map_string_external_enum":{"mks81":0}}
`)

var jsonStringModel3 = []byte(`
{"t_string1":"Hello 1","t_string2":"Hello 2","t_string3":"Hello 3","t_string4":"Hello 4","t_string5":"Hello 5","t_string6":"Hello 6","t_string7":"Hello 7","t_string8":"Hello 8","t_string9":"Hello 9","t_string10":"Hello 10","t_int32":1,"t_int64":2,"t_uint32":3,"t_uint64":4,"t_sint32":5,"t_sint64":6,"t_sfixed32":7,"t_sfixed64":8,"t_fixed32":9,"t_fixed64":10,"t_float":11,"t_double":12,"t_bool":true}
`)

func Benchmark_GoJSON_Marshal_1(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := model2.MarshalJSON()
			if err != nil {
				b.Fatal("gojson marshal error", err)
			}

		}
	})
}

func Benchmark_StdJSON_Marshal_1(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := json.Marshal(&model2Type)
			if err != nil {
				b.Fatal("standard marshal error", err)
			}

		}
	})
}

func Benchmark_JSONIter_Marshal_1(b *testing.B) {
	_json := jsoniter.ConfigCompatibleWithStandardLibrary

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := _json.Marshal(&model2Type)
			if err != nil {
				b.Fatal("standard marshal error", err)
			}

		}
	})
}

func Benchmark_GoJSON_Unmarshal_1(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var model gojsontest.Model2
			err := model.UnmarshalJSON(jsonStringModel2)
			if err != nil {
				b.Fatal("gojson unmarshal error:", err)
			}
		}
	})
}

func Benchmark_StdJSON_Unmarshal_1(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var model Model2Type
			err := json.Unmarshal(jsonStringModel2, &model)
			if err != nil {
				b.Fatal("standard unmarshal error:", err)
			}
		}
	})

}

func Benchmark_JSONIter_Unmarshal_1(b *testing.B) {
	_json := jsoniter.ConfigCompatibleWithStandardLibrary

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var model Model2Type
			err := _json.Unmarshal(jsonStringModel2, &model)
			if err != nil {
				b.Fatal("standard unmarshal error:", err)
			}
		}
	})

}

func Benchmark_GoJSON_Unmarshal_2(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var model gojsontest.Model3
			err := model.UnmarshalJSON(jsonStringModel3)
			if err != nil {
				b.Fatal("gojson unmarshal error:", err)
			}
		}
	})
}

func Benchmark_StdJSON_Unmarshal_2(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var model Model3Type
			err := json.Unmarshal(jsonStringModel3, &model)
			if err != nil {
				b.Fatal("standard unmarshal error:", err)
			}
		}
	})

}

func Benchmark_JSONIter_Unmarshal_2(b *testing.B) {
	_json := jsoniter.ConfigCompatibleWithStandardLibrary

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var model Model3Type
			err := _json.Unmarshal(jsonStringModel3, &model)
			if err != nil {
				b.Fatal("standard unmarshal error:", err)
			}
		}
	})

}
