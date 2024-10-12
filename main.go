package main

/*
#include <stdio.h>
#include <string.h>

typedef struct {
    char* ArcName;
    int OpenMode;
    int OpenResult;
    char* CmtBuf;
    int CmtBufSize;
    int CmtSize;
    int CmtState;
  } TOpenArchiveData;

typedef struct {
    char ArcName[260];
    char FileName[260];
    int Flags;
    int PackSize;
    int UnpSize;
    int HostOS;
    int FileCRC;
    int FileTime;
    int UnpVer;
    int Method;
    int FileAttr;
    char* CmtBuf;
    int CmtBufSize;
    int CmtSize;
    int CmtState;
  } THeaderData;

__attribute__((weak))
void fill_zero(void *buffer) {
    printf("str_size %d \n", sizeof(THeaderData));
    THeaderData* obj = (THeaderData*)(buffer);
    printf("ptr 0x%x  %d  %d \n", buffer, buffer, &obj);
    printf("ArcName %d %s \n", &obj->ArcName, obj->ArcName);
    printf("FileName %d %s \n", &obj->FileName, obj->FileName);
    printf("Flags %d %d \n", &obj->Flags, obj->Flags);
    printf("PackSize %d %d \n", &obj->PackSize, obj->PackSize);
    printf("UnpSize %d %d \n", &obj->UnpSize, obj->UnpSize);
    printf("HostOS %d %d \n", &obj->HostOS, obj->HostOS);
    printf("FileCRC %d %d \n", &obj->FileCRC, obj->FileCRC);
    printf("FileTime %d %d \n", &obj->FileTime, obj->FileTime);
    printf("UnpVer %d %d \n", &obj->UnpVer, obj->UnpVer);
    printf("Method %d %d \n", &obj->Method, obj->Method);
    printf("FileAttr %d %d \n", &obj->FileAttr, obj->FileAttr);
    printf("CmtBuf %d \n", &obj->CmtBuf);
    printf("CmtBufSize %d %d \n", &obj->CmtBufSize, obj->CmtBufSize);
    printf("CmtSize %d %d \n", &obj->CmtSize, obj->CmtSize);
    printf("CmtState %d %d \n", &obj->CmtState, obj->CmtState);
    memset(obj, 0, sizeof(THeaderData));
    //strncpy((char *)obj->ArcName, "hello", 260);
    obj->Flags = 3;
    obj->PackSize = 4;
    obj->UnpSize = 5;
    obj->FileAttr = 0x20;
    obj->CmtSize = 43;
    obj->CmtBufSize = 44;
}
*/
import "C"
import (
	"fmt"
	"os"
	"totalcmd-wcx-demo-go/totalcmd"
	"unsafe"
)

// var openedFilesMap = map[uintptr]*os.File{}
var openedFilesMap = make(map[uintptr]*os.File)

var testFileCounter = 4

/*
OpenArchive

		HANDLE __stdcall OpenArchive (tOpenArchiveData *ArchiveData);

		OpenArchive should return a unique handle representing the archive.
	    The handle should remain valid until CloseArchive is called.
	    If an error occurs, you should return zero, and specify the error by setting OpenResult member of ArchiveData.
		You can use the ArchiveData to query information about the archive being open, and store the information in
	    ArchiveData to some location that can be accessed via the handle.
*/
//export OpenArchive
func OpenArchive(pArchiveData *C.TOpenArchiveData) uintptr {
	//archiveData := (*totalcmd.TOpenArchiveData)(pArchiveData)
	//archiveData := (*C.TOpenArchiveData)(pArchiveData)
	//archiveData := &pArchiveData

	filename := C.GoString((*C.char)(pArchiveData.ArcName))

	fmt.Println("OpenArchive", filename, pArchiveData)

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		pArchiveData.OpenResult = totalcmd.E_EOPEN
		return 0
	}

	fileInfo, err := file.Stat()
	fmt.Println("File data", fileInfo.Name(), fileInfo.Size(), fileInfo.ModTime())
	openedFilesMap[file.Fd()] = file
	return file.Fd()
}

