package grpcs

import (
	"context"
	"fmt"

	"github.com/sy-yoon/krealtors/gms"
	"github.com/sy-yoon/krealtors/gms-realtor/model"
	realtorpb "github.com/sy-yoon/krealtors/protos/v1/realtor"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/types/known/emptypb"
)

type RealtorServer struct {
	realtorpb.RealtorServiceServer
	db *mongo.Database
}

func genFilter(req *realtorpb.ListReItemsRequest) interface{} {
	filter := bson.D{}
	if req.CityId != "" {
		filter = append(filter, bson.E{"cityId", req.CityId})
	}

	if req.Type != "" {
		filter = append(filter, bson.E{"type", req.Type})
	}

	if req.TxType != "" {
		filter = append(filter, bson.E{"txType", req.TxType})
	}

	return filter
}

func (me *RealtorServer) ListReItems(ctx context.Context, req *realtorpb.ListReItemsRequest) (*realtorpb.ListReItemsResponse, error) {
	findOptions := options.Find()

	if req.Query != nil {
		findOptions.SetLimit(int64(req.Query.Limit))
	}

	//for k, v := range req.Filters {
	//    filter = append(filter, bson.E{k,v})
	//}

	//findOptions.SetProjection(bson.M{"_id": 1, "title": 1, "type": 1, "txType": 1, "cityId": 1, "thumbnail": 1, "price": 1})

	/*
		sort := bson.D{}
		for _, example := examples {
			sort = append(sort, bson.E{example, 1})
		}
		findOptions.SetSort();
	*/

	// Here's an array in which you can store the decoded documents
	var reItems []*realtorpb.ReItemHeader

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := me.db.Collection("reis").Find(ctx, genFilter(req), findOptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var item realtorpb.ReItemHeader
		err := cur.Decode(&item)
		if err != nil {
			gms.Logger.Error("DB", "SQL", err)
			return nil, err
		}
		reItems = append(reItems, &item)
	}

	if err := cur.Err(); err != nil {
		gms.Logger.Error("DB", "SQL", err)
		return nil, err
	}

	response := realtorpb.ListReItemsResponse{
		Items: reItems,
	}
	return &response, nil
}

func (me *RealtorServer) GetReItem(ctx context.Context, req *realtorpb.GetReItemRequest) (*realtorpb.ReItem, error) {
	var item realtorpb.ReItem
	objectId, _ := primitive.ObjectIDFromHex(req.Id)
	err := me.db.Collection("reis").FindOne(ctx, bson.M{"_id": objectId}).Decode(&item)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		gms.Logger.Error("DB", "SQL", err)
		return nil, err
	}
	return &item, err
}

func (me *RealtorServer) CreateReItem(ctx context.Context, item *realtorpb.ReItem) (*realtorpb.ReItem, error) {

	reItem := me.AllocNewObjectIdAndImagePath(item)

	_, err := me.db.Collection("reis").InsertOne(ctx, reItem)
	if err != nil {
		gms.Logger.Error("DB", "SQL", err)
		return nil, err
	}
	return item, err
}

func (me *RealtorServer) AllocNewObjectIdAndImagePath(item *realtorpb.ReItem) *model.ReItem {
	var reItem = model.ReItem{
		Id:            primitive.NewObjectID(),
		Title:         item.Title,
		UserId:        item.UserId,
		Type:          item.Type,
		TxType:        item.TxType,
		CityId:        item.CityId,
		Thumbnail:     item.Thumbnail,
		Images:        item.Images,
		Price:         item.Price,
		Facing:        item.Facing,
		Content:       item.Content,
		Address:       item.Address,
		Location:      model.LocationType{Latitude: float64(item.Location.Lat), Longitude: float64(item.Location.Lng)},
		Bedroom:       item.Bedroom,
		Bathroom:      item.Bathroom,
		Parking:       item.Parking,
		Area:          item.Area,
		Furnished:     item.Furnished,
		Garage:        item.Garage,
		AvailableDate: item.AvailableDate,
		Options:       model.OptionsType{Airconditioner: item.Options.Aircon, Refrigerator: item.Options.Refrigerator, WashingMachine: item.Options.Washingmachine, Dishwasher: item.Options.Dishwasher, Microwave: item.Options.Microwave, TV: item.Options.Tv, Doorlock: item.Options.Doorlock},
		CreatedDate:   uint64(item.CreatedDate),
		UpdatedDate:   uint64(item.CreatedDate),
	}

	if reItem.Images == nil || len(reItem.Images) == 0 {
		return &reItem
	}

	id := reItem.Id.Hex()
	item.Id = id
	reItem.Thumbnail = fmt.Sprintf("/realty/%s/%s/0", reItem.CityId, id)
	item.Thumbnail = reItem.Thumbnail
	for i, v := range reItem.Images {
		reItem.Images[i] = fmt.Sprintf("/realty/%s/%s/%s", reItem.CityId, id, v)
		item.Images[i] = reItem.Images[i]
	}

	return &reItem
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

// func generatePaginationQuery(query, sort, nextKey) {
// 	const sortField = sort == null ? null : sort[0];

// 	if (nextKey == null) {
// 	  return { query, nextKeyFn };
// 	}

// 	let paginatedQuery = query;

// 	if (sort == null) {
// 	  paginatedQuery._id = { $gt: nextKey._id };
// 	  return { paginatedQuery, nextKey };
// 	}

// 	const sortOperator = sort[1] === 1 ? "$gt" : "$lt";

// 	const paginationQuery = [
// 	  { [sortField]: { [sortOperator]: nextKey[sortField] } },
// 	  {
// 		$and: [
// 		  { [sortField]: nextKey[sortField] },
// 		  { _id: { [sortOperator]: nextKey._id } }
// 		]
// 	  }
// 	];

// 	if (paginatedQuery.$or == null) {
// 	  paginatedQuery.$or = paginationQuery;
// 	} else {
// 	  paginatedQuery = { $and: [query, { $or: paginationQuery }] };
// 	}

// 	return { paginatedQuery, nextKeyFn };
//   }

//   func nextKeyFn(items) {
// 	if (items.length === 0) {
// 	  return null;
// 	}

// 	const item = items[items.length - 1];

// 	if (sortField == null) {
// 	  return { _id: item._id };
// 	}

// 	return { _id: item._id, [sortField]: item[sortField] };
//   }
