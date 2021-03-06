package grpcs

import (
	"context"

	"github.com/sy-yoon/krealtors/protos/v1/common"
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

	if err := utils.CheckError(me.orm.Table("cntry").First(&country)); err != nil {
		return nil, err
	}

	return &country, nil
}

func (me *RegionServer) ListCountries(ctx context.Context, req *regionpb.ListCountriesRequest) (*regionpb.ListCountriesResponse, error) {
	countries := []*regionpb.Country{}
	// if err := utils.CheckError(me.orm.Table("cntry").Find(&countries)); err != nil {
	// 	return nil, err
	// }

	sql := `
	WITH RECURSIVE region AS(
		SELECT
			id,
			name,
			open,
			currency,
			location,
			1 as type
		FROM
			cntry		
		UNION all
		select 
		   p.id,
		   p.name,
		   p.open,
		   r.currency,
		   p.location,
		   p.type
		  from
		(select 
			id,
			name,
			open,
			location,
			country_id as pid,
			2 as type
		 from 
			 prvnc
		 union all
		 select 
			id,
			name,
			open,
			location,
			province_id as pid,
			3 as type
		 from city) p,
			region r 
		 where r.id = p.pid 
	) SELECT * 
	FROM region
	order by id
	`

	// Raw SQL
	rows, err := me.orm.Raw(sql).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var (
		id       string
		name     string
		open     bool
		currency string
		location *common.GeoLocation
		_type    int32
	)

	var (
		country  *regionpb.Country
		province *regionpb.Province
		city     *regionpb.City
	)

	for rows.Next() {
		rows.Scan(&id, &name, &open, &currency, &location, &_type)

		if _type == 1 {
			country = &regionpb.Country{
				Id:       id,
				Name:     name,
				Currency: currency,
				Location: location,
			}
			countries = append(countries, country)
		} else if _type == 2 {
			province = &regionpb.Province{
				Id:       id,
				Name:     name,
				Location: location,
			}
			country.Provincies = append(country.Provincies, province)
		} else {
			city = &regionpb.City{
				Id:       id,
				Name:     name,
				Location: location,
			}
			province.Cities = append(province.Cities, city)
		}
	}

	response := regionpb.ListCountriesResponse{
		Countries: countries,
	}
	return &response, nil
}

func (me *RegionServer) CreateCountry(ctx context.Context, country *regionpb.Country) (*regionpb.Country, error) {
	if err := utils.CheckError(me.orm.Table("cntry").Create(country)); err != nil {
		return nil, err
	}

	return country, nil
}

func (me *RegionServer) UpdateCountry(ctx context.Context, country *regionpb.Country) (*regionpb.Country, error) {
	if err := utils.CheckError(me.orm.Table("cntry").Save(country)); err != nil {
		return nil, err
	}

	return country, nil
}

func (me *RegionServer) DeleteCountry(ctx context.Context, req *regionpb.DeleteCountryRequest) (*emptypb.Empty, error) {
	var country regionpb.Country
	if err := utils.CheckError(me.orm.Table("cntry").Where("id", req.Id).Delete(&country)); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (me *RegionServer) GetProvince(ctx context.Context, req *regionpb.GetProvinceRequest) (*regionpb.Province, error) {
	province := regionpb.Province{}
	province.Id = req.Id
	if err := utils.CheckError(me.orm.Table("prvnc").First(&province)); err != nil {
		return nil, err
	}

	return &province, nil
}

func (me *RegionServer) ListProvincies(ctx context.Context, req *regionpb.ListProvinciesRequest) (*regionpb.ListProvinciesResponse, error) {
	provincies := []*regionpb.Province{}
	if err := utils.CheckError(me.orm.Table("prvnc").Find(&provincies)); err != nil {
		return nil, err
	}

	response := regionpb.ListProvinciesResponse{
		Provincies: provincies,
	}
	return &response, nil
}

func (me *RegionServer) CreateProvince(ctx context.Context, province *regionpb.Province) (*regionpb.Province, error) {
	if err := utils.CheckError(me.orm.Table("prvnc").Create(province)); err != nil {
		return nil, err
	}

	return province, nil
}

func (me *RegionServer) UpdateProvince(ctx context.Context, province *regionpb.Province) (*regionpb.Province, error) {
	if err := utils.CheckError(me.orm.Table("prvnc").Save(province)); err != nil {
		return nil, err
	}

	return province, nil
}

func (me *RegionServer) DeleteProvince(ctx context.Context, req *regionpb.DeleteProvinceRequest) (*emptypb.Empty, error) {
	var province regionpb.Province
	if err := utils.CheckError(me.orm.Table("prvnc").Where("id", req.Id).Delete(&province)); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (me *RegionServer) GetCity(ctx context.Context, req *regionpb.GetCityRequest) (*regionpb.City, error) {
	city := regionpb.City{}
	city.Id = req.Id

	if err := utils.CheckError(me.orm.Table("city").First(&city)); err != nil {
		return nil, err
	}

	return &city, nil
}

func (me *RegionServer) ListCities(ctx context.Context, req *regionpb.ListCitiesRequest) (*regionpb.ListCitiesResponse, error) {
	cities := []*regionpb.City{}
	if err := utils.CheckError(me.orm.Table("city").Find(&cities)); err != nil {
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
	if err := utils.CheckError(me.orm.Table("city").Save(city)); err != nil {
		return nil, err
	}

	return city, nil
}

func (me *RegionServer) DeleteCity(ctx context.Context, req *regionpb.DeleteCityRequest) (*emptypb.Empty, error) {
	var city regionpb.City
	if err := utils.CheckError(me.orm.Table("city").Where("id", req.Id).Delete(&city)); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
