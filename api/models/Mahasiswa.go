package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Mahasiswa struct {
	ID                         uint32    `json:"id"`
	Nm_pd                      string    `json:"nm_pd"`
	Jk                         string    `json:"jk"`
	Npwp                       string    `json:"npwp"`
	Nik                        string    `json:"nik"`
	Tmpt_lahir                 string    `json:"tmpt_lahir"`
	Tgl_lahir                  string    `json:"tgl_lahir"`
	Id_agama                   string    `json:"id_agama"`
	Id_kk                      string    `json:"id_kk"`
	Jln                        string    `json:"jln"`
	Rt                         string    `json:"rt"`
	Rw                         string    `json:"rw"`
	Nm_dsn                     string    `json:"nm_dsn"`
	Ds_kel                     string    `json:"ds_kel"`
	Id_wil                     string    `json:"id_wil"`
	Kode_pos                   string    `json:"kode_pos"`
	Id_jns_tinggal             string    `json:"id_jns_tinggal"`
	Id_alat_transport          string    `json:"id_alat_transport"`
	No_tel_rmh                 string    `json:"no_tel_rmh"`
	No_hp                      string    `json:"no_hp"`
	Email                      string    `json:"email"`
	A_terima_kps               string    `json:"a_terima_kps"`
	No_kps                     string    `json:"no_kps"`
	Stat_pd                    string    `json:"stat_pd"`
	Nik_ayah                   string    `json:"nik_ayah"`
	Nm_ayah                    string    `json:"nm_ayah"`
	Tgl_lahir_ayah             string    `json:"tgl_lahir_ayah"`
	Id_jenjang_pendidikan_ayah string    `json:"id_jenjang_pendidikan_ayah"`
	Id_kebutuhan_khusus_ayah   string    `json:"id_kebutuhan_khusus_ayah"`
	Id_kebutuhan_khusus_ibu    string    `json:"id_kebutuhan_khusus_ibu"`
	Id_pekerjaan_ayah          string    `json:"id_pekerjaan_ayah"`
	Id_penghasilan_ayah        string    `json:"id_penghasilan_ayah"`
	Nik_ibu                    string    `json:"nik_ibu"`
	Nm_ibu_kandung             string    `json:"nm_ibu_kandung"`
	Tgl_lahir_ibu              string    `json:"tgl_lahir_ibu"`
	Id_jenjang_pendidikan_ibu  string    `json:"id_jenjang_pendidikan_ibu"`
	Id_penghasilan_ibu         string    `json:"id_penghasilan_ibu"`
	Id_pekerjaan_ibu           string    `json:"id_pekerjaan_ibu"`
	Nm_wali                    string    `json:"nm_wali"`
	Tgl_lahir_wali             string    `json:"tgl_lahir_wali"`
	Id_jenjang_pendidikan_wali string    `json:"id_jenjang_pendidikan_wali"`
	Id_pekerjaan_wali          string    `json:"id_pekerjaan_wali"`
	Id_penghasilan_wali        string    `json:"id_penghasilan_wali"`
	Kewarganegaraan            string    `json:"kewarganegaraan"`
	Kode_jurusan               string    `json:"kode_jurusan"`
	Id_jns_daftar              string    `json:"id_jns_daftar"`
	Nipd                       string    `json:"nipd"`
	Tgl_masuk_sp               time.Time `json:"tgl_masuk_sp"`
	Mulai_smt                  string    `json:"mulai_smt"`
	Id_jalur_masuk             string    `json:"id_jalur_masuk"`
	Dosen_pa                   string    `json:"dosen_pa"`
	Password                   string    `json:"password"`
}

func (Mahasiswa) TableName() string {
	return "mhs"
}

func (mhs *Mahasiswa) Prepare() {
	mhs.ID = 0
	mhs.Nipd = html.EscapeString(strings.TrimSpace(mhs.Nipd))
	mhs.Nm_pd = html.EscapeString(strings.TrimSpace(mhs.Nm_pd))
}

func (mhs *Mahasiswa) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if mhs.Nipd == "" {
			return errors.New("Required NIM")
		}
		if mhs.Nm_pd == "" {
			return errors.New("Required Nama Peserta Didik")
		}
		return nil
	default:
		if mhs.Nipd == "" {
			return errors.New("Required NIM")
		}
		if mhs.Nm_pd == "" {
			return errors.New("Required Nama Peserta Didik")
		}
		return nil
	}
}

func (mhs *Mahasiswa) SaveMahasiswa(db *gorm.DB) (*Mahasiswa, error) {

	var err error
	err = db.Debug().Create(&mhs).Error
	if err != nil {
		return &Mahasiswa{}, err
	}
	return mhs, nil
}

func (mhs *Mahasiswa) GetAllMhs(db *gorm.DB) (*[]Mahasiswa, error) {
	var err error
	mahasiswa := []Mahasiswa{}
	err = db.Debug().Model(&User{}).Find(&mahasiswa).Error
	if err != nil {
		return &[]Mahasiswa{}, err
	}
	return &mahasiswa, err
}

func (mhs *Mahasiswa) GetMahasiswaByID(db *gorm.DB, uid uint32) (*Mahasiswa, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&mhs).Error
	if err != nil {
		return &Mahasiswa{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Mahasiswa{}, errors.New("Nim tidak ditemukan")
	}
	return mhs, err
}

func (mhs *Mahasiswa) UpdateMahasiswa(db *gorm.DB, uid uint32) (*Mahasiswa, error) {

	db = db.Debug().Model(&Mahasiswa{}).Where("id = ?", uid).Take(&Mahasiswa{}).UpdateColumns(
		map[string]interface{}{
			"nipd":  mhs.Nipd,
			"nm_pd": mhs.Nm_pd,
		},
	)
	if db.Error != nil {
		return &Mahasiswa{}, db.Error
	}

	return mhs, nil
}

func (mhs *Mahasiswa) DeleteMahasiswa(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&Mahasiswa{}).Where("id = ?", uid).Take(&Mahasiswa{}).Delete(&Mahasiswa{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
