import { z } from "zod";

export type User = {
  id: number;
  full_name: string;
  username: string;
  email: string;
  created_at: string;
  last_seen: string;
}

export type LoginDto = {
  username: string;
  password: string;
}

export const registerDtoSchema = z.object({
  full_name: z.string().min(1, "Полное имя не может быть пустым"),
  username: z.string().min(1, "Имя пользователя не может быть пустым"),
  email: z.string().min(1, "Почта не может быть пустой").email("Почта указана неправильно"),
  password: z.string().min(8, "Пароль должен включать 8-24 символа").max(24, "Пароль должен включать 8-24 символа")
})

export type RegisterDto = typeof registerDtoSchema._type

export type MessageResponse = {
  type: number;
} & (
  | { type: 2, new_message: {chat_id: number, message: Message} }
  | { type: 8, user_online: { user_id: number } }
  | { type: 9, user_offline: { user_id: number } }
)

export type LoginResponse = {
  token: string;
  user: User;
}

export type Chat = {
  id: number;
  type: 'private';
  created_at: string;
  name: string;
  messages: Message[];
  members: User[];
  online: boolean;
}

export type Message = {
  id: number;
  content: string;
  created_at: string;
  user_id: number;
  user_full_name: string;
  chat_id: number;
}