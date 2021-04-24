// 作为访问数据库的函数，理论上是没有能力处理 error，
// 应当将 error 向上传递，交给业务层决定：遇上这个 error 应该重试、直接 drop 掉这个请求等...
func daoLevelGet(req string) (string, error) {
	// handle request

	var name string
	_, err := db.QueryRow(req, 1).Scan(&name)
	return name, errors.Wrapf(err, "get request[%s] from SQL failed", req);
}

func DoWork() {
	// ....

	data, err := daoLevelGet(req)
	if err != nil {
		fmt.Printf("original error: %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("start trace:\n%+v\n", err)

		// 取决于业务需要决定要不要直接退出处理该请求
		return
	}

	// ....
}
