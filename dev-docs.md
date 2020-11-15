# Developer documentation

It copies the package of the desired operating system and prepends the email address in base64 encoding.

This is necessary in order to be able to link users to their devices in the log files, which is otherwise difficult to do.

This is supposed to be used alongside Filebeat, which will parse the base64 encoded email.

This can be found in the [elastic-fleet](https://github.com/kylezs/elastic-fleet/blob/master/templates/fleet-configmap.yaml) repository.

## Building the launcher packages
The launcher packages need to be built (see: https://github.com/kolide/launcher) (on their respective OS's) and then named as so
```
Mac: launcher.darwin-launchd-pkg.pkg
Windows: launcher.windows-service-msi.msi
```

These should then be put in a directory named `launcher-pkgs`

So, the directory tree for this repo should look like this before compiling:
```
.
├── README.md
├── dev-docs.md
├── launcher-pkgs
│   ├── launcher.darwin-launchd-pkg.pkg
│   └── launcher.windows-service-msi.msi
├── output
| ... <any outputs that you create from using the scrip>
├── run
│   ├── mac-launcher-rename
│   └── windows-launcher-rename.exe
└── src
    ├── build.sh
    ├── main.go
    └── rename.go
```

## Compiling

For the list of environments you can compile a binary for:
https://golang.org/doc/install/source#environment

You can use the script in `/src` to compile for windows and mac.

Example: Compiling Windows from something else. 
> This does NOT work for the launcher packages in the step above. You will have to compile those on the respective OSs

```
env GOOS=windows GOARCH=amd64 go build src/main.go -o bin/windows-launcher-rename.exe
```

NB: Requires the `.exe` extension for windows

Then just make sure they are in the `bin` directory.

<hr>



## Dev Notes

Currently `go run main.go` is broken due to filepath issues. But the built binaries work cross platform, no matter where the package is stored.
