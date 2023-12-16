import axios from 'axios'
// import { config } from './env.config'

const api = axios.create({
  baseURL: 'http://localhost:1323/api/v1',
  headers: {
    'Content-Type': 'application/json'
  }
})

// handleError;

// request.interceptors.response.use(function (config) {
//   const token = localStorage.getItem("merchant_token");
//   config.headers.Authorization = token ? `Bearer ${token}` : "";
//   return config;
// });

export default api
