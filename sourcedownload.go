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

package main

import (
	"flag"
	"fmt"
	"github.com/CanonicalLtd/sourcedownload/fetcher"
	url2 "net/url"
	"os"
)

func main() {
	snap, revision, url, p := parseFlags()

	fetch := fetcher.NewFetcher(url, p)
	if err := fetch.Run(snap, revision); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Download complete")
}

func parseFlags() (string, int, *url2.URL, string) {
	var (
		snap         string
		revision     int
		pathDownload string
		url          string
	)

	flag.StringVar(&snap, "snap", "", "The name of the snap to process")
	flag.IntVar(&revision, "revision", 0, "The revision of the snap to process")
	flag.StringVar(&url, "url", "https://sources.iotdevice.io", "The URL of the Compliance Service")
	flag.StringVar(&pathDownload, "path", "download", "Location to store the download files")
	flag.Parse()

	// Validate
	if len(snap) == 0 || revision == 0 {
		fmt.Println("The `-snap` and `-revision` arguments are mandatory")
		os.Exit(1)
	}
	u, err := url2.Parse(url)
	if err != nil {
		fmt.Println("The URL of the compliance service is invalid:", err)
		os.Exit(1)
	}
	return snap, revision, u, pathDownload
}
