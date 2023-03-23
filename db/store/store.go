package store

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ErrorNotFound = errors.New("not found")

type Store struct {
	basename string     // base name of the segment files
	len      uint       // number of segments currently written to
	lim      uint       // maximum length in bytes of each segment
	segments []*Segment // pointers to all segments
}

func NewStore(basename string, lim uint) (*Store, error) {
	segments := make([]*Segment, 0)
	segmentName := formatSegmentName(basename, 0)
	initialSegment, err := newSegment(segmentName, lim)
	if err != nil {
		return nil, err
	}
	segments = append(segments, initialSegment)
	return &Store{
		basename: basename,
		len:      1,
		segments: segments,
	}, nil
}

// sets the key-value in current segment, or a new segment if doing so would put us over lim
func (s *Store) Set(key, line string) error {
	currentSegment := s.segments[s.len-1]
	entryLen := len(key + line + "\n")
	if currentSegment.len+uint(entryLen) > currentSegment.lim {
		err := s.addNewSegment()
		if err != nil {
			return err
		}
		currentSegment = s.segments[s.len-1]
	}
	// if the entry would put us over the limit, we need a new segment
	err := currentSegment.set(key, line)
	if err != nil {
		return err
	}
	return nil
}

// adds an empty segment to the store
func (s *Store) addNewSegment() error {
	formattedName := formatSegmentName(s.basename, s.len)
	newSegment, err := newSegment(formattedName, s.lim)
	if err != nil {
		return err
	}
	s.len++
	s.segments = append(s.segments, newSegment)
	return nil
}

func (s *Store) Get(key string) (string, error) {
	for i := s.len; i > 0; i-- {
		// start at the last (most recent) segment
		currentSegment := s.segments[i-1]
		fmt.Println(i, currentSegment.filename, currentSegment.hash)
		// if it doesn't have what we're looking for, we will need to look back to older segments
		line, err := currentSegment.get(key)
		if err != nil && !errors.Is(err, ErrorNotFound) {
			// continue
			return "", err
		}
		if line != "" {
			return line, nil
		}
	}
	return "", ErrorNotFound
}

// compaction should compact all segments
func (s *Store) Compaction() error {
	for _, segment := range s.segments {
		err := segment.compaction()
		if err != nil {
			return err
		}
	}
	return nil
}

// the merge function combines neighbor segments
// it does not touch the final segment
func (s *Store) Merge() error {
	i, err := s.getLowestSegmentIndex()
	if err != nil {
		return err
	}
	for ; uint(i) < s.len-1; i++ {
		curr := s.segments[i]
		next := s.segments[i+1]
		err := s.mergeSegments(curr, next)
		if err != nil {
			return err
		}
		// is this better handled by mergeSegments?
		err = s.deleteSegment(uint(i))
		if err != nil {
			return err
		}
	}
	return nil
}

// mergeSegments is responsible for merging two compacted segments and deleting the extra
func (s *Store) mergeSegments(curr, next *Segment) error {
	for key := range curr.hash {
		if _, present := next.hash[key]; !present {
			val, err := curr.get(key)
			if err != nil {
				return err
			}
			err = next.set(key, val)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// delete removes the file, decrements the length and updates the segments slice
func (s *Store) deleteSegment(i uint) error {
	segmentToDelete := s.segments[i]
	err := os.Remove(segmentToDelete.filename)
	if err != nil {
		return err
	}
	s.len--
	s.segments = append(s.segments[:i], s.segments[i+1:]...)
	return nil
}

func formatSegmentName(basename string, len uint) string {
	return fmt.Sprintf(basename+"-%d.txt", len)
}

func (s *Store) getLowestSegmentIndex() (int, error) {
	withoutPrefix := strings.TrimPrefix(s.segments[0].filename, s.basename+"-")
	withoutSuffix := strings.TrimSuffix(withoutPrefix, ".txt")
	i, err := strconv.Atoi(withoutSuffix)
	if err != nil {
		return 0, err
	}
	return i, nil
}
