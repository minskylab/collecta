package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/minskylab/collecta/ent"
	"golang.org/x/net/context"
)

const userCtxKey = "user"

// Middleware decodes the share session cookie and packs the session into context
func (collectaAuth *Auth) Middleware() gin.HandlerFunc {
	return func(c  *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader,"Bearer") {
			c.String(http.StatusForbidden, "invalid Token")
			return
		}

		token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))


		userRequester, err := collectaAuth.verifyJWTToken(c.Request.Context(), token)
		if err != nil {
			c.String(http.StatusForbidden, "invalid Token")
			return
		}

		// put it in context
		ctx := context.WithValue(c.Request.Context(), userCtxKey, userRequester)

		// and call the next with our new context

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// UserOfContext finds the user from the context. REQUIRES Middleware to have run.
func (collectaAuth *Auth) UserOfContext(ctx context.Context) *ent.User {
	raw, isOk := ctx.Value(userCtxKey).(*ent.User)
	if !isOk {
		return nil
	}
	return raw
}