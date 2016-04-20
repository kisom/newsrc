package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func writeCSource(opts *opts) error {
	licenseText, ok := licenses[opts.license]
	if !ok {
		return errUnknownLicense
	}

	if _, err := os.Stat(opts.path); err != nil {
		if !os.IsNotExist(err) && !opts.force {
			return err
		}
	} else if !opts.force {
		return errFileExists
	}

	file, err := os.Create(opts.path)
	if err != nil {
		return err
	}
	defer file.Close()

	now := time.Now()
	licenseText[0] = fmt.Sprintf(licenseText[0], now.Year())

	_, err = file.Write([]byte("/*\n"))
	if err != nil {
		return err
	}

	for _, line := range licenseText {
		line = " * " + line + "\n"
		_, err = file.Write([]byte(line))
		if err != nil {
			return err
		}
	}

	_, err = file.Write([]byte(" */\n"))
	if err != nil {
		return err
	}

	path := filepath.Base(opts.path)
	body := fmt.Sprintf(opts.body, path)
	_, err = file.Write([]byte(body))
	return err
}

func writeCCSource(opts *opts) error {
	licenseText, ok := licenses[opts.license]
	if !ok {
		return errUnknownLicense
	}

	if _, err := os.Stat(opts.path); err != nil {
		if !os.IsNotExist(err) && !opts.force {
			return err
		}
	} else if !opts.force {
		return errFileExists
	}

	file, err := os.Create(opts.path)
	if err != nil {
		return err
	}
	defer file.Close()

	now := time.Now()
	licenseText[0] = fmt.Sprintf(licenseText[0], now.Year())

	for _, line := range licenseText {
		line = "// " + line + "\n"
		_, err = file.Write([]byte(line))
		if err != nil {
			return err
		}
	}

	path := filepath.Base(opts.path)
	body := fmt.Sprintf(opts.body, path)
	_, err = file.Write([]byte(body))
	return err
}
