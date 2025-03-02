package totalcmd

const (
	SUCCESS          = 0
	E_END_ARCHIVE    = 10 // No more files in archive
	E_NO_MEMORY      = 11 // Not enough memory
	E_BAD_DATA       = 12 // CRC error in the data of the currently unpacked file
	E_BAD_ARCHIVE    = 13 // The archive as a whole is bad, e.g. damaged headers
	E_UNKNOWN_FORMAT = 14 // Archive format unknown
	E_EOPEN          = 15 // Cannot open existing file
	E_ECREATE        = 16 // Cannot create file
	E_ECLOSE         = 17 // Error closing file
	E_EREAD          = 18 // Error reading from file
	E_EWRITE         = 19 // Error writing to file
	E_SMALL_BUF      = 20 // Buffer too small
	E_EABORTED       = 21 // Function aborted by user
	E_NO_FILES       = 22 // No files found
	E_TOO_MANY_FILES = 23 // Too many files to pack
	E_NOT_SUPPORTED  = 24 // Function not supported
)
