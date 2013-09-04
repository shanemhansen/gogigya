package gogigya

type UserList struct {
	UserCount     int `json:"usersCount"`
	IdentityCount int `json:"identitiesCount"`
	Users         []User
}
type User struct {
	UID        string
	Identities []Identity `json:"identities"`
}
type Identity struct {
	Provider    string `json:"provider"`
	ProviderUID string `json:"providerUID"`
}

type QueryResult struct {
	Results      []Account `json:"results"`
	ObjectsCount int       `json:"objectsCount"`
	TotalCount   int       `json:"totalCount"`
	StatusCode   int       `json:"statusCode"`
	ErrorCode    int       `json:"errorCode"`
	ErrorMessage string    `json:"errorMessage"`
	StatusReason string    `json:"statusReason"`
	CallId       string    `json:"callId"`
}
type Account struct {
	UID      string
	LoginIds LoginIds `json:loginIds`
}
type LoginIds struct {
	Emails []string `json:emails`
}

var Methods []string = []string{
	"socialize.Checkin",
	"socialize.checkin",
	"socialize.deleteAccount",
	"socialize.delUserSettings",
	"socialize.exportUsers",
	"socialize.facebookGraphOperation",
	"socialize.getAlbums",
	"socialize.getContacts",
	"socialize.getFeed",
	"socialize.getFriendsInfo",
	"socialize.getPhotos",
	"socialize.getPlaces",
	"socialize.getRawData",
	"socialize.getReactionsCount",
	"socialize.getSessionInfo",
	"socialize.getTopShares",
	"socialize.getUserInfo",
	"socialize.getUserSettings",
	"socialize.incrementReactionsCount",
	"socialize.logout",
	"socialize.notifyLogin",
	"socialize.notifyRegistration",
	"socialize.publishUserAction",
	"socialize.removeConnection",
	"socialize.sendNotification",
	"socialize.setStatus",
	"socialize.setUID",
	"socialize.setUserInfo",
	"socialize.setUserSettings",
	"socialize.shortenURL",
	"accounts.deleteAccount",
	"accounts.deleteScreenSet",
	"accounts.finalizeRegistration",
	"accounts.getAccountInfo",
	"accounts.getPolicies",
	"accounts.getSchema",
	"accounts.getScreenSets",
	"accounts.importProfilePhoto",
	"accounts.initRegistration",
	"accounts.isAvailableLoginID",
	"accounts.linkAccounts",
	"accounts.login",
	"accounts.logout",
	"accounts.notifyLogin",
	"accounts.publishProfilePhoto",
	"accounts.register",
	"accounts.resendVerificationCode",
	"accounts.resetPassword",
	"accounts.search",
	"accounts.setAccountInfo",
	"accounts.setPolicies",
	"accounts.setSchema",
	"accounts.setScreenSet",
	"comments.deleteComment",
	"comments.flagComment",
	"comments.getCategoryInfo",
	"comments.getCommentCounts",
	"comments.getComments",
	"comments.getStreamInfo",
	"comments.getThread",
	"comments.getTopRatedStreams",
	"comments.getTopStreams",
	"comments.getUserComments",
	"comments.getUserOptions",
	"comments.postComment",
	"comments.setCategoryInfo",
	"comments.setStreamInfo",
	"comments.setUserOptions",
	"comments.subscribe",
	"comments.unsubscribe",
	"comments.updateComment",
	"comments.vote",
	"gm.deleteAction",
	"gm.deleteChallenge",
	"gm.deleteChallengeVariant",
	"gm.getActionConfig",
	"gm.getActionsLog",
	"gm.getChallengeConfig",
	"gm.getChallengeStatus",
	"gm.getChallengeVariants",
	"gm.getGlobalConfig",
	"gm.getTopUsers",
	"gm.notifyAction",
	"gm.redeemPoints",
	"gm.resetLevelStatus",
	"gm.setActionConfig",
	"gm.setChallengeConfig",
	"gm.setGlobalConfig",
	"reports.getAccountsStats",
	"reports.getChatStats",
	"reports.getCommentsStats",
	"reports.getFeedStats",
	"reports.getGMStats",
	"reports.getGMTopUsers",
	"reports.getGMUserStats",
	"reports.getIRank",
	"reports.getReactionsStats",
	"reports.getSocializeStats",
	"chat.getMessages",
	"chat.postMessage",
	"ds.delete",
	"ds.get",
	"ds.getSchema",
	"ds.search",
	"ds.setSchema",
	"ds.store",
	"ids.deleteAccount",
	"ids.getAccountInfo",
	"ids.getSchema",
	"ids.search",
	"ids.setAccountInfo",
	"ids.setSchema",
	"gcs.deleteObjectData",
	"gcs.deleteUserData",
	"gcs.getObjectData",
	"gcs.getSchema",
	"gcs.getUserData",
	"gcs.search",
	"gcs.setObjectData",
	"gcs.setSchema",
	"gcs.setUserData",
}
