package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"sportshot/pkg/utils/global"
	"sportshot/pkg/utils/proto"
)

type EventServer struct {
	proto.UnimplementedEventServiceServer
}

func (s *EventServer) SearchEvents(ctx context.Context, req *proto.SearchEventsRequest) (*proto.EventsReply, error) {
	// initial filter
	filter := bson.D{}

	// condition filter
	if req.LeagueName != "" {
		filter = append(filter, bson.E{Key: "leagueName", Value: req.LeagueName})
	}
	if req.HomeName != "" {
		filter = append(filter, bson.E{Key: "homeName", Value: req.HomeName})
	}
	if req.AwayName != "" {
		filter = append(filter, bson.E{Key: "awayName", Value: req.AwayName})
	}
	if req.SportType != "" {
		filter = append(filter, bson.E{Key: "sportType", Value: req.SportType})
	}

	// default date to avoid select too much
	// but for the convenience of testing, replaced this default temporary
	//if req.StartDate == "" {
	//	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	//	filter = append(filter, bson.E{Key: "date", Value: bson.M{"$gte": yesterday}})
	//} else {
	//	filter = append(filter, bson.E{Key: "date", Value: bson.M{"$gte": req.StartDate}})
	//}
	//if req.EndDate != "" {
	//	filter = append(filter, bson.E{Key: "date", Value: bson.M{"$lte": req.EndDate}})
	//}

	if req.StartDate != "" {
		filter = append(filter, bson.E{Key: "date", Value: bson.M{"$gte": req.StartDate}})
	}
	if req.EndDate != "" {
		filter = append(filter, bson.E{Key: "date", Value: bson.M{"$lte": req.EndDate}})
	}

	zap.S().Infof("search filter: %v", filter)

	// connect to mongodb
	databaseName := "sportevents"
	collectionName := "basketball"
	collection := global.MongodbClient.Database(databaseName).Collection(collectionName)
	zap.S().Infof("connected to mongodb, start to search events")

	// start to search
	zap.S().Infof("start to search by filter : %v", filter)
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	zap.S().Infof("search successfully")
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			zap.S().Warnf("failed to close cursor, %v", err)
		}
	}(cursor, context.TODO())

	// processing data
	var results []*proto.EventInfo
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	// reorg reply
	reply := &proto.EventsReply{
		Events: results,
		Count:  int32(len(results)),
	}

	return reply, nil
}
