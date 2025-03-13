package mongo

import (
	"context"
	"errors"
	"net/http"

	"github.com/EduartePaiva/kubernetes-authentication-microservices/common"
	"github.com/EduartePaiva/kubernetes-authentication-microservices/users-api/db/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	sdk "go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoDB struct {
	client *sdk.Client
}

const (
	mongoDbName   = "kubernetes"
	mongoColUsers = "users"
)

func NewMongoAction(client *sdk.Client) *MongoDB {
	return &MongoDB{client: client}
}

type createUser struct {
	HashedPassword string `bson:"hashed_password"`
	Email          string `bson:"email"`
}

func (m *MongoDB) CreateUser(ctx context.Context, email, hashedPassword string) (models.InsertUserResult, error) {
	res, err := m.client.Database(mongoDbName).Collection(mongoColUsers).InsertOne(ctx, createUser{
		HashedPassword: hashedPassword,
		Email:          email,
	})
	if err != nil {
		return models.InsertUserResult{}, err
	}
	// cast id into a bson.ObjectID
	id := res.InsertedID.(bson.ObjectID)

	return models.InsertUserResult{
		ID:    id.String(),
		Email: email,
	}, nil
}

func (m *MongoDB) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	user := models.User{}
	err := m.client.
		Database(mongoDbName).
		Collection(mongoColUsers).
		FindOne(ctx, bson.D{{Key: "email", Value: email}}).
		Decode(&user)

	if errors.Is(err, sdk.ErrNoDocuments) {
		err = common.HttpError{
			Code:    http.StatusNotFound,
			Message: "user not found",
		}
	}
	return user, err
}
