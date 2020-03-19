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
	"encoding/json"
	"fmt"
	"net/http"
	url2 "net/url"
	"os"
	"path"
	"strings"
)

// FetchService is the interface to fetch the packages
type FetchService interface {
	Run(snap string, revision int) error
}

// Fetcher implements the fetch service
type Fetcher struct {
	URL  *url2.URL
	Path string
}

// NewFetcher creates a new fetcher service
func NewFetcher(url *url2.URL, path string) *Fetcher {
	return &Fetcher{
		URL:  url,
		Path: path,
	}
}

// Run downloads the sources for a snap revision
func (f Fetcher) Run(snap string, revision int) error {
	// Get the details of the snap from the API
	snapResp, err := f.snapDetails(snap, revision)
	if err != nil {
		return err
	}

	// Download the packages for the snap
	for _, pp := range snapResp.Snap.Packages {
		p := path.Join(f.Path, pp.BinaryName, pp.BinaryVersion)
		if err := os.MkdirAll(p, os.ModePerm); err != nil {
			return err
		}

		for _, u := range pp.SourceFileURLs {
			// Get the filename of the download URL
			filename := path.Base(u)
			filepath := path.Join(p, filename)

			// Download the file
			fmt.Printf("\r%s", strings.Repeat(" ", 40))
			fmt.Printf("%s (%s)", pp.BinaryName, pp.BinaryVersion)
			if err := DownloadFile(filepath, u); err != nil {
				return err
			}
		}
	}

	return nil
}

func (f Fetcher) snapDetails(snap string, revision int) (*SnapResponse, error) {
	// Get the details of the snap from the API
	u := fmt.Sprintf("%s/v1/snaps/%s/%d", f.URL.String(), snap, revision)
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		var e ErrorResponse
		json.NewDecoder(resp.Body).Decode(&e)
		return nil, fmt.Errorf(e.Error)
	}

	defer resp.Body.Close()

	// Decode the JSON message
	var s SnapResponse
	if err := json.NewDecoder(resp.Body).Decode(&s); err != nil {
		return nil, err
	}
	return &s, nil
}
