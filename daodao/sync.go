// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package daodao

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"

	"github.com/wuhan005/daodao-sdk/internal/request"
)

type SyncV2Response struct {
	Data struct {
		Account []struct {
			Lists []struct {
				AccountType int    `json:"account_type"`
				BankName    string `json:"bank_name"`
				BillDay     int    `json:"bill_day"`
				BinlogId    int64  `json:"binlog_id"`
				Content     string `json:"content"`
				CreditLimit string `json:"credit_limit"`
				Ctime       int    `json:"ctime"`
				Currency    string `json:"currency"`
				DeviceKey   string `json:"device_key"`
				Dtime       int    `json:"dtime"`
				FromColor   string `json:"from_color"`
				ImgId       int    `json:"img_id"`
				Money       string `json:"money"`
				Mtime       int    `json:"mtime"`
				Name        string `json:"name"`
				ReturnDay   int    `json:"return_day"`
				Sort        int    `json:"sort"`
				ToColor     string `json:"to_color"`
				UserId      int    `json:"user_id"`
				Uuid        string `json:"uuid"`
			} `json:"lists"`
			NextPage struct {
				LastBinlogId int64  `json:"last_binlog_id"`
				Over         bool   `json:"over"`
				Direction    string `json:"direction"`
			} `json:"next_page"`
			Sqltag string `json:"sqltag"`
			Table  string `json:"table"`
		} `json:"account"`
		Category []struct {
			Lists []struct {
				CateId   int    `json:"cate_id"`
				CateName string `json:"cate_name"`
				Content  string `json:"content"`
				Income   int    `json:"income"`
				Sort     int    `json:"sort"`
			} `json:"lists"`
			NextPage struct {
				LastBinlogId int    `json:"last_binlog_id"`
				Over         bool   `json:"over"`
				Direction    string `json:"direction"`
			} `json:"next_page"`
			Sqltag string `json:"sqltag"`
			Table  string `json:"table"`
		} `json:"category"`
		InteractionReplyRecord []struct {
			Lists []struct {
				BinlogId        int64  `json:"binlog_id"`
				ButtonType      string `json:"button_type"`
				ChapterId       int    `json:"chapter_id"`
				Content         string `json:"content"`
				ContentType     string `json:"content_type"`
				CreateTime      int64  `json:"create_time"`
				Ctime           int    `json:"ctime"`
				DeviceKey       string `json:"device_key"`
				Dtime           int    `json:"dtime"`
				EmojiId         int    `json:"emoji_id"`
				Mtime           int    `json:"mtime"`
				PowerConsume    int    `json:"power_consume,omitempty"`
				ReplyContentId  int    `json:"reply_content_id"`
				RoleId          int    `json:"role_id"`
				Sex             int    `json:"sex"`
				ShowPage        int    `json:"show_page"`
				StarAutokids    string `json:"star_autokids"`
				StarId          int    `json:"star_id"`
				Stime           int    `json:"stime,omitempty"`
				StoreId         int    `json:"store_id"`
				TargetId        int64  `json:"target_id"`
				TargetType      int    `json:"target_type,omitempty"`
				TheaterId       int    `json:"theater_id"`
				Type            int    `json:"type"`
				UserId          int    `json:"user_id"`
				UserStarAutokid int64  `json:"user_star_autokid"`
				Uuid            string `json:"uuid"`
				JsonInfo        string `json:"json_info,omitempty"`
			} `json:"lists"`
			NextPage struct {
				LastBinlogId int64  `json:"last_binlog_id"`
				Over         bool   `json:"over"`
				Direction    string `json:"direction"`
			} `json:"next_page"`
			Sqltag string `json:"sqltag"`
			Table  string `json:"table"`
		} `json:"interactionReplyRecord"`
		Record []struct {
			Lists []struct {
				AccountCurrency     string `json:"account_currency"`
				AccountExchangeRate string `json:"account_exchange_rate"`
				AccountId           string `json:"account_id"`
				AccountMoney        string `json:"account_money"`
				BinlogId            int64  `json:"binlog_id"`
				CategoryId          int    `json:"category_id"`
				Content             string `json:"content"`
				CreateTime          int64  `json:"create_time"`
				Ctime               int    `json:"ctime"`
				DeviceKey           string `json:"device_key"`
				Dtime               int    `json:"dtime"`
				Flow                int    `json:"flow"`
				Image               string `json:"image"`
				ImgHeight           string `json:"img_height"`
				ImgWidth            string `json:"img_width"`
				Income              int    `json:"income"`
				JsonInfo            string `json:"json_info"`
				Mtime               int    `json:"mtime"`
				NeedReply           int    `json:"need_reply"`
				Rate                string `json:"rate"`
				RateCurrency        string `json:"rate_currency"`
				RateMoney           string `json:"rate_money"`
				RecordCurrency      string `json:"record_currency"`
				RecordMoney         string `json:"record_money"`
				RecordType          string `json:"record_type"`
				RoleId              int    `json:"role_id"`
				Rtime               int    `json:"rtime"`
				ShowPage            int    `json:"show_page"`
				StarAutokids        string `json:"star_autokids"`
				StarId              int    `json:"star_id"`
				TransferRid         string `json:"transfer_rid"`
				UserId              int    `json:"user_id"`
				Uuid                string `json:"uuid"`
			} `json:"lists"`
			NextPage struct {
				LastBinlogId int64  `json:"last_binlog_id"`
				Over         bool   `json:"over"`
				Direction    string `json:"direction"`
			} `json:"next_page"`
			Sqltag string `json:"sqltag"`
			Table  string `json:"table"`
		} `json:"record"`
		RecordType []struct {
			Lists []struct {
				BinlogId   int64  `json:"binlog_id"`
				CategoryId int    `json:"category_id"`
				CommonSort int    `json:"common_sort"`
				Content    string `json:"content"`
				Ctime      int    `json:"ctime"`
				DeviceKey  string `json:"device_key"`
				Dtime      int    `json:"dtime"`
				ImgId      string `json:"img_id"`
				IsCommon   int    `json:"is_common"`
				Mtime      int    `json:"mtime"`
				Sort       int    `json:"sort"`
				UserId     int    `json:"user_id"`
				Uuid       string `json:"uuid"`
			} `json:"lists"`
			NextPage struct {
				LastBinlogId int64  `json:"last_binlog_id"`
				Over         bool   `json:"over"`
				Direction    string `json:"direction"`
			} `json:"next_page"`
			Sqltag string `json:"sqltag"`
			Table  string `json:"table"`
		} `json:"recordType"`
		ReplyRecord []struct {
			Lists []struct {
				AllowBlack      bool   `json:"allow_black"`
				AllowGuide      bool   `json:"allow_guide"`
				BatchId         string `json:"batch_id"`
				BinlogId        int64  `json:"binlog_id"`
				Blacklist       int    `json:"blacklist"`
				ChapterId       int    `json:"chapter_id"`
				ChatUuid        int64  `json:"chat_uuid"`
				ChooseId        int    `json:"choose_id,omitempty"`
				ClientType      int    `json:"client_type"`
				CreateTime      int64  `json:"create_time"`
				Ctime           int    `json:"ctime"`
				Dtime           int    `json:"dtime"`
				GtplId          string `json:"gtpl_id"`
				IsExclusive     bool   `json:"is_exclusive"`
				IsPicked        int    `json:"is_picked,omitempty"`
				Like            int    `json:"like"`
				Mtime           int    `json:"mtime"`
				NextChoose      string `json:"next_choose"`
				NoteId          string `json:"note_id,omitempty"`
				ReadAt          int    `json:"read_at"`
				RecordId        string `json:"record_id"`
				ReplyContentId  int    `json:"reply_content_id"`
				ReplyPlusId     int    `json:"reply_plus_id"`
				ReplyRuleId     int    `json:"reply_rule_id,omitempty"`
				ReplyStarId     int    `json:"reply_star_id"`
				RoleId          int    `json:"role_id"`
				ShowDate        string `json:"show_date,omitempty"`
				ShowPage        int    `json:"show_page"`
				StoreId         int    `json:"store_id,omitempty"`
				TargetType      int    `json:"target_type"`
				TheaterId       int    `json:"theater_id"`
				Type            string `json:"type"`
				U               int    `json:"u"`
				UserStarAutokid int64  `json:"user_star_autokid"`
				Uuid            string `json:"uuid"`
				Value           string `json:"value"`
				WordId          int    `json:"word_id,omitempty"`
				JsonInfo        string `json:"json_info,omitempty"`
			} `json:"lists"`
			NextPage struct {
				LastBinlogId int64  `json:"last_binlog_id"`
				Over         bool   `json:"over"`
				Direction    string `json:"direction"`
			} `json:"next_page"`
			Sqltag string `json:"sqltag"`
			Table  string `json:"table"`
		} `json:"replyRecord"`
	} `json:"data"`
	Encrypt    int    `json:"encrypt"`
	Message    string `json:"message"`
	Name       string `json:"name"`
	ServerTime int    `json:"server_time"`
	Success    bool   `json:"success"`
}

