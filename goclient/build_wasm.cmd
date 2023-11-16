REM
REM Run the following script in the Windows Command Prompt (Not Powershell) to build the WebAssembly module
REM
REM Dpendencies:
REM
REM Install Scoop, the CLI Windows Installer: (Use PowerShell)
REM Set-ExecutionPolicy RemoteSigned -Scope CurrentUser # Optional: Needed to run a remote script the first time
REM irm get.scoop.sh | iex
REM
REN Run in the Windows Command Prompt
REM Install TinyGO - The Go Compiler
REM scoop install tinygo
REM
REM Add WebAssembly/WASI compiler support
REM scoop install binaryen
REM
REM Set the full path to extism.exe
SET EXTISM="%USERPROFILE%\Downloads\extism-v0.3.3-windows-amd64\extism.exe"

SET GOARCH=amd64
SET GOOS=windows
CALL go run hello.go

SET GOARCH=amd64
SET GOOS=windows
REM CALL go run client.go

DEL /Q main.wasm
ECHO main.wasm deleted
DIR
go mod init goclient
go get github.com/gorilla/websocket
go get github.com/extism/go-pdk@main
tinygo build -o main.wasm -target wasi main.go
ECHO *** main.wasm created ***
DIR

%EXTISM% call main.wasm greet --input "Benjamin" --wasi

REM This WebAssembly format is not supported by Extism
SET GOOS=wasip1
SET GOARCH=wasm
go build -o goWithWebSockets.wasm goWithWebSockets.go

%EXTISM% call goWithWebSockets.wasm _start --input "Benjamin" --wasi

REM tinygo build -o goWithWebSockets.wasm -target wasi goWithWebSockets.go
REM
REM ERROR:
REM
REM wasm-ld: error: C:\Users\tgrau\AppData\Local\Temp\tinygo3115611328\main.o: undefined symbol: crypto/internal/bigmod.addMulVVW1024
REM
