import backend from "@/helpers/backend"
import type { RespData } from '@/types';

export const login = (name: string, passwd: string) => {

    const data = {
        account: name,
        password: passwd,
    }

    // return backend.post<RespData>('/api/user/login',data)

    return backend({
        url: '/api/user/login',
        method: 'post',
        data: data
    })
}


export const logout = () => {

    return backend({
        url: '/api/user/logout',
        method: 'post'
    })
}

export const getUserList = (query: any) => {

    return backend({
        url: '/api/user/login',
        method: 'get',
        data: query
    })
}

//修改用户信息
export const updateUser = (id: string | number, data: any) => {
    // return backend(`/api/users/${id}`, data)
}



