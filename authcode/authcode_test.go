package authcode_test

import (
	. "github.com/kyoukaya/hoxy/authcode"
	"testing"
)

var (
	sMicaQueueEnc = `#fmYI1czkZTxQD6vzNZwNka6TsDL3RO/OPjznarYZD3QtVcwzp60SuvigeTvDbZMRv/zvHGFiTpZtuCpk/vW0bWXl0upozF83TSUZqpPHA7Cbdgz3GWaBfAAMUnKUFsF0rm3fEEQFPft/uIh68++eptmj5RAJmr4yH8A/e+QCMvbcH16fblOCeaZmz7cqjLEQ9N9yj/+69YHdPP55k99lc5PRzOuQB/KSy+M4lyJrdNHhSOMplfqQiViGq/ChMPxbdMu3KW8+GfvPC4nCqIZg8COcnxEc5ww6I+BTvVAC9v98psF1rUBvtaQ/nDCPUv69jgbozdoGB+vYnAdTkRxOEe9J51xjpg19QJdJfDHID3QfdRwtL5w0v4HCG5WK`
	sMicaQueueDec = `{"uid":"773091","sign":"f2457dfa713cfba324d3d6b151069392","is_username_exist":true,"app_guard_id":"554f9e675828a9f66836ad899a50a1b449cbc3aa:1567686532388:TgUBBA==","real_name":0,"authentication_url":"http:\/\/realauth.ucenter.ppgame.com\/authoriz.html?appid={0}&openid={1}&accounttype=1&language=zh"}`
)

// TestAuthcodeDecodeEncode performs decoding of a compressed input, uncompressed re-encoding
// and then another decoding round to ensure that the output is the same.
func TestAuthcodeDecodeEncode(t *testing.T) {
	dec, _, err := Decode(sMicaQueueEnc, DefaultKey)
	// Expect an expired error here if the input string is old.
	if err != nil && err != ErrExpired {
		t.Error(err)
	}

	enc, err := Encode(string(dec), DefaultKey)
	if err != nil {
		t.Error(err)
	}

	dec, _, err = Decode(string(enc), DefaultKey)
	if err != nil {
		t.Error(err)
	}

	if string(dec) != sMicaQueueDec {
		t.Errorf("Decode->Encode->Decode output different from original\nOriginal: %s\nDecoded: %s", sMicaQueueDec, string(dec))
	}
}

// TestAuthcodeGzip extends TestAuthcodeDecodeEncode to test compression before encoding
func TestAuthcodeGzip(t *testing.T) {
	dec, _, err := Decode(sMicaQueueEnc, "")
	// Expect an expired error here if the input string is old.
	if err != nil && err != ErrExpired {
		t.Error(err)
	}

	zipped, err := GzipCompress([]byte(dec))
	if err != nil {
		t.Error(err)
	}

	enc, err := Encode(string(zipped), DefaultKey)
	if err != nil {
		t.Error(err)
	}

	dec, _, err = Decode("#"+string(enc), "")
	if err != nil {
		t.Error(err)
	}

	if string(dec) != sMicaQueueDec {
		t.Errorf("Decode->Encode->Decode output different from original\nOriginal: %s\nDecoded: %s", sMicaQueueDec, string(dec))
	}
}
