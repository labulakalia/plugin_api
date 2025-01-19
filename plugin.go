package plugin_api

import "github.com/labulakalia/plugin_api/plugin"

type IPlugin interface {
	// plugin id
	PluginId() (string, error)
	// get auth type like form edit,qrcode,oauth2
	GetAuthType() (*plugin.AuthType, error)
	// check auth data status
	CheckAuth(*plugin.AuthType) (*plugin.Status, error)
	// get auth data when check auth status is success
	GetAuthData() ([]byte, error)
	// use auth data init auth
	CheckAuthData([]byte) (*plugin.Status, error)
	// plugin auth id,it need unqiue for same driver
	PluginAuthId() (string, error)
	// get dir entry from driver plugin
	GetDirEntry(dir_path string, page, page_size uint64) (*plugin.DirEntry, error)
	// get file entry resource from driver plugin
	GetFileResource(file_path string) (*plugin.FileResource, error)
	// close driver plugin
	Close() error
}

func RegistryPlugin(iPlugin IPlugin) {
	pluginExport = &PluginExport{
		IPlugin: iPlugin,
	}
}
