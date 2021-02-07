package database

import (
	"WinterHomework/utilities"
	"fmt"
)

//获取视频id
func VideoId(bvCode string) int64 {
	var id int64
	pre:=fmt.Sprintf("select id from videos_information where bv_code ='%s",bvCode)
	rows,err:=DB.Query(pre)
	if err != nil {
		utilities.LogError("GetVideoId Error",err)
	}

	if rows.Next(){
		rows.Scan(&id)
	}
	return id
}

//获取视频简要信息
func BriefVideoInformation(start int,end int) []utilities.VideoInformation {
	var (
		tem utilities.VideoInformation
		res []utilities.VideoInformation
	)
	pre:=fmt.Sprintf("select bv_code,video_path,cover_path,title,brief,plays from videos_information order by plays desc limit %d,%d",start,end)
	rows,err:=DB.Query(pre)
	defer rows.Close()
	if err != nil {
		utilities.LogError("GetBriefVideoInformation Error",err)
	}

	for rows.Next(){
		rows.Scan(&tem.BvCode,&tem.VideoPath,&tem.CoverPath,&tem.Title,&tem.Brief,&tem.Plays)
		res=append(res,tem)
	}
	return res
}

//获取视频评论
func VideoComments(videoId int64,limit string) []utilities.MetaComment {
	var(
		temMeta utilities.MetaComment				//存储一级评论的临时体
		temReply utilities.ReplyComment			//存储楼中楼评论的临时体
		res   []utilities.MetaComment			//结果
		temAuthor map[int64]utilities.Author	//用于存储已查询到的用户信息，防止重复查询
	)
	preMeta:=fmt.Sprintf("select id,data_time,content,likes,floor,author_id " +
		"from video_meta_comments where video_id=%d "+limit+" order by id desc",videoId)
	rowsMeta,err:=DB.Query(preMeta)
	defer rowsMeta.Close()
	if err != nil {
		utilities.LogError("GetVideoMetaComments Error",err)
	}

	//获取一级评论
	for rowsMeta.Next(){
		rowsMeta.Scan(&temMeta.Id,&temMeta.Date,&temMeta.Content,&temMeta.Likes,&temMeta.Floor,&temMeta.Author.Id)
		//获取评论用户信息
		//查询是否已存有该ID的用户信息
		value,ok:=temAuthor[temMeta.Author.Id]
		if ok{						//如果查询存在，直接赋值
			temMeta.Author=value
		}else{						//如果不在，进行查找并添加到map中
			temMeta.Author=*commentAuthorInformation(temMeta.Author.Id)
			temAuthor[temMeta.Author.Id]=temMeta.Author
		}

		//获取楼中楼评论
		preReply:=fmt.Sprintf("select id,data_time,content,likes,author_id " +
			"from video_reply_comments where reply_comment_id=%d "+limit+" order by id desc",temMeta.Id)
		rowsReply,err:=DB.Query(preReply)
		if err != nil {
			utilities.LogError("GetVideoReplyComment Error",err)
		}

		for rowsReply.Next(){
			rowsReply.Scan(&temReply.Id,&temReply.Date,&temReply.Content,&temReply.Likes,&temReply.Author.Id,&temReply.ReplyAuthor.Id)
			//获取评论用户信息
			//查询是否已存有该ID的用户信息
			value,ok:=temAuthor[temReply.Author.Id]
			if ok{						//如果查询存在，直接赋值
				temReply.Author=value
			}else{						//如果不在，进行查找并添加到map中
				temReply.Author=*commentAuthorInformation(temReply.Author.Id)
				temAuthor[temReply.Author.Id]=temReply.Author
			}
			temMeta.ReplyComments=append(temMeta.ReplyComments,temReply)
		}
		rowsReply.Close()

		res=append(res,temMeta)
	}
	return res
}
//获取视频评论的附庸组件
//获取视频评论的用户的简要信息，用于获取视频评论时快捷获取其用户信息
func commentAuthorInformation(id int64) *utilities.Author {
	var temAuthor utilities.Author
	temAuthor.Id=id

	preAuthor:=fmt.Sprintf("select name,vip,level from users_information where id = %d",id)
	rowsAuthor,err:=DB.Query(preAuthor)
	defer rowsAuthor.Close()
	if err != nil {
		utilities.LogError("GetCommentAuthor Error",err)
	}

	if rowsAuthor.Next(){
		rowsAuthor.Scan(&temAuthor.Name,&temAuthor.Vip,&temAuthor.Level)
	}

	return &temAuthor
}