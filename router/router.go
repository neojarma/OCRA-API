package router

import (
	channel_controller "ocra_server/controller/channel"
	comment_controller "ocra_server/controller/comment"
	user_controller "ocra_server/controller/user"
	verification_controller "ocra_server/controller/verification"
	video_controller "ocra_server/controller/video"
	auth_middleware "ocra_server/middleware/auth"
	"ocra_server/middleware/renew_session"
	channel_repository "ocra_server/repository/channels"
	comment_repository "ocra_server/repository/comment"
	session_repository "ocra_server/repository/session"
	user_repository "ocra_server/repository/user"
	verification_repository "ocra_server/repository/verification"
	videos_repository "ocra_server/repository/video"
	channel_service "ocra_server/service/channel"
	comment_service "ocra_server/service/comment"
	cookie_service "ocra_server/service/cookie"
	firebase_service "ocra_server/service/firebase"
	mail_service "ocra_server/service/mail"
	session_service "ocra_server/service/session"
	user_service "ocra_server/service/user"
	verification_service "ocra_server/service/verification"
	video_service "ocra_server/service/video"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

func Router(group *echo.Group, db *gorm.DB, dialer *gomail.Dialer, firebaseService firebase_service.FirebaseService) {

	sessionCache := cache.New(time.Hour*72, time.Hour*120)

	repo := session_repository.NewSessionRepository(db)
	service := session_service.NewSessionService(repo, sessionCache)
	middl := renew_session.NewRenewSession(service)

	authMiddleware := auth_middleware.NewAuthMiddleware(service)

	group.Use(middl.RenewSession)

	userRoute(group, db, dialer, sessionCache)
	verifRoute(group, db, dialer)
	videoRoute(group, db, authMiddleware, firebaseService)
	channelRoute(group, db, authMiddleware, firebaseService)
	commentRoute(group, db, authMiddleware)
}

func userRoute(group *echo.Group, db *gorm.DB, dialer *gomail.Dialer, cache *cache.Cache) {
	userRepo := user_repository.NewUserRepository(db)
	mailService := mail_service.NewMailService(dialer)
	verifRepo := verification_repository.NewVerificationRepository(db)
	verifService := verification_service.NewVerificationService(verifRepo, mailService, userRepo)
	sessionRepo := session_repository.NewSessionRepository(db)
	sessionService := session_service.NewSessionService(sessionRepo, cache)
	service := user_service.NewUserService(userRepo, verifService, sessionService)
	cookieService := cookie_service.NewCookieService()
	controller := user_controller.NewUserController(service, cookieService)

	group.POST("/auth/login", controller.Login)
	group.DELETE("/logout", controller.Logout)
	group.POST("/register", controller.Register)
	group.PATCH("/user/:id", controller.UpdateUser)
}

func verifRoute(group *echo.Group, db *gorm.DB, dialer *gomail.Dialer) {
	verifRepo := verification_repository.NewVerificationRepository(db)
	mailService := mail_service.NewMailService(dialer)
	userRepo := user_repository.NewUserRepository(db)
	verifService := verification_service.NewVerificationService(verifRepo, mailService, userRepo)
	controller := verification_controller.NewVerificationController(verifService)

	group.GET("/email-verification", controller.ValidateVerificationToken)
	group.POST("/resend-email-verification", controller.CreateVerificationToken)
}

func videoRoute(group *echo.Group, db *gorm.DB, middleware auth_middleware.AuthMiddleware, firebaseService firebase_service.FirebaseService) {

	repo := videos_repository.NewVideosRepository(db)
	service := video_service.NewVideoService(repo, firebaseService)
	controller := video_controller.NewVideoController(service)

	group.GET("/videos", controller.GetAllVideos)
	group.GET("/video", controller.GetDetailVideos)
	group.POST("/video", controller.CreateVideo)
}

func channelRoute(group *echo.Group, db *gorm.DB, middleware auth_middleware.AuthMiddleware, firebaseService firebase_service.FirebaseService) {
	repo := channel_repository.NewChannelRepository(db)
	service := channel_service.NewChannelService(repo, firebaseService)
	controller := channel_controller.NewChannelController(service)

	group.GET("/channel", controller.DetailChannel)
	group.POST("/channel", controller.CreateChannel)
}

func commentRoute(group *echo.Group, db *gorm.DB, middleware auth_middleware.AuthMiddleware) {
	repo := comment_repository.NewCommentRepository(db)
	service := comment_service.NewCommentService(repo)
	controller := comment_controller.NewCommentController(service)

	group.GET("/comment", controller.GetComment)
	group.POST("/comment", controller.CreateComment)
	group.PATCH("/comment", controller.UpdateComment)
	group.DELETE("/comment", controller.DeleteComment)
}
