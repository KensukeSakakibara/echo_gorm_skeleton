/*
main.go
@import github.com/KensukeSakakibara/echo_gorm_skeleton
@author Kensuke Sakakibara
@since 2021.01.20
@copyright Copyright (c) 2021 Kensuke Sakakibara
*/
package main

import (
	"github.com/KensukeSakakibara/echo_gorm_skeleton/registry"
)

func main() {
	// アプリケーションの初期化
	initInterface := registry.DiInit()
	initInterface.Run()

	// アプリ実行
	routerInterface := registry.DiRouter()
	routerInterface.Run()
}
