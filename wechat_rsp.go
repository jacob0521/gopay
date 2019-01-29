//==================================
//  * Name：Jerry
//  * Tel：18017448610
//  * DateTime：2019/1/13 14:03
//==================================
package gopay

type WeChatUnifiedOrderResponse struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	Appid      string `xml:"appid"`
	MchId      string `xml:"mch_id"`
	DeviceInfo string `xml:"device_info"`
	NonceStr   string `xml:"nonce_str"`
	Sign       string `xml:"sign"`
	ResultCode string `xml:"result_code"`
	ErrCode    string `xml:"err_code"`
	ErrCodeDes string `xml:"err_code_des"`
	TradeType  string `xml:"trade_type"`
	PrepayId   string `xml:"prepay_id"`
	CodeUrl    string `xml:"code_url"`
	MwebUrl    string `xml:"mweb_url"`
}

type WeChatQueryOrderResponse struct {
	ReturnCode         string `xml:"return_code"`
	ReturnMsg          string `xml:"return_msg"`
	Appid              string `xml:"appid"`
	MchId              string `xml:"mch_id"`
	NonceStr           string `xml:"nonce_str"`
	Sign               string `xml:"sign"`
	ResultCode         string `xml:"result_code"`
	ErrCode            string `xml:"err_code"`
	ErrCodeDes         string `xml:"err_code_des"`
	DeviceInfo         string `xml:"device_info"`
	Openid             string `xml:"openid"`
	IsSubscribe        string `xml:"is_subscribe"`
	TradeType          string `xml:"trade_type"`
	TradeState         string `xml:"trade_state"`
	BankType           string `xml:"bank_type"`
	TotalFee           int    `xml:"total_fee"`
	SettlementTotalFee int    `xml:"settlement_total_fee"`
	FeeType            string `xml:"fee_type"`
	CashFee            int    `xml:"cash_fee"`
	CashFeeType        string `xml:"cash_fee_type"`
	CouponFee          int    `xml:"coupon_fee"`
	CouponCount        int    `xml:"coupon_count"`
	CouponType0        string `xml:"coupon_type_0"`
	CouponId0          string `xml:"coupon_id_0"`
	CouponFee0         int    `xml:"coupon_fee_0"`
	TransactionId      string `xml:"transaction_id"`
	OutTradeNo         string `xml:"out_trade_no"`
	Attach             string `xml:"attach"`
	TimeEnd            string `xml:"time_end"`
	TradeStateDesc     string `xml:"trade_state_desc"`
}

type WeChatCloseOrderResponse struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	Appid      string `xml:"appid"`
	MchId      string `xml:"mch_id"`
	DeviceInfo string `xml:"device_info"`
	NonceStr   string `xml:"nonce_str"`
	Sign       string `xml:"sign"`
	ResultCode string `xml:"result_code"`
	ErrCode    string `xml:"err_code"`
	ErrCodeDes string `xml:"err_code_des"`
}

type WeChatRefundResponse struct {
	ReturnCode          string `xml:"return_code"`
	ReturnMsg           string `xml:"return_msg"`
	ResultCode          string `xml:"result_code"`
	ErrCode             string `xml:"err_code"`
	ErrCodeDes          string `xml:"err_code_des"`
	Appid               string `xml:"appid"`
	MchId               string `xml:"mch_id"`
	NonceStr            string `xml:"nonce_str"`
	Sign                string `xml:"sign"`
	TransactionId       string `xml:"transaction_id"`
	OutTradeNo          string `xml:"out_trade_no"`
	OutRefundNo         string `xml:"out_refund_no"`
	RefundId            string `xml:"refund_id"`
	RefundFee           int    `xml:"refund_fee"`
	SettlementRefundFee int    `xml:"settlement_refund_fee"`
	TotalFee            int    `xml:"total_fee"`
	SettlementTotalFee  int    `xml:"settlement_total_fee"`
	FeeType             string `xml:"fee_type"`
	CashFee             int    `xml:"cash_fee"`
	CashFeeType         string `xml:"cash_fee_type"`
	CashRefundFee       int    `xml:"cash_refund_fee"`
	CouponType0         string `xml:"coupon_type_0"`
	CouponRefundFee     int    `xml:"coupon_refund_fee"`
	CouponRefundFee0    int    `xml:"coupon_refund_fee_0"`
	CouponRefundCount   int    `xml:"coupon_refund_count"`
	CouponRefundId0     string `xml:"coupon_refund_id_0"`
}

