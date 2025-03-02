package main

/*
#include "totalcmd/wcxhead.h"
*/
import "C"
import (
	"fmt"
	"os"
	"time"
	"totalcmd-wcx-demo-go/totalcmd"
)

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
func OpenArchive(pArchiveData *C.tOpenArchiveData) uintptr {
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
func ReadHeader(hArcData uintptr, pHeaderData *C.tHeaderData) C.int {
	if hArcData == 0 || pHeaderData == nil {
		return totalcmd.E_EREAD
	}

	file := openedFilesMap[hArcData]
	fmt.Println("ReadHeader", file.Name(), testFileCounter)

	fileName := "example.txt"

	C.strncpy(&pHeaderData.FileName[0], C.CString(fileName), C.size_t(len(fileName)))
	pHeaderData.PackSize = 42
	pHeaderData.UnpSize = 73
	pHeaderData.FileCRC = 0x12345678

	// (year - 1980) << 25 | month << 21 | day << 16 | hour << 11 | minute << 5 | second/2
	//Make sure that:
	//year is in the four digit format between 1980 and 2100
	//month is a number between 1 and 12
	//hour is in the 24 hour format
	fileTime := time.Now()
	pHeaderData.FileTime = C.uint((fileTime.Year()-1980)<<25 | int(fileTime.Month())<<21 | fileTime.Day()<<16 | fileTime.Hour()<<11 | fileTime.Minute()<<5 | fileTime.Second()/2)

	pHeaderData.Method = 0 // Compression method (0 - no compression)
	pHeaderData.FileAttr = totalcmd.READ_ONLY | totalcmd.ARCHIVE
	pHeaderData.Flags = 0 // Флаги (можно задать при необходимости)
	pHeaderData.CmtSize = 0
	pHeaderData.CmtState = 0

	testFileCounter -= 1
	if testFileCounter <= 0 {
		return totalcmd.E_END_ARCHIVE
	}
	return totalcmd.SUCCESS
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
	var path = C.GoString(destPath)
	var name = C.GoString(destName)
	fmt.Println()
	fmt.Println("-- ProcessFile", operation, path, name)

	return totalcmd.SUCCESS
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

		defer closeF()
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
