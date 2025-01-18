package main

import "fmt"

// observer pattern
type Observer interface {
	update(content string, user string)
	getID() string
}
type Follower struct {
	name string
}

func (f *Follower) getID() string {
	return f.name
}
func (f *Follower) update(name, content string) {
	fmt.Println("notify to follower about new content :", content, f.name)

}

// Subject
type Subject interface {
	follow(Observer)
	unfollow(Observer)
	notifyFollower(string)
}

type User struct {
	name string
	list []Observer
}

func newUser(name string) *User {
	return &User{name: name}
}
func (u *User) follow(o Observer) {
	u.list = append(u.list, o)
	fmt.Printf("%s started following %s\n", o.getID(), u.name)

}

func (u *User) unfollow(o Observer) {
	u.list = removeFromSlice(u.list, o)
	fmt.Printf("%s stopped following %s\n", o.getID(), u.name)

}
func (u *User) notifyAll(content string) {
	for _, observer := range u.list {
		observer.update(content, u.name)
	}

}

func (u *User) postUpdate(content string) {
	fmt.Printf("%s posted: %s\n", u.name, content)
	u.notifyAll(content)

}

func removeFromSlice(list []Observer, o Observer) []Observer {

	l := len(list)
	for i, observer := range list {
		if observer.getID() == o.getID() {
			list[l-1], list[i] = list[i], list[l-1]

			return list[:l-1]
		}
	}

	return list
}
func main() {

	user1 := newUser("John")
	user2 := newUser("krish")

	follower1 := &Follower{
		name: "Follower 1",
	}
	follower2 := &Follower{name: "Follower 2"}
	follower3 := &Follower{name: "Follower 3"}

	user1.follow(follower1)
	user1.follow(follower2)
	user2.follow(follower3)

	user1.postUpdate("hello all, posting an content")
	user2.postUpdate("posting another content")

	user1.unfollow(follower1)
}
