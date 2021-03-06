package datamodel

import (
	"time"
)

type (
	// Vehicles represents the structure of our resource
	Vehicles struct {
		ID       int         `bson:"vtsid" json:"vtsid"`
		Sertm    time.Time   `bson:"sertm" json:"sertm"`
		Loctm    time.Time   `bson:"loctm" json:"loctm"`
		Loc      []float64   `bson:"loc" json:"loc"`
		Speed    int         `bson:"speed" json:"speed"`
		Bearing  int         `bson:"bearing" json:"bearing"`
		Appvr    string      `bson:"appvr" json:"appvr"`
		UID      string      `bson:"uid" json:"uid"`
		Btr      int         `bson:"btr" json:"btr"`
		Btrst    string      `bson:"btrst" json:"btrst"`
		Alwspeed int         `bson:"alwspeed" json:"alwspeed"`
		Flag     string      `bson:"flag" json:"flag"`
		Accr     int         `bson:"accr" json:"accr"`
		Alt      int         `bson:"alt" json:"alt"`
		Gpstm    string      `bson:"gpstm" json:"gpstm"`
		Actvt    string      `bson:"actvt" json:"actvt"`
		Acttm    string      `bson:"acttm" json:"acttm"`
		Vhid     string      `bson:"vhid" json:"vhid"`
		Sat      int         `bson:"sat" json:"sat"`
		Oe       int         `bson:"oe" json:"oe"`
		Gp       int         `bson:"gp" json:"gp"`
		Alm      string      `bson:"alm" json:"alm"`
		Chrg     int         `bson:"chrg" json:"chrg"`
		Acc      int         `bson:"acc" json:"acc"`
		Gsmsig   int         `bson:"gsmsig" json:"gsmsig"`
		Histdate time.Time   `bson:"histtm" json:"-"`
		VhNm     string      `bson:"vhname" json:"vno"`
		Vhd      interface{} `bson:"vhd" json:"vhd"`
		Extra    interface{} `bson:"extra" json:"extra"`
		Vhtyp    string      `bson:"vhtyp" json:"ico"`
		Ipadr    string      `bson:"ip" json:"-"`
		IsSpd    bool        `bson:"isp" json:"-"`
		LstSpd   int         `bson:"lstspd" json:"lstspd"`
		LstSpdtm time.Time   `bson:"lstspdtm" json:"lstspdtm"`
		D1       int         `bson:"d1" json:"d1"`
		Sim      string      `bson:"sim" json:"sim"`
		Clients  []string    `bson:"pushcl" json:"pushcl"`
		Block    bool        `bson:"block" json:"block"`
	}
)
