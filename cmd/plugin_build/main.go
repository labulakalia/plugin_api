package main

import (
	"archive/zip"
	"compress/flate"
	"errors"

	"fmt"
	"io"
	"path/filepath"

	"log/slog"
	"os"
	"os/exec"

	"github.com/BurntSushi/toml"
)

type pluginConfig struct {
	Id        string   `toml:"id"`
	Name      string   `toml:"name"`
	Desc      string   `toml:"desc"`
	Author    []string `toml:"author"`
	Version   string   `toml:"version"`
	Icon      string   `toml:"icon"`
	Changelog []string `toml:"changelog"`
}

func main() {
	pluginTomlFile := "plugin.toml"
	pluginFile, err := os.Open("plugin.toml")
	if err != nil {
		slog.Error("can open plugin config file plugin.toml")
		os.Exit(1)
	}
	pluginConfig := &pluginConfig{}
	_, err = toml.NewDecoder(pluginFile).Decode(pluginConfig)
	if err != nil {
		slog.Error("read plugin.toml failed", "err", err)
		os.Exit(1)
	}
	pluginFile.Close()

	if len(pluginConfig.Icon) > 0 {
		_, err = os.Stat(pluginConfig.Icon)
		if err != nil {
			slog.Error("read icon failed", "err", err)
			os.Exit(1)
		}
	}
	_, err = os.Stat("dist")
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			slog.Error("check dist dir failed", "err", err)
			os.Exit(1)
		}
		err = os.MkdirAll("dist", 0755)
		if err != nil {
			slog.Error("mkdir dist dir failed", "err", err)
			os.Exit(1)
		}
	} else {
		_ = os.RemoveAll("./dist/*")
	}

	buildWasmFile := fmt.Sprintf("dist/%s.wasm", pluginConfig.Id)
	defer os.Remove(buildWasmFile)
	cmd := exec.Command("go", "build", "-buildmode=c-shared", "-o", buildWasmFile, ".")
	cmd.Env = append(os.Environ(), "GOOS=wasip1", "GOARCH=wasm")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		slog.Error("exec command failed", "err", err)
		os.Exit(1)
	}

	outCompressFile := filepath.Join("./dist", fmt.Sprintf("%s_%s.zip", pluginConfig.Id, pluginConfig.Version))
	outFile, err := os.Create(outCompressFile)
	if err != nil {
		slog.Error("create file failed", "file", outCompressFile, "err", err)
		os.Exit(1)
	}
	defer outFile.Close()
	zw := zip.NewWriter(outFile)
	zw.RegisterCompressor(zip.Deflate, func(out io.Writer) (io.WriteCloser, error) {
		return flate.NewWriter(out, flate.BestCompression)
	})
	defer zw.Close()

	zipFileMap := map[string]string{
		pluginTomlFile:                          pluginTomlFile,
		fmt.Sprintf("%s.wasm", pluginConfig.Id): buildWasmFile,
		pluginConfig.Icon:                       pluginConfig.Icon,
	}

	for fileName, filePath := range zipFileMap {
		localFile, err := os.Open(filePath)
		if err != nil {
			slog.Error("zip create file failed", "file", filePath, "err", err)
			os.Exit(1)
		}

		zwfile, err := zw.Create(fileName)
		if err != nil {
			slog.Error("zip create file failed", "file", fileName, "err", err)
			os.Exit(1)
		}
		io.Copy(zwfile, localFile)
		localFile.Close()
	}

	slog.Info("build success", "plugin", pluginConfig.Name, "plugin_path", outCompressFile)
}
