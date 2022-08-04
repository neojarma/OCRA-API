package response

// general status
const StatusSuccess = "success"
const StatusFailed = "failed"

// login messages
const MessageSuccessLogin = "success login"
const MessageWrongCredentials = "wrong credential"
const MessageNotVerifedUser = "user email is not verified"

// register
const MessageSuccessRegister = "success register a new user"
const MessageFailedRegister = "failed while creating a new user"
const MessageFailedRegisterEmailExist = "email already exist"

// verify email
const MessageSuccessVerifyEmail = "success verify user email"
const MessageFailedVerifyEmail = "verification token is either expired or invalid"

// not authenticated
const MessageMissingAuthToken = "missing auth token"
const MessageExpiredAuthToken = "token expired"

// resend email verification
const MessageSuccessSendingNewToken = "new token has been sent"
const MessageFailedSendingNewToken = "failed while requesting new token"

// session
const MessageInvalidSession = "invalid session"
const MessageSuccessRenewSession = "success renew session"
const MessageMissingSessionId = "missing session id"
const MessageSuccessLogout = "success logout"
