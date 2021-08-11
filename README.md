# twitter
System Design Example

Designing Twitter
Let's design a Twitter-like social networking service. Users of the service will be able to post tweets, follow other people, and favorite tweets.

We'll cover the following
1. What is Twitter?
2. Requirements and Goals of the System
3. Capacity Estimation and Constraints
4. System APIs
5. High Level System Design
6. Database Schema
7. Data Sharding
8. Cache
9. Timeline Generation
10. Replication and Fault Tolerance
11. Load Balancing
12. Monitoring
13. Extended Requirements

1. What is Twitter?
Twitter is an online social networking service where users post and read short 140-character messages called "tweets." Registered users can post and read tweets, but those who are not registered can only read them. Users access Twitter through their website interface, SMS, or mobile app.

We will be designing a simpler version of Twitter with the following requirements:
Functional Requirements

Users should be able to post new tweets.
A user should be able to follow other users.
Users should be able to mark tweets as favorites.
The service should be able to create and display a user’s timeline consisting of top tweets from all the people the user follows.
Tweets can contain photos and videos.
Non-functional Requirements

Our service needs to be highly available.
Acceptable latency of the system is 200ms for timeline generation.
Consistency can take a hit (in the interest of availability); if a user doesn’t see a tweet for a while, it should be fine.
Extended Requirements

Searching for tweets.
Replying to a tweet.
Trending topics – current hot topics/searches.
Tagging other users.
Tweet Notification.
Who to follow? Suggestions?
Moments.

TODO

