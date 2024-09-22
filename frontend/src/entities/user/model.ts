import { atom } from "nanostores";
import { User } from '@/shared/lib';

export const currentUser = atom<User | undefined>()