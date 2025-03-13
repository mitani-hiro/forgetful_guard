import axios from "axios";

export const apiClient = axios.create({
  baseURL: "http://localhost:8080", // TODO 環境変数で管理
  headers: {
    "Content-Type": "application/json",
  },
  timeout: 5000,
});
