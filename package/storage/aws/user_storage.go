package aws

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/monk78anthony/apiv6/domain"
	"github.com/monk78anthony/apiv6/package/storage"
)

var _ storage.UserStorer = UserStorage{}

type UserStorage struct {
	timeout time.Duration
	client  *dynamodb.DynamoDB
}

func NewUserStorage(session *session.Session, timeout time.Duration) UserStorage {
	return UserStorage{
		timeout: timeout,
		client:  dynamodb.New(session),
	}
}

func (u UserStorage) Insert(ctx context.Context, user storage.User) error {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	item, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		log.Println(err)
		return domain.ErrInternal
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String("users"),
		Item:      item,
		ExpressionAttributeNames: map[string]*string{
			"#uuid": aws.String("uuid"),
		},
		ConditionExpression: aws.String("attribute_not_exists(#uuid)"),
	}

	if _, err := u.client.PutItemWithContext(ctx, input); err != nil {
		log.Println(err)

		if _, ok := err.(*dynamodb.ConditionalCheckFailedException); ok {
			return domain.ErrConflict
		}

		return domain.ErrInternal
	}

	return nil
}

func (u UserStorage) Find(ctx context.Context, uuid string) (storage.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	input := &dynamodb.GetItemInput{
		TableName: aws.String("users"),
		Key: map[string]*dynamodb.AttributeValue{
			"uuid": {S: aws.String(uuid)},
		},
	}

	res, err := u.client.GetItemWithContext(ctx, input)
	if err != nil {
		log.Println(err)

		return storage.User{}, domain.ErrInternal
	}

	if res.Item == nil {
		return storage.User{}, domain.ErrNotFound
	}

	var user storage.User
	if err := dynamodbattribute.UnmarshalMap(res.Item, &user); err != nil {
		log.Println(err)

		return storage.User{}, domain.ErrInternal
	}

	return user, nil
}

func (u UserStorage) Delete(ctx context.Context, uuid string) error {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	input := &dynamodb.DeleteItemInput{
		TableName: aws.String("users"),
		Key: map[string]*dynamodb.AttributeValue{
			"uuid": {S: aws.String(uuid)},
		},
	}

	if _, err := u.client.DeleteItemWithContext(ctx, input); err != nil {
		log.Println(err)

		return domain.ErrInternal
	}

	return nil
}

func (u UserStorage) Update(ctx context.Context, user storage.User) error {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	roles := make([]*dynamodb.AttributeValue, len(user.Roles))
	for i, role := range user.Roles {
		roles[i] = &dynamodb.AttributeValue{S: aws.String(role)}
	}

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String("users"),
		Key: map[string]*dynamodb.AttributeValue{
			"uuid": {S: aws.String(user.UUID)},
		},
		ExpressionAttributeNames: map[string]*string{
			"#name":  aws.String("name"),
			"#level": aws.String("level"),
			"#roles": aws.String("roles"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":name":  {S: aws.String(user.Name)},
			":level": {N: aws.String(fmt.Sprint(user.Grade))},
			":roles": {L: roles},
		},
		UpdateExpression: aws.String("set #name = :name, #grade = :grade, #roles = :roles"),
		ReturnValues:     aws.String("UPDATED_NEW"),
	}

	if _, err := u.client.UpdateItemWithContext(ctx, input); err != nil {
		log.Println(err)

		return domain.ErrInternal
	}

	return nil
}
