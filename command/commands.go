package command

import (
	"fmt"
	"os"
	"path"

	"github.com/sharpvik/vish/public"
)

type CD struct {
	Path string
}

func (cd *CD) Exec() (err error) {
	return os.Chdir(path.Clean(cd.Path))
}

type Version struct{}

func (version Version) Exec() (err error) {
	_, err = fmt.Println(public.ShortInfo())
	return
}

type Exit struct{}

func (exit Exit) Exec() (err error) {
	fmt.Println(public.GoodBye)
	os.Exit(0)
	return
}
