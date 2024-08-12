//go:build integration

// run tests with this command: go test . --tags integration --count=1

package data

import (
	"database/sql"
	_ "embed"
	"fmt"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

//goland:noinspection SpellCheckingInspection
var (
	host     = "localhost"
	user     = "postgres"
	password = "secret"
	dbName   = "celeritas_test"
	port     = "5435"
	dsn      = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable timezone=UTC connect_timeout=5"
)

var dummyUser = User{
	FirstName: "Some",
	LastName:  "Guy",
	Email:     "me@here.com",
	Active:    1,
	Password:  "password",
}

//goland:noinspection GoUnusedGlobalVariable
var models Models
var testDB *sql.DB
var resource *dockertest.Resource
var pool *dockertest.Pool

//go:embed auth_tables.sql
var authTables string

func TestMain(m *testing.M) {
	_ = os.Setenv("DATABASE_TYPE", "postgres")
	_ = os.Setenv("UPPER_DB_LOG", "ERROR")

	p, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("could not connect to docker: %s", err)
	}
	pool = p

	opts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "13.4",
		Env: []string{
			"POSTGRES_USER=" + user,
			"POSTGRES_PASSWORD=" + password,
			"POSTGRES_DB=" + dbName,
		},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: "0.0.0.0", HostPort: port},
			},
		},
	}
	resource, err = pool.RunWithOptions(&opts)
	if err != nil {
		_ = pool.Purge(resource)
		log.Fatalf("could not start resource: %s", err)
	}

	if err := pool.Retry(func() error {
		var err error
		testDB, err = sql.Open("pgx", fmt.Sprintf(dsn, host, port, user, password, dbName))
		if err != nil {
			log.Println("Error connecting to database:", err)
			return err
		}
		return testDB.Ping()
	}); err != nil {
		_ = pool.Purge(resource)
		log.Fatalf("could not connect to docker: %s", err)
	}

	err = createTables(testDB)
	if err != nil {
		log.Fatalf("could not create tables: %s", err)
	}

	models = New(testDB)
	code := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("could not purge resource: %s", err)
	}

	os.Exit(code)

}

func createTables(db *sql.DB) error {
	_, err := db.Exec(authTables)
	if err != nil {
		return err
	}
	return nil
}

func TestUser_Table(t *testing.T) {
	s := models.Users.Table()
	if s != "users" {
		t.Error("expected 'users', got ", s)
	}
}

func TestUser_Insert(t *testing.T) {
	id, err := models.Users.Insert(dummyUser)
	if err != nil {
		t.Error("error inserting user", err)
	}

	if id == 0 {
		t.Error("expected id > 0, got ", id)
	}
}

func TestUser_Get(t *testing.T) {
	u, err := models.Users.Get(1)
	if err != nil {
		t.Error("error getting user", err)
	}
	if u.ID == 0 {
		t.Error("expected id > 0, got ", u.ID)
	}
}

func TestUser_GetAll(t *testing.T) {
	_, err := models.Users.GetAll()
	if err != nil {
		t.Error("error getting user", err)
	}
}

func TestUser_GetByEmail(t *testing.T) {
	u, err := models.Users.GetByEmail("me@here.com")
	if err != nil {
		t.Error("error getting user", err)
	}
	if u.ID == 0 {
		t.Error("expected id > 0, got ", u.ID)
	}
}

func TestUser_Update(t *testing.T) {
	u, err := models.Users.Get(1)
	if err != nil {
		t.Error("error getting user", err)
	}
	u.LastName = "Smith"
	err = u.Update(*u)
	if err != nil {
		t.Error("error updating user", err)
	}
	u, err = models.Users.Get(1)
	if err != nil {
		t.Error("error getting user", err)
	}
	if u.LastName != "Smith" {
		t.Error("expected LastName Smith, got ", u.LastName)
	}
}

func TestUser_PasswordMatches(t *testing.T) {
	u, err := models.Users.Get(1)
	if err != nil {
		t.Error("error getting user", err)
	}

	match, err := u.PasswordMatches("password")
	if err != nil {
		t.Error("error checking match", err)
	}
	if !match {
		t.Error("password do not match, got ", match)
	}

	match, err = u.PasswordMatches("123")
	if err != nil {
		t.Error("error checking match", err)
	}
	if match {
		t.Error("password matches when it should not")
	}
}

