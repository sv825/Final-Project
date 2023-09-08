package main

import (
	"testing"
)

func TestVoteList(t *testing.T) {
	vl := NewVoteList()
	vl.AddVote(Vote{
		VoteID:  1,
		VoterID: 1,
		PollID:  1,
		Choice:  "A",
	})
	votes, err := vl.GetVotes()
	if err != nil {
		t.Errorf("Error getting votes: %v", err)
	}
	if len(votes) != 1 {
		t.Errorf("Expected 1 vote, got %d", len(votes))
	}
	vote, err := vl.GetVote("vote:1")
	if err != nil {
		t.Errorf("Error getting vote with ID 1: %v", err)
	}
	if vote.Choice != "A" {
		t.Errorf("Expected vote choice to be 'A', got '%s'", vote.Choice)
	}
}

func TestVote(t *testing.T) {
	v := Vote{
		VoteID:  1,
		VoterID: 1,
		PollID:  1,
		Choice:  "A",
	}
	if v.VoteID != 1 {
		t.Errorf("Expected VoteID to be 1, got %d", v.VoteID)
	}
	if v.VoterID != 1 {
		t.Errorf("Expected VoterID to be 1, got %d", v.VoterID)
	}
	if v.PollID != 1 {
		t.Errorf("Expected PollID to be 1, got %d", v.PollID)
	}
	if v.Choice != "A" {
		t.Errorf("Expected Choice to be 'A', got '%s'", v.Choice)
	}
}
