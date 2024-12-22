package user

import (
	"context"

	"github.com/cockroachdb/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const searchLimit = 1000

type Repo struct {
	Coll *mongo.Collection
}

type Filter struct {
	Email *string
	Name  *string
	Pass  *string
}

func (r *Repo) GetMany(ctx context.Context, filter Filter) ([]User, error) {
	bsonFlt, lmt := filter.toBSONFilterAndLimit()

	cur, err := r.Coll.Find(ctx, bsonFlt, lmt)
	if err != nil {
		return nil, errors.Wrap(err, "find users in mongo")
	}

	var userDocs []userDoc

	err = cur.All(ctx, &userDocs)
	if err != nil {
		return nil, errors.Wrap(err, "fetch users from mongo")
	}

	return docsToUsers(userDocs), nil
}

func (r *Repo) WriteMany(ctx context.Context, users []User) error {
	docs := usersToDocs(users)

	_, err := r.Coll.InsertMany(ctx, docs)
	if err != nil {
		return errors.Wrap(err, "write users batch to mongo")
	}

	return nil
}

func (f Filter) toBSONFilterAndLimit() (bson.M, *options.FindOptions) {
	filter := bson.M{}
	notEmpty := false

	if f.Email != nil {
		filter["email"] = bson.M{"$regex": *f.Email, "$options": "i"}
		notEmpty = true
	}

	if f.Name != nil {
		filter["name"] = bson.M{"$regex": *f.Name, "$options": "i"}
		notEmpty = true
	}

	if f.Pass != nil {
		filter["pass"] = bson.M{"$regex": *f.Pass, "$options": "i"}
		notEmpty = true
	}

	if !notEmpty {
		return filter, options.Find().SetLimit(searchLimit)
	}

	return filter, nil
}
