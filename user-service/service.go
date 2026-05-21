// package main

// import "errors"

// var users []User
// var nextID = 1

// func init() {
// 	seed := User{
// 		UserID:   nextID,
// 		Name:     "seed",
// 		Email:    "seed@mail.com",
// 		Password: "hashed",
// 		Role:     "customer",
// 	}

// 	users = append(users, seed)
// 	nextID++
// }

// func Register(name, email, password, role string) (User, error) {
// 	for _, u := range users {
// 		if u.Email == email {
// 			return User{}, errors.New("email exists")
// 		}
// 	}

// 	hash, _ := HashPassword(password)

// 	u := User{
// 		UserID:   nextID,
// 		Name:     name,
// 		Email:    email,
// 		Password: hash,
// 		Role:     role,
// 	}

// 	users = append(users, u)
// 	nextID++

// 	return u, nil
// }

// func Login(email, password string) (string, error) {
// 	for _, u := range users {
// 		if u.Email == email && CheckPassword(password, u.Password) {
// 			return GenerateToken(u.UserID, u.Role), nil
// 		}
// 	}
// 	return "", errors.New("invalid login")
// }

// func GetProfile(id int) *User {
// 	for i := range users {
// 		if users[i].UserID == id {
// 			return &users[i]
// 		}
// 	}
// 	return nil
// }

// func UpdateProfile(id int, alamat string, pref string) bool {
// 	for i := range users {
// 		if users[i].UserID == id {
// 			users[i].Alamat = alamat
// 			users[i].Preferensi = pref
// 			return true
// 		}
// 	}
// 	return false
// }

// // package main

// // var users []User
// // var nextID = 1

// // func init() {
// // 	seed := User{
// // 		UserID:   nextID,
// // 		Name:     "seed",
// // 		Email:    "seed@mail.com",
// // 		Password: "hashed",
// // 		Role:     "customer",
// // 	}

// // 	users = append(users, seed)
// // 	nextID++
// // }

// // func Register(name, email, password, role string) (User, error) {
// // 	return User{}, nil
// // }

// // func Login(email, password string) (string, error) {
// // 	return "", nil
// // }

// // func GetProfile(id int) *User {
// // 	return nil
// // }

// // func UpdateProfile(id int, alamat string, pref string) bool {
// // 	return false
// // }



// // package main

// // import "errors"

// // var users []User
// // var nextID = 1

// // func init() {
// // 	seed := User{
// // 		UserID:   nextID,
// // 		Name:     "seed",
// // 		Email:    "seed@mail.com",
// // 		Password: "hashed",
// // 		Role:     "customer",
// // 	}

// // 	users = append(users, seed)
// // 	nextID++
// // }

// // func Register(name, email, password, role string) (User, error) {
// // 	for _, u := range users {
// // 		if u.Email == email {
// // 			return User{}, errors.New("email exists")
// // 		}
// // 	}

// // 	hash, _ := HashPassword(password)

// // 	u := User{
// // 		UserID:   nextID,
// // 		Name:     name,
// // 		Email:    email,
// // 		Password: hash,
// // 		Role:     role,
// // 	}

// // 	users = append(users, u)
// // 	nextID++

// // 	return u, nil
// // }

// // func Login(email, password string) (string, error) {
// // 	for _, u := range users {
// // 		if u.Email == email && CheckPassword(password, u.Password) {
// // 			return GenerateToken(u.UserID, u.Role), nil
// // 		}
// // 	}
// // 	return "", errors.New("invalid login")
// // }

// // func GetProfile(id int) *User {
// // 	for i := range users {
// // 		if users[i].UserID == id {
// // 			return &users[i]
// // 		}
// // 	}
// // 	return nil
// // }

// // func UpdateProfile(id int, alamat string, pref string) bool {
// // 	for i := range users {
// // 		if users[i].UserID == id {
// // 			users[i].Alamat = alamat
// // 			users[i].Preferensi = pref
// // 			return true
// // 		}
// // 	}
// // 	return false
// // }









// package main

// var users []User
// var nextID = 1

