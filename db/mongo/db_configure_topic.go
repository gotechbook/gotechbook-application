package mongo

import "time"

// TbConfigureTopicName topic 配置表
const TbConfigureTopicName = "tb_configure_topic"

type TbConfigureTopic struct {
	ID        string    `bson:"id" json:"ID,omitempty"`            // ID
	Code      int64     `bson:"code" json:"code,omitempty"`        // Code
	Name      string    `bson:"name" json:"name,omitempty"`        // Name
	Describe  string    `bson:"describe"json:"describe,omitempty"` // 描述
	IsRead    bool      `bson:"isRead" json:"isRead,omitempty"`    // 是否只读
	IsDel     bool      `bson:"isDel" json:"isDel,omitempty"`      // 是否删除
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}
