package main

import (
	"github.com/labulakalia/plugin_temp/plugin"
	pluginproto "github.com/labulakalia/plugin_temp/proto"

	"github.com/labulakalia/wazero_net/util"
	"google.golang.org/protobuf/proto"
)

func main() {}

var pluginimpl plugin.IPlugin

func init() {
	pluginimpl = plugin.NewPlugin()
}

// [32 byte][1 err flag][31 length]
func err2Uint64(err error) uint64 {
	errStr := err.Error()
	errLength := uint32(len(errStr) | 1<<31)
	errPtr := uint32(util.StringToPtr(&errStr))
	return util.Uint32ToUint64(errPtr, uint32(errLength))
}

// go::wasmexport id
func id() uint64 {
	id := pluginimpl.Id()
	idPtr := uint32(util.StringToPtr(&id))
	idLen := uint32(len(id))
	return util.Uint32ToUint64(idPtr, idLen)
}

// go::wasmexport get_auth_type
func get_auth_type() uint64 {
	authType := pluginimpl.GetAuthType()
	data, err := proto.Marshal(authType)
	if err != nil {
		return err2Uint64(err)
	}
	return util.Uint32ToUint64(uint32(util.BytesToPtr(data)), uint32(len(data)))
}

// go::wasmexport check_auth
func check_auth(authTypePtr, authTypeLenPtr uint64) uint64 {
	data := util.PtrToBytes(uint32(authTypePtr), uint32(authTypePtr))
	authType := &pluginproto.AuthType{}
	err := proto.Unmarshal(data, authType)
	if err != nil {
		return err2Uint64(err)
	}
	status := pluginimpl.CheckAuth(authType)
	statusData, err := proto.Marshal(status)
	if err != nil {
		return err2Uint64(err)
	}
	return util.Uint32ToUint64(uint32(util.BytesToPtr(statusData)), uint32(len(data)))
}

// go::wasmexport get_auth_data
func get_auth_data() uint64 {
	authData := pluginimpl.GetAuthData()
	return util.Uint32ToUint64(uint32(util.BytesToPtr(authData)), uint32(len(authData)))
}

// go::wasmexport init_auth
func init_auth(raw_auth_dataPtr, raw_auth_dataLen uint64) uint64 {
	rawAuthData := util.PtrToBytes(uint32(raw_auth_dataPtr), uint32(raw_auth_dataLen))
	status := pluginimpl.InitAuth(rawAuthData)
	statusData, err := proto.Marshal(status)
	if err != nil {
		return err2Uint64(err)
	}
	return util.Uint32ToUint64(uint32(util.BytesToPtr(statusData)), uint32(len(statusData)))
}

// go::wasmexport auth_id
func auth_id() uint64 {
	authId := pluginimpl.AuthId()
	return util.Uint32ToUint64(uint32(util.StringToPtr(&authId)), uint32(len(authId)))
}

// go::wasmexport get_dir_entry
func get_dir_entry(dir_pathPtr, dir_pathLen, page, page_size uint64) uint64 {
	dir := util.PtrToString(uint32(dir_pathPtr), uint32(dir_pathLen))
	dirEntry := pluginimpl.GetDirEntry(dir, page, page_size)
	dirEntryData, err := proto.Marshal(dirEntry)
	if err != nil {
		return err2Uint64(err)
	}
	return util.Uint32ToUint64(uint32(util.BytesToPtr(dirEntryData)), uint32(len(dirEntryData)))
}

// go::wasmexport get_file_resource
func get_file_resource(file_pathPtr, file_pathLen uint64) uint64 {
	file_path := util.PtrToString(uint32(file_pathPtr), uint32(file_pathLen))
	fileResource := pluginimpl.GetFileResource(file_path)
	fileResourceData, err := proto.Marshal(fileResource)
	if err != nil {
		return err2Uint64(err)
	}
	return util.Uint32ToUint64(uint32(util.BytesToPtr(fileResourceData)), uint32(len(fileResourceData)))
}

// go::wasmexport close
func close() uint64 {
	pluginimpl.Close()
	return 0
}
