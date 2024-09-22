import { Chat, Message } from '@/shared/lib';

export const getChats = async (limit: number = 10, page: number = 1): Promise<Chat[]> => {
  const res = await fetch(import.meta.env.VITE_API_URL + `/api/chats?limit=${limit}&page=${page}`)
  return res.json()
}

export const getMessages = async (chatId: number, limit: number = 25, page: number = 1): Promise<Message[]> => {
  const res = await fetch(import.meta.env.VITE_API_URL + `/api/chats/messages?chatId=${chatId}&limit=${limit}&page=${page}`)
  return res.json()
}