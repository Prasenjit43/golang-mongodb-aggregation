package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Prasenjit43/golang-mongodb-aggregation/databases"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection
var booksCollection *mongo.Collection

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client := databases.Client
	defer client.Disconnect(ctx)

	userCollection = databases.OpenConnection(client, "aggee", "users")
	booksCollection = databases.OpenConnection(client, "aggee", "books")

	/*Example 01*/
	/*How many users are active*/
	filter := bson.D{{"$match", bson.D{{"isActive", true}}}}
	group := bson.D{
		{"$count", "active_users"},
	}

	pipelineParam := mongo.Pipeline{filter, group}
	displayOutput(userCollection, ctx, &pipelineParam)

	/*Example 02*/
	/*What is the average age of all users*/

	group = bson.D{{
		"$group", bson.D{
			{"_id", nil},
			{"avg_age", bson.D{{"$avg", "$age"}}},
		}}}

	pipelineParam = mongo.Pipeline{group}
	displayOutput(userCollection, ctx, &pipelineParam)

	/*Example 03*/
	/*List out the 5 top most common favorite fruits amoung users*/
	group = bson.D{{
		"$group", bson.D{
			{"_id", "$favoriteFruit"},
			{"count", bson.D{{"$sum", 1}}},
		}}}

	sort := bson.D{{
		"$sort", bson.D{
			{"count", -1},
		}},
	}

	limit := bson.D{{"$limit", 5}}

	pipelineParam = mongo.Pipeline{group, sort, limit}
	displayOutput(userCollection, ctx, &pipelineParam)

	/*Example 04*/
	/*Find the total numbers of males and females*/

	group = bson.D{{
		"$group", bson.D{
			{"_id", "$gender"},
			{"gender_count", bson.D{{"$sum", 1}}},
		}},
	}
	pipelineParam = mongo.Pipeline{group}
	displayOutput(userCollection, ctx, &pipelineParam)

	/*Example 05*/
	/*Which country has the highest no of registered user*/

	group = bson.D{{
		"$group", bson.D{
			{"_id", "$company.location.country"},
			{"user_count", bson.D{{"$sum", 1}}},
		},
	}}

	sort = bson.D{{
		"$sort", bson.D{
			{"user_count", -1},
		},
	}}

	limit = bson.D{{
		"$limit", 1,
	}}

	pipelineParam = mongo.Pipeline{group, sort, limit}
	displayOutput(userCollection, ctx, &pipelineParam)

	/*Example 06*/
	/*List all unique colours present in the collection*/

	group = bson.D{{
		"$group", bson.D{
			{"_id", "$eyeColor"},
		},
	}}

	pipelineParam = mongo.Pipeline{group}
	displayOutput(userCollection, ctx, &pipelineParam)

	/*Example 07*/
	/*What is the average numbers of tags per user*/

	unwind := bson.D{{
		"$unwind", "$tags"},
	}

	group = bson.D{{
		"$group", bson.D{
			{"_id", "$_id"},
			{"total_tags", bson.D{{"$sum", 1}}},
		},
	}}

	groupInOneDoc := bson.D{{
		"$group", bson.D{
			{"_id", nil},
			{"avg", bson.D{{"$avg", "$total_tags"}}},
		},
	}}

	pipelineParam = mongo.Pipeline{unwind, group, groupInOneDoc}
	displayOutput(userCollection, ctx, &pipelineParam)

	/*Another Approach*/

	addExtraField := bson.D{{
		"$addFields", bson.D{
			{"noOfTags", bson.D{
				{"$size", bson.D{
					{"$ifNull", bson.A{"$tags", bson.A{}}},
				}},
			}},
		},
	}}

	groupInOneDoc = bson.D{{
		"$group", bson.D{
			{"_id", nil},
			{"avg", bson.D{
				{"$avg", "$noOfTags"},
			}},
		}},
	}

	pipelineParam = mongo.Pipeline{addExtraField, groupInOneDoc}
	displayOutput(userCollection, ctx, &pipelineParam)

	/*Example 09*/
	/*How many users have enum as one of their tags*/
	filter = bson.D{{
		"$match", bson.D{
			{"tags", bson.D{
				{"$in", bson.A{"enim", "$tags"}},
			}},
		},
	}}

	group = bson.D{{
		"$count", "user_with_enum",
	}}

	pipelineParam = mongo.Pipeline{filter, group}
	displayOutput(userCollection, ctx, &pipelineParam)

	/*Example 10*/
	/*What are the names and age of users who are inactive and have 'velit' as a tag*/

	filter = bson.D{{
		"$match", bson.D{
			{"tags", "velit"},
			{"isActive", false},
		},
	}}

	projection := bson.D{{
		"$project", bson.D{
			{"name", 1},
			{"age", 1},
		},
	}}

	pipelineParam = mongo.Pipeline{filter, projection}
	displayOutput(userCollection, ctx, &pipelineParam)

	/*Example 11*/
	/*Who many users have phone no starting with '+1(940)'*/

	filter = bson.D{{
		"$match", bson.D{
			{"company.phone", bson.D{
				{"$regex", "^\\+1 \\(940\\)"},
			}},
		},
	}}

	group = bson.D{{
		"$count", "user_count",
	}}

	pipelineParam = mongo.Pipeline{filter, group}
	displayOutput(userCollection, ctx, &pipelineParam)

	/*Example 12*/
	/*Who has registered the most recently'*/

	sort = bson.D{{
		"$sort", bson.D{
			{"registered", -1},
		},
	}}

	limit = bson.D{{
		"$limit", 1,
	}}

	projection = bson.D{{
		"$project", bson.D{
			{"name", 1},
			{"gender", 1},
			{"favoriteFruit", 1},
		},
	}}

	pipelineParam = mongo.Pipeline{sort, limit, projection}
	displayOutput(userCollection, ctx, &pipelineParam)

	/*Example 13*/
	/*Categorize user by their favorite fruits*/

	group = bson.D{{
		"$group", bson.D{
			{"_id", "$favoriteFruit"},
			{"users", bson.D{
				{"$push", "$name"},
			}},
		},
	}}

	pipelineParam = mongo.Pipeline{group}
	displayOutput(userCollection, ctx, &pipelineParam)

	/*Example 14*/
	/*How many users have 'ad' as the second tag in their list of tags*/

	filter = bson.D{{
		"$match", bson.D{
			{"tags.1", "ad"},
		},
	}}

	group = bson.D{{
		"$count", "user_with_'ad_tags'",
	}}

	pipelineParam = mongo.Pipeline{filter, group}
	displayOutput(userCollection, ctx, &pipelineParam)

	/*Example 15*/
	/*Find user who have both 'enum' and 'id' as their tags*/

	filter = bson.D{{
		"$match", bson.D{
			{"tags", bson.D{
				{"$all", bson.A{"enim", "id"}},
			}}},
	},
	}

	pipelineParam = mongo.Pipeline{filter}
	displayOutput(userCollection, ctx, &pipelineParam)

	/*Example 16*/
	/*List all companies located in USA with their correcponding user count*/

	filter = bson.D{{
		"$match", bson.D{
			{"company.location.country", "USA"},
		},
	}}

	group = bson.D{{
		"$group", bson.D{
			{"_id", "$company.title"},
			{"count", bson.D{
				{"$sum", 1}},
			},
		},
	}}

	pipelineParam = mongo.Pipeline{filter, group}
	displayOutput(userCollection, ctx, &pipelineParam)

	/*Example 17*/

	booksFilter := bson.D{{
		"$match", bson.D{
			{"title", "The Great Gatsby"},
		},
	}}

	lookup := bson.D{{
		"$lookup", bson.D{
			{"from", "authors"},
			{"localField", "author_id"},
			{"foreignField", "_id"},
			{"as", "author_details"},
		},
	}}

	unwind = bson.D{
		{"$unwind", "$author_details"},
	}

	pipelineParam = mongo.Pipeline{booksFilter, lookup, unwind}
	displayOutput(booksCollection, ctx, &pipelineParam)

}

func displayOutput(collectionName *mongo.Collection, ctx context.Context, pipelineInput *mongo.Pipeline) {
	fmt.Println("************************")
	loadedCursor, err := collectionName.Aggregate(ctx, *pipelineInput)
	if err != nil {
		fmt.Println("Error :", err.Error())
		return
	}

	var loaded []bson.M

	if err = loadedCursor.All(ctx, &loaded); err != nil {

		fmt.Println("Error 111:", err.Error())

		panic(err)
	}

	for _, result := range loaded {
		fmt.Println(result)
	}
}
