package data

import (
	"time"
)

type AuthGroup struct {
	Id   int    `xorm:"not null pk autoincr INT"`
	Name string `xorm:"not null unique VARCHAR(150)"`
}

type AuthGroupPermissions struct {
	Id           int `xorm:"not null pk autoincr INT"`
	GroupId      int `xorm:"not null unique(auth_group_permissions_group_id_permission_id_0cd325b0_uniq) INT"`
	PermissionId int `xorm:"not null index unique(auth_group_permissions_group_id_permission_id_0cd325b0_uniq) INT"`
}

type AuthPermission struct {
	Id            int    `xorm:"not null pk autoincr INT"`
	Name          string `xorm:"not null VARCHAR(255)"`
	ContentTypeId int    `xorm:"not null unique(auth_permission_content_type_id_codename_01ab375a_uniq) INT"`
	Codename      string `xorm:"not null unique(auth_permission_content_type_id_codename_01ab375a_uniq) VARCHAR(100)"`
}

type AuthUser struct {
	Id          int       `xorm:"not null pk autoincr INT"`
	Password    string    `xorm:"not null VARCHAR(128)"`
	LastLogin   time.Time `xorm:"DATETIME(6)"`
	IsSuperuser int       `xorm:"not null TINYINT(1)"`
	Username    string    `xorm:"not null unique VARCHAR(150)"`
	FirstName   string    `xorm:"not null VARCHAR(150)"`
	LastName    string    `xorm:"not null VARCHAR(150)"`
	Email       string    `xorm:"not null VARCHAR(254)"`
	IsStaff     int       `xorm:"not null TINYINT(1)"`
	IsActive    int       `xorm:"not null TINYINT(1)"`
	DateJoined  time.Time `xorm:"not null DATETIME(6)"`
}

type AuthUserGroups struct {
	Id      int `xorm:"not null pk autoincr INT"`
	UserId  int `xorm:"not null unique(auth_user_groups_user_id_group_id_94350c0c_uniq) INT"`
	GroupId int `xorm:"not null index unique(auth_user_groups_user_id_group_id_94350c0c_uniq) INT"`
}

type AuthUserUserPermissions struct {
	Id           int `xorm:"not null pk autoincr INT"`
	UserId       int `xorm:"not null unique(auth_user_user_permissions_user_id_permission_id_14a6b632_uniq) INT"`
	PermissionId int `xorm:"not null index unique(auth_user_user_permissions_user_id_permission_id_14a6b632_uniq) INT"`
}

type CitiesCity struct {
	CityCode     string    `xorm:"not null pk comment('市级城市代码') unique UNSIGNED BIGINT"`
	CityName     string    `xorm:"not null comment('市级城市名称') unique(cities_city_index) VARCHAR(50)"`
	CityLevel    string    `xorm:"not null comment('市级城市标签') unique(cities_city_index) VARCHAR(50)"`
	ProvinceCode string    `xorm:"not null comment('省级城市代码') unique(cities_city_index) index UNSIGNED BIGINT"`
	Updated      time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') TIMESTAMP"`
	Created      time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
}

type CitiesCounty struct {
	CountyCode  string    `xorm:"not null pk comment('县级城市代码') unique UNSIGNED BIGINT"`
	CountyName  string    `xorm:"not null comment('县级城市名称') unique(cities_county_index) VARCHAR(50)"`
	CountyLevel string    `xorm:"not null comment('县级城市标签') unique(cities_county_index) VARCHAR(50)"`
	CityCode    string    `xorm:"not null comment('市级城市代码') index unique(cities_county_index) UNSIGNED BIGINT"`
	Updated     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') TIMESTAMP"`
	Created     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
}

type CitiesProvince struct {
	ProvinceCode  string    `xorm:"not null pk comment('省级城市代码') unique UNSIGNED BIGINT"`
	ProvinceName  string    `xorm:"not null comment('省级城市名称') unique(cities_province_index) unique VARCHAR(50)"`
	ProvinceLevel string    `xorm:"not null comment('省级城市标签') unique(cities_province_index) VARCHAR(50)"`
	Updated       time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') TIMESTAMP"`
	Created       time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
}

