package user_test

import (
	"context"
	"fmt"
	"testing"

	"supercollector/internal/config"
	"supercollector/internal/user"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RepoTestSuite struct {
	suite.Suite
	conn *mongo.Client
	db   *mongo.Database
	coll *mongo.Collection
}

func (s *RepoTestSuite) SetupSuite() {
	ctx := context.Background()
	conf := config.Read()

	conn, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoURI(conf)))
	s.Require().NoError(err)
	s.conn = conn

	db := conn.Database(fmt.Sprintf("supercollector-%s", uuid.New().String()))
	s.db = db

	coll := db.Collection("users-test")
	s.coll = coll

	// fixtures
	_, err = s.coll.InsertOne(ctx, bson.M{"name": "John Snow", "email": "snow@warrior.com", "pass": "wall123"})
	s.Require().NoError(err)

	_, err = s.coll.InsertOne(ctx, bson.M{"name": "john_snow", "email": "john.snow@me.com"})
	s.Require().NoError(err)

	_, err = s.coll.InsertOne(ctx, bson.M{"name": "itssnowing", "email": "ryan.gos@li.ng", "pass": "snow"})
	s.Require().NoError(err)
}

func (s *RepoTestSuite) TearDownSuite() {
	ctx := context.Background()

	err := s.db.Drop(ctx)
	s.Require().NoError(err)
}

func (s *RepoTestSuite) TestGetManyNameFilter() {
	ctx := context.Background()
	repo := user.Repo{Coll: s.coll}

	name := "snow"
	flt := user.Filter{Name: &name}
	want := []user.User{
		{
			Name:        "John Snow",
			Credentials: "snow@warrior.com",
			Pass:        "wall123",
		},
		{
			Name:        "john_snow",
			Credentials: "john.snow@me.com",
		},
		{
			Name:        "itssnowing",
			Credentials: "ryan.gos@li.ng",
			Pass:        "snow",
		},
	}

	got, err := repo.GetMany(ctx, flt)
	s.Require().NoError(err)
	s.ElementsMatch(want, got)
}

func (s *RepoTestSuite) TestGetManyEmailFilter() {
	ctx := context.Background()
	repo := user.Repo{Coll: s.coll}

	email := "snow"
	flt := user.Filter{Email: &email}
	want := []user.User{
		{
			Name:        "John Snow",
			Credentials: "snow@warrior.com",
			Pass:        "wall123",
		},
		{
			Name:        "john_snow",
			Credentials: "john.snow@me.com",
		},
	}

	got, err := repo.GetMany(ctx, flt)
	s.Require().NoError(err)
	s.ElementsMatch(want, got)
}

func TestRepoSuite(t *testing.T) {
	suite.Run(t, new(RepoTestSuite))
}
