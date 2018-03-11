// The MIT License (MIT)
//
// Copyright (c) 2018 Junpei Kawamoto
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"context"
	"os"

	"github.com/jkawamoto/go-pixeldrain/client"
	"github.com/jkawamoto/go-pixeldrain/client/file"
)

// Upload the given file to PixelDrain with the given client and under the given context.
// If cli is nil, the default client will be used. If a name is given, the uploaded file will be renamed name.
// After the upload succeeds, an ID associated with the uploaded file will be returned.
func Upload(ctx context.Context, cli *client.PixelDrain, fp *os.File, name string) (id string, err error) {

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	if cli == nil {
		cli = client.Default
	}
	if name == "" {
		name = fp.Name()
	}

	res, err := cli.File.PostFile(file.NewPostFileParamsWithContext(ctx).WithFile(*fp).WithName(&name))
	if err != nil {
		return
	}

	id = res.Payload.ID
	return

}
