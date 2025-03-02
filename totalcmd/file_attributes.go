package totalcmd

const (
    READ_ONLY = 0x1  // Read-only file
    HIDDEN    = 0x2  // Hidden file
    SYSTEM    = 0x4  // System file
    VOLUME_ID = 0x8  // Volume ID file
    DIRECTORY = 0x10 // Directory
    ARCHIVE   = 0x20 // Archive file
    ANY       = 0x3F // Any file
)
