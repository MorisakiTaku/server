// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

const TableNameEpisodeComment = "chii_ep_comments"

// EpisodeComment mapped from table <chii_ep_comments>
type EpisodeComment struct {
	ID          uint32 `gorm:"column:ep_pst_id;type:mediumint(8) unsigned;primaryKey;autoIncrement:true"`
	TopicID     uint32 `gorm:"column:ep_pst_mid;type:mediumint(8) unsigned;not null"`
	UID         uint32 `gorm:"column:ep_pst_uid;type:mediumint(8) unsigned;not null"`
	Related     uint32 `gorm:"column:ep_pst_related;type:mediumint(8) unsigned;not null"`
	CreatedTime uint32 `gorm:"column:ep_pst_dateline;type:int(10) unsigned;not null"`
	Content     string `gorm:"column:ep_pst_content;type:mediumtext;not null"`
}

// TableName EpisodeComment's table name
func (*EpisodeComment) TableName() string {
	return TableNameEpisodeComment
}