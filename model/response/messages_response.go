package response

// general status
const StatusSuccess = "success"
const StatusFailed = "failed"

const MessageErrorBindingData = "missmatch value type"

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
const MessageUserIsAlreadyVerified = "user is already verified"

// not authenticated
const MessageMissingAuthToken = "missing auth token"
const MessageExpiredAuthToken = "token expired"
const MessageDifferentUserId = "requested id and cookie are not match"

// resend email verification
const MessageSuccessSendingNewToken = "new token has been sent"
const MessageFailedSendingNewToken = "failed while requesting new token"

// session
const MessageInvalidSession = "invalid session"
const MessageSuccessRenewSession = "success renew session"
const MessageMissingSessionId = "missing session id"
const MessageSuccessLogout = "success logout"

// validate input
const MessageInvalidJsonInput = "invalid value or missing property in body request"

// videos
const MessageInvalidParameter = "invalid query parameter"
const MessageSuccesGetAllVideos = "success get all videos"
const MessageSuccesGetVideo = "success get video data"
const MessageMissingThumbnail = "missing thumbnail in form request"
const MessageMissingVideo = "missing video in form request"
const MessageSuccesUploadVideo = "succes upload video"
const MessageNoVideo = "there is no video with this id"
const MessageMissingVideoId = "missing video id"
const MessageSuccessUpdateViews = "success update views count"

// channels
const MessageNoChannelWithID = "theres is no channel with this id"
const MessageUserAlreadyHasAChannel = "user already has a channel"
const MessageSuccessGetDetailChannel = "success get detail channel"
const MessageSuccessCreateChannel = "success create channel"
const MessageInvalidUserId = "there is no user with this id"

// comments
const MessageInvalidChannelId = "there is no channel with this id"
const MessageSuccesGetAllComment = "success get all comment"
const MessageSuccessCreateComment = "success create comment"
const MessageSuccessUpdateComment = "success update comment"
const MessageSuccessDeleteComment = "success delete comment"

// like and dislike
const MessageSuccesLikeVideo = "success like video"
const MessageSuccessDislikeVideo = "success dislike video"
const MessageAlreadyLikeThisVideo = "already like this video"
const MessageAlreadyDislikeThisVideo = "already dislike this video"
const MessageSuccessDeleteLike = "success delete like record"
const MessageSuccessDeleteDislike = "success delete dislike record"

// subscribe
const MessageSuccesSubscribe = "success subsribe channel"
const MessageSuccesUnsubscribe = "success unsubscribe channel"
const MessageUserAlreadySubscribe = "user is already subsribe this channel"

// user
const MessageSuccessUpdateUserData = "resource updated successfully"
const MessageFailedUpdateUserData = "failed updated resource"

// watch later
const MessageFailedWatchLaterRecordExist = "requested data already exist"
const MessageFailedInsertWacthRecord = "invalid channel id or video id"
const MessageSuccessGetAllWatchLaterRecord = "success get all record"
const MessageSuccessCreateWatchLaterRecord = "success create new record"
const MessageSuccessDeleteWatchLaterRecord = "success delete record"
const MessageFailedResourceExist = "resource already exist"

// watch later
const MessageFailedHistoryRecordExist = "requested data already exist"
const MessageFailedInsertHistory = "invalid channel id or video id"
const MessageSuccessGetAllHistoryRecord = "success get all record"
const MessageSuccessCreateHistoryRecord = "success create new record"
const MessageSuccessDeleteHistoryRecord = "success delete record"