// func init() {
// 	seed := User{
// 		UserID:   nextID,
// 		Name:     "seed",
// 		Email:    "seed@mail.com",
// 		Password: "hashed",
// 		Role:     "customer",
// 	}

// 	users = append(users, seed)
// 	nextID++
// }

// func Register(name, email, password, role string) (User, error) {
// 	return User{}, nil
// }

// func Login(email, password string) (string, error) {
// 	return "", nil
// }

// func GetProfile(id int) *User {
// 	return nil
// }

// func UpdateProfile(id int, alamat string, pref string) bool {
// 	return false
// }

// package main

// import "errors"

// var users []User
// var nextID = 1

// func init() {
// 	seed := User{
// 		UserID:   nextID,
// 		Name:     "seed",
// 		Email:    "seed@mail.com",
// 		Password: "hashed",
// 		Role:     "customer",
// 	}

// 	users = append(users, seed)
// 	nextID++
// }

// func Register(name, email, password, role string) (User, error) {

// 	for _, u := range users {
// 		if u.Email == email {
// 			return User{}, errors.New("email exists")
// 		}
// 	}

// 	u := User{
// 		UserID:   nextID,
// 		Name:     name,
// 		Email:    email,
// 		Password: password,
// 		Role:     role,
// 	}

// 	users = append(users, u)

// 	// SIMPAN KE DATABASE
// 	_, err := DB.Exec(
// 		"INSERT INTO users(name,email,password,role) VALUES(?,?,?,?)",
// 		u.Name,
// 		u.Email,
// 		u.Password,
// 		u.Role,
// 	)

// 	if err != nil {
// 		return User{}, err
// 	}

// 	nextID++

// 	return u, nil
// }

// func Login(email, password string) (string, error) {

// 	for _, u := range users {
// 		if u.Email == email && u.Password == password {
// 			return "dummy-token", nil
// 		}
// 	}

// 	return "", errors.New("invalid login")
// }

// func GetProfile(id int) *User {

// 	for i := range users {
// 		if users[i].UserID == id {
// 			return &users[i]
// 		}
// 	}

// 	return nil
// }

// func UpdateProfile(id int, alamat string, pref string) bool {

// 	for i := range users {
// 		if users[i].UserID == id {
// 			users[i].Alamat = alamat
// 			users[i].Preferensi = pref
// 			return true
// 		}
// 	}

// 	return false
// }

package main

import "errors"

func Register(
	name string,
	email string,
	password string,
	role string,
) (User, error) {

	result, err := db.Exec(`
		INSERT INTO users(name,email,password,role)
		VALUES(?,?,?,?)
	`,
		name,
		email,
		password,
		role,
	)

	if err != nil {
		return User{}, err
	}

	id, _ := result.LastInsertId()

	user := User{
		UserID:  int(id),
		Name:    name,
		Email:   email,
		Password: password,
		Role:    role,
	}

	return user, nil
}

func Login(email string, password string) (User, error) {

	var u User

	err := db.QueryRow(`
		SELECT user_id,name,email,password,role
		FROM users
		WHERE email = ?
	`, email).Scan(
		&u.UserID,
		&u.Name,
		&u.Email,
		&u.Password,
		&u.Role,
	)

	if err != nil {
		return User{}, errors.New("login gagal")
	}

	if u.Password != password {
		return User{}, errors.New("password salah")
	}

	return u, nil
}

func GetProfile(id int) *User {

	var u User

	err := db.QueryRow(`
		SELECT user_id,name,email,role,alamat,preferensi
		FROM users
		WHERE user_id = ?
	`, id).Scan(
		&u.UserID,
		&u.Name,
		&u.Email,
		&u.Role,
		&u.Alamat,
		&u.Preferensi,
	)

	if err != nil {
		return nil
	}

	return &u
}

func UpdateProfile(
	userID int,
	alamat string,
	preferensi string,
) bool {

	_, err := db.Exec(`
		UPDATE users
		SET alamat=?, preferensi=?
		WHERE user_id=?
	`,
		alamat,
		preferensi,
		userID,
	)

	return err == nil
}
