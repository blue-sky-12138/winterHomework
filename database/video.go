package database

import (
	"WinterHomework/utilities"
	"database/sql"
	"fmt"
	"time"
)

//获取视频id
func VideosId(bvCode string) int64 {
	var id int64
	pre := fmt.Sprintf("select id from videos_information where bv_code = '%s' ",bvCode)
	rows,err := DB.Query(pre)
	defer rows.Close()
	if err != nil {
		utilities.LogError("GetVideoId Error",err)
	}

	if rows.Next(){
		rows.Scan(&id)
	}
	return id
}

//获取视频地址
//p为分集数
func VideoPath(videoId int64) *[]utilities.VideoPathInformation {
	var (
		res []utilities.VideoPathInformation		//返回结果
		tem utilities.VideoPathInformation			//临时存储获取的地址
	)
	pre := fmt.Sprintf("select video_path,video_name from videos_path where videos_id = %d ",videoId)
	rows,err := DB.Query(pre)
	if err != nil {
		utilities.LogError("GetVideoPath Error",err)
	}

	if rows.Next() {
		rows.Scan(&tem.Path,&tem.Name)
		res = append(res,tem)
	}

	return &res
}

//获取视频的详细信息
func DetailedVideoInformation(bvCode string) *utilities.VideoInformation {
	var (
		res utilities.VideoInformation
		mapAuthor utilities.VideoAuthorInformation		//视频作者信息结构体
		jointWork int									//判断是否为联合投稿
		temTime time.Time								//用于接收mysql中的日期数据
	)

	//获取视频元数据
	preMeta := "select id,cover_path,title,brief,plays,date_time,p,author_id,joint_work " +
		fmt.Sprintf("from videos_information where bv_code = '%s'",bvCode)
	rowsMeta,err := DB.Query(preMeta)
	if err != nil {
		utilities.LogError("GetDetailedVideoMeta Error",err)
	}
	defer rowsMeta.Close()
	if rowsMeta.Next() {
		rowsMeta.Scan(&res.Id,&res.CoverPath,&res.Title,&res.Brief,&res.Plays,&temTime,&res.P,&mapAuthor.Id,&jointWork)
	}
	//获取时间
	res.Date = temTime.Format("2006-01-02 15:04:05")

	//获取视频制作人信息
	if jointWork == 0 {		//如果不是联合投稿
		getCommonVideoAuthorInformation(&mapAuthor)
		res.Author = append(res.Author,mapAuthor)
	}else {					//如果是联合投稿
		getJointVideoAuthorInformation(&res.Id,&res.Author)
	}

	//获取视频的点赞数、投币数、收藏数、分享数
	preCommon := "select sum(o.likes),sum(o.coins),sum(o.collections),sum(o.shares) from users_operate_videos_relationship o " +
		fmt.Sprintf("where o.videos_id = %d",res.Id)
	rowsCommon,err := DB.Query(preCommon)
	defer rowsCommon.Close()
	if err != nil {
		utilities.LogError("GetDetailedVideoCommon Error",err)
	}
	if rowsCommon.Next(){
		rowsCommon.Scan(&res.Likes,&res.Coins,&res.Collections,&res.Shares)
	}

	//获取评论总数
	getCommentsCounts(&res.Id,&res.CommentNumbers)

	return &res
}
//仅在本包使用
//获取视频信息的附属组件
//获取非联合投稿视频的作者信息
func getCommonVideoAuthorInformation(author *utilities.VideoAuthorInformation) {
	pre := fmt.Sprintf("select name,signature,vip,level,head_path from users_information where id = %d",author.Id)
	rows,err := DB.Query(pre)
	if err != nil {
		utilities.LogError("GetCommonVideoAuthorInformation Error",err)
	}

	if rows.Next(){
		rows.Scan(&author.Name,&author.Signature,&author.Vip,&author.Level,&author.HeadPath)
	}
}
//仅在本包使用
//获取视频信息的附属组件
//获取联合投稿视频的作者信息
func getJointVideoAuthorInformation(videoId *int64, author *[]utilities.VideoAuthorInformation) {
	var mapAuthor utilities.VideoAuthorInformation
	pre := "select u.id,u.name,u.vip,u.head_path,t.detail from " +
		"(joint_video_relationship j inner join users_information u " +
		fmt.Sprintf("on j.videos_id = %d and j.authors_id = u.id) ",*videoId) +
		"inner join targets_details t on j.position_id = t.target_id"
	rows,err := DB.Query(pre)
	if err != nil {
		utilities.LogError("GetJointVideoAuthorInformation Error",err)
	}

	for rows.Next(){
		rows.Scan(&mapAuthor.Id,&mapAuthor.Name,&mapAuthor.Vip,&mapAuthor.HeadPath,&mapAuthor.Position)
		*author = append(*author,mapAuthor)
	}
}

