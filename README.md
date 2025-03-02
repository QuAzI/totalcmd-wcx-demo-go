# WCX Packer extension (plug-in) example for TotalCommander in Golang

Each plug-in is created as a C DLL library which exports a defined set of functions. 
The library has the extensions wdx/wfx/wlx/wcx (x32) and wdx64/wfx64/wlx64/wcx64 (AMD64) instead of DLL 
depending on which type of plug-in it represents. 
The set of functions that a plug-in must export is defined by the Total Commander plug-in interface.


## Compilation

Required TDM-GCC-64 v10.3.0 from https://jmeubank.github.io/tdm-gcc/download/ 

## Plugin installer

Just ZIP `pluginst.inf` which contains plugin's description in libraries from build

```
totalcmd-wcx-demo-go.zip
├── pluginst.inf
├── totalcmd-wcx-demo-go.wcx
└── totalcmd-wcx-demo-go.wcx64
```

## Additional materials

[WCX Documentation](https://plugins.ghisler.com/plugins/wcx_ref2.21se_chm.zip)

[Plug-ins and some sources](https://www.ghisler.com/plugins.htm)

[Documentation for Plugin Writers](https://totalcmd.net/directory/developer.html)

[Writing a Total Commander plug-in in Visual Basic (or C#)](https://www.codeproject.com/Articles/33984/Writing-a-Total-Commander-plug-in-in-Visual-Basic)
as a good description of how does it work

[Dynamic-Link Library Best Practices](https://learn.microsoft.com/en-us/windows/win32/dlls/dynamic-link-library-best-practices)

I found `wcxtest` but it doesn't work properly, please use extra copy of Total Commander for tests
