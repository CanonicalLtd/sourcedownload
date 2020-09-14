/*
 * Copyright (C) 2020 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package fetcher

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory. We pass an io.TeeReader
// into Copy() to report progress on the download.
func DownloadFile(filepath string, url string) error {
	// Create the file, but give it a tmp file extension, this means we won't overwrite a
	// file until it's downloaded, but we'll remove the tmp extension once downloaded.
	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		return err
	}

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		out.Close()
		return err
	}
	defer resp.Body.Close()

	// Create our progress reporter and pass it to be used alongside our writer
	counter := &WriteCounter{}
	if _, err = io.Copy(out, io.TeeReader(resp.Body, counter)); err != nil {
		out.Close()
		return err
	}

	// The progress use the same line so print a new line once it's finished downloading
	fmt.Print("\n")

	// Close the file without defer so it can happen before Rename()
	out.Close()

	if err = os.Rename(filepath+".tmp", filepath); err != nil {
		return err
	}
	return nil
}

// DownloadSnapcraft download the snapcraft.yaml for the snap and revision
func DownloadSnapcraft(filepath string, snap Snap) error {
	if err := os.MkdirAll(filepath, os.ModePerm); err != nil {
		return err
	}

	if snap.SnapYAML == "" {
		fmt.Printf("snapcraft.yaml not available for %s (%s)", snap.Name, snap.Revision)
		return nil
	}

	name := fmt.Sprintf("snapcraft_%s_%s.yaml", snap.Name, snap.Revision)
	p := path.Join(filepath, name)

	b, err := base64.StdEncoding.DecodeString(snap.SnapYAML)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(p, b, 0644)
}
