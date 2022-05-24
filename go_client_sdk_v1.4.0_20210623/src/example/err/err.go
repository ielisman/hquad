package err

import (
	"google.golang.org/grpc/status"
	biostarErr "biostar/service/err"
)

func GetMultiError(err error) *biostarErr.MultiErrorResponse {
	st, ok := status.FromError(err)
	if ok && len(st.Details()) == 1 {
		errInfo, ok := st.Details()[0].(*biostarErr.MultiErrorResponse)

		if ok {
			return errInfo
		}
	}

	return nil
}

