package jwt

import (
	"errors"
	"fmt"
	"go-template/internal/configs"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go-template/internal/utils/logutil"
)

// Service Struct，用於產生和驗證 JWT token
type Service struct {
	secretKey     string
	oldSecretKeys []string // 新增欄位儲存舊密鑰
	issuer        string
	expiration    time.Duration
}

// NewService 建立一個新的 JWTService 實例
func NewService(cfg *configs.Config) *Service {
	logutil.Logger.Infof("Initializing JWTService with secret: %s", cfg.JWTSecret) // 新增日誌
	return &Service{
		secretKey:     cfg.JWTSecret,
		oldSecretKeys: cfg.JWTOldSecrets,  // 初始化舊密鑰列表
		issuer:        "go-template",      // JWT 的發行者，可以根據你的應用程式修改
		expiration:    cfg.TokenExpiresIn, // Token 的過期時間
	}
}

// GenerateToken 產生一個 JWT token
// @param userID 要放入 token 中的使用者 ID
// @return string 產生的 JWT token
// @return error 錯誤訊息
func (s *Service) GenerateToken(userID uint) (string, error) {
	// 建立 JWT claims (有效負載)
	claims := &jwt.RegisteredClaims{
		Issuer:    s.issuer,                                         // 設定發行者
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.expiration)), // 設定過期時間
		IssuedAt:  jwt.NewNumericDate(time.Now()),                   // 設定發行時間
		Subject:   strconv.FormatUint(uint64(userID), 10),           // 將使用者 ID 轉換成字串並設定為 Subject
	}

	// 使用 HMAC SHA256 演算法和 secretKey 簽署 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secretKey))
}

// ValidateToken 驗證 JWT token
// @param tokenString 要驗證的 JWT token
// @return uint 使用者 ID
// @return error 錯誤訊息
func (s *Service) ValidateToken(tokenString string) (uint, error) {
	// 先嘗試使用當前密鑰解析 token
	userID, err := s.validateTokenWithSecret(tokenString, s.secretKey)
	if err == nil {
		return userID, nil
	}

	// 嘗試使用舊密鑰解析 token
	for _, oldSecret := range s.oldSecretKeys {
		userID, err := s.validateTokenWithSecret(tokenString, oldSecret)
		if err == nil {
			logutil.Logger.Warnf("Token validated with old secret key")
			return userID, nil
		}
	}

	logutil.Logger.Debugf("Invalid token: %v", err)
	return 0, errors.New("invalid token")
}

// validateTokenWithSecret 使用指定的 secret 驗證 token
// @param tokenString 要驗證的 JWT token
// @param secret 要使用的 secret 密鑰
// @return uint 使用者 ID
func (s *Service) validateTokenWithSecret(tokenString, secret string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 檢查簽名演算法是否符合預期
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// 返回 secretKey 用於驗證簽名
		return []byte(secret), nil
	})

	// 如果解析失敗，返回錯誤
	if err != nil {
		return 0, err
	}

	// 檢查 token 是否有效
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, err := strconv.ParseUint(claims["sub"].(string), 10, 32)
		if err != nil {
			return 0, errors.New("invalid user ID in token")
		}
		return uint(userID), nil
	}

	return 0, errors.New("invalid token")
}
