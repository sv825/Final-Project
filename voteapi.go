package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type Vote struct {
	VoteID  uint
	VoterID uint
	PollID  uint
	Choice  string
}

type VoteList struct {
	rdb *redis.Client
}

// constructor for VoteList struct
func NewVoteList() *VoteList {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return &VoteList{
		rdb: rdb,
	}
}

// Get all vote resources
func (vl *VoteList) GetVotes() ([]Vote, error) {
	keys, err := vl.rdb.Keys(ctx, "vote:*").Result()
	if err != nil {
		return nil, err
	}
	votes := make([]Vote, 0, len(keys))
	for _, key := range keys {
		vote, err := vl.GetVote(key)
		if err != nil {
			return nil, err
		}
		votes = append(votes, vote)
	}
	return votes, nil
}

// Get a single vote resource with voteID=:id.
func (vl *VoteList) GetVote(id string) (Vote, error) {
	val, err := vl.rdb.Get(ctx, id).Result()
	if err != nil {
		return Vote{}, err
	}
	var v Vote
	err = json.Unmarshal([]byte(val), &v)
	if err != nil {
		return Vote{}, err
	}
	return v, nil
}

// POST version adds one to the "database"
func (vl *VoteList) AddVote(v Vote) error {
	key := fmt.Sprintf("vote:%d", v.VoteID)
	val, err := json.Marshal(v)
	if err != nil {
		return err
	}
	err = vl.rdb.Set(ctx, key, val, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

// Returns a "health" record indicating that the vote API is functioning properly and some metadata about the API.
func (vl *VoteList) HealthCheck() string {
	return "API is functioning properly"
}
