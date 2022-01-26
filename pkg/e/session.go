package e

import (
	"sync"
)

type Session struct {
	Cookie      string                 `json:"cookie"`
	ExpireTime  int64                  `json:"expire_time"`
	SessionList map[string]interface{} `json:"session_list"`
	Lock        *sync.Mutex
}

/*
func (s *Session) Set(key string,value interface{}) (err error) {
	// 加锁，防止并行读取
	s.Lock.Lock()

	defer s.Lock.Unlock()
	// 取出session
	sessionStr, err := dao.RedisClient.Get(s.Cookie).Result()
	if err != nil{
		return
	}
	var session Session
	err = json.Unmarshal([]byte(sessionStr), &session)
	if err != nil {
		return
	}
	// 将新的key/value放入session
	session.SessionList[key] = value

	marshalSession , err := json.Marshal(session)



}*/
