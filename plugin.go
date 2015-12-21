package main

import (
	"fmt"

	// CF Pluginを書くために必要なパッケージ
	"github.com/cloudfoundry/cli/plugin"
)

type MyFirstPlugin struct{}

// 実際の Plugin の処理はここに書かれていきます。
// cliConnection は(おおまかに言うと) cf CLI の各種コマンドを呼び出すためのものです。
// args には Plugin を実行する際のコマンドライン引数が入ります。
func (mfp *MyFirstPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	fmt.Println("Hello, CF Plugin!")
}

// Plugin の metadata(名前、バージョン、コマンドの説明等) を設定
func (mfp *MyFirstPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "MyFirstPlugin",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 0,
			Build: 0,
		},
		Commands: []plugin.Command{
			plugin.Command{
				Name:     "my-first-plugin",
				HelpText: "This is a help text.",

				UsageDetails: plugin.Usage{
					Usage: "my-first-plugin\n   cf my-first-plugin",
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(MyFirstPlugin))
}
