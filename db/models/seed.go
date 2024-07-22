package main

import (
	"fmt"

	"github.com/jetnoli/go-admin/db/models/user"
	userCredentials "github.com/jetnoli/go-admin/db/models/user/credentials"
)

func main() {
	fmt.Println("Starting Seeding...")

	user.Seed()
	userCredentials.Seed()

	fmt.Println("Seeding Complete")
}