func TestUser_ResetPassword(t *testing.T) {
	err := models.Users.ResetPassword(1, "new_password")
	if err != nil {
		t.Error("error resetting password", err)
	}

	err = models.Users.ResetPassword(2, "new_password")
	if err == nil {
		t.Error("did not get an error when trying to reset password for non-existent user")
	}
}

func TestUser_Delete(t *testing.T) {
	err := models.Users.Delete(1)
	if err != nil {
		t.Error("error deleting user", err)
	}

	_, err = models.Users.Get(1)
	if err == nil {
		t.Error("did not get an error when trying to delete non-existent user")
	}
}

func TestToken_Table(t *testing.T) {
	s := models.Tokens.Table()
	if s != "tokens" {
		t.Error("expected 'tokens', got ", s)
	}
}

func TestToken_GenerateToken(t *testing.T) {
	id, err := models.Users.Insert(dummyUser)
	if err != nil {
		t.Error("error inserting user", err)
	}

	_, err = models.Tokens.GenerateToken(id, time.Hour*24*365)
	if err != nil {
		t.Error("error generating token", err)
	}
}

func TestToken_Insert(t *testing.T) {
	u, err := models.Users.GetByEmail(dummyUser.Email)
	if err != nil {
		t.Error("error getting user", err)
	}

	token, err := models.Tokens.GenerateToken(u.ID, time.Hour*24*365)
	if err != nil {
		t.Error("error generating token", err)
	}
	err = models.Tokens.Insert(*token, *u)
	if err != nil {
		t.Error("error inserting token", err)
	}
}

func TestToken_GetUserForToken(t *testing.T) {
	token := "abc"
	_, err := models.Tokens.GetUserForToken(token)
	if err == nil {
		t.Error("did not get an error when trying to get user for bad token")
	}

	u, err := models.Users.GetByEmail(dummyUser.Email)
	if err != nil {
		t.Error("error getting user", err)
	}

	_, err = models.Tokens.GetUserForToken(u.Token.PlainText)
	if err != nil {
		t.Error("error getting user with valid token", err)
	}
}

func TestToken_GetTokensForUser(t *testing.T) {
	tokens, err := models.Tokens.GetTokensForUser(1)
	if err != nil {
		t.Error("error getting tokens for user", err)
	}

	if len(tokens) > 0 {
		t.Error("expected no tokens for non-existent user, got ", tokens)
	}
}

func TestToken_Get(t *testing.T) {
	u, err := models.Users.GetByEmail(dummyUser.Email)
	if err != nil {
		t.Error("error getting user", err)
	}
	_, err = models.Tokens.Get(u.Token.ID)
	if err != nil {
		t.Error("error getting token by id", err)
	}
}

func TestToken_GetByToken(t *testing.T) {
	u, err := models.Users.GetByEmail(dummyUser.Email)
	if err != nil {
		t.Error("error getting user", err)
	}

	_, err = models.Tokens.GetByToken(u.Token.PlainText)
	if err != nil {
		t.Error("error getting token by token", err)
	}

	_, err = models.Tokens.GetByToken("123")
	if err == nil {
		t.Error("no error getting non-existent token by token", err)
	}
}

//goland:noinspection SpellCheckingInspection
var authData = []struct {
	name          string
	token         string
	email         string
	errorExpected bool
	message       string
}{
	{"invalid", "abcdefghijklmnopqrstuvwxyz", "a@here.com", true, "invalid token accepted as valid"},
	{"invalid_length", "abcdefghijklmnopqrstuvwxy", "a@here.com", true, "token of wrong length accepted as valid"},
	{"no user", "abcdefghijklmnopqrstuvwxyz", "a@here.com", true, "no user but token accepted as valid"},
	{"valid", "", "me@here.com", false, "valid token reported as invalid"},
}

