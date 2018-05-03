package alicloud

type InstanceType string

const (
	PrivateType  = InstanceType("PRIVATE")
	PublicType   = InstanceType("PUBLIC")
	PrivateType_ = InstanceType("1")
	PublicType_  = InstanceType("0")
)

type DrdsDbEncode string
const(
	UTF8Encode = DrdsDbEncode("utf8")
	GBKEncode = DrdsDbEncode("gbk")
	Latin1Encode = DrdsDbEncode("latin1")
	Utf8mb4Encode = DrdsDbEncode("utf8mb4")
)

type DRDSInstancePayType string

const (
	DRDSInstancePostPayType = DRDSInstancePayType("drdsPost")
)
