package webscreen

import (
	"bytes"
	"github.com/tebeka/selenium"
	"io"
	"os"
)

type Runner struct {
	wd selenium.WebDriver
}

func (r *Runner) WebDriver() selenium.WebDriver {
	return r.wd
}

func (r *Runner) SetSize(width, height int) error {
	return r.wd.ResizeWindow("", width, height)
}

func (r *Runner) Get(url string) error {
	return r.wd.Get(url)
}

func (r *Runner) Screenshot(path string) error {
	if pic, err := r.wd.Screenshot(); err == nil {
		if f, err := os.Create(path); err == nil {
			defer f.Close()
			if _, err := io.Copy(f, bytes.NewReader(pic)); err != nil {
				return err
			}
			return nil
		} else {
			return err
		}
	} else {
		return err
	}
}
