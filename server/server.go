package server

import (
	"net/http"
	"time"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email        string      `json:"email"`
	PasswordHash []byte      `json:"-"`
	CreatedAt    time.Time   `json:"createdAt"`
	Keys         []PublicKey `json:"publicKeys"`
}

type PublicKey struct {
	Public     bool      `json:"public"`
	Value      string    `json:"value"`
	Expiration time.Time `json:"expiresAt"`
}

var MemoryDB map[string]*User

func Serve(port string) {
	MemoryDB = make(map[string]*User)

	m := martini.Classic()

	m.Use(render.Renderer())

	m.Get("/", func(r render.Render) {
		r.JSON(200, struct{}{})
	})

	m.Post("/users", register)

	m.RunOnAddr(":" + port)
}

func register(r render.Render, hr *http.Request) {
	email := hr.FormValue("email")
	password := hr.FormValue("password")

	if _, ok := MemoryDB[email]; ok {
		r.Status(400)
		return
	}
	if email == "" || password == "" {
		r.Status(400)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		panic(err)
	}

	user := &User{
		Email:        email,
		PasswordHash: hashedPassword,
		CreatedAt:    time.Now(),
		Keys:         []PublicKey{},
	}

	MemoryDB[email] = user

	r.JSON(200, user)
}
