import { Component, createSignal, Match, Switch } from 'solid-js';
import { TabButton } from '@/shared/ui/tabs';
import { currentUser, LoginForm, RegisterForm } from '@/entities/user';
import { useStore } from '@nanostores/solid';
import { MainView } from '@/pages/MainView.tsx';


export const HomeView: Component = () => {
  const user = useStore(currentUser)
  const [activeTab, setActiveTab] = createSignal<0 | 1>(0)

  return (
    <Switch>
      <Match when={!user()}>
        <div class='h-[100dvh] flex items-center justify-center'>
          <div class='grid gap-2 border shadow rounded-2xl p-3 min-w-[350px] bg-slate-50 transition-[height]'>
            <h1 class='text-xl text-center'>Secgram</h1>

            <div class='flex items-center gap-2'>
              <TabButton isActive={activeTab() == 0} onClick={() => setActiveTab(0)}>Войти</TabButton>
              <TabButton isActive={activeTab() == 1} onClick={() => setActiveTab(1)}>Регистрация</TabButton>
            </div>

            <Switch>
              <Match when={activeTab() == 0}>
                <LoginForm />
              </Match>
              <Match when={activeTab() == 1}>
                <RegisterForm />
              </Match>
            </Switch>
          </div>
        </div>
      </Match>
      <Match when={user()}>
        <MainView />
      </Match>
    </Switch>
  );
}