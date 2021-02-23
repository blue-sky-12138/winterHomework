package database

import (
	"WinterHomework/utilities"
	"database/sql"
	"fmt"
	"time"
)

//获取视频id。
func VideosId(bvCode string) (int64, error) {
	var id int64
	pre := fmt.Sprintf("select id from videos_information where bv_code = '%s' ",bvCode)
	rows,err := DB.Query(pre)
	defer rows.Close()
	if err != nil {
		utilities.LogError("GetVideoId Error",err)
		return 0, fmt.Errorf("未知错误#vi0018")
	}

	if rows.Next(){
		rows.Scan(&id)
	}
	return id, nil
}

//获取视频地址。
func VideoPath(videoId int64) (*[]utilities.VideoPathInformation, error) {
	var (
		res []utilities.VideoPathInformation		//返回结果
		tem utilities.VideoPathInformation			//临时存储获取的地址
	)
	pre := fmt.Sprintf("select p,video_path,video_name from videos_path where videos_id = %d ",videoId)
	rows,err := DB.Query(pre)
	defer rows.Close()
	if err != nil {
		utilities.LogError("GetVideoPath Error",err)
		return nil, fmt.Errorf("未知错误#vp0039")
	}

	for rows.Next() {
		rows.Scan(&tem.P,&tem.Path,&tem.Name)
		res = append(res,tem)
	}

	return &res, nil
}

//获取视频的详细信息。
func DetailedVideoInformation(bvCode string) (*utilities.DetailedVideoInformation,error) {
	var (
		id int64												//临时存储视频id
		res utilities.DetailedVideoInformation					//返回结果
		mapAuthor utilities.DetailedVideoAuthorInformation		//视频作者信息结构体
		jointWork int											//判断是否为联合投稿
		temTime time.Time										//用于接收mysql中的日期数据
		temTag string							//临时存储tag
	)

	//获取视频元数据
	preMeta := "select v.id,cover_path,title,brief,plays,date_time,p,author_id,joint_work,type,detailed,url " +
		fmt.Sprintf("from videos_information v left join activity_detailed a on v.activity_id = a.activity_id where bv_code = '%s'",bvCode)
	rowsMeta,err := DB.Query(preMeta)
	defer rowsMeta.Close()
	if err != nil {
		utilities.LogError("GetDetailedVideoMeta Error",err)
		return nil, fmt.Errorf("未知错误#dvi0067")
	}

	if rowsMeta.Next() {
		rowsMeta.Scan(&id,&res.CoverPath,&res.Title,&res.Brief,&res.Plays,&temTime,&res.P,&mapAuthor.Id,&jointWork,
			&res.Type,&res.Activity.Name,&res.Activity.Url)
	}
	fmt.Printf("%#v \n",res)

	//获取时间
	res.Date = temTime.Format("2006-01-02 15:04:05")

	//更新视频是否参加了活动
	if res.Activity.Name == "" {
		res.Activity.Join = 0
	}else {
		res.Activity.Join = 1
	}

	//获取视频制作人信息
	if jointWork == 0 {		//如果不是联合投稿
		err = getCommonVideoAuthorInformation(&mapAuthor)
		if err == nil {		//如果返回空，即获取数据成功
			res.Author = append(res.Author,mapAuthor)		//添加数据到结果中
		}else {
			return nil, err
		}
	}else {					//如果是联合投稿
		err = getJointVideoAuthorInformation(&id,&res.Author)
		if err == nil {		//如果返回为空，即获取数据成功
					//无需操作，已在函数中完成添加
		}else {
			return nil, err
		}
	}

	//获取视频的点赞数、投币数、收藏数、分享数
	preCommon := "select sum(o.likes),sum(o.coins),sum(o.collections),sum(o.shares) from users_operate_videos_relationship o " +
		fmt.Sprintf("where o.videos_id = %d",id)
	rowsCommon, err := DB.Query(preCommon)
	defer rowsCommon.Close()
	if err != nil {
		utilities.LogError("GetDetailedVideoCommon Error",err)
		return nil, fmt.Errorf("未知错误#dvi0098")
	}
	if rowsCommon.Next(){
		rowsCommon.Scan(&res.Likes,&res.Coins,&res.Collections,&res.Shares)
	}

	//获取视频tag
	preTag := "select t.detailed from tags_videos_relation v inner join tags_detailed t " +
		fmt.Sprintf("on v.videos_id = %d and v.tags_id = t.tags_id",id)
	rowsTag, err := DB.Query(preTag)
	if err != nil {
		utilities.LogError("GetDetailedVideoTag Error",err)
		return nil, fmt.Errorf("未知错误#dvi0098")
	}
	defer rowsTag.Close()
	for rowsTag.Next() {
		rowsTag.Scan(&temTag)
		res.Tags = append(res.Tags,temTag)
	}

	//获取评论总数
	err = getCommentsCounts(&id,&res.CommentNumbers)
	if err == nil {		//如果返回为空，即获取数据成功
		//无需操作，已在函数中完成添加
	}else {
		return nil, err
	}

	return &res, nil
}
//仅在本包使用。
//获取视频信息的附属组件。
//获取非联合投稿视频的作者信息。
func getCommonVideoAuthorInformation(author *utilities.DetailedVideoAuthorInformation) error {
	pre := fmt.Sprintf("select name,signature,vip,level,head_path from users_information where id = %d",author.Id)
	rows,err := DB.Query(pre)
	defer rows.Close()
	if err != nil {
		utilities.LogError("GetCommonVideoAuthorInformation Error",err)
		return fmt.Errorf("未知错误#gcvai0122")
	}

	if rows.Next(){
		rows.Scan(&author.Name,&author.Signature,&author.Vip,&author.Level,&author.HeadPath)
	}
	return nil
}
//仅在本包使用。
//获取视频信息的附属组件。
//获取联合投稿视频的作者信息。
func getJointVideoAuthorInformation(videoId *int64, author *[]utilities.DetailedVideoAuthorInformation) error {
	var mapAuthor utilities.DetailedVideoAuthorInformation
	pre := "select u.id,u.name,u.vip,u.head_path,t.detail from " +
		"(joint_video_relationship j inner join users_information u " +
		fmt.Sprintf("on j.videos_id = %d and j.authors_id = u.id) ",*videoId) +
		"inner join targets_detailed t on j.position_id = t.target_id"
	rows,err := DB.Query(pre)
	defer rows.Close()
	if err != nil {
		utilities.LogError("GetJointVideoAuthorInformation Error",err)
		return fmt.Errorf("未知错误#gjvai0143")
	}

	for rows.Next(){
		rows.Scan(&mapAuthor.Id,&mapAuthor.Name,&mapAuthor.Vip,&mapAuthor.HeadPath,&mapAuthor.Position)
		*author = append(*author,mapAuthor)
	}
	return nil
}
//仅在本包使用。
//获取视频信息的附属组件。
//获视视频评论数。
func getCommentsCounts(videoId *int64,counts *int64) error {
	pre := "select count(m.id) count_id from videos_meta_comments m " +
		fmt.Sprintf("where m.video_id = %d ",*videoId) +
		"union all select count(r.id) count_id from videos_reply_comments r " +
		fmt.Sprintf("where r.video_id = %d",*videoId)
	rows,err := DB.Query(pre)
	defer rows.Close()
	if err != nil {
		utilities.LogError("GetCommentsCounts Error",err)
		return fmt.Errorf("未知错误#gcc0164")
	}

	var tem int64
	for rows.Next(){		//扫描出的第一行是元评论总数，第二行是楼中楼评论总数
		rows.Scan(&tem)
		*counts += tem
	}
	return nil
}


