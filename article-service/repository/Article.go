package repository

import (
	"article-service/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func collArticles() *mongo.Collection {
	return client.Database("article-alpha").Collection("article")
}

func GetArticles(ctx context.Context, filter bson.D) ([]*model.Article, error) {
	var articles []*model.Article

	if filter == nil {
		filter = bson.D{}
	}

	cursor, err := collArticles().Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &articles); err != nil {
		return nil, err
	}
	return articles, nil
}

func AddArticle(ctx context.Context, article *model.Article) error {
	article.ID = primitive.NewObjectID().Hex()
	_, err := collArticles().InsertOne(ctx, article)
	if err != nil {
		return err
	}
	return nil
}

func DeleteArticle(ctx context.Context, articleId string) error {
	var err error

	objectId, err := primitive.ObjectIDFromHex(articleId)
	if err != nil {
		return err
	}

	_, err = collArticles().DeleteOne(ctx, bson.D{{"_id", objectId.Hex()}})
	if err != nil {
		return err
	}
	return nil
}
