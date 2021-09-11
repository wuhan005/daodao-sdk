// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package daodao

import (
	"net/http"

	"github.com/pkg/errors"

	"github.com/wuhan005/daodao-sdk/internal/request"
)

type InfoResponse struct {
	ServerTime int    `json:"server_time"`
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	Data       struct {
		Info struct {
			UserId                  string `json:"user_id"`
			Headimage               string `json:"headimage"`
			Nick                    string `json:"nick"`
			Gender                  int    `json:"gender"`
			ActivatedRole           int    `json:"activated_role"`
			RecordDays              int    `json:"record_days"`
			ChatBgSrc               string `json:"chat_bg_src"`
			Budget                  string `json:"budget"`
			WeixinNick              string `json:"weixin_nick"`
			Province                string `json:"province"`
			City                    string `json:"city"`
			Location                string `json:"location"`
			TransProvince           string `json:"trans_province"`
			TransCity               string `json:"trans_city"`
			Birthday                string `json:"birthday"`
			Theyear                 string `json:"theyear"`
			Themonth                string `json:"themonth"`
			Theday                  string `json:"theday"`
			ViewNovice              string `json:"view_novice"`
			SignDays                int    `json:"sign_days"`
			LastSignTime            int    `json:"last_sign_time"`
			LastShareTime           int    `json:"last_share_time"`
			Identity                int    `json:"identity"`
			MessageEndTime          int    `json:"message_end_time"`
			SurveyShow              int    `json:"survey_show"`
			SkipMiniShow            int    `json:"skip_mini_show"`
			LikeTimes               int    `json:"like_times"`
			ChatGroupName           string `json:"chat_group_name"`
			ShowStarAutokid         int64  `json:"show_star_autokid"`
			CurrentCurrency         string `json:"current_currency"`
			LikePush                int    `json:"like_push"`
			PendantId               int    `json:"pendant_id"`
			ChatGroupOpen           int    `json:"chat_group_open"`
			ChatGroupTop            int    `json:"chat_group_top"`
			Desc                    string `json:"desc"`
			ChatGroupOpenTimes      int    `json:"chat_group_open_times"`
			DdjzLiteShowStarAutokid int64  `json:"ddjz_lite_show_star_autokid"`
			DdjzLitePendantId       int    `json:"ddjz_lite_pendant_id"`
			PostOfficeBirthday      int    `json:"post_office_birthday"`
			QqNick                  string `json:"qq_nick"`
			WeiboNick               string `json:"weibo_nick"`
			WeiboKey                string `json:"weibo_key"`
			QqKey                   string `json:"qq_key"`
			WeixinKey               string `json:"weixin_key"`
			MpCode                  int    `json:"mp_code"`
			MoExtra                 string `json:"mo_extra"`
			RegisterTime            int    `json:"register_time"`
			ShortUuid               string `json:"short_uuid"`
			InitialStamina          int    `json:"initial_stamina"`
			AuditType               int    `json:"audit_type"`
			AuditPoint              int    `json:"audit_point"`
			SinaKey                 string `json:"sina_key"`
			AppleKey                string `json:"apple_key"`
			IsTeenMode              int    `json:"is_teen_mode"`
			IsVip                   int    `json:"is_vip"`
			FanNum                  int    `json:"fan_num"`
			PostNum                 int    `json:"post_num"`
			FollowNum               int    `json:"follow_num"`
			Background              string `json:"background"`
			AvatarPendant           struct {
				PendantId       int    `json:"pendant_id"`
				PendantName     string `json:"pendant_name"`
				IsVipExclusive  int    `json:"is_vip_exclusive"`
				PendantUrl      string `json:"pendant_url"`
				BackgroundColor string `json:"background_color"`
				IsDefault       int    `json:"is_default"`
			} `json:"avatar_pendant"`
			CertificationInfo struct {
				Icon string `json:"icon"`
				Name string `json:"name"`
				Desc string `json:"desc"`
			} `json:"certification_info"`
		} `json:"info"`
		MoExtra   string `json:"mo_extra"`
		WeixinKey string `json:"weixin_key"`
	} `json:"data"`
	Encrypt int `json:"encrypt"`
}

func (c *Client) Info() (*InfoResponse, error) {
	query := c.DefaultParameters()
	resp, err := request.Get("/user/info").Query(query).Do()
	if err != nil {
		return nil, errors.Wrap(err, "request")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("unexpected response code: %d", resp.StatusCode)
	}

	var infoResponse InfoResponse
	if err := resp.ToJSON(&infoResponse); err != nil {
		return nil, errors.Wrap(err, "to json")
	}
	return &infoResponse, nil
}
