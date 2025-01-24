package main

import (
	"log"

	"github.com/Alam049/golang-campus/internal/configs"
	"github.com/Alam049/golang-campus/internal/handler/memberships"
	membershipRepo "github.com/Alam049/golang-campus/internal/repository/memberships"
	membershipSvc "github.com/Alam049/golang-campus/internal/service/memberships"
	"github.com/Alam049/golang-campus/pkg/internalsql"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolders([]string{"./internal/configs"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatalf("failed to init config: %v", err)
	}
	cfg = configs.Get()
	log.Println("config", cfg)

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Failted to init database", err)
	}

	membershipRepo := membershipRepo.NewRepository(db)

	membershipService := membershipSvc.NewService(membershipRepo)

	membershipHandler := memberships.NewHandler(r, membershipService)
	membershipHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
