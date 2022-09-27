import backend from "@/helpers/backend"



class UserService {

    //登入
    login(name:string,passwd:string): Promise<any> {

        const data = {
            account: name,
            password: passwd,
        }
        return backend.post("/api/user/login",data)
    }

    //登出
    logout(): Promise<any> {
        return backend.post("/api/user/logout")
    }

    //获取当前登陆账号
    info(): Promise<any> {
        return backend.get(`/api/user/info`)
    }

    //修改用户信息
    update(id: string | number, data: any): Promise<any> {
        return backend.put(`/api/users/${id}`,data)
    }


    //获取用户列表
    list(query: any): Promise<any> {
        return backend.get("/api/user/search")
    }



}

export default new UserService();