package errors

const (
    // ErrParamIDEmpty is error code for empty param that retreived from 'context'
    ErrParamIsEmpty = iota + 600

    // ErrParamInvalid is error code for invalid request param
    ErrParamIsInvalid

    // ErrRequestDataInvalid is error code for invalid request data
    ErrRequestDataInvalid
)
const (
    // ErrParamEmpty is error code for empty param that retreived from 'context'
    ErrParamIsEmptyMsg = "request parameter is empty"

    // ErrParamInvalidMsg is error message for invalid request param
    ErrParamIsInvalidMsg = "request parameter is invalid"

    // ErrRequestDataInvalid is error code for invalid request data
    ErrRequestDataInvalidMsg = "request data invalid"
)
