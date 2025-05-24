package route

import (
	"net/http"

	"forum/handlers"
	"forum/middleware"
)

func Route() {
	http.HandleFunc("/", (handlers.HomeHandler))
	http.HandleFunc("/logout", middleware.Auth(handlers.LogOutHandler))

	http.HandleFunc("/login",handlers.LoginShowHandler)
	http.HandleFunc("/register", handlers.RegisterShowHandler)

	http.HandleFunc("/loginAuth", middleware.RateLimitLoginMiddleware( handlers.LoginHandler)) 
	http.HandleFunc("/registerAuth", handlers.RegisterHandler)

	http.HandleFunc("/static/", handlers.StyleHandler)
	http.HandleFunc("/uploads/", handlers.UploadHandler)

	http.HandleFunc("/createPost", middleware.Auth(middleware.RateLimitPostsMiddleware(handlers.CreatePost)))

	http.HandleFunc("/reaction", middleware.Auth(middleware.RateLimitLikesMiddleware(handlers.ReactionHandler)))

	http.HandleFunc("/comment", middleware.Auth(middleware.RateLimitCommentsMiddleware(handlers.CommentHandler)))

	http.HandleFunc("/CommentsLike", middleware.Auth(middleware.RateLimitLikesMiddleware(handlers.CommentsLikeHandler)))

	http.HandleFunc("/filter", handlers.Filter_By_Categorie)

	http.HandleFunc("/myPosts", middleware.Auth(handlers.MyPosts))

	http.HandleFunc("/likedPosts", middleware.Auth(handlers.LikedPosts))
	
	http.HandleFunc("/deletePost", middleware.Auth(handlers.DeleteHandler))
}
