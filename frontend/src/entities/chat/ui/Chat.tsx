import { Component, createSignal, For, Show } from 'solid-js';
import { Button } from '@/shared/ui/button.tsx';
import { currentChat } from '@/entities/chat';
import { generateGradient } from '@/shared/lib';
import { Input } from '@/shared/ui/input.tsx';
import { useStore } from '@nanostores/solid';
import { getMessages } from '@/shared/lib/server';

export interface ChatProps {
  isMobile: boolean,
  ws: WebSocket | undefined,
}

export const Chat: Component<ChatProps> = (props) => {
  const [chatDiv, setChatDiv] = createSignal<HTMLDivElement | undefined>();
  const [message, setMessage] = createSignal("");
  const [page, setPage] = createSignal(1);
  const chat = useStore(currentChat);

  const onScroll = async () => {
    const top = chatDiv()?.scrollTop
    if (top == undefined) return;

    if (top <= 50) {
      const messages = await getMessages(chat()?.id!, 25, page() + 1)
      if (!messages) return
      setPage(i => i+1)

      const newChat = {
        ...(currentChat.get() as any),
        messages: [...messages, ...(currentChat.get()?.messages ?? [])]
      }

      currentChat.set(newChat)
      console.log(messages)
      console.log("do work")
    }
  }

  currentChat.subscribe(() => {
    const scrollTopMax = (chatDiv()?.scrollHeight ?? 0) - (chatDiv()?.clientHeight ?? 0)
    const diff = scrollTopMax - (chatDiv()?.scrollTop ?? 0)
    console.log(diff, chatDiv()?.scrollTop)
    if (diff < 100 || (chatDiv()?.scrollTop == 0 && page() == 1)) {
      chatDiv()?.scrollTo({
        top: chatDiv()?.scrollHeight,
        behavior: "smooth",
      });
    } else if (chatDiv()?.scrollTop == 0 && page() != 1) {
      chatDiv()?.scrollTo({
        top: scrollTopMax / 2,
        behavior: "instant",
      });
    }
  })

  const sendMessage = () => {
    if (!props.ws) return;

    props.ws.send(
      JSON.stringify({
        type: 2,
        new_message: {
          chat_id: chat()?.id,
          message: message(),
        },
      })
    );

    setMessage("");
  };

  return (
    <div class="bg-slate-50 w-[100dvw] h-[100dvh] relative">
      <div class="h-[50px] bg-white border-b flex gap-2 items-center p-1">
        <Show when={props.isMobile}>
          <Button onClick={() => currentChat.set(undefined)}>{"<"}</Button>
        </Show>
        {chat() ? chat()?.name : "Выберите чат"}
      </div>
      <div class="w-full">
        <div
          class="flex flex-col gap-1 overflow-y-auto max-h-[calc(100vh-50px-50px)]"
          ref={setChatDiv}
          onScrollEnd={onScroll}
        >
          <For each={chat()?.messages ?? []}>
            {(message) => (
              <div class="p-2 w-full hover:bg-white rounded-lg flex items-center gap-2">
                <div
                  class="h-12 w-12 flex items-center justify-center text-white rounded-full"
                  style={{
                    background: generateGradient(message.user_full_name),
                  }}
                >
                  {message.user_full_name.at(0)}
                </div>
                <div class="flex flex-col">
                      <span class="text-slate-600 text-sm">
                        {message.user_full_name},{" "}
                        {new Date(message.created_at).toLocaleString()}
                      </span>
                  <p class="text-md">{message.content}</p>
                </div>
              </div>
            )}
          </For>
        </div>

        <Show when={chat()}>
          <div class="absolute flex gap-2 bottom-2 w-full px-2">
            <Input
              value={message()}
              onKeyDown={(e) => {
                if (e.key == "Enter") {
                  e.preventDefault();
                  sendMessage();
                }
              }}
              placeholder="Введите сообщение и нажмите enter"
              onInput={ (e) => setMessage(e.target.value)}
            />
            <Button onClick={() => sendMessage()}>{">"}</Button>
          </div>
        </Show>
      </div>
    </div>
  )
}