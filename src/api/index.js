import http from "../utils/request";

export const getData = () => {
    return http.get('/getData')
}
export const registerUser = (data) => {
    return http.post('/register', data)
}
export const loginUser = (data) => {
    return http.post('/loginUser', data)
}
export const getUser = function (params) {
    if (params.name === "") {
        return http({
            method: "get",
            url: "/getUser",
            params: {
                name: "all"
            },
        })
    } else {
        return http({
            method: "get",
            url: "/getUser",
            params: params,
        })
    }
}
export const addUser = (data) => {
    return http.post('/addUser', data)
}
export const editUser = (data) => {
    return http.post('/editUser', data)
}
export const deleteUser = (data) => {
    return http.post('/deleteUser', data)
}
