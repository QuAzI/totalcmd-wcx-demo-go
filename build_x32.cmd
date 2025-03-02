set GOARCH=386
set CGO_ENABLED=1
set CC=c:\TDM-GCC-64\bin\x86_64-w64-mingw32-gcc.exe

go build -buildmode=c-shared -o build\x32\totalcmd-wcx-demo-go.wcx totalcmd-wcx-demo-go