type Comment struct {
	CommentId           int64     `xorm:"not null pk comment('评论id') BIGINT"`
	UserId              int64     `xorm:"not null comment('评论人userId') index BIGINT"`
	UserName            string    `xorm:"not null default '' comment('评论人名称') VARCHAR(50)"`
	ArticleId           int64     `xorm:"comment('评论的文章id') index BIGINT"`
	ArticleTitle        string    `xorm:"default '' comment('评论的文章标题') VARCHAR(255)"`
	RelationTable       string    `xorm:"comment('关系表，1日记表，2消费表，3学习表，4任务表') VARCHAR(255)"`
	ParentCommentId     int64     `xorm:"comment('父评论id') index BIGINT"`
	ParentCommentUserId int64     `xorm:"not null comment('父评论的用户id') BIGINT"`
	ReplyCommentId      int64     `xorm:"comment('被回复的评论id') BIGINT"`
	ReplyCommentUserId  int64     `xorm:"comment('被回复的评论用户id') BIGINT"`
	CommentLevel        int       `xorm:"not null default 1 comment('评论等级[ 1 一级评论 默认 ，2 二级评论]') INT"`
	Content             string    `xorm:"not null default '' comment('评论的内容') VARCHAR(255)"`
	StatusId            int       `xorm:"not null default 1 comment('状态 (1 有效，0 逻辑删除)') INT"`
	PraiseNum           int       `xorm:"not null default 0 comment('点赞数') INT"`
	TopStatus           int       `xorm:"not null default 0 comment('置顶状态[ 1 置顶，0 不置顶 默认 ]') TINYINT"`
	CreateTime          time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') index TIMESTAMP"`
}

type Cost struct {
	CostId      int64     `xorm:"not null pk autoincr BIGINT"`
	UserId      int64     `xorm:"not null index BIGINT"`
	CostContext string    `xorm:"not null comment('花销目录') VARCHAR(255)"`
	CostMoney   string    `xorm:"not null comment('花销金额') DECIMAL(10,2)"`
	Tag         string    `xorm:"comment('标签分类') VARCHAR(255)"`
	Created     time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	Updated     time.Time `xorm:"DATETIME"`
}

type DailyFiles struct {
	FileId   int64     `xorm:"not null pk autoincr BIGINT"`
	FilePath string    `xorm:"comment('磁盘路径') VARCHAR(255)"`
	FileUrl  string    `xorm:"comment('访问 路径') VARCHAR(500)"`
	FileSize int64     `xorm:"BIGINT"`
	UserId   int64     `xorm:"BIGINT"`
	Created  time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	Updeted  time.Time `xorm:"TIMESTAMP"`
}

