package storage

import (
	"fmt"
	"testing"
)

func TestObtain(t *testing.T) {
	store, err := Obtain()
	if err != nil {
		panic(err)
	}

	store.Env(`ANDROID_HOME`, `/usr/local/android-sdk`)
	store.Env(`ANDROID_HOME`, ``)
	store.Env(`FLUTTER`, `./usr/local/flutter`)

	store.Alias(`ssh_linux`, `ssh user@127.0.0.1`)
	store.Alias(`ssh_linux`, ``)
	store.Alias(`cls`, `clear`)
	store.Alias(`v2`, `HTTPS_PROXY=127.0.0.1:1081 HTTP_PROXY=127.0.0.1:1081`)

	store.Path(`./usr/local/bin`)
	store.Path(`/usr/local/jdk/bin`)
	//store.Path(`@/usr/local/bin`)

	fmt.Println(store.View())

	fmt.Println(store.Flush())
	fmt.Println(store.Profile())

}
