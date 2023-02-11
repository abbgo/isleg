package backApi

import (
	"github/abbgo/isleg/isleg-backend/auth"
	"github/abbgo/isleg/isleg-backend/middlewares"

	adminController "github/abbgo/isleg/isleg-backend/controllers/admin"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(back *gin.RouterGroup) {

	admin := back.Group("/auth")
	{
		{
			// RegisterAdmin admin registrasiya etmek ucin ulanylyar.
			// Admini dine super admin registrasiya edip bilyar. Admin admin registrasiya edip bilenok
			admin.POST("register", middlewares.IsSuperAdmin(), adminController.RegisterAdmin)

			// Adminlerin maglumatlaryny uytgetmek ucin ulanylyar. Adminlerin maglumatlaryny dine super admin
			// uytgedip bilyar. Admin hic bit adminin maglumatlaryny uytgedip bilenok
			admin.PUT("information-of-admin", middlewares.IsSuperAdmin(), adminController.UpdateAdminInformation)

			// Adminlerin parollaryny uytgetmek ucin ulanylyar. Islendik adminin parolyny dine
			// super admin uytgedip bilyar. Admin hic bir adminin parolyny uytgedip bilenok
			admin.PUT("password-of-admin/:id", middlewares.IsSuperAdmin(), adminController.UpdateAdminPassword)

			// LoginAdmin funksiya admin login bolmak ucin ulanylyar.
			admin.POST("login", adminController.LoginAdmin)

			// Adminlerin access tokenin tazelelap refresh bilen access tokeni bile bermek
			// ucin ulanylyar
			admin.POST("refresh", auth.RefreshTokenForAdmin)

			// GetAdmins funksiya hemme adminlerin spisoygyny almak ucin ulanylyar.
			admin.GET("admins/:limit/:page", middlewares.CheckAdmin(), adminController.GetAdmins)

			admin.GET("admin", middlewares.CheckAdmin(), adminController.GetAdmin)

		}
	}

}
