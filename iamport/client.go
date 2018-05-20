package iamport

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/resty.v1"
	"strconv"
	"time"
)

type Client struct {
	Uri       string
	Mid       string
	Secret    string
	ImpKey    string
	ImpSecret string
	Token     string
	Http      *resty.Client
}

type AckResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CancelHistory struct {
	PgTid       string `json:"pg_tid"`
	Amount      int64  `json:"amount"`
	CancelledAt int64  `json:"cancelled_at"`
	ReceiptUrl  string `json:"receipt_url"`
}

type Pagination struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Response struct {
		Total    int                  `json:"total"`
		Previous int                  `json:"previous"`
		Next     int                  `json:"next"`
		List     []PaymentAckResponse `json:"list"`
	} `json:"response"`
}

type PaymentAckResponse struct {
	Amount        int64           `json:"amount"`
	BuyerName     string          `json:"buyer_name"`
	BuyerEmail    string          `json:"buyer_email"`
	BuyerTel      string          `json:"buyer_tel"`
	BuyerAddress  string          `json:"buyer_addr"`
	BuyerPostCode string          `json:"buyer_postcode"`
	CancelAmount  int64           `json:"cancel_amount"`
	CancelHistory []CancelHistory `json:"cancel_history"`
	CancelledAt   int64           `json:"cancelled_at"`
	CardName      string          `json:"card_name"`
	FailError     string          `json:"fail_reason"`
	FailedAt      int64           `json:"failed_at"`
	ImpUid        string          `json:"imp_uid"`
	TLToken       string          `json:"merchant_uid"`
	Name          string          `json:"name"`
	PaidAt        int64           `json:"paid_at"`
	Method        string          `json:"pay_method"`
	PgTid         string          `json:"pg_tid"`
	ReceiptUrl    string          `json:"receipt_url"`
	State         string          `json:"status"`
	BC            string          `json:"vbank_code"`
	BCNum         string          `json:"vbank_num"`
	BCDate        int64           `json:"vbank_date"`
	BCName        string          `json:"vbank_name"`
}

type HistoryAckResponse struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Response struct {
		Amount        int64           `json:"amount"`
		BuyerName     string          `json:"buyer_name"`
		BuyerEmail    string          `json:"buyer_email"`
		BuyerTel      string          `json:"buyer_tel"`
		BuyerAddress  string          `json:"buyer_addr"`
		BuyerPostCode string          `json:"buyer_postcode"`
		CancelAmount  int64           `json:"cancel_amount"`
		CancelHistory []CancelHistory `json:"cancel_history"`
		CancelledAt   int64           `json:"cancelled_at"`
		CardName      string          `json:"card_name"`
		FailError     string          `json:"fail_reason"`
		FailedAt      int64           `json:"failed_at"`
		ImpUid        string          `json:"imp_uid"`
		TLToken       string          `json:"merchant_uid"`
		Name          string          `json:"name"`
		PaidAt        int64           `json:"paid_at"`
		Method        string          `json:"pay_method"`
		PgTid         string          `json:"pg_tid"`
		ReceiptUrl    string          `json:"receipt_url"`
		State         string          `json:"status"`
		BC            string          `json:"vbank_code"`
		BCNum         string          `json:"vbank_num"`
		BCDate        int64           `json:"vbank_date"`
		BCName        string          `json:"vbank_name"`
	} `json:"response"`
}

func NewHttpClient(debug bool) *resty.Client {
	http := resty.New()
	http.SetDebug(debug)
	return http
}

func (c *Client) Authenticate() error {
	data := struct {
		Code     int    `json:"code"`
		Message  string `json:"message"`
		Response struct {
			AccessToken string `json:"access_token"`
			ExpiredAt   int64  `json:"expired_at"`
			Now         int64  `json:"now"`
		} `json:"response"`
	}{}

	_, err := c.Http.R().
		SetFormData(map[string]string{
			"imp_key":    c.ImpKey,
			"imp_secret": c.ImpSecret,
		}).
		SetResult(&data).
		Post(fmt.Sprintf("%s/users/getToken", c.Uri))

	if err != nil {
		return err
	}

	c.Token = data.Response.AccessToken

	return nil
}

