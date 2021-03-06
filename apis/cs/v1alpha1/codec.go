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
		jsoniter.MustGetKind(reflect2.TypeOf(EdgeKubernetesSpecLogConfig{}).Type1()):            EdgeKubernetesSpecLogConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(KubernetesSpecLogConfig{}).Type1()):                KubernetesSpecLogConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(KubernetesNodePoolSpecManagement{}).Type1()):       KubernetesNodePoolSpecManagementCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(KubernetesNodePoolSpecScalingConfig{}).Type1()):    KubernetesNodePoolSpecScalingConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ManagedKubernetesSpecLogConfig{}).Type1()):         ManagedKubernetesSpecLogConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ManagedKubernetesSpecMaintenanceWindow{}).Type1()): ManagedKubernetesSpecMaintenanceWindowCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(EdgeKubernetesSpecLogConfig{}).Type1()):            EdgeKubernetesSpecLogConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(KubernetesSpecLogConfig{}).Type1()):                KubernetesSpecLogConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(KubernetesNodePoolSpecManagement{}).Type1()):       KubernetesNodePoolSpecManagementCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(KubernetesNodePoolSpecScalingConfig{}).Type1()):    KubernetesNodePoolSpecScalingConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ManagedKubernetesSpecLogConfig{}).Type1()):         ManagedKubernetesSpecLogConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ManagedKubernetesSpecMaintenanceWindow{}).Type1()): ManagedKubernetesSpecMaintenanceWindowCodec{},
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
type EdgeKubernetesSpecLogConfigCodec struct {
}

func (EdgeKubernetesSpecLogConfigCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*EdgeKubernetesSpecLogConfig)(ptr) == nil
}

func (EdgeKubernetesSpecLogConfigCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*EdgeKubernetesSpecLogConfig)(ptr)
	var objs []EdgeKubernetesSpecLogConfig
	if obj != nil {
		objs = []EdgeKubernetesSpecLogConfig{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EdgeKubernetesSpecLogConfig{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (EdgeKubernetesSpecLogConfigCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*EdgeKubernetesSpecLogConfig)(ptr) = EdgeKubernetesSpecLogConfig{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []EdgeKubernetesSpecLogConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EdgeKubernetesSpecLogConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*EdgeKubernetesSpecLogConfig)(ptr) = objs[0]
			} else {
				*(*EdgeKubernetesSpecLogConfig)(ptr) = EdgeKubernetesSpecLogConfig{}
			}
		} else {
			*(*EdgeKubernetesSpecLogConfig)(ptr) = EdgeKubernetesSpecLogConfig{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj EdgeKubernetesSpecLogConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EdgeKubernetesSpecLogConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*EdgeKubernetesSpecLogConfig)(ptr) = obj
		} else {
			*(*EdgeKubernetesSpecLogConfig)(ptr) = EdgeKubernetesSpecLogConfig{}
		}
	default:
		iter.ReportError("decode EdgeKubernetesSpecLogConfig", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type KubernetesSpecLogConfigCodec struct {
}

func (KubernetesSpecLogConfigCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*KubernetesSpecLogConfig)(ptr) == nil
}

func (KubernetesSpecLogConfigCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*KubernetesSpecLogConfig)(ptr)
	var objs []KubernetesSpecLogConfig
	if obj != nil {
		objs = []KubernetesSpecLogConfig{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(KubernetesSpecLogConfig{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (KubernetesSpecLogConfigCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*KubernetesSpecLogConfig)(ptr) = KubernetesSpecLogConfig{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []KubernetesSpecLogConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(KubernetesSpecLogConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*KubernetesSpecLogConfig)(ptr) = objs[0]
			} else {
				*(*KubernetesSpecLogConfig)(ptr) = KubernetesSpecLogConfig{}
			}
		} else {
			*(*KubernetesSpecLogConfig)(ptr) = KubernetesSpecLogConfig{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj KubernetesSpecLogConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(KubernetesSpecLogConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*KubernetesSpecLogConfig)(ptr) = obj
		} else {
			*(*KubernetesSpecLogConfig)(ptr) = KubernetesSpecLogConfig{}
		}
	default:
		iter.ReportError("decode KubernetesSpecLogConfig", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type KubernetesNodePoolSpecManagementCodec struct {
}

func (KubernetesNodePoolSpecManagementCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*KubernetesNodePoolSpecManagement)(ptr) == nil
}

func (KubernetesNodePoolSpecManagementCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*KubernetesNodePoolSpecManagement)(ptr)
	var objs []KubernetesNodePoolSpecManagement
	if obj != nil {
		objs = []KubernetesNodePoolSpecManagement{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(KubernetesNodePoolSpecManagement{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (KubernetesNodePoolSpecManagementCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*KubernetesNodePoolSpecManagement)(ptr) = KubernetesNodePoolSpecManagement{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []KubernetesNodePoolSpecManagement

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(KubernetesNodePoolSpecManagement{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*KubernetesNodePoolSpecManagement)(ptr) = objs[0]
			} else {
				*(*KubernetesNodePoolSpecManagement)(ptr) = KubernetesNodePoolSpecManagement{}
			}
		} else {
			*(*KubernetesNodePoolSpecManagement)(ptr) = KubernetesNodePoolSpecManagement{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj KubernetesNodePoolSpecManagement

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(KubernetesNodePoolSpecManagement{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*KubernetesNodePoolSpecManagement)(ptr) = obj
		} else {
			*(*KubernetesNodePoolSpecManagement)(ptr) = KubernetesNodePoolSpecManagement{}
		}
	default:
		iter.ReportError("decode KubernetesNodePoolSpecManagement", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type KubernetesNodePoolSpecScalingConfigCodec struct {
}

func (KubernetesNodePoolSpecScalingConfigCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*KubernetesNodePoolSpecScalingConfig)(ptr) == nil
}

func (KubernetesNodePoolSpecScalingConfigCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*KubernetesNodePoolSpecScalingConfig)(ptr)
	var objs []KubernetesNodePoolSpecScalingConfig
	if obj != nil {
		objs = []KubernetesNodePoolSpecScalingConfig{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(KubernetesNodePoolSpecScalingConfig{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (KubernetesNodePoolSpecScalingConfigCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*KubernetesNodePoolSpecScalingConfig)(ptr) = KubernetesNodePoolSpecScalingConfig{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []KubernetesNodePoolSpecScalingConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(KubernetesNodePoolSpecScalingConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*KubernetesNodePoolSpecScalingConfig)(ptr) = objs[0]
			} else {
				*(*KubernetesNodePoolSpecScalingConfig)(ptr) = KubernetesNodePoolSpecScalingConfig{}
			}
		} else {
			*(*KubernetesNodePoolSpecScalingConfig)(ptr) = KubernetesNodePoolSpecScalingConfig{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj KubernetesNodePoolSpecScalingConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(KubernetesNodePoolSpecScalingConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*KubernetesNodePoolSpecScalingConfig)(ptr) = obj
		} else {
			*(*KubernetesNodePoolSpecScalingConfig)(ptr) = KubernetesNodePoolSpecScalingConfig{}
		}
	default:
		iter.ReportError("decode KubernetesNodePoolSpecScalingConfig", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ManagedKubernetesSpecLogConfigCodec struct {
}

func (ManagedKubernetesSpecLogConfigCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ManagedKubernetesSpecLogConfig)(ptr) == nil
}

func (ManagedKubernetesSpecLogConfigCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ManagedKubernetesSpecLogConfig)(ptr)
	var objs []ManagedKubernetesSpecLogConfig
	if obj != nil {
		objs = []ManagedKubernetesSpecLogConfig{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ManagedKubernetesSpecLogConfig{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ManagedKubernetesSpecLogConfigCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ManagedKubernetesSpecLogConfig)(ptr) = ManagedKubernetesSpecLogConfig{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ManagedKubernetesSpecLogConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ManagedKubernetesSpecLogConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ManagedKubernetesSpecLogConfig)(ptr) = objs[0]
			} else {
				*(*ManagedKubernetesSpecLogConfig)(ptr) = ManagedKubernetesSpecLogConfig{}
			}
		} else {
			*(*ManagedKubernetesSpecLogConfig)(ptr) = ManagedKubernetesSpecLogConfig{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ManagedKubernetesSpecLogConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ManagedKubernetesSpecLogConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ManagedKubernetesSpecLogConfig)(ptr) = obj
		} else {
			*(*ManagedKubernetesSpecLogConfig)(ptr) = ManagedKubernetesSpecLogConfig{}
		}
	default:
		iter.ReportError("decode ManagedKubernetesSpecLogConfig", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ManagedKubernetesSpecMaintenanceWindowCodec struct {
}

func (ManagedKubernetesSpecMaintenanceWindowCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ManagedKubernetesSpecMaintenanceWindow)(ptr) == nil
}

func (ManagedKubernetesSpecMaintenanceWindowCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ManagedKubernetesSpecMaintenanceWindow)(ptr)
	var objs []ManagedKubernetesSpecMaintenanceWindow
	if obj != nil {
		objs = []ManagedKubernetesSpecMaintenanceWindow{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ManagedKubernetesSpecMaintenanceWindow{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ManagedKubernetesSpecMaintenanceWindowCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ManagedKubernetesSpecMaintenanceWindow)(ptr) = ManagedKubernetesSpecMaintenanceWindow{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ManagedKubernetesSpecMaintenanceWindow

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ManagedKubernetesSpecMaintenanceWindow{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ManagedKubernetesSpecMaintenanceWindow)(ptr) = objs[0]
			} else {
				*(*ManagedKubernetesSpecMaintenanceWindow)(ptr) = ManagedKubernetesSpecMaintenanceWindow{}
			}
		} else {
			*(*ManagedKubernetesSpecMaintenanceWindow)(ptr) = ManagedKubernetesSpecMaintenanceWindow{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ManagedKubernetesSpecMaintenanceWindow

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ManagedKubernetesSpecMaintenanceWindow{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ManagedKubernetesSpecMaintenanceWindow)(ptr) = obj
		} else {
			*(*ManagedKubernetesSpecMaintenanceWindow)(ptr) = ManagedKubernetesSpecMaintenanceWindow{}
		}
	default:
		iter.ReportError("decode ManagedKubernetesSpecMaintenanceWindow", "unexpected JSON type")
	}
}
