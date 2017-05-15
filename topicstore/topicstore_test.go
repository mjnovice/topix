package topicstore

import (
	"fmt"
	"testing"
)

var hotTopicLimit = 20
var topicStore = NewTopicStore(hotTopicLimit)

func TestInsert(t *testing.T) {
	topicStore.Insert("hello")
	if topicStore.topics[0].Text != "hello" {
		t.Error("Insert failed.")
	}
}

func TestUpVoteUpVote(t *testing.T) {
	topicStore.UpVote(0)
	if topicStore.topics[0].UpVotes != 1 {
		t.Error("Upvote failed.")
	}
}

func TestUpVoteDownVote(t *testing.T) {
	topicStore.DownVote(0)
	if topicStore.topics[0].DownVotes != 1 {
		t.Error("Downvote failed.")
	}
}

func TestGetAllTopics(t *testing.T) {
	topics := topicStore.GetTopics()
	if !(len(topics) == 1 && topics[0].Text == "hello") {
		t.Error("GetAllTopicsFailed")
	}
}

func TestGetHotTopics(t *testing.T) {
	topics := topicStore.GetHotTopics()
	if !((len(topics) == 1) && (topics[0].Text == "hello")) {
		t.Error("GetHotTopicsFailed")
	}
}

func TestValidateTopicId(t *testing.T) {
	if !(topicStore.validateTopicId(-1) == false && topicStore.validateTopicId(1) == false) {
		t.Error("ValidateTopicId failed.")
	}
}
