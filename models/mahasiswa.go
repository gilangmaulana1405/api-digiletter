package models

type Mahasiswa struct{
	Id int64 `gorm:"primaryKey" json:"id"` 
	Npm int64 `gorm:"type:bigint" json:"npm"` 
	Prodi string `gorm:"type:varchar(255)" json:"prodi"` 
	Status string `gorm:"type:varchar(255)" json:"status"` 
}