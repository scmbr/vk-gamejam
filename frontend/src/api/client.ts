import axios from 'axios';
import type { AuthResponse } from '../types';

const client = axios.create({
  baseURL: '/api', // через proxy — нет проблем с CORS
  withCredentials: true,
});

// Храним accessToken в памяти — не в localStorage
let accessToken: string | null = null;

export const setAccessToken = (token: string) => {
  accessToken = token;
};

export const getAccessToken = () => accessToken;

// Добавляем токен к каждому запросу
client.interceptors.request.use((config) => {
  if (accessToken) {
    config.headers.Authorization = `Bearer ${accessToken}`;
  }
  return config;
});

// Refresh при 401
client.interceptors.response.use(
  (response) => response,
  async (error) => {
    const original = error.config;

    if (error.response?.status === 401 && !original._retry) {
      original._retry = true;

      try {
        const { data } = await axios.post<AuthResponse>(
          'http://localhost:8000/api/auth/refresh',
          {},
          { withCredentials: true }
        );

        setAccessToken(data.accessToken);
        original.headers.Authorization = `Bearer ${data.accessToken}`;
        return client(original);
      } catch {
        // Refresh протух — на логин
        setAccessToken('');
        window.location.href = '/login';
      }
    }

    return Promise.reject(error);
  }
);

export default client;