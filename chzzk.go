package chzzk

func New(NidAuth string, NidSession string) ChzzkClient {
	return ChzzkClient{NidAuth: NidAuth, NidSession: NidSession}
}
