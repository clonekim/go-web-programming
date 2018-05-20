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
	AgreeMkt  int8      `json:"agreeMkt"`
	Image     *string   `json:"image"`
	Sns       *string   `json:"sns"`
	SnsId     *string   `json:"snsId"`
	Job       *string   `json:"job"`
	City      *string   `json:"city"`
	House     *string   `json:"house"`
	Family    *string   `json:"family"`
	SignCH    *string   `json:"signCH"`
	SignPP    *string   `json:"signPP"`
	CreatedAt time.Time `json:"createdAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

type LoginLog struct {
	Id      int       `json:"id" gorm:"primary_key"`
	UserId  int       `json:"userId"`
	LoginAt time.Time `json:"loginAt"`
}

type Role struct {
	Id       int    `json:"id" gorm:"primary_key"`
	UserId   int    `json:"userId"`
	RoleName string `json:"roleName"`
}

type Address struct {
	Id      int    `json:"id" gorm:"primary_key"`
	UserId  int    `json:"userId"`
	Name    string `json:"name"`
	Line1   string `json:"line1"`
	Line2   string `json:"line2"`
	Zipcode string `json:"zipcode"`
}

type Card struct {
	Id        int       `json:"id" gorm:"primary_key"`
	UserId    int       `json:"userId"`
	CCName    string    `json:"ccName" gorm:"column:cc_name"`
	CCPin     string    `json:"ccPin" gorm:"column:cc_pin"`
	CreatedAt time.Time `json:"createdAt`
	DeletedAt time.Time `json:"deletedAt`
}

type Account struct {
	UserId        int    `json:"userId" gorm:"primary_key"`
	AccountName   string `json:"accountName"`
	AccountNum    string `json:"accountNum"`
	AccountHolder string `json:"accountHolder"`
}
type Version struct {
	Id        int       `json:"id"`
	PK        int       `json:"pk"`
	Tab       string    `json:"tab"`
	Raw       string    `json:"raw"`
	UpdaterId int       `json:"updaterId"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Order struct {
	OrderId      int     `json:"orderId" gorm:"primary_key"`
	UserId       int     `json:"userId"`
	UserName     string  `json:"userName"`
	UserTel      string  `json:"userTel"`
	UserMemo     *string `json:"userMemo"`
	Version      int     `json:"version" gorm:"primary_key"`
	Subscription bool    `json:"subscription" gorm:"-"`
}

type History struct {
	Id          int       `json:"id" gorm:"primary_key"`
	OrderId     int       `json:"orderId"`
	UpdaterId   int       `json:"userId"`
	HistoryDate time.Time `json:"historyDate"`
}

type Subscription struct {
	Id           int       `json:"id" gorm:"primary_key"`
	OrderId      int       `json:"orderId"`
	CCPin        string    `json:"ccPin"`
	ScheduleDate *string   `json:"scheduleDate"`
	CreatedAt    time.Time `json:"createdAt"`
	DeletedAt    time.Time `json:"deletedAt"`
}

type Payment struct {
	Id          int        `json:"id" gorm:"primary_key"`
	RelId       int        `json:"relId"`
	OrderId     int        `json:"orderId"`
	OnceToken   string     `json:"onceToken"`
	Method      string     `json:"method"`
	State       string     `json:"state"`
	CCPin       string     `json:"ccPin" gorm:"column:cc_pin"`
	Coupon      string     `json:"coupon"`
	Bc          string     `json:"bc"`
	Amount      int        `json:"amount"`
	Discount    int        `json:"discount"`
	Mons        int        `json:"mons"`
	ReceiptUrl  string     `json:"receiptUrl"`
	ReceiptSn   string     `json:"receiptSn"`
	ErrMsg      string     `json:"errMsg"`
	CustomMsg   string     `json:"customMsg"`
	CreatedAt   time.Time  `json:"createdAt"`
	scheduledAt *KoTime    `json:"scheduledAt"`
	PaidAt      *time.Time `json:"paidAt"`
	EventId     int        `json:"eventId"`
}

type Inventory struct {
	Id               int     `json:"id" gorm:"primary_key"`
	BoxId            int     `json:"boxId"`
	DispNum          int     `json:"dispNum"`
	SkuId            int     `json:"skuId"`
	LastSid          int     `json:"lastSid"`
	OrderId          int     `json:"orderId"`
	InventoryName    *string `json:"inventoryName"`
	InventoryCaption *string `json:"inventoryCaption"`
}

type InventoryState struct {
	Id          int       `json:"id" gorm:"primary_key"`
	InventoryId int       `json:"inventoryId"`
	EventId     int       `json:"eventId"`
	UpdaterId   int       `json:"updaterId"`
	HistoryDate time.Time `json:"historyDate"`
}

type Category struct {
	SkuId     int    `json:"skuId"`
	CatKey    string `json:"catKey"`
	CatName   string `json:"catName"`
	Amount    int    `json:"amount"`
	BuyUnit   string `json:"buyUnit"`
	Min       int    `json:"min"`
	PriceMark int    `json:"priceMark"`
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
	ThreadId   int       `json:"threadId"`
	UserId     int       `json:"userId"`
	PosterName *string   `json:"posterName"`
	PosterTel  *string   `json:"posterTel"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Complete   bool      `json:"complete"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	DeletedAt  time.Time `json:"deletedAt"`
}

type Coupon struct {
	Id            int    `json:"id" gorm:"primary_key"`
	Code          string `json:"code"`
	CouponCaption string `json:"caption"`
	Amount        int    `json:"amount"`
	SkuAllow      int    `json:"skuAllow"`
	CreatedAt     string `json:"createdAt"`
	ExpiredAt     string `json:"expiredAt"`
}

type Feed struct {
	Id        int       `json:"id" gorm:"primary_key"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Visible   string    `json:"visible" gorm:"default:'n'"`
	Url       string    `json:"url"`
	Popup     string    `json:"popup" gorm:"default:'n'"`
	Agent     string    `json:"agent"`
	FeedPos   string    `json:"feedPos"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	ExpiredAt time.Time `json:"expiredAt"`
}

type Attachment struct {
	Id         int       `json:"id" gorm:"primary_key"`
	RefId      int       `json:"refId"`
	RepoType   string    `json:"repoType"`
	Filename   string    `json:"fileName"`
	FileSize   int       `json:"fileSize"`
	MimeType   string    `json:"mimeType"`
	UploadedAt time.Time `json:"uploadedAt"`
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
	OrderName   string        `json:"orderName"`
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

func (s *AdminComment) Scan(value interface{}) error {
	str := string(value.([]byte))
	str = strings.Replace(str, "\n", "<br/>", -1)
	str = strings.Replace(str, "\t", " ", -1)

	if err := json.Unmarshal([]byte(str), &s); err != nil {
		return err
	}
	return nil
}
