package errors

var desc = map[Code]string{
	Null: "",

	InsufficientQuota:    "MONEY_INDUFFICIENT",
	OddsChange:           "ODDS_CHANGE",
	LineChange:           "LINE_CHANGE",
	ProportionChange:     "PROPORTION_CHANGE",
	OnceLimit:            "OVER_LIMIT",
	OfferLimit:           "OVER_LIMIT",
	SetLimit:             "OVER_LIMIT",
	MatchLimit:           "OVER_LIMIT",
	PositionLimit:        "OVER_LIMIT",
	ParamWrong:           "PARAM_ERROR",
	PathWrong:            "PATH_WRONG",
	NeedParam:            "NEED_PARAM",
	MatchQuota:           "SETTING_QUOTA",
	OfferQuota:           "SETTING_QUOTA",
	MatchError:           "MATCH_ERROR",
	ItemCount:            "ITEM_COUNT_ERROR",
	ParlayCount:          "PARLAY_COUNT_ERROR",
	ParlaySameMatch:      "PARLAY_SAME_MATCH",
	InsufficientBank:     "MONEY_INDUFFICIENT",
	LockFail:             "SERVER_ERROR",
	TargetNotAccept:      "NOT_ACCEPT",
	ScoreChange:          "SCORE_CHANGE",
	PositionStationLimit: "OVER_LIMIT",
	PositionPlayerLimit:  "OVER_LIMIT",
	ScoreNotFound:        "SCORE_NOT_FOUND",
	OfferError:           "OFFER_ERROR",
	TargetUndefined:      "TARGET_UNDEFINED",
	AmountTooLow:         "AMOUNT_TOO_LOW",
	OddsTypeWrong:        "ODDS_TYPE_WRONG",
	RateModeWrong:        "RATE_MODE_WRONG",
	OfferNotExist:        "OFFER_NOT_EXIST",
	OddsNotExist:         "ODDS_NOT_EXIST",
	OnceQuota:            "SETTING_QUOTA",
	OrderTimeIsUp:        "SETTING_TIME_IS_UP",
	MultipleFail:         "MULTIPLE_FAIL",
	DuplicateNaming:      "DUPLICATE＿NAMING",
	BeingUsed:            "BEING_USED",

	BindFail:    "BIND_FAIL",
	ServerError: "SERVER_ERROR",

	Validation:   "VALIDATION_ERR",
	UserNotExist: "USER_NOT_EXIST",
}