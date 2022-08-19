package router

import (
	channel_controller "ocra_server/controller/channel"
	choice_controller "ocra_server/controller/choice"
	comment_controller "ocra_server/controller/comment"
	subscriber_controller "ocra_server/controller/subscribe"
	user_controller "ocra_server/controller/user"
	verification_controller "ocra_server/controller/verification"
	video_controller "ocra_server/controller/video"
	auth_middleware "ocra_server/middleware/auth"
	"ocra_server/middleware/renew_session"
	channel_repository "ocra_server/repository/channels"
	comment_repository "ocra_server/repository/comment"
	dislike_repository "ocra_server/repository/dislike"
	like_repository "ocra_server/repository/like"
	session_repository "ocra_server/repository/session"
	subscribe_repository "ocra_server/repository/subscribe"
	user_repository "ocra_server/repository/user"
	verification_repository "ocra_server/repository/verification"
	videos_repository "ocra_server/repository/video"
	channel_service "ocra_server/service/channel"
	choice_service "ocra_server/service/choice"
	comment_service "ocra_server/service/comment"
	cookie_service "ocra_server/service/cookie"
	firebase_service "ocra_server/service/firebase"
	mail_service "ocra_server/service/mail"
	session_service "ocra_server/service/session"
	subscribe_service "ocra_server/service/subscribe"
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

	UserRoute(group, db, dialer, sessionCache, firebaseService, authMiddleware)
	VerifRoute(group, db, dialer)
	VideoRoute(group, db, authMiddleware, firebaseService)
	ChannelRoute(group, db, authMiddleware, firebaseService)
	CommentRoute(group, db, authMiddleware)
	UserChoiceRoute(group, db)
	SubscribeRoute(group, db)
}

func UserRoute(
	group *echo.Group, db *gorm.DB,
	dialer *gomail.Dialer, cache *cache.Cache,
	firebaseService firebase_service.FirebaseService,
	authMiddleware auth_middleware.AuthMiddleware) user_controller.UserController {
	userRepo := user_repository.NewUserRepository(db)
	mailService := mail_service.NewMailService(dialer)
	verifRepo := verification_repository.NewVerificationRepository(db)
	verifService := verification_service.NewVerificationService(verifRepo, mailService, userRepo)
	sessionRepo := session_repository.NewSessionRepository(db)
	sessionService := session_service.NewSessionService(sessionRepo, cache)
	service := user_service.NewUserService(userRepo, verifService, sessionService, firebaseService)
	cookieService := cookie_service.NewCookieService()
	controller := user_controller.NewUserController(service, cookieService)

	group.POST("/auth/login", controller.Login)
	group.DELETE("/logout", controller.Logout)
	group.POST("/register", controller.Register)
	group.PATCH("/user/:id", controller.UpdateUser, authMiddleware.Auth)

	return controller
}

func VerifRoute(group *echo.Group, db *gorm.DB, dialer *gomail.Dialer) verification_controller.VerificationController {
	verifRepo := verification_repository.NewVerificationRepository(db)
	mailService := mail_service.NewMailService(dialer)
	userRepo := user_repository.NewUserRepository(db)
	verifService := verification_service.NewVerificationService(verifRepo, mailService, userRepo)
	controller := verification_controller.NewVerificationController(verifService)

	group.GET("/email-verification", controller.ValidateVerificationToken)
	group.POST("/resend-email-verification", controller.CreateVerificationToken)

	return controller
}

func VideoRoute(group *echo.Group, db *gorm.DB, middleware auth_middleware.AuthMiddleware, firebaseService firebase_service.FirebaseService) video_controller.VideoController {

	likeRepo := like_repository.NewLikeRepository(db)
	dislikeRepo := dislike_repository.NewDislikeRepository(db)
	choiceService := choice_service.NewChoiceService(likeRepo, dislikeRepo)

	subsRepo := subscribe_repository.NewSubsRepository(db)
	subsService := subscribe_service.NewSubsService(subsRepo)

	repo := videos_repository.NewVideosRepository(db)
	service := video_service.NewVideoService(repo, firebaseService, subsService, choiceService)
	controller := video_controller.NewVideoController(service)

	group.GET("/videos", controller.GetAllVideos)
	group.GET("/video", controller.GetDetailVideos)
	group.POST("/video", controller.CreateVideo)

	return controller
}

func ChannelRoute(group *echo.Group, db *gorm.DB, middleware auth_middleware.AuthMiddleware, firebaseService firebase_service.FirebaseService) channel_controller.ChannelController {
	repo := channel_repository.NewChannelRepository(db)
	service := channel_service.NewChannelService(repo, firebaseService)
	controller := channel_controller.NewChannelController(service)

	group.GET("/channel", controller.DetailChannel)
	group.POST("/channel", controller.CreateChannel)

	return controller
}

func CommentRoute(group *echo.Group, db *gorm.DB, middleware auth_middleware.AuthMiddleware) comment_controller.CommentController {
	repo := comment_repository.NewCommentRepository(db)
	service := comment_service.NewCommentService(repo)
	controller := comment_controller.NewCommentController(service)

	group.GET("/comment", controller.GetComment)
	group.POST("/comment", controller.CreateComment)
	group.PATCH("/comment", controller.UpdateComment)
	group.DELETE("/comment", controller.DeleteComment)

	return controller
}

func UserChoiceRoute(group *echo.Group, db *gorm.DB) choice_controller.ChoiceController {
	likeRepo := like_repository.NewLikeRepository(db)
	dislikeRepo := dislike_repository.NewDislikeRepository(db)
	choiceService := choice_service.NewChoiceService(likeRepo, dislikeRepo)
	controller := choice_controller.NewChoiceController(choiceService)

	group.POST("/like", controller.CreateLike)
	group.POST("/dislike", controller.CreateDislike)

	return controller
}

func SubscribeRoute(group *echo.Group, db *gorm.DB) subscriber_controller.SubscribeController {
	repo := subscribe_repository.NewSubsRepository(db)
	service := subscribe_service.NewSubsService(repo)
	controller := subscriber_controller.NewSubsController(service)

	group.POST("/subs", controller.SubscribeChannel)
	group.POST("/unsubs", controller.UnsubscribeChannel)

	return controller
}
