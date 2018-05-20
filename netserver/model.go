package netserver

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type User struct {
	Id        int       `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Nick      string    `json:"nick"`
	Password  *string   `json:"password"`
	Tel       *string   `json:"tel"`
	Email     *string   `json:"email"`
	Sex       *string   `json:"sex"`
	Birth     *string   `json:"birth"`
	AgreeMkt  int8      `json:"agree_mkt"`
	Image     *string   `json:"image"`
	Sns       *string   `json:"sns"`
	SnsId     *string   `json:"sns_id"`
	Job       *string   `json:"job"`
	City      *string   `json:"city"`
	House     *string   `json:"house"`
	Family    *string   `json:"family"`
	SignCh    *string   `json:"sign_ch"`
	SignPp    *string   `json:"sign_pp"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_dt"`
}

type LoginLog struct {
	Id      int       `json:"id" gorm:"primary_key"`
	UserId  int       `json:"user_id"`
	LoginAt time.Time `json:"login_at"`
}

type Role struct {
	Id       int    `json:"id" gorm:"primary_key"`
	UserId   int    `json:"user_id"`
	RoleName string `json:"role_name"`
}

type Address struct {
	Id      int    `json:"id" gorm:"primary_key"`
	UserId  int    `json:"user_id"`
	Name    string `json:"name"`
	Line1   string `json:"line1"`
	Line2   string `json:"line2"`
	Zipcode string `json:"zipcode"`
}

type Card struct {
	Id        int       `json:"id" gorm:"primary_key"`
	UserId    int       `json:"user_id"`
	CCName    string    `json:"cc_name" gorm:"column:cc_name"`
	CCPin     string    `json:"cc_pin" gorm:"column:cc_pin"`
	CreatedAt time.Time `json:"created_at`
	DeletedAt time.Time `json:"deleted_at`
}

type Account struct {
	UserId        int    `json:"user_id" gorm:"primary_key"`
	AccountName   string `json:"account_name"`
	AccountNum    string `json:"account_num"`
	AccountHolder string `json:"account_holder"`
}

