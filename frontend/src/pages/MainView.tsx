import { Input } from '@/shared/ui/input';
import { createSignal, For, onMount, Show } from 'solid-js';
import { getChats, getMessages, logout } from '@/shared/lib/server';
import { Chat, chats, currentChat } from '@/entities/chat';
import { useStore } from '@nanostores/solid';
import { Chat as ChatType, generateGradient, MessageResponse } from '@/shared/lib';
import { currentUser } from '@/entities/user';
import { Button } from '@/shared/ui/button.tsx';

export const MainView = () => {
  const [isConnected, setIsConnected] = createSignal(false);
  const chatList = useStore(chats);
  const chat = useStore(currentChat);
  const user = useStore(currentUser);
  let ws: WebSocket | undefined = undefined;


  function connectWS() {
    ws?.close();
    ws = undefined

    ws = new WebSocket(
      import.meta.env.VITE_API_URL.replace("https://", "wss://") + "/api/ws"
    );
    ws.onopen = () => {
      setIsConnected(true);
    };

    ws.onclose = () => {
      setIsConnected(false);
      ws = undefined;
      setTimeout(() => connectWS(), 1000);
    };

    ws.onmessage = (e) => {
      const data = JSON.parse(e.data) as MessageResponse;

      switch (data.type) {
        case 2:
          if (data.new_message?.chat_id == chat()?.id) {
            const curr = currentChat.get();
            const newChat: ChatType = {
              ...(curr as ChatType),
              messages: [...(curr?.messages ?? []), data.new_message?.message!],
            };
            currentChat.set(newChat);
          }
          break;
        case 8:
          // online
          let c = chats.get().map(chat => {
            if (chat.type != 'private') return chat
            const s = chat.members[0]
            if (s.id == data.user_online.user_id) {
              chat.online = true
            }
            return chat
          })
          chats.set(c)
          break
        case 9:
          // offline
          let c2 = chats.get().map(chat => {
            if (chat.type != 'private') return chat
            const s = chat.members[0]
            if (s.id == data.user_offline.user_id) {
              chat.online = false
            }
            return chat
          })
          chats.set(c2)

          break
      }
    };
  }



  (async () => {
    connectWS();
    const fetchedChats = await getChats();
    chats.set(fetchedChats ?? []);
  })();

  const [isMobile, setIsMobile] = createSignal(false);
  onMount(() => {
    setIsMobile(window.innerWidth < 768);
  });

  return (
    <div class="flex">
      <Show when={!isMobile() || !chat()}>
        <div class="border-r flex flex-col h-[100dvh] md:min-w-[350px] max-md:w-[100dvw]">
          <div class="h-[50px] p-1 flex gap-1 items-center w-full border-b bg-slate-50">
            <Input placeholder="Поиск..." />
            <div style={{ color: isConnected() ? "green" : "red" }}>///</div>
          </div>
          <div class="p-1 h-full">
            <For each={chatList()}>
              {(iterChat) => (
                <button
                  class={
                    "border p-1 rounded-lg flex items-center gap-2 w-full hover:bg-slate-100 transition " +
                    (chat()?.id == iterChat.id ? "bg-green-50" : "bg-slate-50")
                  }
                  onClick={async () => {
                    const messages = await getMessages(iterChat.id);
                    currentChat.set({ ...iterChat, messages });
                  }}
                >
                  <div
                    class="rounded-full p-4 w-12 h-12 flex items-center justify-center text-white"
                    style={{ background: generateGradient(iterChat.name) }}
                  >
                    {iterChat.name.at(0)}
                  </div>
                  <span>{iterChat.name}</span>
                  <span class='w-2 h-2 rounded-full' style={{background: iterChat.online ? 'green' : 'red'}}></span>
                </button>
              )}
            </For>
          </div>
          <div class="bg-slate-50 border-t p-2 flex justify-between items-center">
            <div class="flex items-center gap-1">
              <div
                class="h-12 w-12 flex items-center justify-center text-white rounded-full"
                style={{ background: generateGradient(user()?.full_name!) }}
              >
                {user()?.full_name.at(0)}
              </div>
              {user()?.full_name}
            </div>
            <Button
              onClick={async () => {
                await logout();
                currentUser.set(undefined);
              }}
            >
              Выйти
            </Button>
          </div>
        </div>
      </Show>
      <Show when={!isMobile() || chat()}><Chat isMobile={isMobile()} ws={ws} /></Show>
    </div>
  );
};
