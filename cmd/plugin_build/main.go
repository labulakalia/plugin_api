package main

import (
	"archive/zip"
	"compress/flate"
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
	pluginConfig := &pluginConfig{}
	_, err := toml.DecodeFile(pluginTomlFile, pluginConfig)
	if err != nil {
		slog.Error("read plugin.toml failed", "err", err)
		os.Exit(1)
	}
	if len(pluginConfig.Icon) > 0 {
		_, err = os.Stat(pluginConfig.Icon)
		if err != nil {
			slog.Error("read icon failed", "err", err)
			os.Exit(1)
		}
	}
	err = os.RemoveAll("./dist/*")
	if err != nil {
		slog.Error("clean dist dir failed", "err", err)
		os.Exit(1)
	}
	err = os.MkdirAll("dist", 0755)
	if err != nil {
		slog.Error("mkdir dist dir failed", "err", err)
		os.Exit(1)
	}
	buildWasmFile := fmt.Sprintf("%s.wasm", pluginConfig.Id)
	defer os.Remove(buildWasmFile)

	buildCmd := fmt.Sprintf("GOOS=wasip1 GOARCH=wasm go build -buildmode=c-shared -o '%s' .", buildWasmFile)

	cmd := exec.Command("bash", "-c", buildCmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		slog.Error("exec command failed", "err", err)
		os.Exit(1)
	}
	pwd, err := os.Getwd()
	if err != nil {
		slog.Error("get pwd failed", "err", err)
		os.Exit(1)
	}
	outCompressFile := filepath.Join(pwd, "dist", fmt.Sprintf("%s_%s.zip", pluginConfig.Id, pluginConfig.Version))
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

	fileList := []string{
		pluginTomlFile,
		fmt.Sprintf("%s.wasm", pluginConfig.Id),
		pluginConfig.Icon,
	}

	for _, file := range fileList {
		localFile, err := os.Open(file)
		if err != nil {
			slog.Error("zip create file failed", "file", file, "err", err)
			os.Exit(1)
		}

		zwfile, err := zw.Create(file)
		if err != nil {
			slog.Error("zip create file failed", "file", file, "err", err)
			os.Exit(1)
		}
		io.Copy(zwfile, localFile)
		localFile.Close()
	}
	slog.Info("build success", "plugin", pluginConfig.Name, "plugin_path", outCompressFile)
}
