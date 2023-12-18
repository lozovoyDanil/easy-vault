package repository_test

import (
	"context"
	"reflect"
	"testing"

	"main.go/internal/model"
	"main.go/internal/repository"
	_ "modernc.org/sqlite"
)

func TestAuthSQLite_GetUser(t *testing.T) {
	// Вставка тестовых данных
	testUser := model.User{
		FullName: "John Doe",
		Email:    "johndoe@example.com",
		Password: "password123",
		Role:     "user",
		IsBanned: true,
	}
	_, err := testDB.NewInsert().Model(&testUser).Exec(context.Background())
	if err != nil {
		t.Error(err)
	}

	type args struct {
		username string
		password string
	}
	tests := []struct {
		name    string
		r       *repository.AuthSQLite
		args    args
		want    model.User
		wantErr bool
	}{
		{
			name: "successful fetch",
			r:    repository.NewAuthSQLite(testDB),
			args: args{
				username: "johndoe@example.com",
				password: "password123",
			},
			want:    testUser,
			wantErr: false,
		},
		{
			name: "failed fetch",
			r:    repository.NewAuthSQLite(testDB),
			args: args{
				username: "johndoe@example.com",
				password: "password12",
			},
			want:    model.User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.GetUser(tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthSQLite.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthSQLite.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthSQLite_UserIsBanned(t *testing.T) {
	testUser := model.User{
		FullName: "John Doe",
		Email:    "email.com",
		Password: "password123",
		Role:     "user",
		IsBanned: false,
	}
	_, err := testDB.NewInsert().Model(&testUser).Exec(context.Background())
	if err != nil {
		t.Error(err)
	}

	type args struct {
		userId int
	}
	tests := []struct {
		name    string
		r       *repository.AuthSQLite
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "successful fetch",
			r:    repository.NewAuthSQLite(testDB),
			args: args{
				userId: 1,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "successful fetch 2",
			r:    repository.NewAuthSQLite(testDB),
			args: args{
				userId: 2,
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "failed fetch",
			r:    repository.NewAuthSQLite(testDB),
			args: args{
				userId: 3,
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.UserIsBanned(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthSQLite.UserIsBanned() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AuthSQLite.UserIsBanned() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthSQLite_CreateUser(t *testing.T) {
	type args struct {
		user model.User
	}
	tests := []struct {
		name    string
		r       *repository.AuthSQLite
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "successful create",
			r:    repository.NewAuthSQLite(testDB),
			args: args{
				user: model.User{
					FullName: "John Doe",
					Email:    "email1.com",
					Password: "password123",
					Role:     "user",
					IsBanned: false,
				},
			},
			want:    3,
			wantErr: false,
		},
		{
			name: "failed create",
			r:    repository.NewAuthSQLite(testDB),
			args: args{
				user: model.User{
					FullName: "John Doe",
					Email:    "email.com",
					Password: "",
					Role:     "user",
					IsBanned: false,
				},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.CreateUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthSQLite.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AuthSQLite.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthSQLite_UserInfo(t *testing.T) {
	type args struct {
		userId int
	}
	tests := []struct {
		name    string
		r       *repository.AuthSQLite
		args    args
		want    model.User
		wantErr bool
	}{
		{
			name: "successful fetch",
			r:    repository.NewAuthSQLite(testDB),
			args: args{
				userId: 1,
			},
			want: model.User{
				Id:       1,
				FullName: "John Doe",
				Email:    "johndoe@example.com",
				Password: "password123",
				Role:     "user",
				IsBanned: true,
			},
			wantErr: false,
		},
		{
			name: "failed fetch",
			r:    repository.NewAuthSQLite(testDB),
			args: args{
				userId: 150,
			},
			want:    model.User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.UserInfo(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthSQLite.UserInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthSQLite.UserInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthSQLite_EditUser(t *testing.T) {
	type args struct {
		userId int
		input  model.UpdateUserInput
	}
	tests := []struct {
		name    string
		r       *repository.AuthSQLite
		args    args
		wantErr bool
	}{
		{
			name: "successful edit",
			r:    repository.NewAuthSQLite(testDB),
			args: args{
				userId: 1,
				input: model.UpdateUserInput{
					FullName: stringPtr("Jimmy Doe"),
					Email:    stringPtr("jimmy.com"),
				},
			},
			wantErr: false,
		},
		{
			name: "failed edit",
			r:    repository.NewAuthSQLite(testDB),
			args: args{
				userId: 1,
				input:  model.UpdateUserInput{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.EditUser(tt.args.userId, tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("AuthSQLite.EditUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func stringPtr(s string) *string {
	return &s
}
