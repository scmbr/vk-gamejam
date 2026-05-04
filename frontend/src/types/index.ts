export interface User {
  id: number;
  email: string;
  hasChildProfile: boolean;
  createdAt: string;
}

export interface AuthResponse {
  accessToken: string;
}