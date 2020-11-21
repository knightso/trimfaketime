package trimfaketime

import (
	"bufio"
	"bytes"
	"io"
)

var playbackHeaderPrefix = []byte{0, 0, 'P', 'B'}

func TrimFakeTime(dst io.Writer, src io.Reader) error {

	r := bufio.NewReader(src)

	for {
		data, err := r.ReadBytes(0)
		if err != nil {
			if _, err := dst.Write(data); err != nil {
				return err
			}

			if err == io.EOF {
				return nil
			}

			return err
		}

		if _, err := dst.Write(data[:len(data)-1]); err != nil {
			return err
		}

		if err := r.UnreadByte(); err != nil {
			return err
		}

		header, err := r.Peek(4)
		if err != nil {
			return err
		}

		if bytes.Equal(header, playbackHeaderPrefix) {
			// skip playback header
			var header [16]byte
			if n, err := r.Read(header[:]); err != nil {
				if _, err := dst.Write(header[:n]); err != nil {
					return err
				}

				if err == io.EOF {
					return nil
				}

				return err
			}
		}
	}
}
