package adxhsmarket

import (
	"context"
	"net/http"
)

type CreativityNoteCreateRequest struct {
	AdvertiserId             int64    `json:"advertiser_id"`                        // 是	广告主Id
	UnitId                   int64    `json:"unit_id"`                              // 是	单元Id
	CreativityName           string   `json:"creativity_name"`                      // 是	创意名称
	NoteId                   string   `json:"note_id"`                              // 是	笔记id
	ClickUrls                []string `json:"click_urls,omitempty"`                 // 否	点击链接, 唤端场景不传
	ExpoUrls                 []string `json:"expo_urls,omitempty"`                  // 否	曝光链接，唤端场景不传
	CustomMask               int      `json:"custom_mask,omitempty"`                // 否	自选封面，仅固定首图为封面传2，笔记全部图片均可作为封面传1，（默认0）	智能创意相关字段，支持控制原笔记的图片是否可参与创意优选
	CustomTitle              int      `json:"custom_title,omitempty"`               // 否	是否自提标题 是传1，否传2，（默认0）	智能创意相关字段，支持自定义更多的标题来参与创意优选
	TitleFills               []string `json:"title_fills,omitempty"`                // 否	自提标题，和笔记原标题共计上限10条。	是否自提标题为“2”时，传了无效
	MaskGen                  int      `json:"mask_gen,omitempty"`                   // 否	是否开启自动优化封面 是传1，否传2，（默认0）	智能创意相关字段，开启后，系统将通过对人群特征和海量搜索词的学习自动优化封面，使展现的封面更贴近用户兴趣意图，从而提升投放效果。
	TitleGen                 int      `json:"title_gen,omitempty"`                  // 否	是否开启自动优化标题 是传1，否传2，（默认0）	智能创意相关字段，开启后，系统将通过对人群特征和海量搜索词的学习自动优化标题，使展现的标题更贴近用户兴趣意图，从而提升投放效果。
	MaskPerfer               int      `json:"mask_perfer,omitempty"`                // 否	是否开启封面优选。默认0，开启传1	即将下线，请尽快切换使用新增的「智能创意」相关字段； 新字段不为默认值时，该字段的传参会自动失效
	TitleMaskPerfer          int      `json:"title_mask_perfer,omitempty"`          // 否	是否开启标题优选，默认0，开启传1	即将下线，请尽快切换到新增的字段下； 新字段不为默认值时，该字段的传参会自动失效
	ConversionType           int      `json:"conversion_type"`                      // 是	组件类型，0: 无组件，2: 落地页组件，3: 私信组件，5: poi门店组件，8: 搜索组件，9: 小程序组件，10: 留资组件，11：唤端组件
	JumpURL                  string   `json:"jump_url,omitempty"`                   // 否	落地页/外链url
	LandingPageType          int      `json:"landing_page_type,omitempty"`          // 否	落地页链接类型，1：站内落地页，2：外链
	BarContent               string   `json:"bar_content,omitempty"`                // 否	按钮文案内容，包括： 落地页组件类型：立即参与、立即购买、立即领取、立即预约； 私信组件类型：立即咨询、立即参与、立即领取、立即预约； 唤端组件类型：了解详情，购买同款，领取补贴，优惠下单 小程序组件：立即参与、立即购买、立即领取、立即预约，立即咨询 如果是搜索组件则是搜索词，1~20个字符
	ConversionComponentTypes []int64  `json:"conversion_component_types,omitempty"` // 否	组件位置，0：默认位置，1：置顶评论	两者可以同时选择，默认位置具体为：落地页组件的是图片底表，私信组件的是互动胶囊，搜索组件的是图片或视频；底部唤端组件是互动胶囊小程序组件是互动胶囊
	Comment                  string   `json:"comment,omitempty"`                    // 否	置顶评论文案	当conversion_component_types包含置顶评论时，评论文案
	AppCompIcon              string   `json:"app_comp_icon,omitempty"`              // 否	唤端下，商品主图的地址，必传
	FallBackJumpURL          string   `json:"fall_back_jump_url,omitempty"`         // 否	唤端下，兜底链接，必选
	QualInfo                 QualInfo `json:"qual_info"`                            // 是	创意对应的资质信息,信息见：/api/open/jg/data/qual/info	传了能加快审核速度
	MiniProgramPath          string   `json:"mini_program_path,omitempty"`          // 否	小程序组件链接
}

// CreativityNoteCreate 创建笔记创意
func (s *CreativityService) CreativityNoteCreate(ctx context.Context, req *CreativityNoteCreateRequest, options ...RequestOption) (*CreateCreativityResponse, error) {
	path := "/api/open/jg/creativity/note/create"

	response, err := s.client.Request(ctx, http.MethodPost, path, req, nil, options...)
	if err != nil {
		return nil, err
	}

	result := &CreateCreativityResponse{}
	if err = s.client.JSONUnmarshalBody(response, result); err != nil {
		return nil, err
	}
	return result, nil
}
