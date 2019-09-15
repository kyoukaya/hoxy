// Authcode.go contains modified procedures written by ihciah in
// https://github.com/ihciah/GFHelper/blob/master/cipher/authcode.go

// Copyright 2019 ihciah <ihciah@gmail.com>
//
// Licensed under the GNU Affero General Public License, Version 3.0
// (the "License"); you may not use this file except in compliance with the
// License.
// You may obtain a copy of the License at
//
//     https://www.gnu.org/licenses/agpl-3.0.html
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package authcode

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"crypto/rc4"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var (
	// ErrExpired is returned when an expired message is decoded. Typically not a fatal error.
	ErrExpired = errors.New("deadline exceed")
	// ErrWrongMac is returned when the MAC of the decoded body does not match the MAC prepended to the packet.
	ErrWrongMac = errors.New("wrong message authentication code")
)

const (
	ttl int64 = 3600
	// DefaultKey contains the initial key used when authenticating with the server.
	DefaultKey = "yundoudou"
)

// Return a md5 string(lower case) of a input string
func md5fromString(s string) string {
	return md5fromBytes([]byte(s))
}

// Return a md5 string of input bytes
func md5fromBytes(b []byte) string {
	hasher := md5.New()
	hasher.Write(b)
	return hex.EncodeToString(hasher.Sum(nil))
}

// Return the RC4 encrypted/decrypted bytes of a input bytes and a key string
func transformRC4(data []byte, key string) (output []byte, err error) {
	var c *rc4.Cipher
	if c, err = rc4.NewCipher([]byte(key)); err != nil {
		return nil, err
	}
	output = make([]byte, len(data))
	c.XORKeyStream(output, data)
	return
}

// GzipCompress returns gzip compressed bytes
func GzipCompress(inputs []byte) ([]byte, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	defer gz.Close()

	if _, err := gz.Write(inputs); err != nil {
		return nil, err
	}

	err := gz.Flush()
	return buf.Bytes(), err
}

// GzipDecompress returns gzip decompressed bytes
func GzipDecompress(inputs []byte) ([]byte, error) {
	buf := bytes.NewBuffer(inputs)
	r, err := gzip.NewReader(buf)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	plain, err := ioutil.ReadAll(r)
	return plain, err
}

// Encrypt the body string to bytes
func encode(body, key string) ([]byte, error) {
	hash := md5fromString(key)
	x := md5fromString(hash[:16])
	realKeyX := md5fromString(hash[16:])
	realKey := realKeyX + md5fromString(realKeyX)
	expiry := strconv.FormatInt(time.Now().Unix()+ttl, 10)
	mac := md5fromString(body + x)[:16]
	packet := expiry + mac + body
	return transformRC4([]byte(packet), realKey)
}

// Encode encrypts the body string to base64 string
func Encode(body, key string) (string, error) {
	b64 := base64.StdEncoding
	b, err := encode(body, key)
	if err != nil {
		return "", err
	}
	return b64.EncodeToString(b), nil
}

// Decrypt the enc bytes, returning the decoded bytes, and expiry.
func decode(enc []byte, key string) ([]byte, int64, error) {
	hash := md5fromString(key)
	realKeyX := md5fromString(hash[16:])
	realKey := realKeyX + md5fromString(realKeyX)
	decrypted, err := transformRC4(enc, realKey)
	if err != nil {
		return []byte{}, -1, err
	}
	deadline, err := strconv.ParseInt(string(decrypted[:10]), 10, 64)
	if err != nil {
		return []byte{}, deadline, err
	}
	var retErr error
	if time.Now().Unix() > deadline {
		retErr = ErrExpired
	}

	body := decrypted[26:]
	x := md5fromString(hash[:16])
	if mac := string(decrypted[10:26]); md5fromBytes(append(body, []byte(x)...))[:16] != mac {
		return []byte{}, deadline, ErrWrongMac
	}
	return body, deadline, retErr
}

// Decode decrypts the enc b64 string
func Decode(data, key string) ([]byte, int64, error) {
	var compressed bool
	var enc []byte
	var dec []byte
	var err error
	var expiry int64

	if key == "" {
		key = DefaultKey
	}

	if len(data) == 0 {
		return []byte{}, -1, nil
	}

	data = strings.Trim(data, "\n\r ")

	// Data we receive from the servers are always preceded by a '#' character.
	// Data that the client sends is always URL encoded
	if strings.HasPrefix(data, "#") {
		compressed = true
		data = data[1:]
	} else {
		compressed = false
		data, err = url.PathUnescape(data)
		if err != nil {
			return []byte{}, -1, err
		}
	}

	var base64coder *base64.Encoding
	if data[len(data)-1] == '=' {
		base64coder = base64.StdEncoding
	} else {
		base64coder = base64.RawStdEncoding
	}

	if enc, err = base64coder.DecodeString(data); err != nil {
		return []byte{}, -1, err
	}

	dec, expiry, err = decode(enc, key)

	if compressed {
		var err2 error
		dec, err2 = GzipDecompress(dec)
		if err != nil && err2 != nil {
			// Concatenate error
			err = fmt.Errorf("%s\n%s", err.Error(), err2.Error())
		}
	}

	return dec, expiry, err
}
