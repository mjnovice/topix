A dull website for adding topics, with upvotes/downvotes, with in-memory data store.
Deployed at, https://powerful-refuge-23022.herokuapp.com/

TopicStore:
To display top 20 topics, I am using a map to keep track.
For storing all the topics I am using a slice analogous to an array.
Upon each Upvote/Downvote/Newtopic I check if the map can be updated and do as 
applicable, thus saving me from sorting the entire slice each time.


Assumptions:
1. The maximum number of topics that can be created is 2^64 i.e 18446744073709551616, 
given the machine on which this is hosted has the value of strconv.IntSize == 64.

