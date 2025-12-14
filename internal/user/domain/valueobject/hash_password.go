package valueobject

import (
	user_errors "yoyoj-go-rag/internal/user/domain/errors"

	"golang.org/x/crypto/bcrypt"
)

type HashedPassword struct {
	hash string
}

func NewHashedPassword(plainPassword string) (*HashedPassword, error) {
	// bcrypt.DefaultCost = 10, 可以根据需要调整 (范围: 4-31)
	// 成本越高，计算越慢，安全性越高
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(plainPassword),
		bcrypt.DefaultCost, // 或使用 bcrypt.Cost(12) 提高安全性
	)
	if err != nil {
		return nil, user_errors.ErrPasswordHashFailed
	}
	return &HashedPassword{hash: string(hash)}, nil
}

// FromHash 从已有哈希创建（用于从数据库读取）
func FromHash(hash string) *HashedPassword {
	return &HashedPassword{hash: hash}
}

// Verify 验证明文密码是否匹配
func (hp *HashedPassword) Verify(plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hp.hash), []byte(plainPassword))
	return err == nil
}

// String 返回哈希字符串（用于存储到数据库）
func (hp *HashedPassword) String() string {
	return hp.hash
}