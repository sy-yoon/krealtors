package grpcs

import (
	"context"

	"github.com/sy-yoon/krealtors/gms"
	regionpb "github.com/sy-yoon/krealtors/protos/v1/region"
	"google.golang.org/protobuf/types/known/emptypb"
	"xorm.io/xorm"
)

type RegionServer struct {
	regionpb.RegionServiceServer
	orm *xorm.Engine
}

func (me *RegionServer) AddDBContext(orm interface{}) {
	me.orm = orm.(*xorm.Engine)
}

func (me *RegionServer) GetCountry(ctx context.Context, req *regionpb.GetCountryRequest) (*regionpb.Country, error) {
	country := regionpb.Country{}
	country.CountyId = req.CountryId
	_, err := me.orm.Get(&country)
	if err != nil {
		gms.Logger.Error("DB", "SQL", err)
		return nil, err
	}
	return &country, nil
}

func (me *RegionServer) ListCountries(ctx context.Context, req *regionpb.ListCountriesRequest) (*regionpb.ListCountriesResponse, error) {
	countries := []*regionpb.Country{}
	err := me.orm.Find(&countries)
	if err != nil {
		gms.Logger.Error("DB", "SQL", err)
		return nil, err
	}

	response := regionpb.ListCountriesResponse{
		Countries: countries,
	}
	return &response, nil
}

func (me *RegionServer) CreateCountry(ctx context.Context, country *regionpb.Country) (*regionpb.Country, error) {
	_, err := me.orm.InsertOne(country)
	if err != nil {
		gms.Logger.Error("DB", "SQL", err)
		return nil, err
	}
	return country, nil
}

func (me *RegionServer) UpdateCountry(ctx context.Context, country *regionpb.Country) (*regionpb.Country, error) {
	_, err := me.orm.Update(country)
	if err != nil {
		gms.Logger.Error("DB", "SQL", err)
		return nil, err
	}
	return country, nil
}

func (me *RegionServer) DeleteCountry(ctx context.Context, req *regionpb.DeleteCountryRequest) (*emptypb.Empty, error) {
	var country regionpb.Country
	_, err := me.orm.Where("cntry_id", req.CountryId).Delete(&country)
	if err != nil {
		gms.Logger.Error("DB", "SQL", err)
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (me *RegionServer) GetProvince(ctx context.Context, req *regionpb.GetProvinceRequest) (*regionpb.Province, error) {
	province := regionpb.Province{}
	province.ProvinceId = req.ProvinceId
	_, err := me.orm.Get(&province)
	if err != nil {
		gms.Logger.Error("DB", "SQL", err)
		return nil, err
	}
	return &province, nil
}

func (me *RegionServer) ListProvincies(ctx context.Context, req *regionpb.ListProvinciesRequest) (*regionpb.ListProvinciesResponse, error) {
	provincies := []*regionpb.Province{}
	err := me.orm.Find(&provincies)
	if err != nil {
		gms.Logger.Error("DB", "SQL", err)
		return nil, err
	}

	response := regionpb.ListProvinciesResponse{
		Provincies: provincies,
	}
	return &response, nil
}

func (me *RegionServer) CreateProvince(ctx context.Context, province *regionpb.Province) (*regionpb.Province, error) {
	_, err := me.orm.InsertOne(province)
	if err != nil {
		gms.Logger.Error("DB", "SQL", err)
		return nil, err
	}
	return province, nil
}

func (me *RegionServer) UpdateProvince(ctx context.Context, province *regionpb.Province) (*regionpb.Province, error) {
	_, err := me.orm.Update(province)
	if err != nil {
		gms.Logger.Error("DB", "SQL", err)
		return nil, err
	}
	return province, nil
}

func (me *RegionServer) DeleteProvince(ctx context.Context, req *regionpb.DeleteProvinceRequest) (*emptypb.Empty, error) {
	var province regionpb.Province
	_, err := me.orm.Where("prvnc_id", req.ProvinceId).Delete(&province)
	if err != nil {
		gms.Logger.Error("DB", "SQL", err)
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (me *RegionServer) GetCity(ctx context.Context, req *regionpb.GetCityRequest) (*regionpb.City, error) {
	city := regionpb.City{}
	city.CityId = req.CityId
	_, err := me.orm.Get(&city)
	if err != nil {
		gms.Logger.Error("DB", "SQL", err)
		return nil, err
	}
	return &city, nil
}

func (me *RegionServer) ListCities(ctx context.Context, req *regionpb.ListCitiesRequest) (*regionpb.ListCitiesResponse, error) {
	cities := []*regionpb.City{}
	err := me.orm.Find(&cities)
	if err != nil {
		gms.Logger.Error("DB", "SQL", err)
		return nil, err
	}

	response := regionpb.ListCitiesResponse{
		Cities: cities,
	}
	return &response, nil
}

func (me *RegionServer) CreateCity(ctx context.Context, city *regionpb.City) (*regionpb.City, error) {
	_, err := me.orm.InsertOne(city)
	if err != nil {
		gms.Logger.Error("DB", "SQL", err)
		return nil, err
	}
	return city, nil
}

func (me *RegionServer) UpdateCity(ctx context.Context, city *regionpb.City) (*regionpb.City, error) {
	_, err := me.orm.Update(city)
	if err != nil {
		gms.Logger.Error("DB", "SQL", err)
		return nil, err
	}
	return city, nil
}

func (me *RegionServer) DeleteCity(ctx context.Context, req *regionpb.DeleteCityRequest) (*emptypb.Empty, error) {
	var city regionpb.City
	_, err := me.orm.Where("city_id", req.CityId).Delete(&city)
	if err != nil {
		gms.Logger.Error("DB", "SQL", err)
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
