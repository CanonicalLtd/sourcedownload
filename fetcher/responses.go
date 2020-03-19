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

// ErrorResponse is the error message response from the API
type ErrorResponse struct {
	Error string `json:"error"`
}

// SnapResponse is the response from fetching a snap revision
type SnapResponse struct {
	Snap Snap `json:"snap"`
}

// Snap holds the details of a snap
type Snap struct {
	Name         string    `json:"name"`
	Revision     string    `json:"revision"`
	Arch         string    `json:"arch"`
	Type         string    `json:"snap_type"`
	PackageCount int       `json:"packages_count"`
	Packages     []Package `json:"packages"`
	SnapYAML     string    `json:"snapcraft"`
}

// Package is the package details
type Package struct {
	Arch           string      `json:"arch"`
	BinaryName     string      `json:"binary_name"`
	BinaryVersion  string      `json:"binary_version"`
	SourceName     string      `json:"source_name"`
	SourceVersion  string      `json:"source_version"`
	SourceFileURLs []string    `json:"source_file_urls"`
	Copyrights     []Copyright `json:"copyrights"`
}

// Copyright is the license details
type Copyright struct {
	Base64 string `json:"copyright"`
}