func TestToken_AuthenticateToken(t *testing.T) {
	for _, tt := range authData {
		token := ""
		if tt.email == dummyUser.Email {
			user, err := models.Users.GetByEmail(tt.email)
			if err != nil {
				t.Error("error getting user", err)
			}
			token = user.Token.PlainText
		} else {
			token = tt.token
		}

		req, _ := http.NewRequest("GET", "/", nil)
		req.Header.Add("Authorization", "Bearer "+token)

		_, err := models.Tokens.AuthenticateToken(req)
		if err == nil && tt.errorExpected {
			t.Errorf("%s: %s", tt.name, tt.message)
		} else if !tt.errorExpected && err != nil {
			t.Errorf("%s: %s- %s", tt.name, tt.message, err)
		} else {
			t.Logf("passed %s", tt.name)
		}

	}
}

func TestToken_Delete(t *testing.T) {
	u, err := models.Users.GetByEmail(dummyUser.Email)
	if err != nil {
		t.Error("error getting user", err)
	}

	err = models.Tokens.DeleteByToken(u.Token.PlainText)
	if err != nil {
		t.Error("error deleting token", err)
	}
}

func TestToken_Expired_Token(t *testing.T) {
	// insert a token
	u, err := models.Users.GetByEmail(dummyUser.Email)
	if err != nil {
		t.Error("error getting user", err)
	}

	token, err := models.Tokens.GenerateToken(u.ID, -time.Minute*10)
	if err != nil {
		t.Error("error generating token", err)
	}

	err = models.Tokens.Insert(*token, *u)
	if err != nil {
		t.Error("error inserting token", err)
	}

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Authorization", "Bearer "+token.PlainText)

	_, err = models.Tokens.AuthenticateToken(req)
	if err == nil {
		t.Error("did not get an error when trying to expire token")
	}
}

func TestToken_BadHeader(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	_, err := models.Tokens.AuthenticateToken(req)
	if err == nil {
		t.Error("did not get an error when trying to authenticate request without required header")
	}

	req, _ = http.NewRequest("GET", "/", nil)
	req.Header.Add("Authorization", "abc ")
	_, err = models.Tokens.AuthenticateToken(req)
	if err == nil {
		t.Error("did not get an error when trying to authenticate request with bad required header")
	}

	newUser := User{
		FirstName: "temp",
		LastName:  "temp_last",
		Email:     "you@there.com",
		Active:    1,
		Password:  "abc",
	}
	id, err := models.Users.Insert(newUser)
	if err != nil {
		t.Error("error inserting new user", err)
	}
	token, err := models.Tokens.GenerateToken(id, 1*time.Hour)
	if err != nil {
		t.Error("error generating token", err)
	}
	err = models.Tokens.Insert(*token, newUser)
	if err != nil {
		t.Error("error inserting new user", err)
	}
	err = models.Users.Delete(id)
	if err != nil {
		t.Error("error deleting user", err)
	}

	req, _ = http.NewRequest("GET", "/", nil)
	req.Header.Add("Authorization", "Bearer "+token.PlainText)
	_, err = models.Tokens.AuthenticateToken(req)
	if err == nil {
		t.Error("did not get an error when trying to authenticate request with token for deleted user")
	}
}

func TestToken_DeleteNonExistentToken(t *testing.T) {
	err := models.Tokens.DeleteByToken("abc")
	if err != nil {
		t.Error("error deleting non-existent token", err)
	}
}

func TestToken_ValidToken(t *testing.T) {
	u, err := models.Users.GetByEmail(dummyUser.Email)
	if err != nil {
		t.Error("error getting user", err)
	}
	newToken, err := models.Tokens.GenerateToken(u.ID, 24*time.Hour)
	if err != nil {
		t.Error("error generating token", err)
	}
	err = models.Tokens.Insert(*newToken, *u)
	if err != nil {
		t.Error("error inserting new user", err)
	}

	okay, err := models.Tokens.ValidToken(newToken.PlainText)
	if err != nil {
		t.Error("error validating token", err)
	}
	if !okay {
		t.Error("token should be valid")
	}

	okay, _ = models.Tokens.ValidToken("abc")
	if okay {
		t.Error("token should not be valid")
	}

	u, err = models.Users.GetByEmail(dummyUser.Email)
	if err != nil {
		t.Error("error getting user", err)
	}
	err = models.Tokens.Delete(u.Token.ID)
	if err != nil {
		t.Error("error deleting user", err)
	}
	okay, err = models.Tokens.ValidToken(u.Token.PlainText)
	if err == nil {
		t.Error("no error validating deleted token")
	}
	if okay {
		t.Error("token should not be valid")
	}
}
