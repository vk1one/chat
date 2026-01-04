// model/user.go
package model

type User struct {
    ID        int64     `json:"id"          db:"id"`
    Email     string    `json:"email"       db:"email"`
    Password  string    `json:"-"           db:"password_hash"` // хранится хэш
    CreatedAt time.Time `json:"created_at"  db:"created_at"`
}
// model/conversation.go
package model

type ConversationType int

const (
    PrivateConv ConversationType = iota + 1
    GroupConv
)

type Conversation struct {
    ID          int64            `json:"id"          db:"id"`
    Type        ConversationType `json:"type"        db:"type"`
    Title       string           `json:"title,omitempty" db:"title"` // только для групповых чатов
    CreatedAt   time.Time        `json:"created_at"  db:"created_at"`
}
// model/message.go
package model

type Message struct {
    ID          int64     `json:"id"          db:"id"`
    ConvID      int64     `json:"conv_id"     db:"conversation_id"`
    SenderID    int64     `json:"sender_id"   db:"sender_id"`
    Content     string    `json:"content"     db:"content"`
    CreatedAt   time.Time `json:"created_at"  db:"created_at"`
}
