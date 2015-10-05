package goutil

import (
	"net/http"
	"os"
)

var (
	skipCompress = map[string]int{
		"application/x-bzip2":                     0,
		"application/x-gzip":                      0,
		"application/x-lzip":                      0,
		"application/x-lzma":                      0,
		"application/x-lzop":                      0,
		"application/x-xz":                        0,
		"application/x-compress":                  0,
		"application/x-7z-compressed":             0,
		"application/x-alz-compressed":            0,
		"application/vnd.android.package-archive": 0,
		"application/x-arj":                       0,
		"application/x-b1":                        0,
		"application/vnd.ms-cab-compressed":       0,
		"application/x-cfs-compressed":            0,
		"application/x-dar":                       0,
		"application/x-dgc-compressed":            0,
		"application/x-apple-diskimage":           0,
		"application/x-gca-compressed":            0,
		"application/x-lzh":                       0,
		"application/x-lzx":                       0,
		"application/x-rar-compressed":            0,
		"application/x-stuffit":                   0,
		"application/x-stuffitx":                  0,
		"application/x-gtar":                      0,
		"application/zip":                         0,
		"application/x-zoo":                       0,
	}
)

// compressible checks a MIME type agains a list of known compressed formats in
// order to determine if a file can likely be compressed.
func compressible(in string) bool {
	_, ok := skipCompress[in]
	return !ok
}

// FileMIME returns a file's MIME type.
func FileMIME(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	// determine mime
	b := make([]byte, 512) // all that's needed for mime per doco
	_, err = f.Read(b)
	if err != nil {
		return "", err
	}
	return http.DetectContentType(b), nil
}

// FileCompressible returns a filename's MIME type and if it is likely if it
// can be compressed.
func FileCompressible(filename string) (string, bool, error) {
	mime, err := FileMIME(filename)
	if err != nil {
		return "", false, err
	}

	return mime, compressible(mime), nil
}