//获取视频的简要信息。
//target为搜索关键词，limit为额外的附加限制。
func BriefVideoInformation(target string, limit string) (*[]utilities.BriefVideoInformation, error) {
	var (
		temInf utilities.BriefVideoInformation		//存储单条信息的临时存储变量
		res []utilities.BriefVideoInformation		//返回结果
		temDateTime time.Time						//存储时间的临时存储变量
	)
	pre:="select v.bv_code,v.cover_path,v.title,v.plays,v.date_time,v.author_id,u.name " +
		"from videos_information v join users_information u on v.author_id = u.id " +
		fmt.Sprintf("where match(v.title) against ('%s' in natural language mode ) ",target) + limit
	rows,err:=DB.Query(pre)
	defer rows.Close()
	if err != nil {
		utilities.LogError("GetBriefVideoInformation Error",err)
		return nil, fmt.Errorf("未知错误#bvi0194")
	}

	for rows.Next(){
		rows.Scan(&temInf.BvCode,&temInf.CoverPath,&temInf.Title,&temInf.Plays,&temDateTime,&temInf.Author.Id,&temInf.Author.Name)
		//获取视频发布时间
		temInf.Date = temDateTime.Format("2006-01-02 15:04:05")

		res=append(res,temInf)
	}
	return &res, nil
}

