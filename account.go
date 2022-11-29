package wxpay

type Account struct {
	appID     string
	mchID     string
	apiKey    string
	certData  []byte
	isSandbox bool
}

// 创建微信支付账号
func NewAccount(appID string, mchID string, apiKey string, isSanbox bool) *Account {
	return &Account{
		appID:     appID,
		mchID:     mchID,
		apiKey:    apiKey,
		isSandbox: isSanbox,
	}
}

// 设置证书
func (a *Account) SetCertData(certData []byte) {
	a.certData = certData
}
