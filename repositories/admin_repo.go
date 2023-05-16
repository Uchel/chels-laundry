package repositories

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Uchel/chels-laundry/models"
)

type AdminLaundryRepo interface {
	GetAll() any
	GetByEmail(email string) any
	GetPhotoByEmail(email string) any
	Create(newAdmin *models.Admin) string
	Update(admin *models.Admin, email string) string
	UpdatePhoto(photo string, email string) string
	Delete(email string) string
}

type adminLaundryRepo struct {
	db *sql.DB
}

// ============================================= Get All===============================================
func (r *adminLaundryRepo) GetAll() any {
	var admins []models.Admin

	query := "SELECT id,name, email,phone,photo from admin_laundry order by id asc"
	rows, err := r.db.Query(query)

	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var admin models.Admin

		if err := rows.Scan(&admin.ID, &admin.Name, &admin.Email, &admin.Phone, &admin.Photo); err != nil {
			log.Println(err)
		}

		admins = append(admins, admin)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}

	if len(admins) == 0 {
		return "no data"
	}

	return admins

}

// ============================================= Get By Email=================================================
func (r *adminLaundryRepo) GetByEmail(email string) any {
	var admin models.Admin

	query := "SELECT  name, phone from admin_laundry WHERE email = $1"

	row := r.db.QueryRow(query, email)

	if err := row.Scan(&admin.Name, &admin.Phone); err != nil {
		log.Println(err)
	}

	if admin.Name == "" {
		return "admin not found"
	}

	return admin

}

// ============================================= Get Photo By Email ===============================================

func (r *adminLaundryRepo) GetPhotoByEmail(email string) any {
	var admin models.Admin

	query := "SELECT  photo from admin_laundry WHERE email = $1"

	row := r.db.QueryRow(query, email)

	if err := row.Scan(&admin.Photo); err != nil {
		log.Println(err)
	}

	if admin.Photo == "" {
		return "admin not found"
	}

	return admin.Photo

}

// ============================================= Create ===============================================
func (r *adminLaundryRepo) Create(newAdmin *models.Admin) string {
	query := "INSERT INTO admin_laundry ( id, name, email,password,phone,photo) VALUES($1,$2,$3,$4,$5,$6)"
	_, err := r.db.Exec(query, newAdmin.ID, newAdmin.Name, newAdmin.Email, newAdmin.Password, newAdmin.Phone, newAdmin.Photo)

	if err != nil {
		log.Println(err)
		return "failed to create admin"
	}

	return "admin created successfully"
}

// =============================================Update By email ===============================================
func (r *adminLaundryRepo) Update(admin *models.Admin, email string) string {
	res := r.GetByEmail(email)
	if res == "admin not found" {
		return res.(string)
	}

	query := "UPDATE admin_laundry SET  password = $1  WHERE email = $2 ;"
	_, err := r.db.Exec(query, admin.Password, email)

	if err != nil {
		log.Println(err)
		return "failed to update admin"
	}

	return fmt.Sprintf("admin with email %s updated successfully", email)

}

// ============================================= Photo By Email ===============================================
func (r *adminLaundryRepo) UpdatePhoto(photo string, email string) string {
	res := r.GetByEmail(email)
	if res == "admin not found" {
		return res.(string)
	}

	query := "UPDATE admin_laundry SET photo = $1 WHERE email = $2 ;"
	_, err := r.db.Exec(query, photo, email)

	if err != nil {
		log.Println(err)
		return "failed to update admin"
	}

	return fmt.Sprintf("admin with email %s updated successfully", email)

}

// ============================================= Delete By Email ===============================================
func (r *adminLaundryRepo) Delete(email string) string {
	res := r.GetByEmail(email)
	if res == "admin not found" {
		return res.(string)
	}

	query := "DELETE FROM admin_laundry WHERE email = $1"
	_, err := r.db.Exec(query, email)

	if err != nil {
		log.Println(err)
		return "failed to delete admin"
	}

	return fmt.Sprintf("admin with email %s deleted successfully", email)
}

func NewExampleAdminRepo(db *sql.DB) AdminLaundryRepo {
	repo := new(adminLaundryRepo)

	repo.db = db

	return repo
}
