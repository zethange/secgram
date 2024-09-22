import { atom } from "nanostores";
import { Chat } from '@/shared/lib';

export const chats = atom<Chat[]>([])
export const currentChat = atom<Chat | undefined>()