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

type SetupService struct {
	Group           *echo.Group
	Db              *gorm.DB
	Dialer          *gomail.Dialer
	FirebaseService firebase_service.FirebaseService
	SessionCache    *cache.Cache
	AuthMiddleware  auth_middleware.AuthMiddleware
	RenewMiddleware renew_session.RenewSession
}

func Router(setup *SetupService) {

	setup.SessionCache = cache.New(time.Hour*72, time.Hour*120)

	repo := session_repository.NewSessionRepository(setup.Db)
	service := session_service.NewSessionService(repo, setup.SessionCache)
	setup.RenewMiddleware = renew_session.NewRenewSession(service)

	setup.AuthMiddleware = auth_middleware.NewAuthMiddleware(service)

	setup.Group.Use(setup.RenewMiddleware.RenewSession)

	UserRoute(setup)
	VerifRoute(setup)
	VideoRoute(setup)
	ChannelRoute(setup)
	CommentRoute(setup)
	UserChoiceRoute(setup)
	SubscribeRoute(setup)
}

func UserRoute(setup *SetupService) user_controller.UserController {
	userRepo := user_repository.NewUserRepository(setup.Db)
	mailService := mail_service.NewMailService(setup.Dialer)
	verifRepo := verification_repository.NewVerificationRepository(setup.Db)
	verifService := verification_service.NewVerificationService(verifRepo, mailService, userRepo)
	sessionRepo := session_repository.NewSessionRepository(setup.Db)
	sessionService := session_service.NewSessionService(sessionRepo, setup.SessionCache)
	service := user_service.NewUserService(userRepo, verifService, sessionService, setup.FirebaseService)
	cookieService := cookie_service.NewCookieService()
	controller := user_controller.NewUserController(service, cookieService)

	setup.Group.POST("/auth/login", controller.Login)
	setup.Group.DELETE("/logout", controller.Logout)
	setup.Group.POST("/register", controller.Register)
	setup.Group.PATCH("/user/:id", controller.UpdateUser, setup.AuthMiddleware.Auth)

	return controller
}

func VerifRoute(setup *SetupService) verification_controller.VerificationController {
	verifRepo := verification_repository.NewVerificationRepository(setup.Db)
	mailService := mail_service.NewMailService(setup.Dialer)
	userRepo := user_repository.NewUserRepository(setup.Db)
	verifService := verification_service.NewVerificationService(verifRepo, mailService, userRepo)
	controller := verification_controller.NewVerificationController(verifService)

	setup.Group.GET("/email-verification", controller.ValidateVerificationToken)
	setup.Group.POST("/resend-email-verification", controller.CreateVerificationToken)

	return controller
}

func VideoRoute(setup *SetupService) video_controller.VideoController {

	likeRepo := like_repository.NewLikeRepository(setup.Db)
	dislikeRepo := dislike_repository.NewDislikeRepository(setup.Db)
	choiceService := choice_service.NewChoiceService(likeRepo, dislikeRepo)

	subsRepo := subscribe_repository.NewSubsRepository(setup.Db)
	subsService := subscribe_service.NewSubsService(subsRepo)

	repo := videos_repository.NewVideosRepository(setup.Db)
	service := video_service.NewVideoService(repo, setup.FirebaseService, subsService, choiceService)
	controller := video_controller.NewVideoController(service)

	setup.Group.GET("/videos", controller.GetAllVideos)
	setup.Group.GET("/video", controller.GetDetailVideos)
	setup.Group.POST("/video", controller.CreateVideo)

	return controller
}

func ChannelRoute(setup *SetupService) channel_controller.ChannelController {
	repo := channel_repository.NewChannelRepository(setup.Db)
	service := channel_service.NewChannelService(repo, setup.FirebaseService)
	controller := channel_controller.NewChannelController(service)

	setup.Group.GET("/channel", controller.DetailChannel)
	setup.Group.POST("/channel", controller.CreateChannel)

	return controller
}

func CommentRoute(setup *SetupService) comment_controller.CommentController {
	repo := comment_repository.NewCommentRepository(setup.Db)
	service := comment_service.NewCommentService(repo)
	controller := comment_controller.NewCommentController(service)

	setup.Group.GET("/comment", controller.GetComment)
	setup.Group.POST("/comment", controller.CreateComment)
	setup.Group.PATCH("/comment", controller.UpdateComment)
	setup.Group.DELETE("/comment", controller.DeleteComment)

	return controller
}

func UserChoiceRoute(setup *SetupService) choice_controller.ChoiceController {
	likeRepo := like_repository.NewLikeRepository(setup.Db)
	dislikeRepo := dislike_repository.NewDislikeRepository(setup.Db)
	choiceService := choice_service.NewChoiceService(likeRepo, dislikeRepo)
	controller := choice_controller.NewChoiceController(choiceService)

	setup.Group.POST("/like", controller.CreateLike)
	setup.Group.POST("/dislike", controller.CreateDislike)
	setup.Group.DELETE("/like", controller.DeleteLike)
	setup.Group.DELETE("/dislike", controller.DeleteDislike)

	return controller
}

func SubscribeRoute(setup *SetupService) subscriber_controller.SubscribeController {
	repo := subscribe_repository.NewSubsRepository(setup.Db)
	service := subscribe_service.NewSubsService(repo)
	controller := subscriber_controller.NewSubsController(service)

	setup.Group.POST("/subs", controller.SubscribeChannel)
	setup.Group.POST("/unsubs", controller.UnsubscribeChannel)

	return controller
}
