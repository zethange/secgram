import { LoginDto, LoginResponse, RegisterDto, User } from '@/shared/lib';

export const getCurrentUser = async (): Promise<User> => {
  const res = await fetch(import.meta.env.VITE_API_URL + '/api/users/me')
  if (!res.ok) throw new Error("not authorized")
  return res.json()
}

export const login = async (loginDto: LoginDto): Promise<LoginResponse> => {
  const res = await fetch(import.meta.env.VITE_API_URL + '/api/auth/login', {
    method: "POST",
    body: JSON.stringify(loginDto)
  })
  if (!res.ok) throw new Error('login or password is incorrect')

  return res.json();
}

export const logout = async (): Promise<void> => {

}

export const register = async (registerDto: RegisterDto): Promise<User> => {
  const res = await fetch(import.meta.env.VITE_API_URL + '/api/auth/register', {
    method: "POST",
    body: JSON.stringify(registerDto)
  })

  return res.json()
}