package store

import (
	"bufio"
	"bytes"
	"errors"
	"os"
	"strings"
)

type Segment struct {
	filename string         // formatted off the store's basename
	hash     map[string]int // key -> byte offset of current segment
	len      int            // current length in bytes
	lim      int            // maximum length before a new segment must be made
}

// assumes formatted name already handled
func newSegment(segmentName string, lim int) (*Segment, error) {
	f, err := os.Create(segmentName)
	if err != nil {
		return nil, err
	}
	f.Close()
	return &Segment{
		filename: segmentName,
		hash:     make(map[string]int),
		len:      0,
		lim:      lim,
	}, nil
}

// relies on Store to call it only when there will be space according to the len and lim
func (s *Segment) set(key, line string) error {
	f, err := os.OpenFile(s.filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		return err
	}
	entry := []byte(key + line + "\n")
	size := len(entry)
	s.hash[key] = s.len
	s.len += size
	defer f.Close()
	_, err = f.Write(entry)
	if err != nil {
		return err
	}
	return nil
}

func (s *Segment) get(key string) (string, error) {
	f, err := os.Open(s.filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	// set seek
	var val string
	scanner := bufio.NewScanner(f)
	if offset, ok := s.hash[key]; !ok {
		return "", ErrorNotFound
	} else {
		_, err = f.Seek(int64(offset), 0)
		if err != nil {
			return "", err
		}
	}

	// scan and trim line
	scanner.Scan()
	line := scanner.Text()
	val = strings.TrimPrefix(line, key)
	err = scanner.Err()
	if err != nil {
		return "", err
	}
	if val != "" {
		return val, nil
	}
	return "", ErrorNotFound
}

// segment compaction removes duplicate keys, leaving behind only the latest
func (s *Segment) compaction() error {
	var b bytes.Buffer
	var newLen int
	newHash := make(map[string]int)
	for key := range s.hash {
		val, err := s.get(key)
		if err != nil && !errors.Is(err, ErrorNotFound) {
			return err
		}
		entry := key + val + "\n"
		newHash[key] = newLen
		newLen += len(entry)
		b.WriteString(entry)
	}
	f, err := os.Create(s.filename)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(b.Bytes())
	if err != nil {
		return err
	}
	s.hash = newHash
	s.len = newLen
	return nil
}