//仅在本包使用
//获取视频信息的附属组件
//获视视频评论数
func getCommentsCounts(videoId *int64,counts *int64) {
	pre := "select count(m.id) count_id from videos_meta_comments m " +
		fmt.Sprintf("where m.video_id = %d ",*videoId) +
		"union all select count(r.id) count_id from videos_reply_comments r " +
		fmt.Sprintf("where r.video_id = %d",*videoId)
	rows,err := DB.Query(pre)
	defer rows.Close()
	if err != nil {
		utilities.LogError("GetCommentsCounts Error",err)
	}

	var tem int64
	for rows.Next(){		//扫描出的第一行是元评论总数，第二行是楼中楼评论总数
		rows.Scan(&tem)
		*counts += tem
	}
}


//获取视频的简要信息
//start表示从第几条开始获取，count表示一共获取几条
//如果count为零，则返回
func BriefVideoInformation(start int,count int) *[]utilities.VideoInformation {
	var (
		tem utilities.VideoInformation
		res []utilities.VideoInformation
	)
	pre:="select v.bv_code,v.cover_path,v.title,v.plays,v.date_time,v.author_id,u.name" +
		"from videos_information v join users_information u on v.author_id = u.id " +
		fmt.Sprintf("order by plays desc limit %d,%d",start,count)
	rows,err:=DB.Query(pre)
	defer rows.Close()
	if err != nil {
		utilities.LogError("GetBriefVideoInformation Error",err)
	}

	for rows.Next(){
		rows.Scan(&tem.BvCode,&tem.CoverPath,&tem.Title,&tem.Brief,&tem.Plays)
		res=append(res,tem)
	}
	return &res
}

