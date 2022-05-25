package grpcs

import (
	"context"

	regionpb "github.com/sy-yoon/krealtors/protos/v1/region"
	"github.com/sy-yoon/krealtors/utils"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type RegionServer struct {
	regionpb.RegionServiceServer
	orm *gorm.DB
}

func (me *RegionServer) AddDBContext(orm interface{}) {
	me.orm = orm.(*gorm.DB)
}

func (me *RegionServer) GetCountry(ctx context.Context, req *regionpb.GetCountryRequest) (*regionpb.Country, error) {
	country := regionpb.Country{}
	country.Id = req.Id

	if err := utils.CheckError(me.orm.First(&country)); err != nil {
		return nil, err
	}

	return &country, nil
}

func (me *RegionServer) ListCountries(ctx context.Context, req *regionpb.ListCountriesRequest) (*regionpb.ListCountriesResponse, error) {
	countries := []*regionpb.Country{}
	if err := utils.CheckError(me.orm.Find(countries)); err != nil {
		return nil, err
	}

	response := regionpb.ListCountriesResponse{
		Countries: countries,
	}
	return &response, nil
}

func (me *RegionServer) CreateCountry(ctx context.Context, country *regionpb.Country) (*regionpb.Country, error) {
	if err := utils.CheckError(me.orm.Create(country)); err != nil {
		return nil, err
	}

	return country, nil
}

func (me *RegionServer) UpdateCountry(ctx context.Context, country *regionpb.Country) (*regionpb.Country, error) {
	if err := utils.CheckError(me.orm.Save(country)); err != nil {
		return nil, err
	}

	return country, nil
}

func (me *RegionServer) DeleteCountry(ctx context.Context, req *regionpb.DeleteCountryRequest) (*emptypb.Empty, error) {
	var country regionpb.Country
	if err := utils.CheckError(me.orm.Where("cntry_id", req.Id).Delete(&country)); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (me *RegionServer) GetProvince(ctx context.Context, req *regionpb.GetProvinceRequest) (*regionpb.Province, error) {
	province := regionpb.Province{}
	province.Id = req.Id
	if err := utils.CheckError(me.orm.First(&province)); err != nil {
		return nil, err
	}

	return &province, nil
}

func (me *RegionServer) ListProvincies(ctx context.Context, req *regionpb.ListProvinciesRequest) (*regionpb.ListProvinciesResponse, error) {
	provincies := []*regionpb.Province{}
	if err := utils.CheckError(me.orm.Find(provincies)); err != nil {
		return nil, err
	}

	response := regionpb.ListProvinciesResponse{
		Provincies: provincies,
	}
	return &response, nil
}

func (me *RegionServer) CreateProvince(ctx context.Context, province *regionpb.Province) (*regionpb.Province, error) {
	if err := utils.CheckError(me.orm.Create(province)); err != nil {
		return nil, err
	}

	return province, nil
}

func (me *RegionServer) UpdateProvince(ctx context.Context, province *regionpb.Province) (*regionpb.Province, error) {
	if err := utils.CheckError(me.orm.Save(province)); err != nil {
		return nil, err
	}

	return province, nil
}

func (me *RegionServer) DeleteProvince(ctx context.Context, req *regionpb.DeleteProvinceRequest) (*emptypb.Empty, error) {
	var province regionpb.Province
	if err := utils.CheckError(me.orm.Where("prvnc_id", req.Id).Delete(&province)); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type Result struct {
	GeoLocation string
}

func (me *RegionServer) GetCity(ctx context.Context, req *regionpb.GetCityRequest) (*regionpb.City, error) {
	city := regionpb.City{}
	city.Id = req.Id

	if err := utils.CheckError(me.orm.First(&city)); err != nil {
		return nil, err
	}

	return &city, nil
}

func (me *RegionServer) ListCities(ctx context.Context, req *regionpb.ListCitiesRequest) (*regionpb.ListCitiesResponse, error) {
	cities := []*regionpb.City{}
	if err := utils.CheckError(me.orm.Find(cities)); err != nil {
		return nil, err
	}

	response := regionpb.ListCitiesResponse{
		Cities: cities,
	}
	return &response, nil
}

func (me *RegionServer) CreateCity(ctx context.Context, city *regionpb.City) (*regionpb.City, error) {
	if err := utils.CheckError(me.orm.Create(city)); err != nil {
		return nil, err
	}

	return city, nil
}

func (me *RegionServer) UpdateCity(ctx context.Context, city *regionpb.City) (*regionpb.City, error) {
	if err := utils.CheckError(me.orm.Save(city)); err != nil {
		return nil, err
	}

	return city, nil
}

func (me *RegionServer) DeleteCity(ctx context.Context, req *regionpb.DeleteCityRequest) (*emptypb.Empty, error) {
	var city regionpb.City
	if err := utils.CheckError(me.orm.Where("city_id", req.Id).Delete(&city)); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
