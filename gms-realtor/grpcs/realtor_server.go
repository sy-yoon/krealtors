package grpcs

import (
	"context"

	realtorpb "github.com/sy-yoon/krealtors/protos/v1/realtor"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/emptypb"
)

type RealtorServer struct {
	realtorpb.RealtorServiceServer
	db *mongo.Database
}

func (me *RealtorServer) ListReItems(ctx context.Context, req *realtorpb.ListReItemsRequest) (*realtorpb.ListReItemsResponse, error) {
	return nil, nil
}

func (me *RealtorServer) GetReItem(ctx context.Context, req *realtorpb.GetReItemRequest) (*realtorpb.ReItem, error) {
	var item realtorpb.ReItem
	err := me.db.Collection("properties").FindOne(ctx, bson.M{"_id": req.Id}).Decode(&item)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &item, err
}

func (me *RealtorServer) CreateReItem(ctx context.Context, item *realtorpb.ReItem) (*realtorpb.ReItem, error) {
	_,err :=me.db.Collection("homes").InsertOne(ctx, item)
	if err != nil {
		return nil, err
	}
	return item, err
}

func (me *RealtorServer) UpdateReItem(ctx context.Context, req *realtorpb.ReItem) (*realtorpb.ReItem, error) {
	return nil, nil
}

func (me *RealtorServer) DeleteReItem(ctx context.Context, req *realtorpb.DeleteReItemRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (me *RealtorServer) AddDBContext(db interface{}) {
	me.db = db.(*mongo.Database)
}