//获取视频评论。
func VideoComments(videoId int64,limit string) (*[]utilities.MetaComment, error) {
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
		return nil, fmt.Errorf("未知错误#vc0230")
	}

	for rowsMeta.Next(){
		rowsMeta.Scan(&temMeta.Id,&temTime,&temMeta.Content,&temMeta.Floor,&temMeta.Author.Id)
		//获取时间
		temMeta.Date = temTime.Format("2006-01-02 15:04:05")
		//获取评论用户信息
		//查询是否已存有该ID的用户信息
		value,ok = mapAuthor[temMeta.Author.Id]
		if ok{						//如果查询存在，直接赋值
			temMeta.Author = value
		}else{						//如果不在，进行查找并添加到map中
			err = commentAuthorInformation(&temMeta.Author)
			if err == nil {		//如果返回为空，即正常
				mapAuthor[temMeta.Author.Id] = temMeta.Author
			}else {
				return nil, err
			}
		}

		//获取点赞总数和热度指数
		temMeta.Likes, temMeta.Heat, err = commentLikes(0,&temMeta.Id)
		if err != nil {		//如果返回不为空
			return nil, err
		}

		//获取楼中楼评论
		preReply := "select id,date_time,content,author_id,reply_author_id from videos_reply_comments " +
			fmt.Sprintf("where reply_comment_id = %d ",temMeta.Id) + " order by id asc"
		rowsReply,err = DB.Query(preReply)
		if err != nil {
			utilities.LogError("GetVideoReplyComment Error",err)
			return nil, fmt.Errorf("未知错误#vc0262")
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
				err = commentAuthorInformation(&temReply.Author)
				if err == nil {		//如果返回为空，即返回正常
					mapAuthor[temReply.Author.Id] = temReply.Author
				}else {
					return nil, err
				}
			}
			//获取被评论者用户信息
			//查询是否已存有该ID的用户信息
			value,ok = mapAuthor[temReply.ReplyAuthor.Id]
			if ok{						//如果查询存在，直接赋值
				temReply.ReplyAuthor = value
			}else{						//如果不在，进行查找并添加到map中
				err = commentAuthorInformation(&temReply.ReplyAuthor)
				if err != nil {
					return nil, err
				}else {
					mapAuthor[temReply.ReplyAuthor.Id] = temReply.ReplyAuthor
				}
			}

			if err == nil {		//如果返回为空，即正常
				//添加到回复切片
				temMeta.ReplyComments = append(temMeta.ReplyComments,temReply)
			}else {
				return nil, fmt.Errorf("未知错误#vc0297")
			}
		}
		rowsReply.Close()

		//添加到结果切片
		res = append(res,temMeta)
	}
	return &res, nil
}
//仅在本包使用。
//获取视频评论的附属组件。
//获取视频评论的用户的简要信息，用于获取视频评论时快捷获取其用户信息。
func commentAuthorInformation(author *utilities.CommentsAuthorInformation) error {
	preAuthor := fmt.Sprintf("select name,vip,level from users_information where id = %d",author.Id)
	rowsAuthor,err := DB.Query(preAuthor)
	defer rowsAuthor.Close()
	if err != nil {
		utilities.LogError("GetCommentAuthor Error",err)
		return fmt.Errorf("未知错误#cai0316")
	}

	if rowsAuthor.Next(){
		rowsAuthor.Scan(&author.Name,&author.Vip,&author.Level)
	}
	return nil
}
//仅在本包使用。
//获取视频评论的附属组件。
//获取视频评论的点赞数和热度数。
//commentType中，0为元评论，1为楼中楼评论。
func commentLikes(commentType int, commentId *int64) (int, int, error) {
	var (
		likes int
		dislikes int
	)
	pre := "select sum(if(likes = 1, 1, 0)),sum(if(likes = -1, 1, 0)) from likes_videos_comments_relationship " +
		fmt.Sprintf("where comments_id = %d and comments_type = %d and likes = 1",*commentId,commentType)
	rows,err := DB.Query(pre)
	defer rows.Close()
	if err != nil {
		utilities.LogError("GetCommentLikes Error",err)
		return 0, 0, fmt.Errorf("未知错误#cl0336")
	}

	if rows.Next(){
		rows.Scan(&likes,&dislikes)
	}

	return likes, likes - dislikes, nil
}

