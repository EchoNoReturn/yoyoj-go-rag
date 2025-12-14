package entity_test

import (
	"testing"
	"yoyoj-go-rag/internal/user/domain/entity"
	"yoyoj-go-rag/internal/user/domain/valueobject"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		name         string
		username     string
		userEmail    string
		userPassword string
		want         bool
	}{
		{
			name:         "Valid User",
			username:     "testuser",
			userEmail:    "testuser@example.com",
			userPassword: "ValidPass123",
			want:         false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := entity.NewUser(tt.username, tt.userEmail, tt.userPassword)
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if user.Username != tt.username {
				t.Errorf("expected username %s, got %s", tt.username, user.Username)
			}
			if user.Email != tt.userEmail {
				t.Errorf("expected email %s, got %s", tt.userEmail, user.Email)
			}

			hashedPassword, err := valueobject.NewHashedPassword(tt.userPassword)
			if err != nil {
				t.Fatalf("failed to hash password for comparison: %v", err)
			}
			if !hashedPassword.Verify(tt.userPassword) {
				t.Errorf("password verification failed")
			}
		})
	}
}

func TestValidatePasswordType(t *testing.T) {
	tests := []struct {
		name     string
		password string
		valid    bool
	}{
		{"Short password", "Short1", false},
		{"Too long password", "ThisIsAVeryLongPasswordThatExceedsTheMaximumAllowedLength1234567890ThisIsAVeryLongPasswordThatExceedsTheMaximumAllowedLength1234567890ThisIsAVeryLongPasswordThatExceedsTheMaximumAllowedLength1234567890ThisIsAVeryLongPasswordThatExceedsTheMaximumAllowedLength1234567890ThisIsAVeryLongPasswordThatExceedsTheMaximumAllowedLength1234567890ThisIsAVeryLongPasswordThatExceedsTheMaximumAllowedLength1234567890", false},
		{"Valid password", "ValidPass123", true},
		{"Invalid characters", "Invalid#Pass!", false},
	}

	for _, tt := range tests {
		t.Helper()
		t.Run(tt.name, func(t *testing.T) {
			valid, err := entity.ValidPasswordType(tt.password)
			if valid != tt.valid {
				t.Errorf("expected valid=%v, got %v", tt.valid, valid)
			}
			if !tt.valid && err == nil {
				t.Errorf("expected an error for invalid password, got nil")
			}
		})
	}
}

func TestUser_SamePassword(t *testing.T) {
	t.Helper()
	tests := []struct {
		name         string
		password		 string
		comparePassword string
		want         bool
	}{
		{
			name:         "Same Passwords",
			password:		 "Password123",
			comparePassword: "Password123",
			want:         true,
		},
		{
			name:         "Different Passwords",
			password:		 "Password123",
			comparePassword: "Password124",
			want:         false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := entity.NewUser("testuser", "testuser@example.com", tt.password)
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if user.ValidatePassword(tt.comparePassword) != tt.want {
				t.Errorf("expected SamePassword to return %v, got %v", tt.want, user.ValidatePassword(tt.comparePassword))
			}
		})
	}
}