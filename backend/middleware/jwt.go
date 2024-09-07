package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"time"
)

// Valida y retorna el token si es válido, o un error si no lo es
func parseJWT(c echo.Context) (*jwt.Token, error) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusFound, "/login?error=expired session")
	}

	tokenString := cookie.Value
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, echo.NewHTTPError(http.StatusUnauthorized, "Unexpected signing method")
		}
		return []byte("your-secret-key"), nil
	})

	return token, err
}

// Verifica si el token ha expirado
func isTokenExpired(claims jwt.MapClaims) bool {
	if exp, ok := claims["exp"].(float64); ok {
		return time.Unix(int64(exp), 0).Before(time.Now())
	}
	return false
}

// Valida el token y redirige si es inválido o expirado
func JWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token, err := parseJWT(c)
			if err != nil {
				return err
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				if isTokenExpired(claims) {
					return c.Redirect(http.StatusFound, "/login?error=expired session")
				}

				c.Set("user_id", claims["user_id"])
				c.Set("username", claims["username"])
				c.Set("email", claims["email"])
				c.Set("role", claims["role"])
			} else {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token claims")
			}

			return next(c)
		}
	}
}

// Verifica que el rol del usuario coincida con el requerido
func roleIsAuthorized(c echo.Context, requiredRole string) bool {
	role := c.Get("role")
	return strings.ToLower(role.(string)) == strings.ToLower(requiredRole)
}

// Middleware para proteger rutas según el rol
func RoleMiddleware(requiredRole string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !roleIsAuthorized(c, requiredRole) {
				return c.Redirect(http.StatusFound, "/"+requiredRole+"?error=unauthorized")
			}
			return next(c)
		}
	}
}