//获取视频的弹幕。
//p为分集数。
func VideoBarrages(videoId int64,p int) (*[]utilities.VideoBarrage, error) {
	var (
		res []utilities.VideoBarrage  		//返回结果
		temBarrage	utilities.VideoBarrage	//临时存储获取的弹幕信息
		temDate time.Time					//临时存储DateTime
	)
	pre := "select id,date_time,video_time,users_id,content,type,size,pattern,color " +
		fmt.Sprintf("from videos_barrages where videos_id = %d and p = %d",videoId,p)
	rows,err := DB.Query(pre)
	defer rows.Close()
	if err != nil {
		utilities.LogError("GetVideoBarrages Error",err)
		return nil, fmt.Errorf("未知错误#vb0359")
	}

	for rows.Next() {
		rows.Scan(&temBarrage.Id,&temDate,&temBarrage.VideoTime,&temBarrage.UsersId,&temBarrage.Content,&temBarrage.Type,
			&temBarrage.Size,&temBarrage.Pattern,&temBarrage.Color)
		//获取DateTime
		temBarrage.DateTime = temDate.Format("2006-01-02 15:04:05")

		res = append(res,temBarrage)
	}

	return &res, nil
}

//获取用户对于视频的操作。
func VideoOperateInformation(userId int64) (*utilities.OperateVideoInformation ,error) {
	var res utilities.OperateVideoInformation
	pre := fmt.Sprintf("select like,coin,collect,share from video_operation where user_id = %d",userId)
	rows,err := DB.Query(pre)
	defer rows.Close()
	if err != nil {
		utilities.LogError("GetVideoOperateInformation Error",err)
		return nil, fmt.Errorf("未知错误#voi0381")
	}
	
	if rows.Next() {
		rows.Scan(&res.Like,&res.Coin,&res.Collect,&res.Share)
	}
	return &res, nil
}

//更新用户对于视频的操作。
func UpdateVideoOperation(userId int64, videoId int64, focus string, detailed string) error {
	var pre string
	ok, err := checkOperation(userId, videoId)
	if err != nil {
		return err
	}

	if ok {		//如果存在记录，则直接更新
		pre = fmt.Sprintf("update users_operate_videos_relationship set %s = %s where videos_id = %d and users_id = %d",
		focus, detailed, videoId, userId)
	}else {		//否则，新插入记录
		pre = fmt.Sprintf("insert users_operate_videos_relationship (%s,videos_id,users_id) value (%s,%d,%d)",
		focus, detailed, videoId, userId)
	}	

	stmt,err := DB.Prepare(pre)
	defer stmt.Close()
	if err != nil {
		fmt.Println(pre)
		utilities.LogError("PrepareUpdateVideoOperation Error",err)
		return fmt.Errorf("未知错误#uvo0444")
	}

	_, err = stmt.Exec()
	if err != nil {
		utilities.LogError("ExecuteUpdateVideoOperation Error",err)
		return err
	}

	return nil
}
//仅在本包使用
//用于判断用户是否进行过操作
func checkOperation(userId int64, videoId int64) (bool, error) {
	pre := fmt.Sprintf("select id from users_operate_videos_relationship where videos_id = %d and users_id = %d",videoId, userId)
	rows,err := DB.Query(pre)
	defer rows.Close()
	if err != nil {
		utilities.LogError("CheckOperation Error",err)
		return false,fmt.Errorf("未知错误#uvo0463")
	}

	if rows.Next(){
		return true, nil
	}

	return false, nil
}

//添加元评论留言。
func AddMetaComment(data *utilities.NewComment) error {
	//获取楼层
	preFloor :=fmt.Sprintf("select floor from videos_meta_comments where video_id = %d order by floor desc limit 1",data.VideoId)
	rows, err := DB.Query(preFloor)
	defer rows.Close()
	if err != nil {
		utilities.LogError("GetPrefloor Error",err)
		return err
	}
	if rows.Next() {
		var floor int
		rows.Scan(&floor)
		data.MetaFloor = floor
	}

	//正式添加评论
	preAdd := fmt.Sprintf("insert videos_meta_comments (video_id,date_time,content,floor,author_id) value (%d,'%s','%s',%d,%d)",
		data.VideoId, data.DateTime, data.Content, data.MetaFloor, data.AuthorId)
	stmt, err := DB.Prepare(preAdd)
	defer stmt.Close()
	if err != nil {
		utilities.LogError("PrepareAddMetaComment Error",err)
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		utilities.LogError("ExecuteAddMetaComment Error",err)
		return err
	}

	return nil
}

//添加回复留言。
func AddReplyComment(data *utilities.NewComment) error {
	pre := fmt.Sprintf("insert videos_reply_comments (video_id,reply_comment_id,date_time,content,author_id,reply_author_id) value (%d,%d,'%s','%s',%d,%d)",
		data.VideoId, data.ReplyCommentId, data.DateTime, data.Content, data.AuthorId, data.ReplyAuthorId)
	stmt, err := DB.Prepare(pre)
	defer stmt.Close()
	if err != nil {
		utilities.LogError("PrepareAddReplyCommentError",err)
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		utilities.LogError("ExecuteAddReplyComment Error",err)
		return err
	}

	return nil
}