type Diary struct {
	DiaryId      int64     `xorm:"not null pk autoincr comment('主键') BIGINT"`
	DiaryContext string    `xorm:"comment('日记内容:文本') LONGTEXT"`
	DiaryHtml    string    `xorm:"comment('日记内容:html') LONGTEXT"`
	DialyDate    time.Time `xorm:"comment('日记时间') DATETIME"`
	MoodType     int       `xorm:"INT"`
	Mood         string    `xorm:"comment('心情，原始的，实际去关联label表') VARCHAR(255)"`
	Temperature  string    `xorm:"comment('气温') DECIMAL(3,1)"`
	WeatherType  int       `xorm:"comment('心情') INT"`
	Weather      string    `xorm:"comment('天气') VARCHAR(255)"`
	StatusId     int       `xorm:"default 0 comment('日记状态，公开1，私有0，删除9') SMALLINT"`
	UserId       int64     `xorm:"index BIGINT"`
	Created      time.Time `xorm:"default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	Updated      time.Time `xorm:"comment('更新时间') DATETIME"`
}

type DjangoAdminLog struct {
	Id            int       `xorm:"not null pk autoincr INT"`
	ActionTime    time.Time `xorm:"not null DATETIME(6)"`
	ObjectId      string    `xorm:"LONGTEXT"`
	ObjectRepr    string    `xorm:"not null VARCHAR(200)"`
	ActionFlag    string    `xorm:"not null UNSIGNED SMALLINT"`
	ChangeMessage string    `xorm:"not null LONGTEXT"`
	ContentTypeId int       `xorm:"index INT"`
	UserId        int       `xorm:"not null index INT"`
}

type DjangoContentType struct {
	Id       int    `xorm:"not null pk autoincr INT"`
	AppLabel string `xorm:"not null unique(django_content_type_app_label_model_76bd3d3b_uniq) VARCHAR(100)"`
	Model    string `xorm:"not null unique(django_content_type_app_label_model_76bd3d3b_uniq) VARCHAR(100)"`
}

type DjangoMigrations struct {
	Id      int       `xorm:"not null pk autoincr INT"`
	App     string    `xorm:"not null VARCHAR(255)"`
	Name    string    `xorm:"not null VARCHAR(255)"`
	Applied time.Time `xorm:"not null DATETIME(6)"`
}

type DjangoSession struct {
	SessionKey  string    `xorm:"not null pk VARCHAR(40)"`
	SessionData string    `xorm:"not null LONGTEXT"`
	ExpireDate  time.Time `xorm:"not null index DATETIME(6)"`
}

type Event struct {
	EventId      int64     `xorm:"not null pk autoincr BIGINT"`
	EventContext string    `xorm:"not null comment('事件内容') TEXT"`
	TimeCost     string    `xorm:"comment('消耗时间，分钟') DECIMAL(10,2)"`
	UserId       int64     `xorm:"not null index BIGINT"`
	PlanTime     time.Time `xorm:"comment('事件实际发生时间，可当作计划事件') TIMESTAMP"`
	Remind       int       `xorm:"default b'0' comment('是否需要提醒') BIT(1)"`
	RemindTime   time.Time `xorm:"comment('提醒时间') TIMESTAMP"`
	Completed    int       `xorm:"comment('是否完成') BIT(1)"`
	Emergency    int       `xorm:"default 1 comment('紧急程度1，2，3，4升') INT"`
	Importance   int       `xorm:"default 1 comment('重要程度1，2，3，4升') INT"`
	Deleted      int       `xorm:"BIT(1)"`
	StatusId     int       `xorm:"default 1 comment('状态，9删除') INT"`
	Created      time.Time `xorm:"default CURRENT_TIMESTAMP TIMESTAMP"`
	Updated      time.Time `xorm:"TIMESTAMP"`
	TipEmail     string    `xorm:"comment('提醒邮箱') VARCHAR(255)"`
}

type Label struct {
	LabelId       int64     `xorm:"not null pk autoincr BIGINT"`
	LabelContext  string    `xorm:"not null comment('标签内容') VARCHAR(255)"`
	LabelType     int       `xorm:"comment('标签类型,1心情 ，2消费，3暂定') INT"`
	RelationId    int64     `xorm:"comment('关联表主键') BIGINT"`
	RelationTable int       `xorm:"default 1 comment('关系表，1日记表，2消费表，3学习表，4游记') INT"`
	UserId        int64     `xorm:"BIGINT"`
	Created       time.Time `xorm:"comment('创建时间') TIMESTAMP"`
}

type Learn struct {
	LearnId      int64     `xorm:"not null pk autoincr BIGINT"`
	LearnTitle   string    `xorm:"not null comment('学习标题') VARCHAR(255)"`
	LearnHtml    string    `xorm:"LONGTEXT"`
	LearnContext string    `xorm:"not null comment('学习内容') LONGTEXT"`
	Dir          int       `xorm:"default 1 comment('文件夹') INT"`
	Tag          string    `xorm:"comment('学习表情') VARCHAR(255)"`
	StatusId     int       `xorm:"INT"`
	UserId       int64     `xorm:"BIGINT"`
	Updated      time.Time `xorm:"TIMESTAMP"`
	Created      time.Time `xorm:"default CURRENT_TIMESTAMP TIMESTAMP"`
}

type SpringScheduledTask struct {
	CronId         int    `xorm:"not null pk autoincr comment('主键id') INT"`
	CronKey        string `xorm:"not null comment('定时任务完整类名') unique unique VARCHAR(128)"`
	CronExpression string `xorm:"not null comment('cron表达式') VARCHAR(20)"`
	TaskExplain    string `xorm:"not null default '' comment('任务描述') VARCHAR(50)"`
	StatusId       int    `xorm:"not null default 1 comment('状态,1:正常;2:停用') INT"`
}

type TravelNote struct {
	TravelId     int64     `xorm:"not null pk autoincr BIGINT"`
	TravelTitle  string    `xorm:"comment('标题') VARCHAR(255)"`
	TravelImg    int64     `xorm:"BIGINT"`
	TravelHtml   string    `xorm:"LONGTEXT"`
	TravelNote   string    `xorm:"not null LONGTEXT"`
	UserId       int64     `xorm:"not null BIGINT"`
	ProvinceCode int64     `xorm:"BIGINT"`
	CityCode     int64     `xorm:"BIGINT"`
	Longitude    string    `xorm:"comment('经度') DECIMAL(20,7)"`
	Latitude     string    `xorm:"comment('纬度') DECIMAL(20,7)"`
	Tag          string    `xorm:"VARCHAR(255)"`
	Updated      time.Time `xorm:"TIMESTAMP"`
	Created      time.Time `xorm:"TIMESTAMP"`
}

type Uid struct {
	Uid int `xorm:"not null pk INT"`
}

type UserInfo struct {
	UserId   int64     `xorm:"not null pk autoincr BIGINT"`
	UserName string    `xorm:"VARCHAR(255)"`
	Phone    string    `xorm:"VARCHAR(255)"`
	Email    string    `xorm:"unique VARCHAR(255)"`
	Avatar   string    `xorm:"comment('头像') VARCHAR(255)"`
	NickName string    `xorm:"comment('昵称') VARCHAR(255)"`
	DailyPwd string    `xorm:"comment('日记密码') VARCHAR(255)"`
	Pwd      string    `xorm:"comment('登录密码') VARCHAR(255)"`
	Salt     string    `xorm:"VARCHAR(255)"`
	StatusId int       `xorm:"default 1 comment('用户状态：1可用，0不可用') SMALLINT"`
	Created  time.Time `xorm:"DATETIME"`
}