/*
SetChangeVolProc

		This function allows you to notify user about changing a volume when packing files.

		void __stdcall SetChangeVolProc (HANDLE hArcData, tChangeVolProc pChangeVolProc1);

		pChangeVolProc1 contains a pointer to a function that you may want to call when notifying user to change volume
		(e.g. inserting another diskette). You need to store the value at some place if you want to use it; you can
	    use hArcData that you have returned by OpenArchive to identify that place.
*/
//export SetChangeVolProc
func SetChangeVolProc(hArcData uintptr, pChangeVolProc1 uintptr) {
	fmt.Println("SetChangeVolProc")
}

/*
ReadHeader

		Totalcmd calls ReadHeader to find out what files are in the archive.

		int __stdcall ReadHeader (HANDLE hArcData, tHeaderData *HeaderData);

		ReadHeader is called as long as it returns zero (as long as the previous call to this function returned zero).
	    Each time it is called, HeaderData is supposed to provide Totalcmd with information about the next file contained in
	    the archive. When all files in the archive have been returned, ReadHeader should return E_END_ARCHIVE which will
	    prevent ReaderHeader from being called again.
	    If an error occurs, ReadHeader should return one of the error values or 0 for no error.
		hArcData contains the handle returned by OpenArchive. The programmer is encouraged to store other information in the
	    location that can be accessed via this handle.
	    For example, you may want to store the position in the archive when returning files information in ReadHeader.
		In short, you are supposed to set at least PackSize, UnpSize, FileTime, and FileName members of tHeaderData.
	    Totalcmd will use this information to display content of the archive when the archive is viewed as a directory.
*/
//export ReadHeader
func ReadHeader(hArcData uintptr, pHeaderData *C.THeaderData) C.int {
	file := openedFilesMap[hArcData]
	fmt.Println("ReadHeader", file.Name(), testFileCounter, pHeaderData)

	//headerData := (*totalcmd.THeaderData)(pHeaderData)
	//headerData := (*C.THeaderData)(pHeaderData)
	//headerData := (*c_structs.THeaderData)(pHeaderData)

	//var s string = C.GoStringN((*C.char)(&pHeaderData.FileName), 260)

	structSize := unsafe.Sizeof(*pHeaderData)
	fmt.Println("-- ssize", structSize)

	poi := unsafe.Pointer(pHeaderData)
	C.fill_zero(poi)

	//cFileName := (*C.char)(unsafe.Pointer(&pHeaderData.FileName))
	// var s string = C.GoStringN(cFileName, 260)
	//slice := unsafe.Slice((*C.char)(pHeaderData).FileName, 260)
	//s := string(slice)
	//fmt.Println("-- Filename", s)
	fmt.Println("-- Size", pHeaderData.UnpSize)

	//headerData := (*totalcmd.tHeaderData)(pHeaderData)
	//headerData.UnpSize = C.int(42)
	//headerData.PackSize = C.int(42)
	//headerData.CmtSize = C.int(42)

	// filename := C.GoString((*C.char)(archiveData.ArcName))

	//buf := make([]byte, 256)
	//C.test((*C.char)(unsafe.Pointer(&buf[0])), C.int(len(buf)))

	//mySlice := unsafe.Slice((*byte)(unsafe.Pointer(&headerData.FileName)), 256)
	// and if you need an array type, the slice can be converted
	//myArray := ([256]byte)(mySlice)

	//fmt.Println("-- ArcName", C.GoString((*C.char)(&headerData.FileName[0])))

	//var arr [256]byte
	//copy(arr[:], C.GoBytes(unsafe.Pointer(&C.my_buf), C.BUF_SIZE))
	//copy(headerData.FileName, C.GoBytes(unsafe.Pointer(&headerData.FileName), 256))

	//fmt.Println("-- ArcName", *(*[256]C.char)(unsafe.Pointer(headerData.FileName)))
	//fmt.Println("-- Flags", headerData.Flags)
	//fmt.Println("-- Size", int(headerData.UnpSize))

	//binary.Write(headerData.ArcName, binary.LittleEndian, data)
	//headerData.ArcName
	//headerData.ArcName = C.GoString()
	//C.strcpy((*C.char)(unsafe.Pointer(&headerData.ArcName[0])), (*C.char)(unsafe.Pointer(&bts[0])))

	// fmt.Println("ReadHeader", unsafe.Slice(&headerData.FileName, len(headerData.FileName)))

	testFileCounter -= 1
	if testFileCounter <= 0 {
		return totalcmd.E_END_ARCHIVE
	}
	return totalcmd.SUCCESS
}

