import axios from 'axios'

// create an axios instance
const service = axios.create({
    baseURL: process.env.VUE_APP_BASE_API, // url = base url + request url
    // withCredentials: true, // send cookies when cross-domain requests
    timeout: 10000, // request timeout
    headers: {
        'Content-Type': 'application/json'
    }
  })
  

export function UserList() {
    return service({
        url: '/api/v1/users',
        method: 'get',
    })
}
  
export function addUser(data) {
    return service({
        url: '/api/v1/users',
        method: 'post',
        data: data
    })
}

export function updateUser(data) {
    return service({
      url: '/api/v1/users',
      method: 'put',
      data: data
    })
}

export function deleteUser(id) {
    return service({
      url: '/api/v1/users/' + id,
      method: 'delete'
    })
}