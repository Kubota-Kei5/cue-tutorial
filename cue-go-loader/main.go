package main

import (
	"fmt"
	"log"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"
)

func main() {
	// CUE のコンテキストを作成（cue.Context{} の代わりに cuecontext.New() を使用）
	ctx := cuecontext.New()

	// CUE ファイルをロード
	insts := load.Instances([]string{"./config/config.cue"}, nil)
	if len(insts) == 0 || insts[0] == nil {
		log.Fatal("Failed to load CUE file")
	}

	// インスタンスのエラー確認
	if insts[0].Err != nil {
		log.Fatalf("Error loading CUE: %v", insts[0].Err)
	}

	// CUE インスタンスをビルド
	value := ctx.BuildInstance(insts[0])

	// エラー確認
	if value.Err() != nil {
		log.Fatalf("Failed to build CUE value: %v", value.Err())
	}

	// `app.name` の値を取得
	appName := value.LookupPath(cue.ParsePath("app.name"))
	if appName.Err() != nil {
		log.Fatalf("Failed to find 'app.name': %v", appName.Err())
	}

	fmt.Println("App Name:", appName)
}
