# WCX Packer extension (plug-in) example for TotalCommander in Golang

Each plug-in is created as a C DLL library which exports a defined set of functions. 
The library has the extensions wdx/wfx/wlx/wcx instead of DLL depending on which type of plug-in it represents. 
The set of functions that a plug-in must export is defined by the Total Commander plug-in interface.

## Compilation

Setup TDM-GCC-64 v 10.3.0 from https://jmeubank.github.io/tdm-gcc/download/ 

Set environment
```env
GOARCH=386
CGO_ENABLED=1
CC=c:\TDM-GCC-64\bin\x86_64-w64-mingw32-gcc.exe
```

Build
```shell
go build -buildmode=c-shared -o totalcmd-wcx-demo-go.wcx totalcmd-wcx-demo-go.go
```

## Additional materials

[WCX Documentation](https://plugins.ghisler.com/plugins/wcx_ref2.21se_chm.zip)

[Plug-ins and some sources](https://www.ghisler.com/plugins.htm)

[Documentation for Plugin Writers](https://totalcmd.net/directory/developer.html)

[Writing a Total Commander plug-in in Visual Basic (or C#)](https://www.codeproject.com/Articles/33984/Writing-a-Total-Commander-plug-in-in-Visual-Basic)
as a good description of how does it work

[Dynamic-Link Library Best Practices](https://learn.microsoft.com/en-us/windows/win32/dlls/dynamic-link-library-best-practices)
