package main

import (
	"testing"
)

func TestPollList(t *testing.T) {
	pl := NewPollList()
	pl.AddPoll(Poll{
		PollID:    1,
		Title:     "Test Poll",
		Questions: []string{"Question 1", "Question 2"},
	})
	polls, err := pl.GetPolls()
	if err != nil {
		t.Errorf("Error getting polls: %v", err)
	}
	if len(polls) != 1 {
		t.Errorf("Expected 1 poll, got %d", len(polls))
	}
	poll, err := pl.GetPoll("poll:1")
	if err != nil {
		t.Errorf("Error getting poll with ID 1: %v", err)
	}
	if poll.Title != "Test Poll" {
		t.Errorf("Expected poll title to be 'Test Poll', got '%s'", poll.Title)
	}
	if len(poll.Questions) != 2 {
		t.Errorf("Expected poll to have 2 questions, got %d", len(poll.Questions))
	}
}

func TestPoll(t *testing.T) {
	p := Poll{
		PollID:    1,
		Title:     "Test Poll",
		Questions: []string{"Question 1", "Question 2"},
	}
	if p.PollID != 1 {
		t.Errorf("Expected PollID to be 1, got %d", p.PollID)
	}
	if p.Title != "Test Poll" {
		t.Errorf("Expected Title to be 'Test Poll', got '%s'", p.Title)
	}
	if len(p.Questions) != 2 {
		t.Errorf("Expected Questions to have length 2, got %d", len(p.Questions))
	}
}
