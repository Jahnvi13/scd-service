package proto

import (
	"testing"
)

func TestJobMessage(t *testing.T) {
	job := &Job{
		Id:      "job1",
		Version: 1,
	}
	if job.Id != "job1" {
		t.Errorf("Expected Id to be job1, got %s", job.Id)
	}
}
