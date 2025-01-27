package main

import (
	"log"

	"github.com/Alam049/golang-campus/internal/configs"
	"github.com/Alam049/golang-campus/internal/handler/memberships"
	"github.com/Alam049/golang-campus/internal/handler/posts"
	membershipRepo "github.com/Alam049/golang-campus/internal/repository/memberships"
	postRepo "github.com/Alam049/golang-campus/internal/repository/posts"
	membershipSvc "github.com/Alam049/golang-campus/internal/service/memberships"
	postSvc "github.com/Alam049/golang-campus/internal/service/posts"
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

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Failted to init database", err)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	membershipRepo := membershipRepo.NewRepository(db)
	postRepo := postRepo.NewRepository(db)

	membershipService := membershipSvc.NewService(cfg, membershipRepo)
	postService := postSvc.NewService(cfg, postRepo)

	membershipHandler := memberships.NewHandler(r, membershipService)
	membershipHandler.RegisterRoute()

	postHandler := posts.NewHandler(r, postService)
	postHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