type WeChatQueryRefundResponse struct {
	ReturnCode           string `xml:"return_code"`
	ReturnMsg            string `xml:"return_msg"`
	ResultCode           string `xml:"result_code"`
	ErrCode              string `xml:"err_code"`
	ErrCodeDes           string `xml:"err_code_des"`
	Appid                string `xml:"appid"`
	MchId                string `xml:"mch_id"`
	NonceStr             string `xml:"nonce_str"`
	Sign                 string `xml:"sign"`
	TotalRefundCount     int    `xml:"total_refund_count"`
	TransactionId        string `xml:"transaction_id"`
	OutTradeNo           string `xml:"out_trade_no"`
	TotalFee             int    `xml:"total_fee"`
	SettlementTotalFee   int    `xml:"settlement_total_fee"`
	FeeType              string `xml:"fee_type"`
	CashFee              int    `xml:"cash_fee"`
	RefundCount          int    `xml:"refund_count"`
	OutRefundNo0         string `xml:"out_refund_no_0"`
	RefundId0            string `xml:"refund_id_0"`
	RefundChannel0       string `xml:"refund_channel_0"`
	RefundFee0           int    `xml:"refund_fee_0"`
	SettlementRefundFee0 int    `xml:"settlement_refund_fee_0"`
	CouponType00         string `xml:"coupon_type_0_0"`
	CouponRefundFee0     int    `xml:"coupon_refund_fee_0"`
	CouponRefundCount0   int    `xml:"coupon_refund_count_0"`
	CouponRefundId00     string `xml:"coupon_refund_id_0_0"`
	CouponRefundFee00    int    `xml:"coupon_refund_fee_0_0"`
	RefundStatus0        string `xml:"refund_status_0"`
	RefundAccount0       string `xml:"refund_account_0"`
	RefundRecvAccout0    string `xml:"refund_recv_accout_0"`
	RefundSuccessTime0   string `xml:"refund_success_time_0"`
}

type getSignKeyResponse struct {
	ReturnCode     string `xml:"return_code"`
	ReturnMsg      string `xml:"return_msg"`
	Retmsg         string `xml:"retmsg"`
	Retcode        string `xml:"retcode"`
	MchId          string `xml:"mch_id"`
	SandboxSignkey string `xml:"sandbox_signkey"`
}

type WeChatNotifyResponse struct {
	ReturnCode         string `xml:"return_code"`
	ReturnMsg          string `xml:"return_msg"`
	ResultCode         string `xml:"result_code"`
	ErrCode            string `xml:"err_code"`
	ErrCodeDes         string `xml:"err_code_des"`
	Appid              string `xml:"appid"`
	MchId              string `xml:"mch_id"`
	DeviceInfo         string `xml:"device_info"`
	NonceStr           string `xml:"nonce_str"`
	Sign               string `xml:"sign"`
	SignType           string `xml:"sign_type"`
	Openid             string `xml:"openid"`
	IsSubscribe        string `xml:"is_subscribe"`
	TradeType          string `xml:"trade_type"`
	TradeState         string `xml:"trade_state"`
	BankType           string `xml:"bank_type"`
	TotalFee           int    `xml:"total_fee"`
	SettlementTotalFee int    `xml:"settlement_total_fee"`
	FeeType            string `xml:"fee_type"`
	CashFee            int    `xml:"cash_fee"`
	CashFeeType        string `xml:"cash_fee_type"`
	CouponFee          int    `xml:"coupon_fee"`
	CouponCount        int    `xml:"coupon_count"`
	CouponType0        string `xml:"coupon_type_0"`
	CouponId0          string `xml:"coupon_id_0"`
	CouponFee0         int    `xml:"coupon_fee_0"`
	TransactionId      string `xml:"transaction_id"`
	OutTradeNo         string `xml:"out_trade_no"`
	Attach             string `xml:"attach"`
	TimeEnd            string `xml:"time_end"`
}