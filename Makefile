mac:
	GOOS=darwin GOARCH=amd64 CGO_CFLAGS="-mmacosx-version-min=10.12" CGO_LDFLAGS="-mmacosx-version-min=10.12" go build  -o bin/beetle
	appify -name beetle -icon ../beatles-client-lib/resource/beatlesicon.png ./bin/beetle
	rm -f *.dmg
	create-dmg --volname "Beetle Installer" --window-pos 200 120 --window-size 800 400  --app-drop-link 600 185  Beetle.dmg beetle.app
