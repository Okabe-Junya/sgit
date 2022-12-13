package object

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/Okabe-Junya/sgit/pkg/dir"
)

func compress(r io.Reader) (io.Reader, error) {
	buf := new(bytes.Buffer)
	w := zlib.NewWriter(buf)
	defer w.Close()

	if _, err := io.Copy(w, r); err != nil {
		return nil, err
	}
	return buf, nil
}

func StoreBlob(content string) {
	header := "blob " + strconv.Itoa(len(content)) + "\x00"

	// sha1 ハッシュを計算する
	store := header + content
	h := sha1.New()
	h.Write([]byte(store))
	hs := h.Sum(nil)
	sha1_store := hex.EncodeToString(hs)

	// zlib でファイルを圧縮する
	zr, err := compress(bytes.NewReader([]byte(store)))
	if err != nil {
		fmt.Println("Error: ", err)
	}
	b := zr.(*bytes.Buffer)

	// ファイルを作成する
	dir_path := sha1_store[0:2]
	file_path := sha1_store[2:]
	dir.CreateDir(string(dir_path))
	full_path := string(dir_path) + "/" + string(file_path)

	f, err := os.Create(full_path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write(b.Bytes())
}
