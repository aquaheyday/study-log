import axios, { AxiosInstance } from "axios";

interface ApiConfig {
  baseURL: string;
}

const API_CONFIG: ApiConfig = {
  baseURL: "http://localhost:8080/v1",
};

const apiClient: AxiosInstance = axios.create({
  baseURL: API_CONFIG.baseURL,
  timeout: 10000,
  headers: {
    "Content-Type": "application/json",
    Authorization: `Bearer ${localStorage.getItem("token")}`,
  },
});

apiClient.interceptors.request.use(
    (config) => {
      // 요청 전에 실행: 토큰 동적 설정
      const token = localStorage.getItem("token");
      if (token) {
        config.headers.Authorization = `Bearer ${token}`;
      }
      return config;
    },
    (error) => {
      // 요청 오류 처리
      console.error("Request Error:", error);
      return Promise.reject(error);
    }
  );

  apiClient.interceptors.response.use(
    (response) => {
      return response;
    },
    (error) => {
      if (error.response && error.response.status === 401) {
        window.location.href = "/expired-token";
      } else {
        console.error("Response Error:", error);
      }
      return Promise.reject(error);
    }
  );

export default apiClient;
export { API_CONFIG };
