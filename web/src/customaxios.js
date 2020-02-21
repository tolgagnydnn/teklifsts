import axios from 'axios'


const instance = axios.create({
  baseURL:"http://api.teklifsts.tk/v1/"
});


export default instance;
