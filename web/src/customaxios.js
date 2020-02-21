import axios from 'axios'

const instance = axios.create({
  baseURL: "http://api.teklifsts.tk/v1/"
  //baseURL: "http://127.0.0.1:8090/v1/"
});

export default instance;
