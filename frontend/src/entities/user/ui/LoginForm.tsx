import { Input } from '@/shared/ui/input';
import { Button } from '@/shared/ui/button';
import { createStore } from 'solid-js/store';
import { login } from '@/shared/lib/server';
import { currentUser } from '@/entities/user';
import { createSignal } from 'solid-js';

export const LoginForm = () => {
  const [data, setData] = createStore({
    username: '',
    password: ''
  })
  const [error, setError] = createSignal("")

  const onSubmit = async () => {
    setError("")
    try {
      const res = await login(data)
      currentUser.set(res.user)}
    catch (e) {
      setError("Неправильный логин или пароль")
    }
  }

  return (
    <>
      <Input placeholder='Имя пользователя...' value={data.username} onChange={e => setData('username', e.target.value)} />
      <Input placeholder='Пароль...' type='password' value={data.password} onChange={e => setData('password', e.target.value)} />
      <Button class='bg-white' onClick={onSubmit}>Войти</Button>
      {error}
    </>
  )
}