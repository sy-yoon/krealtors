#golang download
#curl -O -L https://go.dev/dl/go1.18.1.linux-amd64.tar.gz



# krealtors-dev
#ssh jake@192.168.140.130


# krealtors-prod
#ssh -i "AWS-ECF.pem" ec2-user@ec2-54-215-238-23.us-west-1.compute.amazonaws.com


# Gorm GeoLocation
```
func (x GeoLocation) GormDataType() string {
	return "point"
  }

func (x GeoLocation) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "Point(?,?)",
		Vars: []interface{}{x.Lat, x.Lng},
	}
}

func (x *GeoLocation) Scan(v interface{}) error {
	var data []byte
    switch v := v.(type) {
    case []byte:
        data = v
    case string:
        data = []byte(v)
    case nil:
        return nil
    default:
        return errors.New("(*Point).Scan: unsupported data type")
    }

    if len(data) == 0 {
        return nil
    }

	var err error
    data = data[1 : len(data)-1] // drop the surrounding parentheses
    for i := 0; i < len(data); i++ {
        if data[i] == ',' {
            if x.Lat, err = strconv.ParseFloat(string(data[:i]), 64); err != nil {
                return err
            }

            if x.Lng, err = strconv.ParseFloat(string(data[i+1:]), 64); err != nil {
                return err
            }
            break
        }
    }
	return nil
}
```