import axios from 'axios'
import { isHttps } from '@/utils/helper'

const apiHost = import.meta.env.VITE_API_HOST

const api = axios.create({
  // baseURL: isHttps() ? 'https://' : 'http://' + apiHost,
  baseURL: apiHost,
  headers: {
    'Content-Type': 'application/json'
    // 'ngrok-skip-browser-warning': true
  }
})

// handleError;

// request.interceptors.response.use(function (config) {
//   const token = localStorage.getItem("merchant_token");
//   config.headers.Authorization = token ? `Bearer ${token}` : "";
//   return config;
// });

export default api
