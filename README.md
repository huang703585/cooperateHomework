# cooperateHomework
综合设计

;input:client(java)->server(c/c++)
;output:server->client
Link
(((date: java.util.Date) (isLinked: boolean))
 (data->isLinked))

SignIn
(((UserName: String) (PassWord: String) (SuccessSignIn: bool))
 (UserName+PassWord->SuccessSignIn))

RequestAnnouncement
(((Announcement: String))
 ((void)->Announcement))

RequestMaintenance
(((Success: boolean))
 ((void)->Success))

RequestQuery
(((Result: float))
 ((void)->Result))

RequestContact
(((Success: boolean))
 ((void)->Success))