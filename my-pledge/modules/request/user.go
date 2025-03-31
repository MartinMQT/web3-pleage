package request

type Login struct {
	Name     string `from:"name" binding:"必须的"`
	Password string `from:"password" binding:"必须的"`
}
