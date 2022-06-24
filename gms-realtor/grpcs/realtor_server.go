package grpcs

import (
	"context"
	"fmt"
	"strings"

	realtorpb "github.com/sy-yoon/krealtors/protos/v1/realtor"
	"github.com/sy-yoon/krealtors/utils"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type RealtorServer struct {
	realtorpb.RealtorServiceServer
	orm *gorm.DB
}

func genFilter(req *realtorpb.ListReItemHeaderRequest) map[string]interface{} {
	conditions := make(map[string]interface{})
	if req.CityId != "" {
		conditions["city_id"] = req.CityId
	}

	if req.TxType != 0 {
		conditions["tx_type"] = req.TxType
	}

	if req.ReType != 0 {
		conditions["re_type"] = req.ReType
	}

	return conditions
}

func (me *RealtorServer) ListReItemHeader(ctx context.Context, req *realtorpb.ListReItemHeaderRequest) (*realtorpb.ListReItemHeaderResponse, error) {
	var reItems []*realtorpb.ReItemHeader
	var sb strings.Builder
	params := genFilter(req)
	for key, _ := range params {
		sb.WriteString(" and ")
		sb.WriteString(fmt.Sprintf("%s = @%s", key, key))
	}

	conditions := sb.String()
	if len(conditions) > 0 {
		conditions = conditions[4 : len(conditions)]
	}

	if err := utils.CheckError(me.orm.Table("reis").Where(conditions, params).Find(&reItems)); err != nil {
		return nil, err
	}

	response := realtorpb.ListReItemHeaderResponse{
		Items: reItems,
	}
	return &response, nil
}

func (me *RealtorServer) GetReItem(ctx context.Context, req *realtorpb.GetReItemRequest) (*realtorpb.ReItem, error) {
	var item realtorpb.ReItem
	item.Id = req.Id
	if err := utils.CheckError(me.orm.Table("reis").First(&item)); err != nil {
		return nil, err
	}
	return &item, nil
}

func (me *RealtorServer) CreateReItem(ctx context.Context, item *realtorpb.ReItem) (*realtorpb.ReItem, error) {

	if err := utils.CheckError(me.orm.Table("reis").Omit("created_date", "updated_date").Create(item)); err != nil {
		return nil, err
	}

	return item, nil
}

func (me *RealtorServer) UpdateReItem(ctx context.Context, req *realtorpb.ReItem) (*realtorpb.ReItem, error) {
	if err := utils.CheckError(me.orm.Save(req)); err != nil {
		return nil, err
	}

	return req, nil
}

func (me *RealtorServer) DeleteReItem(ctx context.Context, req *realtorpb.DeleteReItemRequest) (*emptypb.Empty, error) {
	reItem := realtorpb.ReItem{
		Id: req.Id,
	}
	if err := utils.CheckError(me.orm.Table("reis").Delete(&reItem)); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (me *RealtorServer) AddDBContext(db interface{}) {
	me.orm = db.(*gorm.DB)
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

// func (x GeoLocation) GormDataType() string {
// 	return "point"
//   }

// func (x GeoLocation) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
// 	return clause.Expr{
// 		SQL:  "Point(?,?)",
// 		Vars: []interface{}{x.Lat, x.Lng},
// 	}
// }

// func (x *GeoLocation) Scan(v interface{}) error {
// 	var data []byte
//     switch v := v.(type) {
//     case []byte:
//         data = v
//     case string:
//         data = []byte(v)
//     case nil:
//         return nil
//     default:
//         return errors.New("(*Point).Scan: unsupported data type")
//     }

//     if len(data) == 0 {
//         return nil
//     }

// 	var err error
//     data = data[1 : len(data)-1] // drop the surrounding parentheses
//     for i := 0; i < len(data); i++ {
//         if data[i] == ',' {
//             if x.Lat, err = strconv.ParseFloat(string(data[:i]), 64); err != nil {
//                 return err
//             }

//             if x.Lng, err = strconv.ParseFloat(string(data[i+1:]), 64); err != nil {
//                 return err
//             }
//             break
//         }
//     }
// 	return nil
// }
