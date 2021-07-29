/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

func GetEncoder() map[string]jsoniter.ValEncoder {
	return map[string]jsoniter.ValEncoder{
		jsoniter.MustGetKind(reflect2.TypeOf(StoreSpecEncryptConf{}).Type1()):            StoreSpecEncryptConfCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(StoreSpecEncryptConfUserCmkInfo{}).Type1()): StoreSpecEncryptConfUserCmkInfoCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(StoreIndexSpecFullText{}).Type1()):          StoreIndexSpecFullTextCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(StoreSpecEncryptConf{}).Type1()):            StoreSpecEncryptConfCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(StoreSpecEncryptConfUserCmkInfo{}).Type1()): StoreSpecEncryptConfUserCmkInfoCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(StoreIndexSpecFullText{}).Type1()):          StoreIndexSpecFullTextCodec{},
	}
}

func getEncodersWithout(typ string) map[string]jsoniter.ValEncoder {
	origMap := GetEncoder()
	delete(origMap, typ)
	return origMap
}

func getDecodersWithout(typ string) map[string]jsoniter.ValDecoder {
	origMap := GetDecoder()
	delete(origMap, typ)
	return origMap
}

// +k8s:deepcopy-gen=false
type StoreSpecEncryptConfCodec struct {
}

func (StoreSpecEncryptConfCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*StoreSpecEncryptConf)(ptr) == nil
}

func (StoreSpecEncryptConfCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*StoreSpecEncryptConf)(ptr)
	var objs []StoreSpecEncryptConf
	if obj != nil {
		objs = []StoreSpecEncryptConf{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(StoreSpecEncryptConf{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (StoreSpecEncryptConfCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*StoreSpecEncryptConf)(ptr) = StoreSpecEncryptConf{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []StoreSpecEncryptConf

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(StoreSpecEncryptConf{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*StoreSpecEncryptConf)(ptr) = objs[0]
			} else {
				*(*StoreSpecEncryptConf)(ptr) = StoreSpecEncryptConf{}
			}
		} else {
			*(*StoreSpecEncryptConf)(ptr) = StoreSpecEncryptConf{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj StoreSpecEncryptConf

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(StoreSpecEncryptConf{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*StoreSpecEncryptConf)(ptr) = obj
		} else {
			*(*StoreSpecEncryptConf)(ptr) = StoreSpecEncryptConf{}
		}
	default:
		iter.ReportError("decode StoreSpecEncryptConf", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type StoreSpecEncryptConfUserCmkInfoCodec struct {
}

func (StoreSpecEncryptConfUserCmkInfoCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*StoreSpecEncryptConfUserCmkInfo)(ptr) == nil
}

func (StoreSpecEncryptConfUserCmkInfoCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*StoreSpecEncryptConfUserCmkInfo)(ptr)
	var objs []StoreSpecEncryptConfUserCmkInfo
	if obj != nil {
		objs = []StoreSpecEncryptConfUserCmkInfo{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(StoreSpecEncryptConfUserCmkInfo{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (StoreSpecEncryptConfUserCmkInfoCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*StoreSpecEncryptConfUserCmkInfo)(ptr) = StoreSpecEncryptConfUserCmkInfo{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []StoreSpecEncryptConfUserCmkInfo

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(StoreSpecEncryptConfUserCmkInfo{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*StoreSpecEncryptConfUserCmkInfo)(ptr) = objs[0]
			} else {
				*(*StoreSpecEncryptConfUserCmkInfo)(ptr) = StoreSpecEncryptConfUserCmkInfo{}
			}
		} else {
			*(*StoreSpecEncryptConfUserCmkInfo)(ptr) = StoreSpecEncryptConfUserCmkInfo{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj StoreSpecEncryptConfUserCmkInfo

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(StoreSpecEncryptConfUserCmkInfo{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*StoreSpecEncryptConfUserCmkInfo)(ptr) = obj
		} else {
			*(*StoreSpecEncryptConfUserCmkInfo)(ptr) = StoreSpecEncryptConfUserCmkInfo{}
		}
	default:
		iter.ReportError("decode StoreSpecEncryptConfUserCmkInfo", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type StoreIndexSpecFullTextCodec struct {
}

func (StoreIndexSpecFullTextCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*StoreIndexSpecFullText)(ptr) == nil
}

func (StoreIndexSpecFullTextCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*StoreIndexSpecFullText)(ptr)
	var objs []StoreIndexSpecFullText
	if obj != nil {
		objs = []StoreIndexSpecFullText{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(StoreIndexSpecFullText{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (StoreIndexSpecFullTextCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*StoreIndexSpecFullText)(ptr) = StoreIndexSpecFullText{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []StoreIndexSpecFullText

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(StoreIndexSpecFullText{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*StoreIndexSpecFullText)(ptr) = objs[0]
			} else {
				*(*StoreIndexSpecFullText)(ptr) = StoreIndexSpecFullText{}
			}
		} else {
			*(*StoreIndexSpecFullText)(ptr) = StoreIndexSpecFullText{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj StoreIndexSpecFullText

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(StoreIndexSpecFullText{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*StoreIndexSpecFullText)(ptr) = obj
		} else {
			*(*StoreIndexSpecFullText)(ptr) = StoreIndexSpecFullText{}
		}
	default:
		iter.ReportError("decode StoreIndexSpecFullText", "unexpected JSON type")
	}
}
