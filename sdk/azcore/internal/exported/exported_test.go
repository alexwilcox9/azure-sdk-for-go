//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package exported

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestNopCloser(t *testing.T) {
	nc := NopCloser(strings.NewReader("foo"))
	if err := nc.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestHasStatusCode(t *testing.T) {
	if HasStatusCode(nil, http.StatusAccepted) {
		t.Fatal("unexpected success")
	}
	if HasStatusCode(&http.Response{}) {
		t.Fatal("unexpected success")
	}
	if HasStatusCode(&http.Response{StatusCode: http.StatusBadGateway}, http.StatusBadRequest) {
		t.Fatal("unexpected success")
	}
	if !HasStatusCode(&http.Response{StatusCode: http.StatusOK}, http.StatusAccepted, http.StatusOK, http.StatusNoContent) {
		t.Fatal("unexpected failure")
	}
}

func TestPayload(t *testing.T) {
	const val = "payload"
	resp := &http.Response{
		Body: io.NopCloser(strings.NewReader(val)),
	}
	b, err := Payload(resp)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != val {
		t.Fatalf("got %s, want %s", string(b), val)
	}
	b, err = Payload(resp)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != val {
		t.Fatalf("got %s, want %s", string(b), val)
	}
}

func TestNopClosingBytesReader(t *testing.T) {
	const val1 = "the data"
	ncbr := &nopClosingBytesReader{s: []byte(val1)}
	b, err := io.ReadAll(ncbr)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != val1 {
		t.Fatalf("got %s, want %s", string(b), val1)
	}
	const val2 = "something else"
	ncbr.Set([]byte(val2))
	b, err = io.ReadAll(ncbr)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != val2 {
		t.Fatalf("got %s, want %s", string(b), val2)
	}
	if err = ncbr.Close(); err != nil {
		t.Fatal(err)
	}
	// seek to beginning and read again
	i, err := ncbr.Seek(0, io.SeekStart)
	if err != nil {
		t.Fatal(err)
	}
	if i != 0 {
		t.Fatalf("got %d, want %d", i, 0)
	}
	b, err = io.ReadAll(ncbr)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != val2 {
		t.Fatalf("got %s, want %s", string(b), val2)
	}
	// seek to middle from the end
	i, err = ncbr.Seek(-4, io.SeekEnd)
	if err != nil {
		t.Fatal(err)
	}
	if l := int64(len(val2)) - 4; i != l {
		t.Fatalf("got %d, want %d", l, i)
	}
	b, err = io.ReadAll(ncbr)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != "else" {
		t.Fatalf("got %s, want %s", string(b), "else")
	}
	// underflow
	_, err = ncbr.Seek(-int64(len(val2)+1), io.SeekCurrent)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
}
