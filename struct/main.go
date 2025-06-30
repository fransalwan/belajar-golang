package main

import (
	"struct/management"
)

func main() {
	user := management.User{1, "John", "Doe", "john.doe@example.com", true}

	user2 := management.User{2, "Jane", "Smith", "jane.smith@example.com", false}

	group := management.Group{
		Name:        "Gamer",
		Admin:       user,
		Users:       []management.User{user, user2},
		IsAvailable: true,
	}

	group.DisplayGroup()

	// users := []User{user, user2}

	// group := Group{"Gamer", user, users, true}

	// displayUser1 := displayUser(user)
	// displayUser2 := displayUser(user2)

	// fmt.Println(displayUser1)
	// fmt.Println(displayUser2)

	// user2 := User{2, "Jane", "Smith", "jane.smith@example.com", false}
	// user := User{ID: 1, FirstName: "John", LastName: "Doe", Email: "john.doe@example.com", IsActive: true}

	// user.ID = 1
	// user.FirstName = "John"
	// user.LastName = "Doe"
	// user.Email = "john.doe@example.com"
	// user.IsActive = true

	// user2 := User{ID: 2, FirstName: "Jane", LastName: "Smith", Email: "jane.smith@example.com", IsActive: false}
	// user2.ID = 2
	// user2.FirstName = "Jane"
	// user2.LastName = "Smith"
	// user2.Email = "jane.smith@example.com"
	// user2.IsActive = false
	// displayGroup(group)

}

// func displayGroup(group Group) {
// 	fmt.Printf("Name : %s", group.Name)
// 	fmt.Println("")
// 	fmt.Printf("Member count: %d", len(group.Users))
// 	fmt.Println("")

// 	for _, user := range group.Users {
// 		fmt.Println(user.FirstName)

// 	}
// }

// func displayUser(user User) string {
// 	return fmt.Sprintf("ID: %d, Name: %s %s, Email: %s, Active: %t",
// 		user.ID, user.FirstName, user.LastName, user.Email, user.IsActive)

// }
