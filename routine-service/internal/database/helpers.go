package database

import (
	"bytes"
	"encoding/gob"
	"fmt"
	pb "github.com/flavioesteves/wizer-app/routine/proto"
)

func encodeExercise(exercise []*pb.Exercise) ([]byte, error) {
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	err := enc.Encode(exercise)
	if err != nil {
		return nil, fmt.Errorf("failed to encode exercise: %w", err)
	}
	return buf.Bytes(), nil
}

func decodeExercise(data []byte) ([]*pb.Exercise, error) {
	var exercises []*pb.Exercise
	dec := gob.NewDecoder(bytes.NewReader(data))
	err := dec.Decode(&exercises)
	if err != nil {
		return nil, fmt.Errorf("failed to decode exercise: %w", err)
	}
	return exercises, nil
}
