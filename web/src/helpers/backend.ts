import axios from "axios";

export default axios.create({
    baseURL:'',
    timeout: 10000,
    withCredentials: true,
    headers:{
        "Content-type":"application/json"
    }
});