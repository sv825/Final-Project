package main

import (
	"encoding/json"
	"fmt"
)

type Poll struct {
	PollID    uint
	Title     string
	Questions []string
}

type PollList struct {
	rdb *redis.Client
}

// constructor for PollList struct
func NewPollList() *PollList {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return &PollList{
		rdb: rdb,
	}
}

// Get all poll resources
func (pl *PollList) GetPolls() ([]Poll, error) {
	keys, err := pl.rdb.Keys(ctx, "poll:*").Result()
	if err != nil {
		return nil, err
	}
	polls := make([]Poll, 0, len(keys))
	for _, key := range keys {
		poll, err := pl.GetPoll(key)
		if err != nil {
			return nil, err
		}
		polls = append(polls, poll)
	}
	return polls, nil
}

// Get a single poll resource with pollID=:id.
func (pl *PollList) GetPoll(id string) (Poll, error) {
	val, err := pl.rdb.Get(ctx, id).Result()
	if err != nil {
		return Poll{}, err
	}
	var p Poll
	err = json.Unmarshal([]byte(val), &p)
	if err != nil {
		return Poll{}, err
	}
	return p, nil
}

// POST version adds one to the "database"
func (pl *PollList) AddPoll(p Poll) error {
	key := fmt.Sprintf("poll:%d", p.PollID)
	val, err := json.Marshal(p)
	if err != nil {
		return err
	}
	err = pl.rdb.Set(ctx, key, val, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

// Returns a "health" record indicating that the poll API is functioning properly and some metadata about the API.
func (pl *PollList) HealthCheck() string {
	return "API is functioning properly"
}
