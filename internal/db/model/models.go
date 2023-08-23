package model

type TableTimeYM struct {
	TableTime string `orm:"column(table_time)"`
	Year      int    `orm:"column(year);"`
	Month     int    `orm:"column(month);"`
}

type TableTimeYMD struct {
	TableTime string `orm:"column(table_time)"`
	Year      int    `orm:"column(year);"`
	Month     int    `orm:"column(month);"`
	Day       int    `orm:"column(day);"`
}

type TableTimeYMDHMS struct {
	TableTime string `orm:"column(table_time);null"`
	Year      int    `orm:"column(year);null"`
	Month     int    `orm:"column(month);null"`
	Day       int    `orm:"column(day);null"`
	Hour      int    `orm:"column(hour);null"`
	Minute    int    `orm:"column(minute);null"`
	Second    int    `orm:"column(second);null"`
}