type Version struct {
	Id        int       `json:"id"`
	PK        int       `json:"pk"`
	Tab       string    `json:"tab"`
	Raw       string    `json:"raw"`
	UpdaterId int       `json:"updater_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Order struct {
	OrderId      int     `json:"order_id" gorm:"primary_key"`
	UserId       int     `json:"user_id"`
	UserName     string  `json:"user_name"`
	UserTel      string  `json:"user_tel"`
	UserMemo     *string `json:"user_memo"`
	Version      int     `json:"version" gorm:"primary_key"`
	Subscription bool    `json:"subscription" gorm:"-"`
}

type History struct {
	Id          int       `json:"id" gorm:"primary_key"`
	OrderId     int       `json:"order_id"`
	UpdaterId   int       `json:"user_id"`
	HistoryDate time.Time `json:"history_date"`
}

type Subscription struct {
	Id           int       `json:"id" gorm:"primary_key"`
	OrderId      int       `json:"order_d"`
	CCPin        string    `json:"cc_pin" gorm:"column:cc_pin"`
	ScheduleDate *string   `json:"schedule_date"`
	CreatedAt    time.Time `json:"created_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}

type Payment struct {
	Id          int        `json:"id" gorm:"primary_key"`
	RelId       int        `json:"rel_id"`
	OrderId     int        `json:"order_id"`
	OnceToken   string     `json:"once_token"`
	Method      string     `json:"method"`
	State       string     `json:"state"`
	CCPin       string     `json:"cc_pin" gorm:"column:cc_pin"`
	Coupon      string     `json:"coupon"`
	Bc          string     `json:"bc"`
	Amount      int        `json:"amount"`
	Discount    int        `json:"discount"`
	Mons        int        `json:"mons"`
	ReceiptUrl  string     `json:"receipt_url"`
	ReceiptSn   string     `json:"receipt_sn"`
	ErrMsg      string     `json:"err_msg"`
	CustomMsg   string     `json:"custom_msg"`
	CreatedAt   time.Time  `json:"created_at"`
	scheduledAt *KoTime    `json:"scheduled_at"`
	PaidAt      *time.Time `json:"paid_at"`
	EventId     int        `json:"event_id"`
}

type Inventory struct {
	Id               int     `json:"id" gorm:"primary_key"`
	BoxId            int     `json:"box_id"`
	DispNum          int     `json:"disp_num"`
	SkuId            int     `json:"sku_id"`
	LastSid          int     `json:"last_sid"`
	OrderId          int     `json:"order_id"`
	InventoryName    *string `json:"inventory_name"`
	InventoryCaption *string `json:"inventory_caption"`
}

type InventoryState struct {
	Id          int       `json:"id" gorm:"primary_key"`
	InventoryId int       `json:"inventory_id"`
	EventId     int       `json:"event_id"`
	UpdaterId   int       `json:"updater_id"`
	HistoryDate time.Time `json:"history_date"`
}

type Category struct {
	SkuId     int    `json:"sku_id"`
	CatKey    string `json:"cat_key"`
	CatName   string `json:"cat_name"`
	Amount    int    `json:"amount"`
	BuyUnit   string `json:"buyunit"`
	Min       int    `json:"min"`
	PriceMark int    `json:"pricemark"`
	Image     string `json:"image"`
	Ximage    string `json:"ximage"`
	Comment   string `json:"comment"`
}

type AdminComment struct {
	Id        int    `json:"id" gorm:"primary_key"`
	OrderId   int    `json:"orderId"`
	UserId    int    `json:"userId"`
	Nick      string `json:"nick"`
	Comment   string `json:"comment"`
	CreatedAt string `json:"createdAt"`
}

type Ask struct {
	Id         int       `json:"id" gorm:"primary_key"`
	ThreadId   int       `json:"thread_id"`
	UserId     int       `json:"user_id"`
	PosterName *string   `json:"poster_name"`
	PosterTel  *string   `json:"poster_tel"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Complete   bool      `json:"complete"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}

type Coupon struct {
	Id            int    `json:"id" gorm:"primary_key"`
	Code          string `json:"code"`
	CouponCaption string `json:"caption"`
	Amount        int    `json:"amount"`
	SkuAllow      int    `json:"sku_allow"`
	CreatedAt     string `json:"created_at"`
	ExpiredAt     string `json:"expired_at"`
}

type Feed struct {
	Id        int       `json:"id" gorm:"primary_key"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Visible   string    `json:"visible" gorm:"default:'n'"`
	Url       string    `json:"url"`
	Popup     string    `json:"popup" gorm:"default:'n'"`
	Agent     string    `json:"agent"`
	FeedPos   string    `json:"feed_pos"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

type Attachment struct {
	Id         int       `json:"id" gorm:"primary_key"`
	RefId      int       `json:"ref_id"`
	RepoType   string    `json:"repo_type"`
	Filename   string    `json:"filename" gorm:"column:filename"`
	FileSize   int       `json:"filenize" gorm:"column:filesize"`
	MimeType   string    `json:"mimetype" gorm:"column:minetype"`
	UploadedAt time.Time `json:"uploaded_at"`
}

type Pagination struct {
	Count int         `json:"count"`
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
	Rows  interface{} `json:"rows"`
}

type OrderFilterRow struct {
	Page        int           `query:"page"  json:"omit" gorm:"-"`
	Limit       int           `query:"limit" json:"omit" gorm:"-"`
	OrderId     int           `query:"orderId"  json:"orderId" gorm:"primary_key"`
	OrderName   string        `json:"order_name"`
	OrderType   *string       `query:"orderType" json:"orderType"`
	UserId      int           `query:"userId"   json:"userId"`
	UserName    *string       `query:"userName" json:"userName"`
	UserTel     *string       `query:"userTel"  json:"userTel"`
	UserMemo    *string       `json:"userMemo"`
	OrderedAt   *string       `query:"orderedAt"  json:"orderedAt"`
	CancelledAt *string       `query:"cancelledAt" json:"cancelledAt"`
	ToName      *string       `json:"toName"`
	Tel         *string       `query:"tel" json:"tel"`
	Address     *string       `json:"address"`
	DateAt      *KoTime       `json:"dateAt"`
	DateTime    *int          `json:"dateTime"`
	SubscbId    *int          `json:"subscbId"`
	SubscbDate  *string       `query:"subscbDate" json:"subscbDate"`
	SubscbDay   *int          `query:"subscbDay" json:"subscbDay"`
	OrderState  *string       `json:"orderState"`
	CancelState *string       `json:"cancelState"`
	Comment     *AdminComment `json:"adminComment" gorm:"column:admin_comment"`
}

func (OrderFilterRow) TableName() string {
	return "T_ORDER"
}

//Custom struct marshal, unmarshal
const KO_DATE_FORMAT = "2006/01/02 15:04"

type KoTime struct {
	time.Time
}

func (k *KoTime) UnmarshalJSON(b []byte) error {
	input := string(b)
	input = strings.Trim(input, `"`)

	t, err := time.Parse(KO_DATE_FORMAT, input)

	if err != nil {
		return err
	}

	k.Time = t
	return nil
}

func (k KoTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + k.Time.Format(KO_DATE_FORMAT) + `"`), nil
}

func (k *KoTime) Scan(value interface{}) error {

	var err error

	switch x := value.(type) {
	case time.Time:
		k.Time = x
	case nil:
		return nil
	default:
		err = fmt.Errorf("cannot scan type %T into %v", value, value)
	}

	return err

}

func (k *KoTime) Value() (driver.Value, error) {
	if k == nil {
		return nil, nil
	}

	return k.Time, nil

}

//TODO htmlescape적용할 것!
func (s *AdminComment) Scan(value interface{}) error {
	str := string(value.([]byte))
	str = strings.Replace(str, "\n", "<br/>", -1)
	str = strings.Replace(str, "\t", " ", -1)

	if err := json.Unmarshal([]byte(str), &s); err != nil {
		return err
	}
	return nil
}