func strcpyToC(dst *C.char, src string) {
	n := len(src)

	ds := unsafe.Slice((*byte)(unsafe.Pointer(dst)), n+1)

	copy(ds, src)
	ds[n] = 0
}

/*
ProcessFile

		ProcessFile should unpack the specified file or test the integrity of the archive.

		int __stdcall ProcessFile (HANDLE hArcData, int Operation, char *DestPath, char *DestName);

		ProcessFile should return zero on success, or one of the error values otherwise.
		hArcData contains the handle previously returned by you in OpenArchive. Using this, you should be able to find out
	    information (such as the archive filename) that you need for extracting files from the archive.
		Unlike PackFiles, ProcessFile is passed only one filename. Either DestName contains the full path and file name and
	    DestPath is NULL, or DestName contains only the file name and DestPath the file path.
	    This is done for compatibility with unrar.dll.
		When Total Commander first opens an archive, it scans all file names with OpenMode==PK_OM_LIST, so ReadHeader() is
	    called in a loop with calling ProcessFile(...,PK_SKIP,...). When the user has selected some files and started to
	    decompress them, Total Commander again calls ReadHeader() in a loop.
	    For each file which is to be extracted, Total Commander calls ProcessFile() with Operation==PK_EXTRACT immediately
	    after the ReadHeader() call for this file. If the file needs to be skipped, it calls it with Operation==PK_SKIP.
		Each time DestName is set to contain the filename to be extracted, tested, or skipped. To find out what operation
	    out of these last three you should apply to the current file within the archive,
	    Operation is set to one of the following:
	      Constant Value Description
	      PK_SKIP 0 Skip this file
	      PK_TEST 1 Test file integrity
	      PK_EXTRACT 2 Extract to disk
*/
//export ProcessFile
func ProcessFile(hArcData uintptr, operation int, destPath *C.char, destName *C.char) int {
	const (
		PK_SKIP    = 0
		PK_TEST    = 1
		PK_EXTRACT = 2
	)

	var path = C.GoString(destPath)
	var name = C.GoString(destName)
	fmt.Println()
	fmt.Println("-- ProcessFile", operation, path, name)

	return totalcmd.SUCCESS
}

/*
SetProcessDataProc

	This function allows you to notify user about the progress when you un/pack files.

	void __stdcall SetProcessDataProc (HANDLE hArcData, tProcessDataProc pProcessDataProc);

	pProcessDataProc contains a pointer to a function that you may want to call when notifying user about the progress
	being made when you pack or extract files from an archive. You need to store the value at some place if you want to use it;
	you can use hArcData that you have returned by OpenArchive to identify that place.
*/
//export SetProcessDataProc
func SetProcessDataProc(hArcData uintptr, pProcessDataProc uintptr) {
	fmt.Println("SetProcessDataProc")
}

/*
CloseArchive

	CloseArchive should perform all necessary operations when an archive is about to be closed.

	int __stdcall CloseArchive (HANDLE hArcData);

	CloseArchive should return zero on success, or one of the error values otherwise.
	It should free all the resources associated with the open archive.
	The parameter hArcData refers to the value returned by a programmer within a previous call to OpenArchive.
*/
//export CloseArchive
func CloseArchive(hArcData uintptr) int {
	fmt.Println("CloseArchive")
	file := openedFilesMap[hArcData]
	if file != nil {
		file.Sync()
		closeF := func() {
			err := file.Close()
			if err != nil {
				fmt.Errorf("Could not close archive: %w", err)
			}
		}

		defer closeF() // defer
		delete(openedFilesMap, hArcData)
		fmt.Println("CloseArchive success")
		return totalcmd.SUCCESS
	}

	fmt.Println("Nothing to close?")
	return totalcmd.E_ECLOSE
}

func main() {
	// Need a main function to make CGO compile package as C shared library
}

// go build -tags lib -buildmode=c-shared -o golib.a main.go
