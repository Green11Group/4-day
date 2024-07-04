package servicegarden

import (
	"context"
	"database/sql"
	g "garden/genproto/gardenmangement"
	"garden/storage/postgres"
	"log"
)

type GardenServer struct {
	g.UnimplementedGardenServiceServer
	db     *sql.DB
	garden *postgres.GardenRepo
}

func NewGardenServer(db *sql.DB, garden *postgres.GardenRepo) *GardenServer {
	return &GardenServer{db: db, garden: garden}
}

func (s *GardenServer) CreateGarden(ctx context.Context, req *g.CreateGardenRequest) (*g.CreateGardenResponse, error) {
	err := s.garden.CreateGarden(req)
	if err != nil {
		log.Fatalf("Error creating garden: %v", err)
		return nil, err
	}
	return &g.CreateGardenResponse{Message: "Data has been inserted", Success: true}, nil
}

func (s *GardenServer) GetGarden(ctx context.Context, req *g.GetGardenRequest) (*g.GetGardenResponse, error) {
	garden, err := s.garden.GetGarden(req)
	if err != nil {
		log.Fatalf("11111\n: %v", err)
		return nil, err
	}
	return garden, nil
}

func (s *GardenServer) UpdateGarden(ctx context.Context, req *g.UpdateGardenRequest) (*g.UpdateGardenResponse, error) {
	err := s.garden.UpdateGarden(req)
	if err != nil {
		log.Fatalf("2222\n: %v", err)
		return nil, err
	}

	return &g.UpdateGardenResponse{Message: "Information has been updated"}, nil
}

func (s *GardenServer) DeleteGareden(ctx context.Context, req *g.DeleteGardenRequest) (*g.DeleteGardenResponse, error) {
	err := s.garden.DeleteGarden(req)
	// fmt.Println(req)
	if err != nil {
		log.Fatalf("3333\n: %v", err)
		return nil, err
	}
	return &g.DeleteGardenResponse{Message: "Data deleted"}, nil
}

func (s *GardenServer) GetUserGardens(ctx context.Context, req *g.GetUserGardensRequest) (*g.GetUserGardensResponse, error) {
	gardenuser, err := s.garden.GetUserGardens(req)
	if err != nil {
		log.Fatalf("4444\n: %v", err)
		return nil, err
	}
	return gardenuser, nil
}

func (s *GardenServer) CreatePlant(ctx context.Context,req *g.CreatePlantRequest)(*g.CreatePlantReponse,error){
	err:=s.garden.CreatePlant(req)
	if err!=nil{
		log.Fatal("Error creating plant?")
		return nil,err
	}

	return &g.CreatePlantReponse{Message: "Data has been inserted"},nil
}

func (s *GardenServer) GetPlant(ctx context.Context,req *g.GetPlantRequest) (*g.GetPlantResponse,error){
	plants,err:=s.garden.GetPlant(req)
	if err!=nil{
		log.Fatal("Could not get plants")
		return nil,err
	}

	return plants,nil
}

func (s *GardenServer) UpdatePlant(ctx context.Context,req *g.UpdatePlantRequest) (*g.UpdatePlantResponse,error){
	err:=s.garden.UpdatePlant(req)
	if err!=nil{
		log.Fatal("Update error plants?")
		return nil,err
	}

	return &g.UpdatePlantResponse{Message: "information has been updated"},nil
}

func (s *GardenServer) DeletePlant(ctx context.Context,req *g.DeletePlantRequest) (*g.DeletePlantResponse,error){
	err:=s.garden.DeletePlant(req)
	if err!=nil{
		return nil,err
	}

	return &g.DeletePlantResponse{Message: "Data deleted"},nil
}

func (s *GardenServer) CreateCareLog(ctx context.Context,req *g.CreateCareLogRequest) (*g.CreateCareLogResponse,error){
	err:=s.garden.CreateCareLog(req)
	if err!=nil{
		return nil,err
	}
	return &g.CreateCareLogResponse{Message: "A new care schedule has been created"},nil
}

func (s *GardenServer) GetCareLog(ctx context.Context,req *g.GetCareLogRequest) (*g.GetCareLogResponse,error){
	carelog,err:=s.garden.GetCareLog(req)
	if err!=nil{
		return nil,err
	}
	return carelog,nil
}