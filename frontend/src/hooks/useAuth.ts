import { useState } from 'react';
import client, { setAccessToken } from '../api/client';
import type { AuthResponse } from '../types';

export const useAuth = () => {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const login = async (email: string, password: string): Promise<string | null> => {
    setLoading(true);
    setError(null);
    try {
      const { data } = await client.post<AuthResponse>('/auth/login', { email, password });
      setAccessToken(data.accessToken);
      return data.accessToken;
    } catch {
      setError('Неверный email или пароль');
      return null;
    } finally {
      setLoading(false);
    }
  };

  const register = async (email: string, password: string): Promise<string | null> => {
    setLoading(true);
    setError(null);
    try {
      const { data } = await client.post<AuthResponse>('/auth/register', { email, password });
      setAccessToken(data.accessToken);
      return data.accessToken;
    } catch {
      setError('Ошибка регистрации');
      return null;
    } finally {
      setLoading(false);
    }
  };

  const logout = async () => {
    try {
      await client.post('/auth/logout');
    } finally {
      setAccessToken('');
      window.location.href = '/login';
    }
  };

  return { login, register, logout, loading, error };
};