type block struct {
	SqlTag string      `json:"sqltag"`
	Values blockValues `json:"values"`
	Page   blockPage   `json:"page"`
}

type blockValues struct {
	LastBinlogId string `json:"last_binlog_id"`
	UserId       string `json:"userId"`
}

type blockPage struct {
	LastBinlogId string `json:"last_binlog_id"`
	Direction    string `json:"direction"`
	Over         bool   `json:"over"`
}

type SyncV2Block string

var (
	InteractionReplyRecordPage  SyncV2Block = "interactionReplyRecord_page"
	InteractionReplyRecordPageM SyncV2Block = "interactionReplyRecord_page_m"
	ReplyRecordPage             SyncV2Block = "reply_record_page"
	ReplyRecordPageM            SyncV2Block = "reply_record_page_m"
	RecordType                  SyncV2Block = "recordType"
	Category                    SyncV2Block = "category"
	RecordPage                  SyncV2Block = "record_page"
	Account                     SyncV2Block = "account_20190918"
)

func (c *Client) SyncV2(syncV2Blocks ...SyncV2Block) (*SyncV2Response, error) {
	blocks := make([]block, 0, 8)
	for _, b := range syncV2Blocks {
		blocks = append(blocks, block{
			SqlTag: string(b),
			Values: blockValues{LastBinlogId: "-1", UserId: c.userID},
			Page:   blockPage{LastBinlogId: "-1"},
		})
	}

	queryJSON, err := json.Marshal(blocks)
	if err != nil {
		return nil, errors.Wrap(err, "json encode")
	}

	query := c.DefaultParameters()
	query.Set("is_new_user", "1")
	query.Set("json", string(queryJSON))
	resp, err := request.Get("/sync/v2").Query(query).Do()
	if err != nil {
		return nil, errors.Wrap(err, "request")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("unexpected response code: %d", resp.StatusCode)
	}

	var syncV2Response SyncV2Response
	if err := resp.ToJSON(&syncV2Response); err != nil {
		return nil, errors.Wrap(err, "to json")
	}
	return &syncV2Response, nil
}