//获取视频评论
func VideoComments(videoId int64,limit string) *[]utilities.MetaComment {
	var(
		temMeta utilities.MetaComment									//存储一级评论的临时体
		temReply utilities.ReplyComment									//存储楼中楼评论的临时体
		temTime time.Time												//临时存储日期
		res []utilities.MetaComment										//结果
		mapAuthor = make(map[int64]utilities.CommentsAuthorInformation)	//用于存储已查询到的用户信息，防止重复查询
	)
	var (				//直接声明变量，后续操作就不用短变量声明而是直接赋值
		err error
		rowsReply *sql.Rows							//楼中楼评论sql语句
		value utilities.CommentsAuthorInformation	//获取字典内内容
		ok bool										//判断获取字典内用户信息是否成功
	)

	//获取一级评论
	preMeta := "select id,date_time,content,floor,author_id from videos_meta_comments " +
		fmt.Sprintf("where video_id = %d ",videoId) + limit + " order by id desc"
	rowsMeta,err := DB.Query(preMeta)
	defer rowsMeta.Close()
	if err != nil {
		utilities.LogError("GetVideoMetaComments Error",err)
	}

	for rowsMeta.Next(){
		rowsMeta.Scan(&temMeta.Id,&temTime,&temMeta.Content,&temMeta.Floor,&temMeta.Author.Id)
		//获取时间
		temMeta.Date = temTime.Format("2006-01-02 15:04:05")
		//获取评论用户信息
		//查询是否已存有该ID的用户信息
		value,ok = mapAuthor[temMeta.Author.Id]
		if ok{						//如果查询存在，直接赋值
			temMeta.Author=  value
		}else{						//如果不在，进行查找并添加到map中
			commentAuthorInformation(&temMeta.Author)
			mapAuthor[temMeta.Author.Id] = temMeta.Author
		}
		//获取点赞总数
		temMeta.Likes=*commentLikes(0,&temMeta.Id)

		//获取楼中楼评论
		preReply := "select id,date_time,content,author_id,reply_author_id from videos_reply_comments " +
			fmt.Sprintf("where reply_comment_id = %d ",temMeta.Id) + " order by id desc"
		rowsReply,err = DB.Query(preReply)
		if err != nil {
			utilities.LogError("GetVideoReplyComment Error",err)
		}

		for rowsReply.Next(){
			rowsReply.Scan(&temReply.Id,&temTime,&temReply.Content,&temReply.Author.Id,&temReply.ReplyAuthor.Id)
			//获取时间
			temReply.Date = temTime.Format("2006-01-02 15:04:05")
			//获取评论用户信息
			//查询是否已存有该ID的用户信息
			value,ok = mapAuthor[temReply.Author.Id]
			if ok{						//如果查询存在，直接赋值
				temReply.Author = value
			}else{						//如果不在，进行查找并添加到map中
				commentAuthorInformation(&temReply.Author)
				mapAuthor[temReply.Author.Id] = temReply.Author
			}
			//获取被评论者用户信息
			//查询是否已存有该ID的用户信息
			value,ok = mapAuthor[temReply.ReplyAuthor.Id]
			if ok{						//如果查询存在，直接赋值
				temReply.ReplyAuthor = value
			}else{						//如果不在，进行查找并添加到map中
				commentAuthorInformation(&temReply.ReplyAuthor)
				mapAuthor[temReply.ReplyAuthor.Id] = temReply.ReplyAuthor
			}
			//获取点赞总数
			temReply.Likes = *commentLikes(1,&temReply.Id)
			//添加到回复切片
			temMeta.ReplyComments = append(temMeta.ReplyComments,temReply)
		}
		rowsReply.Close()

		//添加到结果切片
		res = append(res,temMeta)
	}
	return &res
}
//仅在本包使用
//获取视频评论的附属组件
//获取视频评论的用户的简要信息，用于获取视频评论时快捷获取其用户信息
func commentAuthorInformation(author *utilities.CommentsAuthorInformation)  {
	preAuthor := fmt.Sprintf("select name,vip,level from users_information where id = %d",author.Id)
	rowsAuthor,err := DB.Query(preAuthor)
	defer rowsAuthor.Close()
	if err != nil {
		utilities.LogError("GetCommentAuthor Error",err)
	}

	if rowsAuthor.Next(){
		rowsAuthor.Scan(&author.Name,&author.Vip,&author.Level)
	}
}
//仅在本包使用
//获取视频评论的附属组件
//获取视频评论的点赞数
//commentType中，0为元评论，1为楼中楼评论
func commentLikes(commentType int, commentId *int64) *int64 {
	var sum int64
	pre := "select sum(likes) as likes_sum from likes_videos_comments_relationship " +
		fmt.Sprintf("where comments_id = %d and comments_type = %d and likes = 1",*commentId,commentType)
	rows,err := DB.Query(pre)
	defer rows.Close()
	if err != nil {
		utilities.LogError("GetCommentLikes Error",err)
	}

	if rows.Next(){
		rows.Scan(&sum)
	}
	return &sum
}

//获取视频的弹幕
//p为分集数
func VideoBarrages(videoId int64,p int) *[]utilities.VideoBarrage {
	var (
		res []utilities.VideoBarrage  		//返回结果
		temBarrage	utilities.VideoBarrage	//临时存储获取的弹幕信息
		temDate time.Time					//临时存储DateTime
	)
	pre := "select id,date_time,video_time,users_id,content,type,size,pattern,color " +
		fmt.Sprintf("from videos_barrages where videos_id = %d and p = %d",videoId,p)
	rows,err := DB.Query(pre)
	if err != nil {
		utilities.LogError("GetVideoBarrages Error",err)
	}

	for rows.Next() {
		rows.Scan(&temBarrage.Id,&temDate,&temBarrage.VideoTime,&temBarrage.UsersId,&temBarrage.Content,&temBarrage.Type,
			&temBarrage.Size,&temBarrage.Pattern,&temBarrage.Color)
		//获取DateTime
		temBarrage.DateTime = temDate.Format("2006-01-02 15:04:05")

		res = append(res,temBarrage)
	}

	return &res
}