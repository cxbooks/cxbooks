import axios from "axios";

export default axios.create({
    baseURL:``,
    withCredentials: true,
    headers:{
        "Content-type":"application/json"
    }
});