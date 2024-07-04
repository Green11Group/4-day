package postgres

import (
	"database/sql"
	"fmt"
	g "garden/genproto/gardenmangement"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
)

type GardenRepo struct {
	db *sql.DB
}

func NewGardenRepo(db *sql.DB) *GardenRepo {
	return &GardenRepo{db: db}
}

func (gar *GardenRepo) CreateGarden(req *g.CreateGardenRequest) error {
	id := uuid.NewString()
	newtime := time.Now()
	_, err := gar.db.Exec(`
		INSERT INTO gardens (id, user_id, name, type, area_sqm, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		id, req.Garden.UserId, req.Garden.Name, req.Garden.Type, req.Garden.AreaSq, newtime, newtime)
	if err != nil {
		log.Fatalf("Error inserting data: %v", err)
		return err
	}
	return nil
}

func (gar *GardenRepo) GetGarden(req *g.GetGardenRequest) (*g.GetGardenResponse, error) {
	garden := g.GetGardenResponse{}
	row := gar.db.QueryRow("select user_id, name, type, area_sqm from gardens where id=$1", req.Id)
	// fmt.Println(row)
	err := row.Scan(&garden.UserId, &garden.Name, &garden.Type, &garden.AreaSqm)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User profile not found?")
		}
		return nil, err
	}
	return &garden, nil
}

func (gar *GardenRepo) UpdateGarden(req *g.UpdateGardenRequest) error {
	newtime := time.Now()
	_, err := gar.db.Exec("update gardens set name=$1,type=$2,area_sqm=$3,updated_at=$4 where id=$5 and user_id=$6", req.Garden.Name, req.Garden.Type, req.Garden.AreaSq, newtime, req.Garden.Id, req.Garden.UserId)
	if err != nil {
		log.Fatal("Update error")
		return err
	}

	return nil
}

func (gar *GardenRepo) DeleteGarden(req *g.DeleteGardenRequest) error {
	newtime := time.Now()
	trimmedID := strings.TrimSpace(req.Id)
	_, err := gar.db.Exec("update gardens set deleted_at=$1 where id=$2", newtime, trimmedID)
	if err != nil {
		log.Fatal("Error deleting data")
		return err
	}
	return nil
}

func (gar *GardenRepo) GetUserGardens(req *g.GetUserGardensRequest) (*g.GetUserGardensResponse, error) {
	var gardenuser g.GetUserGardensResponse
	rows, err := gar.db.Query("select id,user_id,name,type,area_sqm from gardens where user_id=$1 and deleted_at is null", req.UserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var garden g.Garden
		err := rows.Scan(&garden.Id,&garden.UserId,&garden.Name, &garden.Type, &garden.AreaSq)
		if err != nil {
			return nil, err
		}
		gardenuser.Gardens = append(gardenuser.Gardens, &garden)
	}

	return &gardenuser, nil
}

func (gar *GardenRepo) CreatePlant(req *g.CreatePlantRequest) error{
	id:=uuid.NewString()
	newtime:=time.Now().Format("2006-01-02 15:04:05")

	_, err := gar.db.Exec("INSERT INTO plants(id, garden_id, species, quantity, planting_date, status, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8)",
        id, req.Plant.GardenId, req.Plant.Species, req.Plant.Quantity, req.Plant.PlantingDate, req.Plant.Status, newtime, newtime)
	if err!=nil{
		log.Fatal("Error inserting data",err)
		return err
	}
	return nil
}

func (gar *GardenRepo) GetPlant(req *g.GetPlantRequest) (*g.GetPlantResponse,error){
	var plants g.GetPlantResponse

	rows,err:=gar.db.Query("select id,garden_id,species,quantity,planting_date,status from plants where garden_id=$1 and deleted_at is null",req.GardenId)
	if err!=nil{
		return nil,err
	}
	defer rows.Close()

	for rows.Next() {
		var plant g.Plant
		err := rows.Scan(&plant.Id, &plant.GardenId, &plant.Species, &plant.Quantity, &plant.PlantingDate, &plant.Status)
		if err != nil {
			return nil, err
		}
		plants.Plants=append(plants.Plants,&plant)
	}
	return &plants,nil
}

func (gar *GardenRepo) UpdatePlant(req *g.UpdatePlantRequest) error{
	newtime:=time.Now()
	_,err:=gar.db.Exec("update plants set garden_id=$1,species=$2,quantity=$3,planting_date=$4,status=$5,updated_at=$7 where id=$6",req.Plant.GardenId,req.Plant.Species,req.Plant.Quantity,req.Plant.PlantingDate,req.Plant.Status,req.Plant.Id,newtime)
	if err!=nil{
		log.Fatal("Update error?",err)
		return err
	}
	return nil
}

func (gar *GardenRepo) DeletePlant(req *g.DeletePlantRequest) error{
	newtime:=time.Now()
	_,err:=gar.db.Exec("update plants set deleted_at=$1 where id=$2",newtime,req.Id)
	if err!=nil{
		return err
	}

	return nil
}

func (gar *GardenRepo) CreateCareLog(req *g.CreateCareLogRequest) error{
	id:=uuid.NewString()
	newtime:=time.Now()
	_,err:=gar.db.Exec("insert into care_logs(id,plant_id,action,notes,logged_at) values($1,$2,$3,$4,$5)",id,req.CareLog.PlantId,req.CareLog.Action,req.CareLog.Notes,newtime)
	if err!=nil{
		log.Fatal("Inserting data?",err)
		return err
	}
	return nil
}

func (gar *GardenRepo) GetCareLog(req *g.GetCareLogRequest) (*g.GetCareLogResponse,error){
	var carelog g.GetCareLogResponse

	rows,err:=gar.db.Query("SELECT id, plant_id, action, notes FROM care_logs WHERE plant_id = $1",req.PlantId)
	if err!=nil{
		return nil,err
	}
	defer rows.Close()

	for rows.Next(){
		var care g.CareLog
		err:=rows.Scan(&care.Id,&care.PlantId,&care.Action,&care.Notes)
		if err!=nil{
			return nil,err
		}
		carelog.CareLog=append(carelog.CareLog, &care)
	}

	return &carelog,nil
}