package topicstore

import (
	"errors"
)

type TopicStoreElement struct {
	Text      string
	UpVotes   int64
	DownVotes int64
}

type TopicStore struct {
	topics        []TopicStoreElement
	hotTopics     map[int]int64
	hotTopicLimit int
	access        chan bool
}

func NewTopicStore(hotTopicLimit int) *TopicStore {
	t := &TopicStore{
		hotTopics:     make(map[int]int64),
		access:        make(chan bool, 1),
		hotTopicLimit: hotTopicLimit,
	}
	t.access <- true
	return t
}

func (t *TopicStore) validateTopicId(id int) bool {
	if (id < 0) || (id > len(t.topics)-1) {
		return false
	}
	return true
}

func (t *TopicStore) Insert(topicText string) {
	topic := TopicStoreElement{
		Text:      topicText,
		UpVotes:   0,
		DownVotes: 0,
	}
	<-t.access
	//Entering critical section
	t.topics = append(t.topics, topic)
	lastInsertId := len(t.topics) - 1

	//case where there are no hot topics yet.
	if len(t.hotTopics) < t.hotTopicLimit {
		t.hotTopics[lastInsertId] = topic.UpVotes
	}
	//Exiting critical section
	t.access <- true
}

func (t *TopicStore) GetTopics() []TopicStoreElement {
	return t.topics
}

func (t *TopicStore) UpVote(currentTopicId int) error {
	if !t.validateTopicId(currentTopicId) {
		return errors.New("Bad topic id")
	}
	t.topics[currentTopicId].UpVotes += 1
	<-t.access
	//Entering critical section
	//updating hot topics
	for topicId, numUpVotes := range t.hotTopics {
		if numUpVotes < t.topics[topicId].UpVotes {
			delete(t.hotTopics, topicId)
			t.hotTopics[currentTopicId] = t.topics[currentTopicId].UpVotes
			break
		}
	}
	//Exiting critical section
	t.access <- true
	return nil
}

func (t *TopicStore) DownVote(currentTopicId int) error {
	if !t.validateTopicId(currentTopicId) {
		return errors.New("Bad topic id")
	}
	t.topics[currentTopicId].DownVotes += 1
	return nil
}

func (t *TopicStore) GetHotTopics() []TopicStoreElement {
	result := []TopicStoreElement{}
	<-t.access
	//Entering critical section
	for topicId, _ := range t.hotTopics {
		result = append(result, t.topics[topicId])
	}
	//Exiting critical section
	t.access <- true
	return result
}
