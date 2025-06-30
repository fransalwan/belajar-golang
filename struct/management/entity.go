package management

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

func (user User) Display() string {
	return fmt.Sprintf("ID: %d, Name: %s %s, Email: %s, Active: %t",
		user.ID, user.FirstName, user.LastName, user.Email, user.IsActive)
}

type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

func (group Group) DisplayGroup() {
	fmt.Printf("Name: %s\n", group.Name)
	fmt.Printf("Member count: %d\n", len(group.Users))

	fmt.Println("Users name :")
	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}

}
