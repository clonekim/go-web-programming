package iamport

import (
	"fmt"
	"testing"
	//"time"
)

var client *Client

func TestA(t *testing.T) {

	client = &Client{
		Uri:       "https://api.iamport.kr",
		Mid:       "nictest04m",
		Secret:    "b+zhZ4yOZ7FsH8pm5lhDfHZEb79tIwnjsdA0FBXh86yLc6BJeFVrZFXhAoJ3gEWgrWwN+lJMV0W4hvDdbe4Sjw==",
		ImpKey:    "9845496902780316",
		ImpSecret: "2yfGHeyD5XfAWTz19mPfH0TszVvo3yE67UNO8lseiHPfAphKL7Zt7TWHOr6S0LY7FMWO7o3q5TFGiXq3",
		Http:      NewHttpClient(false),
	}

	err := client.Authenticate()

	if err != nil {
		t.Error(err.Error())
		return
	}

	fmt.Println("Token:", client.Token)
}

func TestHistory(t *testing.T) {
	fmt.Println("History")
	list, err := client.History("daehee-6148")

	if err != nil {
		t.Error(err.Error())
		return
	}

	fmt.Println(list)

}

func TestConfirm(t *testing.T) {
	fmt.Println("Confirm")
	list, err := client.ConfirmAll("daehee-23456789")

	if err != nil {
		t.Error(err.Error())
		return
	}

	fmt.Println(list)
}

/*
func TestF(t *testing.T) {
	err := client.Cancel("daehee-12345678", 1000)

	if err != nil {
		t.Error(err.Error())
		return
	}

	fmt.Println("Cancel OK")
}

/*
func TestB(t *testing.T) {
	err := client.InstallBilingKey("daehee-6148", "6258-1797-0050-6148", "2022-02", "770402", "13", "테스트1", "01027293094")

	if err != nil {
		t.Error(err.Error())
		return
	}

	fmt.Println("biling key ok")
}


func TestC(t *testing.T) {

	//	t2, _ := time.Parse("2006-01-02 15:04:05 MST", "2018-04-26 11:00:00 KST")

	err := client.Checkout(
		"daehee-6148",
		"daehee-23456789",
		"TEST01",
		10000)

	if err != nil {
		t.Error(err.Error())
		return
	}

	fmt.Println("CheckOut OK")

}


func TestD(t *testing.T) {
	err := client.Cancel("daehee-23456789", 1000)

	if err != nil {
		t.Error(err.Error())
		return
	}

	fmt.Println("Cancel OK1")
}

func TestE(t *testing.T) {
	err := client.Cancel("daehee-23456789", 1000)

	if err != nil {
		t.Error(err.Error())
		return
	}

	fmt.Println("Cancel OK2")
}

/*
func TestD(t *testing.T) {

	err := client.UnSchedule("daehee-6148", "daehee-123456")

	if err != nil {
		t.Error(err.Error())
		return
	}

	fmt.Println("UnSchedule OK")

}
*/

/*

func TestE(t *testing.T) {
	var err error
	err = client.UninstallBilingKey("daehee-6148")
	if err != nil {
		t.Error(err.Error())
		return
	}

	fmt.Println("Uninstall OK 1")

	err = client.UninstallBilingKey("daehee_928671xY")
	if err != nil {
		t.Error(err.Error())
		return
	}

	fmt.Println("Uninstall OK 2")
}
*/
