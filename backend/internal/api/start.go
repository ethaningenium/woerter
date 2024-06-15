package api

import (
	"log"
	"net/http"
	"word/config"
	"word/internal/handler"
	"word/internal/middleware"
	"word/internal/repository"
)

func Start() {
	//Инициализируем конфиг
	cfg := config.New()
	log.Printf("Mode: %s", cfg.Mode)

	//Инициализируем репозиторий
	repo := repository.New(cfg)
	repo.Migrate()
	defer repo.CloseConnection()

	//Инициализируем хендлеры
	handler := handler.New(cfg, repo)

	//Инициализация миддлвейров
	m := middleware.New()

	//Создаем сервер
	server := http.NewServeMux()
	server.HandleFunc("/", handler.Proxy)

	//Создаем пулл запросов
	api := http.NewServeMux()
	api.HandleFunc("/hello", handler.HelloWorld)
	api.HandleFunc("POST /word", m.With(handler.WordCreate, m.Info))
	api.HandleFunc("/word", m.With(handler.WordGetAll, m.Info))
	server.Handle("/api/v1/", http.StripPrefix("/api/v1", api))

	//Создаем пулл google oauth
	gouth := http.NewServeMux()
	gouth.HandleFunc("/login", handler.GoogleLogin)
	gouth.HandleFunc("/callback", handler.GoogleCallback)
	gouth.HandleFunc("/me", m.With(handler.MeGet, m.Info))
	server.Handle("/oauth/google/", http.StripPrefix("/oauth/google", gouth))

	//Слушаем порт
	log.Printf("Starting server on port %s", cfg.Port)
	http.ListenAndServe(cfg.Port, server)
}