func (c *Client) MakeCCPin(ccpin, number, expiry, birth, pwd2digit, name, tel string) error {
	data := AckResponse{}

	resp, err := c.Http.R().
		SetHeader("Authorization", c.Token).
		SetFormData(map[string]string{
			"customer_uid": ccpin,
			"card_number":  number,
			"expiry":       expiry,
			"birth":        birth,
			"pwd_2digit": func() string {
				if len(birth) == 10 {
					return ""
				}
				return pwd2digit
			}(),
			"customer_name": name,
			"customer_tel":  tel,
		}).
		Post(fmt.Sprintf("%s/subscribe/customers/%s", c.Uri, ccpin))

	if err != nil {
		return err
	}

	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		return err
	}

	if data.Code != 0 {
		return errors.New(data.Message)
	}

	return nil
}

func (c *Client) RevokeCCPin(ccpin string) error {

	data := AckResponse{}

	resp, err := c.Http.R().
		SetHeader("Authorization", c.Token).
		Delete(fmt.Sprintf("%s/subscribe/customers/%s", c.Uri, ccpin))

	if err != nil {
		return err
	}

	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		return err
	}

	if data.Code != 0 {
		return errors.New(data.Message)
	}

	return nil

}

func (c *Client) Checkout(ccpin string, tlToken string, name string, amount int64) error {
	t := time.Now()
	t.Add(1 * time.Minute)

	return c.Schedule(ccpin, tlToken, name, amount, t)

}

func (c *Client) Schedule(ccpin string, tlToken string, name string, amount int64, scheduleAt time.Time) error {

	data := AckResponse{}

	resp, err := c.Http.R().
		SetHeader("Authorization", c.Token).
		SetBody(map[string]interface{}{
			"customer_uid": ccpin,
			"schedules": []map[string]interface{}{
				map[string]interface{}{
					"merchant_uid": tlToken,
					"amount":       amount,
					"name":         name,
					"schedule_at":  scheduleAt.Unix(),
				}},
		}).
		Post(fmt.Sprintf("%s/subscribe/payments/schedule", c.Uri))

	if err != nil {
		return err
	}

	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		return err
	}

	if data.Code != 0 {
		return errors.New(data.Message)
	}

	return nil
}

func (c *Client) UnSchedule(ccpin string, tlToken string) error {

	data := AckResponse{}

	resp, err := c.Http.R().
		SetHeader("Authorization", c.Token).
		SetFormData(map[string]string{
			"customer_uid": ccpin,
			"merchant_uid": tlToken,
		}).
		Post(fmt.Sprintf("%s/subscribe/payments/unschedule", c.Uri))

	if err != nil {
		return err
	}

	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		return err
	}

	if data.Code != 0 {
		return errors.New(data.Message)
	}

	return nil
}

func (c *Client) Cancel(tlToken string, amount int64) error {
	data := AckResponse{}

	body := map[string]string{}
	body["merchant_uid"] = tlToken

	if amount > 0 {
		body["amount"] = strconv.FormatInt(amount, 10)
	}

	resp, err := c.Http.R().
		SetHeader("Authorization", c.Token).
		SetFormData(body).
		Post(fmt.Sprintf("%s/payments/cancel", c.Uri))

	if err != nil {
		return err
	}

	if resp.StatusCode() == 400 {
		return errors.New("404 error")
	}

	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		return err
	}

	if data.Code != 0 {
		return errors.New(data.Message)
	}

	return nil

}

func (c *Client) Confirm(ccpin string) error {
	data := AckResponse{}

	resp, err := c.Http.R().
		SetHeader("Authorization", c.Token).
		Get(fmt.Sprintf("%s/payments/%s", c.Uri, ccpin))

	if err != nil {
		return err
	}

	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		return err
	}

	if data.Code != 0 {
		return errors.New(data.Message)
	}

	return nil
}

func (c *Client) ConfirmAll(tlToken string) (*Pagination, error) {
	data := Pagination{}

	resp, err := c.Http.R().
		SetHeader("Authorization", c.Token).
		Get(fmt.Sprintf("%s/payments/findAll/%s", c.Uri, tlToken))

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		return nil, err
	}

	if data.Code != 0 {
		return nil, errors.New(data.Message)
	}

	return &data, nil
}

func (c *Client) History(ccpin string) (*Pagination, error) {
	data := Pagination{}

	resp, err := c.Http.R().
		SetHeader("Authorization", c.Token).
		Get(fmt.Sprintf("%s/subscribe/customers/%s/payments", c.Uri, ccpin))

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		return nil, err
	}

	if data.Code != 0 {
		return nil, errors.New(data.Message)
	}

	return &data, nil
}
