package main

import (
	"fmt"
	"net/http"
)

// ハンドラ（handler）という言葉は、イベントがトリガー（きっかけ）として起動されるコールバック関数でよく使われる
func handler(writer http.ResponseWriter, request *http.Request) {
	// 接頭辞F が付いているものは、書き込み先を明示的に指定可能
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func main() {
	// 指定されたパターン（ここではルートURL（/））が呼び出された時、ハンドラが起動されるように設定する
	// （内部的には ServeMux が起動している。ServeMux は HTTPリクエストマルチプレクサで、リクエストで飛んできたURLが登録済みのパターンにマッチした時に指定されたハンドラを呼び出す）
	http.HandleFunc("/", handler)
	// ポート8080を監視するようにサーバを起動する
	http.ListenAndServe(":8080", nil)
}
