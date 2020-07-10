package target

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func FindAllFiles(pathname string, flist *[]string) error {
	ps := string(os.PathSeparator)
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		return err
	}
	for _, fi := range rd {
		fname := filepath.Join(pathname, fi.Name())
		if fi.IsDir() {
			err = FindAllFiles(fname, flist)
			if err != nil {
				return err
			}
			fname += ps
		}

		*flist = append(*flist, fname)
	}
	return nil
}

func RemovePrefix(fnarr []string, prefix string) []string {
	pr := 0
	ps := string(os.PathSeparator)
	tmparr := make([]string, len(fnarr))

	if prefix != "." && prefix != "."+ps {
		pr = len(prefix)
	}
	for i := 0; i < len(fnarr); i++ {
		tmparr[i] = fnarr[i][pr:]
	}
	return tmparr
}
